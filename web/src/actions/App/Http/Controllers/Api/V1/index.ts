import AuthController from './AuthController'
import CompanyController from './CompanyController'
import PersonnelController from './PersonnelController'
import BaseDataController from './BaseDataController'
import UserController from './UserController'
import RoleController from './RoleController'
import CompanyGroupController from './CompanyGroupController'
import PermissionController from './PermissionController'
import EnterpriseTypeController from './EnterpriseTypeController'
import OperationalTypeController from './OperationalTypeController'
import RecordTypeController from './RecordTypeController'
import ProcessTypeController from './ProcessTypeController'
import SpecialTypeController from './SpecialTypeController'
import SpaceWorkTypeController from './SpaceWorkTypeController'
import SafetyManagementTypeController from './SafetyManagementTypeController'
import DangerousCraftTypeController from './DangerousCraftTypeController'
import RegulationLibraryTypeController from './RegulationLibraryTypeController'
import ChecklistLibraryController from './ChecklistLibraryController'
import BuildingController from './BuildingController'
import JobManagementController from './JobManagementController'
import SafetyArchitectureController from './SafetyArchitectureController'
import ProductionStandardizationController from './ProductionStandardizationController'
import OccupationalHealthSafetyController from './OccupationalHealthSafetyController'
import SafetyManagementController from './SafetyManagementController'
import CertificateStatusController from './CertificateStatusController'
import EliminationControlStaffController from './EliminationControlStaffController'
import SkillUnitController from './SkillUnitController'
import DesignUnitController from './DesignUnitController'
import EvaluationUnitController from './EvaluationUnitController'
import ChangeManageController from './ChangeManageController'
import UploadController from './UploadController'
import UploadOssController from './UploadOssController'
import FilingPermissionsController from './FilingPermissionsController'
import SpecialEquipmentController from './SpecialEquipmentController'
import SpecialEquipmentStaffController from './SpecialEquipmentStaffController'
import DangerousEquipmentController from './DangerousEquipmentController'
import HazardousChemicalController from './HazardousChemicalController'
import HazardousProcessController from './HazardousProcessController'
import MajorHazardController from './MajorHazardController'
import MajorHazardousProcessController from './MajorHazardousProcessController'
import SpaceWorkController from './SpaceWorkController'
import OccupationalHazardController from './OccupationalHazardController'
import HealthSystemController from './HealthSystemController'
import HealthFileController from './HealthFileController'
import HealthReportController from './HealthReportController'
import HealthCheckupController from './HealthCheckupController'
import HygieneEvaluationController from './HygieneEvaluationController'
import FireSafetyController from './FireSafetyController'
import FireEquipmentController from './FireEquipmentController'
import FireFacilityController from './FireFacilityController'
import EmergencyDrillController from './EmergencyDrillController'
import EmergencySuppliesController from './EmergencySuppliesController'
import EmergencyTeamController from './EmergencyTeamController'
import ThreeTogetherController from './ThreeTogetherController'
import InternalTrainingController from './InternalTrainingController'
import ExternalTrainingController from './ExternalTrainingController'
import OtherTrainingController from './OtherTrainingController'
import TrainingFileController from './TrainingFileController'
import TrainingVideoController from './TrainingVideoController'
import MeetingController from './MeetingController'
import MeetingSignUpController from './MeetingSignUpController'
import InternalMeetingController from './InternalMeetingController'
import OfficeNotifyController from './OfficeNotifyController'
import NewNotifyController from './NewNotifyController'
import SecurityInvestmentController from './SecurityInvestmentController'
import FundingExecutionController from './FundingExecutionController'
import HiddenDangerInfoController from './HiddenDangerInfoController'
import RectificationSheetController from './RectificationSheetController'
import HazardController from './HazardController'
import RiskLogController from './RiskLogController'
import InherentRiskLogController from './InherentRiskLogController'
import FloatRiskLogController from './FloatRiskLogController'
import RiskDistributionMapController from './RiskDistributionMapController'
import MonthlyInvestigationController from './MonthlyInvestigationController'
import BlocMonthlyInvestigationsController from './BlocMonthlyInvestigationsController'
import AssessDetectController from './AssessDetectController'
import DailyCheckTemplatesController from './DailyCheckTemplatesController'
import DailyChecksController from './DailyChecksController'
import SystemChecksController from './SystemChecksController'
import NormChecksController from './NormChecksController'
import AccidentCheckTemplatesController from './AccidentCheckTemplatesController'
import AccidentChecksController from './AccidentChecksController'
import AddPointsTemplatesController from './AddPointsTemplatesController'
import AddPointsController from './AddPointsController'
import AssessmentSummaryController from './AssessmentSummaryController'
import CheckStatisticsController from './CheckStatisticsController'
import AssessmentTemplatesController from './AssessmentTemplatesController'
import AccidentBulletinsController from './AccidentBulletinsController'
import InformationDraftsController from './InformationDraftsController'
import ExpertInspectionsController from './ExpertInspectionsController'
import RectificationSheetsController from './RectificationSheetsController'
import MonthStatisticsController from './MonthStatisticsController'
import YearStatisticsController from './YearStatisticsController'
import NoticeStatsController from './NoticeStatsController'
import BlocExpertsController from './BlocExpertsController'
import BlocRisksController from './BlocRisksController'
import BlocNoticesController from './BlocNoticesController'
import BlocAccidentBulletinsController from './BlocAccidentBulletinsController'
import TodosController from './TodosController'
import HomesController from './HomesController'
import CompanyMapController from './CompanyMapController'
import EnterpriseReportController from './EnterpriseReportController'
import OnlineInspectionController from './OnlineInspectionController'
import FocusCompaniesController from './FocusCompaniesController'
import RegulationLibraryController from './RegulationLibraryController'
import RisksController from './RisksController'
import MajorHazardMapController from './MajorHazardMapController'
import RectificationSheetCountController from './RectificationSheetCountController'
import SpecialEquipmentStatisticsController from './SpecialEquipmentStatisticsController'

