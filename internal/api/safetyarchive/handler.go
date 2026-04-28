package safetyarchive

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/storage"
)

// Handler bundles all sub-resources rendered by the 安全档案 navigation
// section: 5 top-level modules + 6 sub-tabs that the personnel /
// qualifications / safety-investment pages call as separate endpoints.
// One package keeps the dependency wiring tight (router.go injects a
// single Handler) while each module gets its own file for readability.
type Handler struct {
	archSvc              *service.SafetyArchitectureService
	filingSvc            *service.FilingPermissionService
	buildingSvc          *service.BuildingService
	investmentSvc        *service.SecurityInvestmentService
	personnelSvc         *service.PersonnelService
	certStatusSvc        *service.CertificateStatusService
	productionStdSvc     *service.ProductionStandardizationService
	ohsSvc               *service.OccupationalHealthSafetyService
	fundingExecSvc       *service.FundingExecutionService
	safetyMgmtStaffSvc   *service.SafetyManagementStaffService
	specialEquipStaffSvc *service.SpecialEquipmentStaffService
	elimCtrlStaffSvc     *service.EliminationControlStaffService

	// OSS is optional — only safety-architecture pre-signs annex URLs;
	// nil means the handler returns the raw stored URL untouched.
	oss *storage.Client
}

// HandlerDeps groups the 12 service deps so the constructor stays
// readable. Pass `nil` for OSS when uploads aren't configured.
type HandlerDeps struct {
	Architecture          *service.SafetyArchitectureService
	Filing                *service.FilingPermissionService
	Building              *service.BuildingService
	Investment            *service.SecurityInvestmentService
	Personnel             *service.PersonnelService
	CertStatus            *service.CertificateStatusService
	ProductionStd         *service.ProductionStandardizationService
	OHS                   *service.OccupationalHealthSafetyService
	FundingExec           *service.FundingExecutionService
	SafetyMgmtStaff       *service.SafetyManagementStaffService
	SpecialEquipStaff     *service.SpecialEquipmentStaffService
	EliminationCtrlStaff  *service.EliminationControlStaffService
	OSS                   *storage.Client
}

func NewHandler(d HandlerDeps) *Handler {
	return &Handler{
		archSvc:              d.Architecture,
		filingSvc:            d.Filing,
		buildingSvc:          d.Building,
		investmentSvc:        d.Investment,
		personnelSvc:         d.Personnel,
		certStatusSvc:        d.CertStatus,
		productionStdSvc:     d.ProductionStd,
		ohsSvc:               d.OHS,
		fundingExecSvc:       d.FundingExec,
		safetyMgmtStaffSvc:   d.SafetyMgmtStaff,
		specialEquipStaffSvc: d.SpecialEquipStaff,
		elimCtrlStaffSvc:     d.EliminationCtrlStaff,
		oss:                  d.OSS,
	}
}

// flexInt64 unmarshals JSON numbers, JSON strings, null, and empty
// string into an int64. Front-end Select option values are stringified
// before submit, so request fields that the backend models as int64
// need this tolerance — same trick as in basedata package, kept private
// to avoid leaking JSON quirks into the service layer.
type flexInt64 int64

func (f *flexInt64) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" {
		*f = 0
		return nil
	}
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		inner := s[1 : len(s)-1]
		if inner == "" {
			*f = 0
			return nil
		}
		v, err := strconv.ParseInt(inner, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid integer: %q", inner)
		}
		*f = flexInt64(v)
		return nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*f = flexInt64(v)
	return nil
}

func writeAppError(c *gin.Context, err error) {
	if ae, ok := errs.As(err); ok {
		switch ae.Code {
		case errs.CodeNotFound:
			response.NotFound(c, ae.Message)
		case errs.CodeUnprocessable:
			response.Failed(c, http.StatusUnprocessableEntity, ae.Message)
		case errs.CodeForbidden:
			response.Forbidden(c, ae.Message)
		default:
			response.Internal(c, ae.Message)
		}
		return
	}
	response.Internal(c, err.Error())
}
