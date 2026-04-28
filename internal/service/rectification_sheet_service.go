// Package service — 整改通知单 + 隐患分类统计 Excel 生成。
//
// 跟 Laravel RectificationSheetsController 对应：
//   buildRectificationSheet  → BuildRectificationSheet（单 expert_company 维度）
//   buildHazardStatistics    → BuildHazardStatistics  （单 company 维度，可按 type 过滤）
//
// 用 excelize 实现 PhpSpreadsheet 的边框 + 合并单元格 + 行高 + 图片嵌入。
// 图片走 OSS signed URL；OSS 未配置或下载失败时跳过该单元格的图片，文本仍正常写入。
package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/storage"
)

// RectificationSheetService 实例化时不持有数据 — 调用时按需走 db / oss 拉取。
type RectificationSheetService struct {
	db          *gorm.DB
	oss         *storage.Client // 可为 nil；nil 时图片直接跳过
	storageRoot string          // <root>/rectification-sheets/YYYY/MM/<file>
}

func NewRectificationSheetService(db *gorm.DB, oss *storage.Client, storageRoot string) *RectificationSheetService {
	return &RectificationSheetService{db: db, oss: oss, storageRoot: storageRoot}
}

// 隐患七大类别 — 跟 Laravel 的 HIDDEN_TYPES 常量保持一致。
var hiddenTypes = []string{"危化品", "电气", "机械", "安全设施", "消防", "职业健康", "其他"}

// ── 整改通知单（单 expert_company） ───────────────────────────────────

