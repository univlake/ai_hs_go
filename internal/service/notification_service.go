package service

import (
	"context"
	"strings"
	"time"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/repository"
)

// NotificationService 聚合三个模块（通知公告 / 会议通知 / 信息稿管理）。
//
// 三个模块都遵循「同一张表服务两个视图」的模式：生产视图（receive-only 前端）
// 跟集团视图（管理员前端）共享同一份 list/show/create/update/delete，前端
// 仅靠 receiveOnly 隐藏写入控件。后端只需在 list 时按当前 company_id 同时匹配
// (a) 我作为发起企业  (b) 我作为接收企业（scope.company_id）。
//
// 子表（kids）整体替换语义跟 risk / accident_emergency 模块一致：
// 传 nil 切片 = 不动；传非 nil（含空切片）= 整体替换。
type NotificationService struct {
	db          *gorm.DB
	companyRepo *repository.CompanyRepo
}

func NewNotificationService(db *gorm.DB, companyRepo *repository.CompanyRepo) *NotificationService {
	return &NotificationService{db: db, companyRepo: companyRepo}
}

// ── 通用：接收企业 kid（office / new / meeting 共用） ─────────────────

// ScopeKid — 创建/更新时前端传过来的 kid 条目。
type ScopeKid struct {
	CompanyID        int64
	CompanyName      *string
	CompanyGroupID   *int64
	CompanyGroupName *string
}

// ── 任务型通知（office_notifies） ─────────────────────────────────────

type OfficeNotifyListParams struct {
	Page              int
	PerPage           int
	CompanyID         int64
	NotifyCompanyID   int64
	Name              string
}

func (s *NotificationService) ListOfficeNotify(ctx context.Context, p OfficeNotifyListParams) ([]model.OfficeNotify, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.OfficeNotify{})
	if p.CompanyID > 0 {
		q = q.Where(
			"company_id = ? OR id IN (?)",
			p.CompanyID,
			s.db.WithContext(ctx).Model(&model.OfficeNotifyScope{}).
				Select("office_notify_id").Where("company_id = ?", p.CompanyID),
		)
	}
	if p.NotifyCompanyID > 0 {
		q = q.Where("id IN (?)",
			s.db.WithContext(ctx).Model(&model.OfficeNotifyScope{}).
				Select("office_notify_id").Where("company_id = ?", p.NotifyCompanyID),
		)
	}
	if v := strings.TrimSpace(p.Name); v != "" {
		q = q.Where("name ILIKE ?", "%"+v+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.OfficeNotify
	if err := q.Preload("OfficeNotifyScope").
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *NotificationService) GetOfficeNotify(ctx context.Context, id int64) (*model.OfficeNotify, error) {
	var m model.OfficeNotify
	if err := s.db.WithContext(ctx).Preload("OfficeNotifyScope").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

type OfficeNotifyInput struct {
	Notify *model.OfficeNotify
	Kids   []ScopeKid
}

func (s *NotificationService) CreateOfficeNotify(ctx context.Context, in OfficeNotifyInput) error {
	if in.Notify.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Notify).Error; err != nil {
			return wrapWrite(err)
		}
		return s.replaceOfficeNotifyScopes(tx, in.Notify.ID, in.Kids, in.Notify.UserID, false)
	})
}

func (s *NotificationService) UpdateOfficeNotify(
	ctx context.Context,
	id int64,
	updates map[string]any,
	kids []ScopeKid,
	replaceKids bool,
	uid int64,
) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.OfficeNotify{}).Where("id = ?", id).Updates(updates)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "通知公告不存在")
		}
		if !replaceKids {
			return nil
		}
		return s.replaceOfficeNotifyScopes(tx, id, kids, uid, true)
	})
}

