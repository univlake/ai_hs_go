package queue

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"ai_hs_go/internal/model"
)

type Handlers struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

// HandleSendNotification mirrors SendNotificationJob: persists NewNotify +
// per-user scope rows in a single transaction.
func (h *Handlers) HandleSendNotification(ctx context.Context, t *asynq.Task) error {
	var p SendNotificationPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	dataJSON, _ := json.Marshal(p.Data)
	return h.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		notify := &model.NewNotify{
			Title: p.Title,
			Body:  &p.Body,
			Data:  dataJSON,
		}
		if err := tx.Create(notify).Error; err != nil {
			return err
		}
		scopes := make([]model.NewNotifyScope, 0, len(p.UserIDs))
		for _, uid := range p.UserIDs {
			scopes = append(scopes, model.NewNotifyScope{NotifyID: notify.ID, UserID: uid})
		}
		if len(scopes) > 0 {
			if err := tx.Create(&scopes).Error; err != nil {
				return err
			}
		}
		h.Logger.Info("notification dispatched",
			zap.Int64("notify_id", notify.ID),
			zap.Int("recipients", len(p.UserIDs)),
		)
		return nil
	})
}

// HandleExportExcel is a stub for the excel-export pipeline. Full implementation
// loads rows, calls into the excel package, uploads to OSS, then chains a
// notification. Phase 6 ships in-process exports for MVP; this hook is here so
// the worker binary has both task types registered.
func (h *Handlers) HandleExportExcel(ctx context.Context, t *asynq.Task) error {
	var p ExportExcelPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	h.Logger.Info("export task received (stub)",
		zap.String("resource", p.Resource),
		zap.Int64("company_id", p.CompanyID),
		zap.Int64("notify_user_id", p.NotifyUserID),
	)
	// TODO(phase 6.5): branch on p.Resource → call into internal/excel + storage.
	return nil
}

func (h *Handlers) Register(mux *asynq.ServeMux) {
	mux.HandleFunc(TypeSendNotification, h.HandleSendNotification)
	mux.HandleFunc(TypeExportExcel, h.HandleExportExcel)
}