// GenerateRectificationSheet — 生成单企业整改通知单 .xlsx，落盘后写 expert_companies.rectification_sheet_url。
//
// 权限：handler 层负责（计划方 / 已指派专家可生成 + 下载，受检企业只可下载）。
func (s *RectificationSheetService) GenerateRectificationSheet(ctx context.Context, expertCompanyID int64) (*model.ExpertCompany, error) {
	var ec model.ExpertCompany
	err := s.db.WithContext(ctx).
		Preload("ExpertInspection").
		Preload("CheckCompany").
		Preload("ExpertInspectionUsers.ExpertUser").
		First(&ec, expertCompanyID).Error
	if err != nil {
		return nil, mapNotFound(err)
	}

	var dangers []model.HiddenDangerInfo
	if ec.ExpertInspectionID != 0 {
		if err := s.db.WithContext(ctx).
			Where("expert_inspection_id = ? AND expert_company_id = ?", ec.ExpertInspectionID, ec.ID).
			Order("id ASC").
			Find(&dangers).Error; err != nil {
			return nil, err
		}
	}
	// 加载附件（attachAnnexes 在 hidden_danger_service 里是非导出方法，这里复用同样的 SQL 拉一次）。
	if err := s.attachDangerAnnexes(ctx, dangers); err != nil {
		return nil, err
	}

	checkCompanyName := ""
	supervisor := ""
	supervisorTel := ""
	if ec.CheckCompany != nil {
		checkCompanyName = ec.CheckCompany.Title
		if ec.CheckCompany.SuperviseName != nil {
			supervisor = *ec.CheckCompany.SuperviseName
		}
		if ec.CheckCompany.SuperviseTel != nil {
			supervisorTel = *ec.CheckCompany.SuperviseTel
		}
	}
	checkTime := ""
	if ec.CheckCompanyTime != nil {
		checkTime = time.Time(*ec.CheckCompanyTime).Format("2006-01-02 15:04:05")
	}

	f := excelize.NewFile()
	defer f.Close()
	const sheet = "整改通知单"
	idx, err := f.NewSheet(sheet)
	if err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建工作表失败", err)
	}
	f.SetActiveSheet(idx)
	logXlsxErr(f.DeleteSheet("Sheet1"), "DeleteSheet")

	s.applyBaseStyle(f, sheet)

	count := len(dangers)
	totalRows := count + 5
	if totalRows < 4 {
		totalRows = 4
	}
	s.applyOuterBorder(f, sheet, "A1", fmt.Sprintf("H%d", totalRows))

	logXlsxErr(f.MergeCell(sheet, "A1", "H1"), "MergeCell")
	logXlsxErr(f.MergeCell(sheet, "A2", "H2"), "MergeCell")
	logXlsxErr(f.SetRowHeight(sheet, 1, 40), "SetRowHeight")
	logXlsxErr(f.SetRowHeight(sheet, 2, 26), "SetRowHeight")
	logXlsxErr(f.SetRowHeight(sheet, 3, 36), "SetRowHeight")

	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Family: "宋体", Size: 20, Bold: true},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
	})
	logXlsxErr(f.SetCellStyle(sheet, "A1", "A1", titleStyle), "SetCellStyle")
	subtitleStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Family: "宋体", Size: 11},
		Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center", WrapText: true},
	})
	logXlsxErr(f.SetCellStyle(sheet, "A2", "A2", subtitleStyle), "SetCellStyle")

	logXlsxErr(f.SetCellValue(sheet, "A1", "现场隐患描述及整改反馈表（整改期限15个工作日）"), "SetCellValue")
	logXlsxErr(f.SetCellValue(sheet, "A2", fmt.Sprintf(
		"企业名称：%s        联系人：%s        联系方式：%s        检查时间：%s",
		checkCompanyName, supervisor, supervisorTel, checkTime,
	)), "SetCellValue")
	headers := []string{"序号", "不符合项描述", "整改前图片", "整改依据", "整改建议", "隐患类别", "整改后图片", "整改情况说明"}
	for i, h := range headers {
		col, _ := excelize.CoordinatesToCellName(i+1, 3)
		logXlsxErr(f.SetCellValue(sheet, col, h), "SetCellValue")
	}

	for i, d := range dangers {
		row := i + 4
		logXlsxErr(f.SetRowHeight(sheet, row, 70), "SetRowHeight")
		logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("A%d", row), i+1), "SetCellValue")
		logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("B%d", row), strings.TrimSpace(d.Area+":"+d.Description)), "SetCellValue")
		logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("D%d", row), derefString(d.Legalises)), "SetCellValue")
		logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("E%d", row), derefString(d.Proposal)), "SetCellValue")
		logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("F%d", row), d.HiddenType+"/"+d.HiddenLevel), "SetCellValue")
		logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("H%d", row), derefString(d.ReviewDescription)), "SetCellValue")

		s.embedFirstAnnex(f, sheet, fmt.Sprintf("C%d", row), d.HiddenExamineAnnex)
		s.embedFirstAnnex(f, sheet, fmt.Sprintf("G%d", row), d.HiddenRectifyAnnex)
	}

	auditRow := count + 4
	if auditRow < 5 {
		auditRow = 5
	}
	logXlsxErr(f.MergeCell(sheet, fmt.Sprintf("A%d", auditRow), fmt.Sprintf("H%d", auditRow)), "MergeCell")
	logXlsxErr(f.MergeCell(sheet, fmt.Sprintf("A%d", auditRow+1), fmt.Sprintf("H%d", auditRow+1)), "MergeCell")
	logXlsxErr(f.SetRowHeight(sheet, auditRow, 180), "SetRowHeight")
	auditStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Family: "宋体", Size: 11},
		Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center", WrapText: true},
	})
	logXlsxErr(f.SetCellStyle(sheet, fmt.Sprintf("A%d", auditRow), fmt.Sprintf("A%d", auditRow+1), auditStyle), "SetCellStyle")

	planCompanyTitle := s.fetchCompanyTitle(ctx, ec.ExpertInspection)
	narrative := s.buildAuditNarrative(&ec, dangers, planCompanyTitle, checkCompanyName)
	logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("A%d", auditRow), narrative), "SetCellValue")
	logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("A%d", auditRow+1), "备注：隐患分七类（危化品、电气、机械、安全设施、消防、职业健康、其他）。"), "SetCellValue")

	relPath, err := s.saveAs(f, "整改通知单-"+checkCompanyName)
	if err != nil {
		return nil, err
	}
	if err := s.db.WithContext(ctx).Model(&ec).
		Update("rectification_sheet_url", relPath).Error; err != nil {
		return nil, wrapWrite(err)
	}
	if err := s.db.WithContext(ctx).First(&ec, expertCompanyID).Error; err != nil {
		return nil, err
	}
	return &ec, nil
}