func (s *NotificationService) replaceOfficeNotifyScopes(tx *gorm.DB, notifyID int64, kids []ScopeKid, uid int64, wipe bool) error {
	if wipe {
		if err := tx.Where("office_notify_id = ?", notifyID).Delete(&model.OfficeNotifyScope{}).Error; err != nil {
			return err
		}
	}
	if len(kids) == 0 {
		return nil
	}
	seen := make(map[int64]struct{}, len(kids))
	scopes := make([]model.OfficeNotifyScope, 0, len(kids))
	for _, k := range kids {
		if k.CompanyID == 0 {
			continue
		}
		if _, dup := seen[k.CompanyID]; dup {
			continue
		}
		seen[k.CompanyID] = struct{}{}
		scope := model.OfficeNotifyScope{
			CompanyName:      k.CompanyName,
			CompanyGroupID:   k.CompanyGroupID,
			CompanyGroupName: k.CompanyGroupName,
			OfficeNotifyID:   notifyID,
			CompanyID:        k.CompanyID,
			FeedbackStatus:   "未反馈",
		}
		scope.UserID = uid
		scope.LastUserID = uid
		scopes = append(scopes, scope)
	}
	if len(scopes) == 0 {
		return nil
	}
	return wrapWrite(tx.Create(&scopes).Error)
}

func (s *NotificationService) DeleteOfficeNotify(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("office_notify_id = ?", id).Delete(&model.OfficeNotifyScope{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.OfficeNotify{}, id).Error
	})
}

// PendingOfficeNotifyScopes — 仪表板：本企业尚未反馈的任务型通知。
func (s *NotificationService) PendingOfficeNotifyScopes(ctx context.Context, companyID int64) ([]model.OfficeNotifyScope, error) {
	var rows []model.OfficeNotifyScope
	err := s.db.WithContext(ctx).
		Where("company_id = ?", companyID).
		Where("feedback_status IS NULL OR feedback_status <> ?", "已反馈").
		Preload("OfficeNotify", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, name, comment, feedback_time, alert_others, annex_url, excel_url")
		}).
		Order("id DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// SubmitOfficeNotifyFeedback — 接收企业上传反馈附件并标记已反馈。
func (s *NotificationService) SubmitOfficeNotifyFeedback(ctx context.Context, scopeID int64, annexURL string, uid int64) (*model.OfficeNotifyScope, error) {
	var scope model.OfficeNotifyScope
	if err := s.db.WithContext(ctx).First(&scope, scopeID).Error; err != nil {
		return nil, mapNotFound(err)
	}
	today := model.Date(time.Now())
	updates := map[string]any{
		"feedback_annex_url": annexURL,
		"feedback_time":      today,
		"feedback_status":    "已反馈",
		"last_user_id":       uid,
	}
	if err := s.db.WithContext(ctx).Model(&scope).Updates(updates).Error; err != nil {
		return nil, wrapWrite(err)
	}
	if err := s.db.WithContext(ctx).First(&scope, scopeID).Error; err != nil {
		return nil, err
	}
	return &scope, nil
}

// ── 通知型通知（new_notifies） ────────────────────────────────────────

type NewNotifyListParams struct {
	Page            int
	PerPage         int
	CompanyID       int64
	NotifyCompanyID int64
	Name            string
}

func (s *NotificationService) ListNewNotify(ctx context.Context, p NewNotifyListParams) ([]model.NewNotify, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.NewNotify{})
	if p.CompanyID > 0 {
		q = q.Where(
			"company_id = ? OR id IN (?)",
			p.CompanyID,
			s.db.WithContext(ctx).Model(&model.NewNotifyScope{}).
				Select("new_notify_id").Where("company_id = ?", p.CompanyID),
		)
	}
	if p.NotifyCompanyID > 0 {
		q = q.Where("id IN (?)",
			s.db.WithContext(ctx).Model(&model.NewNotifyScope{}).
				Select("new_notify_id").Where("company_id = ?", p.NotifyCompanyID),
		)
	}
	if v := strings.TrimSpace(p.Name); v != "" {
		q = q.Where("name ILIKE ?", "%"+v+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.NewNotify
	if err := q.Preload("NewNotifyScope").
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *NotificationService) GetNewNotify(ctx context.Context, id int64) (*model.NewNotify, error) {
	var m model.NewNotify
	if err := s.db.WithContext(ctx).Preload("NewNotifyScope").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

type NewNotifyInput struct {
	Notify *model.NewNotify
	Kids   []ScopeKid
}

func (s *NotificationService) CreateNewNotify(ctx context.Context, in NewNotifyInput) error {
	if in.Notify.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Notify).Error; err != nil {
			return wrapWrite(err)
		}
		return s.replaceNewNotifyScopes(tx, in.Notify.ID, in.Kids, in.Notify.UserID, false)
	})
}

