package queue

import (
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeSendNotification = "notification:send"
	TypeExportExcel      = "excel:export"
	TypeCalculateRisk    = "risk:calculate"

	QueueDefault = "default"
	QueueExports = "exports"
)

type SendNotificationPayload struct {
	UserIDs []int64                `json:"user_ids"`
	Title   string                 `json:"title"`
	Body    string                 `json:"body"`
	Data    map[string]any         `json:"data"`
	LogName string                 `json:"log_name,omitempty"`
}

func NewSendNotificationTask(p SendNotificationPayload) (*asynq.Task, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeSendNotification, body, asynq.Queue(QueueDefault), asynq.MaxRetry(3)), nil
}

type ExportExcelPayload struct {
	Resource     string         `json:"resource"`     // e.g. special_equipment
	CompanyID    int64          `json:"company_id"`
	RequestedBy  int64          `json:"requested_by"`
	NotifyUserID int64          `json:"notify_user_id"`
	Filters      map[string]any `json:"filters"`
}

func NewExportExcelTask(p ExportExcelPayload) (*asynq.Task, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeExportExcel, body, asynq.Queue(QueueExports), asynq.MaxRetry(3)), nil
}

type CalculateRiskPayload struct {
	Content string `json:"content"`
}

// NewCalculateRiskTask 每天凌晨重算所有企业风险值。
// MaxRetry=0 避免重算失败后产生重试风暴；Unique(24h) 等价于 Laravel 的 withoutOverlapping。
func NewCalculateRiskTask(p CalculateRiskPayload) (*asynq.Task, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCalculateRisk, body,
		asynq.Queue(QueueDefault),
		asynq.MaxRetry(0),
		asynq.Unique(24*time.Hour),
	), nil
}
