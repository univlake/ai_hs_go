package queue

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handlers struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

// HandleSendNotification — 历史占位 worker。原实现写入 new_notifies/new_notify_scopes，
// 现在这两张表已重新分配给业务侧的「通知公告 - 通知型」模块（schema 完全不同），
// 队列侧通知通道改为只记录日志，等真正接入推送/邮件时在此处接通。
//
// 没有任何业务路径调用 NewSendNotificationTask；队列只是 worker 启动时注册的两个
// 任务类型之一，保留是为了兼容已部署 worker 的注册表。
func (h *Handlers) HandleSendNotification(_ context.Context, t *asynq.Task) error {
	var p SendNotificationPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	h.Logger.Info("notification task received (no-op)",
		zap.String("title", p.Title),
		zap.Int("recipients", len(p.UserIDs)),
		zap.String("log_name", p.LogName),
	)
	return nil
}

// HandleExportExcel is a stub for the excel-export pipeline. Full implementation
// loads rows, calls into the excel package, uploads to OSS, then chains a
// notification. Phase 6 ships in-process exports for MVP; this hook is here so
// the worker binary has both task types registered.
func (h *Handlers) HandleExportExcel(_ context.Context, t *asynq.Task) error {
	var p ExportExcelPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	h.Logger.Info("export task received (stub)",
		zap.String("resource", p.Resource),
		zap.Int64("company_id", p.CompanyID),
		zap.Int64("notify_user_id", p.NotifyUserID),
	)
	return nil
}

func (h *Handlers) Register(mux *asynq.ServeMux) {
	mux.HandleFunc(TypeSendNotification, h.HandleSendNotification)
	mux.HandleFunc(TypeExportExcel, h.HandleExportExcel)
}