func (s *NotificationService) UpdateNewNotify(
	ctx context.Context,
	id int64,
	updates map[string]any,
	kids []ScopeKid,
	replaceKids bool,
	uid int64,
) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.NewNotify{}).Where("id = ?", id).Updates(updates)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "新通知不存在")
		}
		if !replaceKids {
			return nil
		}
		return s.replaceNewNotifyScopes(tx, id, kids, uid, true)
	})
}

func (s *NotificationService) replaceNewNotifyScopes(tx *gorm.DB, notifyID int64, kids []ScopeKid, uid int64, wipe bool) error {
	if wipe {
		if err := tx.Where("new_notify_id = ?", notifyID).Delete(&model.NewNotifyScope{}).Error; err != nil {
			return err
		}
	}
	if len(kids) == 0 {
		return nil
	}
	seen := make(map[int64]struct{}, len(kids))
	scopes := make([]model.NewNotifyScope, 0, len(kids))
	for _, k := range kids {
		if k.CompanyID == 0 {
			continue
		}
		if _, dup := seen[k.CompanyID]; dup {
			continue
		}
		seen[k.CompanyID] = struct{}{}
		scope := model.NewNotifyScope{
			CompanyName:      k.CompanyName,
			CompanyGroupID:   k.CompanyGroupID,
			CompanyGroupName: k.CompanyGroupName,
			NewNotifyID:      notifyID,
			CompanyID:        k.CompanyID,
			ReadStatus:       "未读",
		}
		scope.UserID = uid
		scope.LastUserID = uid
		scopes = append(scopes, scope)
	}
	if len(scopes) == 0 {
		return nil
	}
	return wrapWrite(tx.Create(&scopes).Error)
}

func (s *NotificationService) DeleteNewNotify(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("new_notify_id = ?", id).Delete(&model.NewNotifyScope{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.NewNotify{}, id).Error
	})
}

func (s *NotificationService) PendingNewNotifyScopes(ctx context.Context, companyID int64) ([]model.NewNotifyScope, error) {
	var rows []model.NewNotifyScope
	err := s.db.WithContext(ctx).
		Where("company_id = ?", companyID).
		Where("read_status IS NULL OR read_status <> ?", "已读").
		Preload("NewNotify", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, name, comment, alert_others, annex_url")
		}).
		Order("id DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *NotificationService) MarkNewNotifyRead(ctx context.Context, scopeID int64, uid int64) (*model.NewNotifyScope, error) {
	var scope model.NewNotifyScope
	if err := s.db.WithContext(ctx).First(&scope, scopeID).Error; err != nil {
		return nil, mapNotFound(err)
	}
	today := model.Date(time.Now())
	updates := map[string]any{
		"read_time":    today,
		"read_status":  "已读",
		"last_user_id": uid,
	}
	if err := s.db.WithContext(ctx).Model(&scope).Updates(updates).Error; err != nil {
		return nil, wrapWrite(err)
	}
	if err := s.db.WithContext(ctx).First(&scope, scopeID).Error; err != nil {
		return nil, err
	}
	return &scope, nil
}

// ── 会议通知（meetings + meeting_notices + meeting_sign_ups） ─────────