// RectificationSheetFile — 下载场景：返回绝对路径 + 文件名。
func (s *RectificationSheetService) RectificationSheetFile(ctx context.Context, expertCompanyID int64) (abs string, name string, ec *model.ExpertCompany, err error) {
	var row model.ExpertCompany
	if err := s.db.WithContext(ctx).
		Preload("ExpertInspection").
		Preload("ExpertInspectionUsers").
		First(&row, expertCompanyID).Error; err != nil {
		return "", "", nil, mapNotFound(err)
	}
	if row.RectificationSheetURL == nil || *row.RectificationSheetURL == "" {
		return "", "", &row, errs.New(errs.CodeUnprocessable, "整改通知单尚未生成")
	}
	abs = filepath.Join(s.storageRoot, *row.RectificationSheetURL)
	if _, err := os.Stat(abs); err != nil {
		return "", "", &row, errs.New(errs.CodeNotFound, "整改通知单文件不存在")
	}
	return abs, filepath.Base(abs), &row, nil
}

// ── 隐患分类统计（按 company） ────────────────────────────────────────

type HazardStatisticsParams struct {
	CompanyID int64
	StartTime string // YYYY-MM-DD HH:mm:ss 或 YYYY-MM-DD
	EndTime   string
	Type      string // "" 或 "all" → 全部 7 类
}