const V1 = {
    AuthController: Object.assign(AuthController, AuthController),
    CompanyController: Object.assign(CompanyController, CompanyController),
    PersonnelController: Object.assign(PersonnelController, PersonnelController),
    BaseDataController: Object.assign(BaseDataController, BaseDataController),
    UserController: Object.assign(UserController, UserController),
    RoleController: Object.assign(RoleController, RoleController),
    CompanyGroupController: Object.assign(CompanyGroupController, CompanyGroupController),
    PermissionController: Object.assign(PermissionController, PermissionController),
    EnterpriseTypeController: Object.assign(EnterpriseTypeController, EnterpriseTypeController),
    OperationalTypeController: Object.assign(OperationalTypeController, OperationalTypeController),
    RecordTypeController: Object.assign(RecordTypeController, RecordTypeController),
    ProcessTypeController: Object.assign(ProcessTypeController, ProcessTypeController),
    SpecialTypeController: Object.assign(SpecialTypeController, SpecialTypeController),
    SpaceWorkTypeController: Object.assign(SpaceWorkTypeController, SpaceWorkTypeController),
    SafetyManagementTypeController: Object.assign(SafetyManagementTypeController, SafetyManagementTypeController),
    DangerousCraftTypeController: Object.assign(DangerousCraftTypeController, DangerousCraftTypeController),
    RegulationLibraryTypeController: Object.assign(RegulationLibraryTypeController, RegulationLibraryTypeController),
    ChecklistLibraryController: Object.assign(ChecklistLibraryController, ChecklistLibraryController),
    BuildingController: Object.assign(BuildingController, BuildingController),
    JobManagementController: Object.assign(JobManagementController, JobManagementController),
    SafetyArchitectureController: Object.assign(SafetyArchitectureController, SafetyArchitectureController),
    ProductionStandardizationController: Object.assign(ProductionStandardizationController, ProductionStandardizationController),
    OccupationalHealthSafetyController: Object.assign(OccupationalHealthSafetyController, OccupationalHealthSafetyController),
    SafetyManagementController: Object.assign(SafetyManagementController, SafetyManagementController),
    CertificateStatusController: Object.assign(CertificateStatusController, CertificateStatusController),
    EliminationControlStaffController: Object.assign(EliminationControlStaffController, EliminationControlStaffController),
    SkillUnitController: Object.assign(SkillUnitController, SkillUnitController),
    DesignUnitController: Object.assign(DesignUnitController, DesignUnitController),
    EvaluationUnitController: Object.assign(EvaluationUnitController, EvaluationUnitController),
    ChangeManageController: Object.assign(ChangeManageController, ChangeManageController),
    UploadController: Object.assign(UploadController, UploadController),
    UploadOssController: Object.assign(UploadOssController, UploadOssController),
    FilingPermissionsController: Object.assign(FilingPermissionsController, FilingPermissionsController),
    SpecialEquipmentController: Object.assign(SpecialEquipmentController, SpecialEquipmentController),
    SpecialEquipmentStaffController: Object.assign(SpecialEquipmentStaffController, SpecialEquipmentStaffController),
    DangerousEquipmentController: Object.assign(DangerousEquipmentController, DangerousEquipmentController),
    HazardousChemicalController: Object.assign(HazardousChemicalController, HazardousChemicalController),
    HazardousProcessController: Object.assign(HazardousProcessController, HazardousProcessController),
    MajorHazardController: Object.assign(MajorHazardController, MajorHazardController),
    MajorHazardousProcessController: Object.assign(MajorHazardousProcessController, MajorHazardousProcessController),
    SpaceWorkController: Object.assign(SpaceWorkController, SpaceWorkController),
    OccupationalHazardController: Object.assign(OccupationalHazardController, OccupationalHazardController),
    HealthSystemController: Object.assign(HealthSystemController, HealthSystemController),
    HealthFileController: Object.assign(HealthFileController, HealthFileController),
    HealthReportController: Object.assign(HealthReportController, HealthReportController),
    HealthCheckupController: Object.assign(HealthCheckupController, HealthCheckupController),
    HygieneEvaluationController: Object.assign(HygieneEvaluationController, HygieneEvaluationController),
    FireSafetyController: Object.assign(FireSafetyController, FireSafetyController),
    FireEquipmentController: Object.assign(FireEquipmentController, FireEquipmentController),
    FireFacilityController: Object.assign(FireFacilityController, FireFacilityController),
    EmergencyDrillController: Object.assign(EmergencyDrillController, EmergencyDrillController),
    EmergencySuppliesController: Object.assign(EmergencySuppliesController, EmergencySuppliesController),
    EmergencyTeamController: Object.assign(EmergencyTeamController, EmergencyTeamController),
    ThreeTogetherController: Object.assign(ThreeTogetherController, ThreeTogetherController),
    InternalTrainingController: Object.assign(InternalTrainingController, InternalTrainingController),
    ExternalTrainingController: Object.assign(ExternalTrainingController, ExternalTrainingController),
    OtherTrainingController: Object.assign(OtherTrainingController, OtherTrainingController),
    TrainingFileController: Object.assign(TrainingFileController, TrainingFileController),
    TrainingVideoController: Object.assign(TrainingVideoController, TrainingVideoController),
    MeetingController: Object.assign(MeetingController, MeetingController),
    MeetingSignUpController: Object.assign(MeetingSignUpController, MeetingSignUpController),
    InternalMeetingController: Object.assign(InternalMeetingController, InternalMeetingController),
    OfficeNotifyController: Object.assign(OfficeNotifyController, OfficeNotifyController),
    NewNotifyController: Object.assign(NewNotifyController, NewNotifyController),
    SecurityInvestmentController: Object.assign(SecurityInvestmentController, SecurityInvestmentController),
    FundingExecutionController: Object.assign(FundingExecutionController, FundingExecutionController),
    HiddenDangerInfoController: Object.assign(HiddenDangerInfoController, HiddenDangerInfoController),
    RectificationSheetController: Object.assign(RectificationSheetController, RectificationSheetController),
    HazardController: Object.assign(HazardController, HazardController),
    RiskLogController: Object.assign(RiskLogController, RiskLogController),
    InherentRiskLogController: Object.assign(InherentRiskLogController, InherentRiskLogController),
    FloatRiskLogController: Object.assign(FloatRiskLogController, FloatRiskLogController),
    RiskDistributionMapController: Object.assign(RiskDistributionMapController, RiskDistributionMapController),
    MonthlyInvestigationController: Object.assign(MonthlyInvestigationController, MonthlyInvestigationController),
    BlocMonthlyInvestigationsController: Object.assign(BlocMonthlyInvestigationsController, BlocMonthlyInvestigationsController),
    AssessDetectController: Object.assign(AssessDetectController, AssessDetectController),
    DailyCheckTemplatesController: Object.assign(DailyCheckTemplatesController, DailyCheckTemplatesController),
    DailyChecksController: Object.assign(DailyChecksController, DailyChecksController),
    SystemChecksController: Object.assign(SystemChecksController, SystemChecksController),
    NormChecksController: Object.assign(NormChecksController, NormChecksController),
    AccidentCheckTemplatesController: Object.assign(AccidentCheckTemplatesController, AccidentCheckTemplatesController),
    AccidentChecksController: Object.assign(AccidentChecksController, AccidentChecksController),
    AddPointsTemplatesController: Object.assign(AddPointsTemplatesController, AddPointsTemplatesController),
    AddPointsController: Object.assign(AddPointsController, AddPointsController),
    AssessmentSummaryController: Object.assign(AssessmentSummaryController, AssessmentSummaryController),
    CheckStatisticsController: Object.assign(CheckStatisticsController, CheckStatisticsController),
    AssessmentTemplatesController: Object.assign(AssessmentTemplatesController, AssessmentTemplatesController),
    AccidentBulletinsController: Object.assign(AccidentBulletinsController, AccidentBulletinsController),
    InformationDraftsController: Object.assign(InformationDraftsController, InformationDraftsController),
    ExpertInspectionsController: Object.assign(ExpertInspectionsController, ExpertInspectionsController),
    RectificationSheetsController: Object.assign(RectificationSheetsController, RectificationSheetsController),
    MonthStatisticsController: Object.assign(MonthStatisticsController, MonthStatisticsController),
    YearStatisticsController: Object.assign(YearStatisticsController, YearStatisticsController),
    NoticeStatsController: Object.assign(NoticeStatsController, NoticeStatsController),
    BlocExpertsController: Object.assign(BlocExpertsController, BlocExpertsController),
    BlocRisksController: Object.assign(BlocRisksController, BlocRisksController),
    BlocNoticesController: Object.assign(BlocNoticesController, BlocNoticesController),
    BlocAccidentBulletinsController: Object.assign(BlocAccidentBulletinsController, BlocAccidentBulletinsController),
    TodosController: Object.assign(TodosController, TodosController),
    HomesController: Object.assign(HomesController, HomesController),
    CompanyMapController: Object.assign(CompanyMapController, CompanyMapController),
    EnterpriseReportController: Object.assign(EnterpriseReportController, EnterpriseReportController),
    OnlineInspectionController: Object.assign(OnlineInspectionController, OnlineInspectionController),
    FocusCompaniesController: Object.assign(FocusCompaniesController, FocusCompaniesController),
    RegulationLibraryController: Object.assign(RegulationLibraryController, RegulationLibraryController),
    RisksController: Object.assign(RisksController, RisksController),
    MajorHazardMapController: Object.assign(MajorHazardMapController, MajorHazardMapController),
    RectificationSheetCountController: Object.assign(RectificationSheetCountController, RectificationSheetCountController),
    SpecialEquipmentStatisticsController: Object.assign(SpecialEquipmentStatisticsController, SpecialEquipmentStatisticsController),
}

export default V1