type MeetingListParams struct {
	Page            int
	PerPage         int
	CompanyID       int64
	NotifyCompanyID int64
	Name            string
	StartDate       string
	EndDate         string
}

func (s *NotificationService) ListMeeting(ctx context.Context, p MeetingListParams) ([]model.Meeting, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.Meeting{})
	if p.CompanyID > 0 {
		q = q.Where(
			"company_id = ? OR id IN (?)",
			p.CompanyID,
			s.db.WithContext(ctx).Model(&model.MeetingNotice{}).
				Select("meeting_id").Where("company_id = ?", p.CompanyID),
		)
	}
	if p.NotifyCompanyID > 0 {
		q = q.Where("id IN (?)",
			s.db.WithContext(ctx).Model(&model.MeetingNotice{}).
				Select("meeting_id").Where("company_id = ?", p.NotifyCompanyID),
		)
	}
	if v := strings.TrimSpace(p.Name); v != "" {
		q = q.Where("name ILIKE ?", "%"+v+"%")
	}
	if v := strings.TrimSpace(p.StartDate); v != "" {
		q = q.Where("meeting_date >= ?", v)
	}
	if v := strings.TrimSpace(p.EndDate); v != "" {
		q = q.Where("meeting_date <= ?", v)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.Meeting
	if err := q.Preload("MeetingNotices").
		Order("meeting_date DESC NULLS LAST").Order("id DESC").
		Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *NotificationService) GetMeeting(ctx context.Context, id int64) (*model.Meeting, error) {
	var m model.Meeting
	if err := s.db.WithContext(ctx).
		Preload("MeetingNotices.Company").
		Preload("MeetingSignUps").
		First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

type MeetingInput struct {
	Meeting *model.Meeting
	Kids    []ScopeKid
}

func (s *NotificationService) CreateMeeting(ctx context.Context, in MeetingInput) error {
	if in.Meeting.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Meeting).Error; err != nil {
			return wrapWrite(err)
		}
		return s.replaceMeetingNotices(tx, in.Meeting.ID, in.Kids, in.Meeting.UserID, false)
	})
}

func (s *NotificationService) UpdateMeeting(
	ctx context.Context,
	id int64,
	updates map[string]any,
	kids []ScopeKid,
	replaceKids bool,
	uid int64,
) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.Meeting{}).Where("id = ?", id).Updates(updates)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "会议不存在")
		}
		if !replaceKids {
			return nil
		}
		return s.replaceMeetingNotices(tx, id, kids, uid, true)
	})
}

func (s *NotificationService) replaceMeetingNotices(tx *gorm.DB, meetingID int64, kids []ScopeKid, uid int64, wipe bool) error {
	if wipe {
		if err := tx.Where("meeting_id = ?", meetingID).Delete(&model.MeetingNotice{}).Error; err != nil {
			return err
		}
	}
	if len(kids) == 0 {
		return nil
	}
	seen := make(map[int64]struct{}, len(kids))
	notices := make([]model.MeetingNotice, 0, len(kids))
	for _, k := range kids {
		if k.CompanyID == 0 {
			continue
		}
		if _, dup := seen[k.CompanyID]; dup {
			continue
		}
		seen[k.CompanyID] = struct{}{}
		notice := model.MeetingNotice{
			CompanyName:      k.CompanyName,
			MeetingID:        meetingID,
			CompanyID:        k.CompanyID,
			CompanyGroupID:   k.CompanyGroupID,
			CompanyGroupName: k.CompanyGroupName,
		}
		notice.UserID = uid
		notice.LastUserID = uid
		notices = append(notices, notice)
	}
	if len(notices) == 0 {
		return nil
	}
	return wrapWrite(tx.Create(&notices).Error)
}

func (s *NotificationService) DeleteMeeting(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("meeting_id = ?", id).Delete(&model.MeetingNotice{}).Error; err != nil {
			return err
		}
		if err := tx.Where("meeting_id = ?", id).Delete(&model.MeetingSignUp{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.Meeting{}, id).Error
	})
}