// GenerateHazardStatistics — 一次性生成 .xlsx 并返回绝对路径供 handler 下载后删除。
//
// 跟 Laravel `exportStatistics` 一样：按公司聚合「该公司发起的所有 expert_inspections
// 下的全部隐患」，按 hidden_type 分桶。
func (s *RectificationSheetService) GenerateHazardStatistics(ctx context.Context, p HazardStatisticsParams) (abs string, err error) {
	if p.CompanyID == 0 {
		return "", errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	var co model.Company
	_ = s.db.WithContext(ctx).First(&co, p.CompanyID).Error
	companyTitle := co.Title

	var inspectionIDs []int64
	if err := s.db.WithContext(ctx).Model(&model.ExpertInspection{}).
		Where("company_id = ?", p.CompanyID).
		Pluck("id", &inspectionIDs).Error; err != nil {
		return "", err
	}

	types := hiddenTypes
	if t := strings.TrimSpace(p.Type); t != "" && t != "all" {
		types = []string{t}
	}

	buckets := make(map[string][]model.HiddenDangerInfo, len(types))
	bucketCount := 0
	totalCount := 0
	for _, t := range types {
		var rows []model.HiddenDangerInfo
		if len(inspectionIDs) > 0 {
			q := s.db.WithContext(ctx).
				Where("expert_inspection_id IN ?", inspectionIDs).
				Where("hidden_type = ?", t)
			if p.StartTime != "" {
				q = q.Where("created_at >= ?", p.StartTime)
			}
			if p.EndTime != "" {
				q = q.Where("created_at <= ?", p.EndTime)
			}
			if err := q.Order("id ASC").Find(&rows).Error; err != nil {
				return "", err
			}
			if err := s.attachDangerAnnexes(ctx, rows); err != nil {
				return "", err
			}
		}
		buckets[t] = rows
		if len(rows) > 0 {
			bucketCount++
			totalCount += len(rows)
		}
	}

	f := excelize.NewFile()
	defer f.Close()
	const sheet = "隐患排查分类统计表"
	idx, err := f.NewSheet(sheet)
	if err != nil {
		return "", errs.Wrap(errs.CodeInternal, "创建工作表失败", err)
	}
	f.SetActiveSheet(idx)
	logXlsxErr(f.DeleteSheet("Sheet1"), "DeleteSheet")

	s.applyBaseStyle(f, sheet)

	totalRows := totalCount + bucketCount + 3
	if totalRows >= 3 {
		s.applyOuterBorder(f, sheet, "A1", fmt.Sprintf("H%d", totalRows))
	}

	logXlsxErr(f.MergeCell(sheet, "A1", "H1"), "MergeCell")
	logXlsxErr(f.MergeCell(sheet, "A2", "H2"), "MergeCell")
	logXlsxErr(f.SetRowHeight(sheet, 1, 40), "SetRowHeight")
	logXlsxErr(f.SetRowHeight(sheet, 2, 26), "SetRowHeight")
	logXlsxErr(f.SetRowHeight(sheet, 3, 36), "SetRowHeight")

	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Family: "宋体", Size: 20, Bold: true},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
	})
	logXlsxErr(f.SetCellStyle(sheet, "A1", "A1", titleStyle), "SetCellStyle")
	subtitleStyle, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center", WrapText: true},
	})
	logXlsxErr(f.SetCellStyle(sheet, "A2", "A2", subtitleStyle), "SetCellStyle")

	logXlsxErr(f.SetCellValue(sheet, "A1", "隐患排查分类统计表 · "+companyTitle), "SetCellValue")
	subtitle := fmt.Sprintf("统计时间：%s ~ %s", emptyDash(p.StartTime), emptyDash(p.EndTime))
	if t := strings.TrimSpace(p.Type); t != "" && t != "all" {
		subtitle += "　|　类别：" + t
	}
	logXlsxErr(f.SetCellValue(sheet, "A2", subtitle), "SetCellValue")

	headers := []string{"序号", "不符合项描述", "整改前图片", "整改依据", "整改建议", "隐患类别", "整改后图片", "整改情况说明"}
	for i, h := range headers {
		col, _ := excelize.CoordinatesToCellName(i+1, 3)
		logXlsxErr(f.SetCellValue(sheet, col, h), "SetCellValue")
	}

	row := 4
	for _, t := range types {
		bucket := buckets[t]
		if len(bucket) == 0 {
			continue
		}
		logXlsxErr(f.MergeCell(sheet, fmt.Sprintf("A%d", row), fmt.Sprintf("H%d", row)), "MergeCell")
		logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("A%d", row), t+"类隐患"), "SetCellValue")
		logXlsxErr(f.SetRowHeight(sheet, row, 36), "SetRowHeight")
		row++

		for i, d := range bucket {
			logXlsxErr(f.SetRowHeight(sheet, row, 70), "SetRowHeight")
			logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("A%d", row), i+1), "SetCellValue")
			logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("B%d", row), strings.TrimSpace(d.Area+":"+d.Description)), "SetCellValue")
			logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("D%d", row), derefString(d.Legalises)), "SetCellValue")
			logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("E%d", row), derefString(d.Proposal)), "SetCellValue")
			logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("F%d", row), d.HiddenType+"/"+d.HiddenLevel), "SetCellValue")
			logXlsxErr(f.SetCellValue(sheet, fmt.Sprintf("H%d", row), derefString(d.ReviewDescription)), "SetCellValue")
			s.embedFirstAnnex(f, sheet, fmt.Sprintf("C%d", row), d.HiddenExamineAnnex)
			s.embedFirstAnnex(f, sheet, fmt.Sprintf("G%d", row), d.HiddenRectifyAnnex)
			row++
		}
	}

	relPath, err := s.saveAs(f, "隐患分类统计-"+companyTitle)
	if err != nil {
		return "", err
	}
	return filepath.Join(s.storageRoot, relPath), nil
}

// ── helpers ─────────────────────────────────────────────────────────