func (s *NotificationService) NotifyMeetingNotice(ctx context.Context, noticeID int64, uid int64) (*model.MeetingNotice, error) {
	var notice model.MeetingNotice
	if err := s.db.WithContext(ctx).First(&notice, noticeID).Error; err != nil {
		return nil, mapNotFound(err)
	}
	now := model.DateTime(time.Now())
	if err := s.db.WithContext(ctx).Model(&notice).Updates(map[string]any{
		"notify_time":  now,
		"last_user_id": uid,
	}).Error; err != nil {
		return nil, wrapWrite(err)
	}
	if err := s.db.WithContext(ctx).First(&notice, noticeID).Error; err != nil {
		return nil, err
	}
	return &notice, nil
}

func (s *NotificationService) ConfirmMeetingNotice(ctx context.Context, noticeID int64, uid int64) (*model.MeetingNotice, error) {
	var notice model.MeetingNotice
	if err := s.db.WithContext(ctx).First(&notice, noticeID).Error; err != nil {
		return nil, mapNotFound(err)
	}
	now := model.DateTime(time.Now())
	if err := s.db.WithContext(ctx).Model(&notice).Updates(map[string]any{
		"confirm_time": now,
		"last_user_id": uid,
	}).Error; err != nil {
		return nil, wrapWrite(err)
	}
	if err := s.db.WithContext(ctx).First(&notice, noticeID).Error; err != nil {
		return nil, err
	}
	return &notice, nil
}

func (s *NotificationService) PendingMeetingNotices(ctx context.Context, companyID int64) ([]model.MeetingNotice, error) {
	var rows []model.MeetingNotice
	err := s.db.WithContext(ctx).
		Where("company_id = ?", companyID).
		Where("confirm_time IS NULL").
		Where("notify_time IS NULL OR notify_time <= ?", time.Now()).
		Preload("Meeting", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, name, meeting_date, sign_end_date, place, theme")
		}).
		Order("notify_time DESC NULLS LAST").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// ── 会议报名（meeting_sign_ups） ──────────────────────────────────────

type MeetingSignUpListParams struct {
	Page      int
	PerPage   int
	MeetingID int64
	CompanyID int64
}

func (s *NotificationService) ListMeetingSignUp(ctx context.Context, p MeetingSignUpListParams) ([]model.MeetingSignUp, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.MeetingSignUp{})
	if p.MeetingID > 0 {
		q = q.Where("meeting_id = ?", p.MeetingID)
	}
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.MeetingSignUp
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *NotificationService) GetMeetingSignUp(ctx context.Context, id int64) (*model.MeetingSignUp, error) {
	var m model.MeetingSignUp
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *NotificationService) CreateMeetingSignUp(ctx context.Context, m *model.MeetingSignUp) error {
	if m.CompanyID == 0 || m.MeetingID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id 或 meeting_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *NotificationService) UpdateMeetingSignUp(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.MeetingSignUp{}, id, updates)
}

func (s *NotificationService) DeleteMeetingSignUp(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.MeetingSignUp{}, id).Error
}

// ── 信息稿（information_drafts） ──────────────────────────────────────
//
// 信息稿件管理的 list 接收 scope 参数：
//   "published" — 我发布的（company_id = 我；含已上报的稿件，由前端隐藏写入）
//   "received"  — 我收到的（下级上报 next_reviewer = 我；上级已发布在祖先链上 status=1）
//   ""          — 通用：自己稿件 + 上级已发布稿件
//
// 集团视图（receiveOnly 前端）跟生产视图共享同一组路由，前端切换 scope 即可。
// dashboardReadings / dashboardPending 给仪表板用，跟 Laravel 控制器同名方法对齐。

type InformationDraftListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Scope     string
	Status    string // "" 表示不过滤；前端目前不传该参数
	Name      string
}

func (s *NotificationService) ListInformationDraft(ctx context.Context, p InformationDraftListParams) ([]model.InformationDraft, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.InformationDraft{})
	if p.CompanyID > 0 {
		ancestors := s.ancestorChain(ctx, p.CompanyID)
		switch p.Scope {
		case "published":
			q = q.Where("company_id = ?", p.CompanyID)
		case "received":
			upstream := filterOut(ancestors, p.CompanyID)
			q = q.Where("company_id <> ?", p.CompanyID).
				Where(s.db.Where("next_reviewer = ?", p.CompanyID).
					Or(s.db.Where("next_reviewer IN ?", upstream).Where("status = ?", 1)))
		default:
			q = q.Where(
				s.db.Where("next_reviewer = ?", p.CompanyID).
					Or(s.db.Where("next_reviewer IN ?", ancestors).
						Where("next_reviewer <> ?", p.CompanyID).
						Where("status = ?", 1)),
			)
		}
	}
	if v := strings.TrimSpace(p.Status); v != "" {
		q = q.Where("status = ?", v)
	}
	if v := strings.TrimSpace(p.Name); v != "" {
		q = q.Where("name ILIKE ?", "%"+v+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.InformationDraft
	if err := q.Order("is_top DESC").Order("id DESC").
		Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *NotificationService) GetInformationDraft(ctx context.Context, id int64) (*model.InformationDraft, error) {
	var m model.InformationDraft
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *NotificationService) CreateInformationDraft(ctx context.Context, m *model.InformationDraft) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *NotificationService) UpdateInformationDraft(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.InformationDraft{}, id, updates)
}

func (s *NotificationService) DeleteInformationDraft(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.InformationDraft{}, id).Error
}

// DashboardInformationDraftReadings — 仪表板「最新动态」：
// next_reviewer 在本企业祖先链上、status=1 的稿件（含本企业发布的、上级发布的、
// 下级上报由上级复核发布的）。
func (s *NotificationService) DashboardInformationDraftReadings(ctx context.Context, companyID int64, limit int) ([]model.InformationDraft, error) {
	if limit <= 0 || limit > 100 {
		limit = 6
	}
	ancestors := s.ancestorChain(ctx, companyID)
	var rows []model.InformationDraft
	err := s.db.WithContext(ctx).
		Where("next_reviewer IN ?", ancestors).
		Where("status = ?", 1).
		Order("is_top DESC").Order("id DESC").
		Limit(limit).Find(&rows).Error
	return rows, err
}

// DashboardInformationDraftPending — 仪表板「待办」：下级上报且未处理的稿件。
func (s *NotificationService) DashboardInformationDraftPending(ctx context.Context, companyID int64, limit int) ([]model.InformationDraft, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	var rows []model.InformationDraft
	err := s.db.WithContext(ctx).
		Where("next_reviewer = ?", companyID).
		Where("company_id <> ?", companyID).
		Where("status = ?", 0).
		Order("id DESC").
		Limit(limit).Find(&rows).Error
	return rows, err
}

// ── 私有 helpers ─────────────────────────────────────────────────────

// ancestorChain — 跟 Laravel BlocNoticesController.ancestorChain / Information-
// DraftsController.ancestorChain 同语义：从 companyID 出发，沿 parent_id 上溯，
// 返回 [companyID, parent, grandparent, ...]。出错时退化为只含 companyID 的列表，
// 仪表板查询不至于全失败。
func (s *NotificationService) ancestorChain(ctx context.Context, companyID int64) []int64 {
	chain := []int64{companyID}
	current := companyID
	for i := 0; i < 32; i++ {
		c, err := s.companyRepo.GetByID(ctx, current)
		if err != nil || c == nil || c.ParentID == nil || *c.ParentID == 0 {
			break
		}
		current = *c.ParentID
		chain = append(chain, current)
	}
	return chain
}

func filterOut(ids []int64, exclude int64) []int64 {
	out := make([]int64, 0, len(ids))
	for _, id := range ids {
		if id != exclude {
			out = append(out, id)
		}
	}
	return out
}