func (s *RectificationSheetService) applyBaseStyle(f *excelize.File, sheet string) {
	// 列宽贴 PhpSpreadsheet 数值（PhpSpreadsheet 列宽单位 ≈ 字符宽，excelize 一致）。
	logXlsxErr(f.SetColWidth(sheet, "A", "A", 9.08), "SetColWidth")
	logXlsxErr(f.SetColWidth(sheet, "B", "B", 32.08), "SetColWidth")
	logXlsxErr(f.SetColWidth(sheet, "C", "C", 24.08), "SetColWidth")
	logXlsxErr(f.SetColWidth(sheet, "D", "D", 28.08), "SetColWidth")
	logXlsxErr(f.SetColWidth(sheet, "E", "E", 28.08), "SetColWidth")
	logXlsxErr(f.SetColWidth(sheet, "F", "F", 20.08), "SetColWidth")
	logXlsxErr(f.SetColWidth(sheet, "G", "G", 24.08), "SetColWidth")
	logXlsxErr(f.SetColWidth(sheet, "H", "H", 24.08), "SetColWidth")
	logXlsxErr(f.SetPageLayout(sheet, &excelize.PageLayoutOptions{
		Orientation: stringPtr("portrait"),
		Size:        intPtr(9), // A4
	}), "SetPageLayout")
}

func (s *RectificationSheetService) applyOuterBorder(f *excelize.File, sheet, from, to string) {
	border, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
		Font:      &excelize.Font{Family: "宋体", Size: 11},
	})
	logXlsxErr(f.SetCellStyle(sheet, from, to, border), "SetCellStyle")
}

// embedFirstAnnex — 取附件列表第一项嵌入到目标单元格。
//
// 行为对齐 PhpWord MemoryDrawing(146, 82) — 固定像素尺寸（约 1.5×0.85 英寸，跟整改单
// 列宽 24+ 行高 70 匹配）。excelize 用 ScaleX/ScaleY 缩放原图，所以先 image.DecodeConfig
// 量出原始宽高，再算一个 (146/w, 82/h) 的缩放因子。
//
// OSS 未配置 / 网络失败 / 解码失败 → 静默跳过，不影响其他单元格 + 文本仍写入。
func (s *RectificationSheetService) embedFirstAnnex(f *excelize.File, sheet, cell string, annexes []model.HiddenUploadAnnex) {
	if len(annexes) == 0 {
		return
	}
	raw := strings.TrimSpace(annexes[0].AnnexURL)
	if raw == "" {
		return
	}
	buf, err := s.fetchImage(raw)
	if err != nil || len(buf) == 0 {
		return
	}
	cfg, _, err := image.DecodeConfig(bytes.NewReader(buf))
	if err != nil {
		return
	}
	const targetW, targetH = 146.0, 82.0
	scaleX := targetW / float64(cfg.Width)
	scaleY := targetH / float64(cfg.Height)
	ext := guessImageExt(raw)
	logXlsxErr(f.AddPictureFromBytes(sheet, cell, &excelize.Picture{
		Extension: ext,
		File:      buf,
		Format: &excelize.GraphicOptions{
			OffsetX:         4,
			OffsetY:         4,
			ScaleX:          scaleX,
			ScaleY:          scaleY,
			LockAspectRatio: false,
			Positioning:     "oneCell",
		},
	}), "AddPictureFromBytes")
}

// fetchImage — 真实 HTTP 抓图。OSS 已配置且 URL 是 OSS 路径走 SignURL，
// 否则直接 GET（兼容已经签好的 URL 或公网图片）。
func (s *RectificationSheetService) fetchImage(raw string) ([]byte, error) {
	url := raw
	if s.oss != nil {
		signed, err := s.oss.SignURL(raw, 600*time.Second, nil)
		if err == nil && signed != "" {
			url = signed
		}
	}
	httpClient := &http.Client{Timeout: 10 * time.Second}
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

// guessImageExt — 从 URL path 段（去掉 query string）猜扩展。OSS 签名 URL 带 ?... 后缀，
// 直接 strings.Contains(".jpg") 会误命中 query string。
func guessImageExt(raw string) string {
	pathOnly := raw
	if i := strings.IndexAny(raw, "?#"); i >= 0 {
		pathOnly = raw[:i]
	}
	pathOnly = strings.ToLower(pathOnly)
	switch {
	case strings.HasSuffix(pathOnly, ".png"):
		return ".png"
	case strings.HasSuffix(pathOnly, ".jpeg"):
		return ".jpeg"
	case strings.HasSuffix(pathOnly, ".gif"):
		return ".gif"
	default:
		return ".jpg"
	}
}

func (s *RectificationSheetService) attachDangerAnnexes(ctx context.Context, rows []model.HiddenDangerInfo) error {
	if len(rows) == 0 {
		return nil
	}
	ids := make([]int64, 0, len(rows))
	idx := make(map[int64]int, len(rows))
	for i, r := range rows {
		ids = append(ids, r.ID)
		idx[r.ID] = i
	}
	var annexes []model.HiddenUploadAnnex
	if err := s.db.WithContext(ctx).
		Where("hidden_danger_info_id IN ?", ids).
		Order("id ASC").
		Find(&annexes).Error; err != nil {
		return err
	}
	for _, a := range annexes {
		i, ok := idx[a.HiddenDangerInfoID]
		if !ok {
			continue
		}
		switch a.IsReview {
		case 1:
			rows[i].HiddenExamineAnnex = append(rows[i].HiddenExamineAnnex, a)
		case 2:
			rows[i].HiddenReviewAnnex = append(rows[i].HiddenReviewAnnex, a)
		case 3:
			rows[i].HiddenRectifyAnnex = append(rows[i].HiddenRectifyAnnex, a)
		}
	}
	return nil
}

func (s *RectificationSheetService) fetchCompanyTitle(ctx context.Context, ins *model.ExpertInspection) string {
	if ins == nil {
		return ""
	}
	var co model.Company
	if err := s.db.WithContext(ctx).Select("title").First(&co, ins.CompanyID).Error; err != nil {
		return ""
	}
	return co.Title
}

func (s *RectificationSheetService) buildAuditNarrative(ec *model.ExpertCompany, dangers []model.HiddenDangerInfo, planCompanyTitle, checkCompanyName string) string {
	var b strings.Builder
	b.WriteString(ec.CreatedAt.Format("2006-01-02 15:04:05"))
	b.WriteString("， ")
	b.WriteString(planCompanyTitle)
	b.WriteString("发出检查通知")

	expertNames := make([]string, 0, len(ec.ExpertInspectionUsers))
	for _, u := range ec.ExpertInspectionUsers {
		var name string
		if u.ExpertUser != nil && u.ExpertUser.Nickname != "" {
			name = u.ExpertUser.Nickname
		} else {
			name = fmt.Sprintf("专家#%d", u.ExpertUserID)
		}
		expertNames = append(expertNames, name)
		if u.ConfirmTime != nil {
			b.WriteString("\n")
			b.WriteString(time.Time(*u.ConfirmTime).Format("2006-01-02 15:04:05"))
			b.WriteString("， ")
			b.WriteString(name)
			b.WriteString("专家确认接收")
		} else {
			b.WriteString("\n")
			b.WriteString(name)
			b.WriteString("专家未接收")
		}
	}
	expertCSV := strings.Join(expertNames, "、")

	if ec.ConfirmTime != nil {
		b.WriteString("\n")
		b.WriteString(time.Time(*ec.ConfirmTime).Format("2006-01-02 15:04:05"))
		b.WriteString("， ")
		b.WriteString(checkCompanyName)
		b.WriteString("确认接收")
	} else {
		b.WriteString("\n")
		b.WriteString(checkCompanyName)
		b.WriteString("未接收")
	}

	if len(dangers) > 0 {
		first := dangers[0]
		for _, d := range dangers {
			if d.CreatedAt.Before(first.CreatedAt) {
				first = d
			}
		}
		b.WriteString("\n")
		b.WriteString(first.CreatedAt.Format("2006-01-02 15:04:05"))
		b.WriteString("， ")
		b.WriteString(expertCSV)
		b.WriteString("专家执行检查任务，下发隐患整改单")

		total := len(dangers)
		rectified := 0
		var latestCorrect *model.Date
		for _, d := range dangers {
			if d.CorrectDate != nil {
				rectified++
				if latestCorrect == nil || time.Time(*d.CorrectDate).After(time.Time(*latestCorrect)) {
					latestCorrect = d.CorrectDate
				}
			}
		}
		if rectified == total && total > 0 && latestCorrect != nil {
			b.WriteString("\n")
			b.WriteString(time.Time(*latestCorrect).Format("2006-01-02"))
			b.WriteString("， ")
			b.WriteString(checkCompanyName)
			b.WriteString("完成整改，反馈整改情况")
		}

		var reviewed []model.HiddenDangerInfo
		for _, d := range dangers {
			if d.ReviewTime != nil {
				reviewed = append(reviewed, d)
			}
		}
		if len(reviewed) > 0 {
			sort.Slice(reviewed, func(i, j int) bool { return reviewed[i].ReviewTime.Before(*reviewed[j].ReviewTime) })
			b.WriteString("\n")
			b.WriteString(reviewed[0].ReviewTime.Format("2006-01-02 15:04:05"))
			b.WriteString("， ")
			b.WriteString(expertCSV)
			b.WriteString("专家进行整改确认")
			if len(reviewed) == total {
				last := reviewed[len(reviewed)-1]
				b.WriteString("\n")
				b.WriteString(last.ReviewTime.Format("2006-01-02 15:04:05"))
				b.WriteString("， ")
				b.WriteString(expertCSV)
				b.WriteString("专家完成整改确认")
			}
		}

		archived := 0
		for _, d := range dangers {
			if d.ArchiveUserID != nil {
				archived++
			}
		}
		if archived == total && total > 0 {
			latest := dangers[0]
			for _, d := range dangers {
				if d.UpdatedAt.After(latest.UpdatedAt) {
					latest = d
				}
			}
			b.WriteString("\n")
			b.WriteString(latest.UpdatedAt.Format("2006-01-02 15:04:05"))
			b.WriteString("， ")
			b.WriteString(planCompanyTitle)
			b.WriteString("确认归档")
		}
	}
	return b.String()
}

func (s *RectificationSheetService) saveAs(f *excelize.File, baseName string) (string, error) {
	now := time.Now()
	relDir := filepath.Join("rectification-sheets", now.Format("2006"), now.Format("01"))
	absDir := filepath.Join(s.storageRoot, relDir)
	if err := os.MkdirAll(absDir, 0o755); err != nil {
		return "", errs.Wrap(errs.CodeInternal, "创建目录失败", err)
	}
	suffix, err := randomSuffix(3)
	if err != nil {
		return "", err
	}
	filename := fmt.Sprintf("%s-%s-%s.xlsx", baseName, now.Format("20060102"), suffix)
	relPath := filepath.Join(relDir, filename)
	absPath := filepath.Join(s.storageRoot, relPath)

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return "", errs.Wrap(errs.CodeInternal, "序列化 xlsx 失败", err)
	}
	if err := os.WriteFile(absPath, buf.Bytes(), 0o644); err != nil {
		return "", errs.Wrap(errs.CodeInternal, "写文件失败", err)
	}
	return relPath, nil
}

func randomSuffix(n int) (string, error) {
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

func derefString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func emptyDash(s string) string {
	if strings.TrimSpace(s) == "" {
		return "—"
	}
	return s
}

func stringPtr(s string) *string { return &s }
func intPtr(v int) *int          { return &v }

// logXlsxErr surfaces an excelize call's error in logs without aborting the
// workbook build. Generators here intentionally tolerate cosmetic failures
// (cell styles, row heights, merges) so a partial sheet still reaches the
// user — but silently swallowing every error hid genuine SetCellValue
// failures that lost data. Genuine fatal points (NewSheet, Write) still
// propagate via err returns.
func logXlsxErr(err error, where string) {
	if err != nil {
		log.Printf("rectification_sheet: %s: %v", where, err)
	}
}
