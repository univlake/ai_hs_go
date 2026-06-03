package excel

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"

	"ai_hs_go/internal/model"
)

var specialEquipmentHeaders = []string{
	"设备名称", "设备类型", "使用位置",
	"规格型号1", "规格型号2", "规格型号3", "规格型号4",
	"使用证编号", "投用时间", "有效期", "检验单位", "维护情况", "备注",
}

// One row of realistic-looking sample data so users can see the expected
// shape (date format, units, Chinese values) before filling in their own.
// The trailing comment cell flags the row as something to delete before
// import — the importer doesn't auto-skip it because we still want the
// row to appear if the user keeps it.
//
// 设备类型 must match a row in special_types.title — "压力容器(含气瓶)"
// is one of the seeded categories whose 规格型号1..3 map to
// 设计压力 / 容积 / 盛放介质 in the UI's SPEC_CONFIGS.
var specialEquipmentExample = []any{
	"储气罐", "压力容器(含气瓶)", "1号车间",
	"1.6 (设计压力 MPa)", "5 (容积 m³)", "压缩空气 (盛放介质)", "",
	"PV-2024-001", "2024-03-15", "2026-03-14",
	"省特种设备检验研究院", "每季度保养",
	"示例数据 — 导入前请删除本行",
}

// SpecialEquipmentTemplate writes an Excel template with a header row and
// one styled example data row.
func SpecialEquipmentTemplate(w io.Writer) error {
	f := excelize.NewFile()
	defer f.Close()
	sheet := "特种设备"
	f.SetSheetName("Sheet1", sheet)
	if err := writeHeaders(f, sheet, specialEquipmentHeaders); err != nil {
		return err
	}
	if err := writeExampleRow(f, sheet, specialEquipmentExample); err != nil {
		return err
	}
	_, err := f.WriteTo(w)
	return err
}

// SpecialEquipmentExport writes existing rows. typeTitleByID resolves
// special_type_id → title so the 设备类型 column stays readable; pass nil
// to leave it blank.
func SpecialEquipmentExport(w io.Writer, rows []model.SpecialEquipment, typeTitleByID map[int64]string) error {
	f := excelize.NewFile()
	defer f.Close()
	sheet := "特种设备"
	f.SetSheetName("Sheet1", sheet)
	if err := writeHeaders(f, sheet, specialEquipmentHeaders); err != nil {
		return err
	}
	for i, r := range rows {
		row := i + 2
		typeTitle := ""
		if r.SpecialTypeID != nil && typeTitleByID != nil {
			typeTitle = typeTitleByID[*r.SpecialTypeID]
		}
		values := []any{
			deref(r.EquipmentName), typeTitle, deref(r.UtilizedLocation),
			deref(r.NormModel1), deref(r.NormModel2), deref(r.NormModel3), deref(r.NormModel4),
			deref(r.SerialNumber),
			derefDate(r.UseTime), derefDate(r.ValidPeriod),
			deref(r.InspectUnit), deref(r.MaintainInfo), deref(r.Comment),
		}
		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			_ = f.SetCellValue(sheet, cell, v)
		}
	}
	_, err := f.WriteTo(w)
	return err
}

// SpecialEquipmentImport parses uploaded Excel file (skipping header row) into
// model rows. Caller is responsible for assigning company_id / created_by.
// typesByTitle maps the 设备类型 cell to special_type_id; pass nil to skip
// the lookup (rows will have a nil SpecialTypeID).
func SpecialEquipmentImport(r io.Reader, typesByTitle map[string]int64) ([]model.SpecialEquipment, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	f, err := excelize.OpenReader(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheet := f.GetSheetName(0)
	rowsIter, err := f.Rows(sheet)
	if err != nil {
		return nil, err
	}
	defer rowsIter.Close()

	var rows []model.SpecialEquipment
	rowIdx := 0
	for rowsIter.Next() {
		rowIdx++
		cells, err := rowsIter.Columns()
		if err != nil {
			return nil, err
		}
		if rowIdx == 1 {
			continue
		}
		if isEmptyRow(cells) {
			continue
		}
		row := model.SpecialEquipment{
			EquipmentName:    optStr(cells, 0),
			SpecialTypeID:    lookupTypeID(cells, 1, typesByTitle),
			UtilizedLocation: optStr(cells, 2),
			NormModel1:       optStr(cells, 3),
			NormModel2:       optStr(cells, 4),
			NormModel3:       optStr(cells, 5),
			NormModel4:       optStr(cells, 6),
			SerialNumber:     optStr(cells, 7),
			UseTime:          optDate(cells, 8),
			ValidPeriod:      optDate(cells, 9),
			InspectUnit:      optStr(cells, 10),
			MaintainInfo:     optStr(cells, 11),
			Comment:          optStr(cells, 12),
		}
		rows = append(rows, row)
	}
	return rows, nil
}

// lookupTypeID resolves the type-title cell to a special_type_id. Returns
// nil when the cell is empty, the lookup map is missing, or the title isn't
// in the dictionary — the row still imports, just without a type.
func lookupTypeID(cells []string, i int, typesByTitle map[string]int64) *int64 {
	if i >= len(cells) {
		return nil
	}
	title := strings.TrimSpace(cells[i])
	if title == "" || typesByTitle == nil {
		return nil
	}
	if id, ok := typesByTitle[title]; ok {
		return &id
	}
	return nil
}

// writeHeaders writes a styled header row.
func writeHeaders(f *excelize.File, sheet string, headers []string) error {
	headerStyle, err := f.NewStyle(&excelize.Style{
		Fill:      excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{"#3B82F6"}},
		Font:      &excelize.Font{Color: "#FFFFFF", Bold: true},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		return err
	}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		_ = f.SetCellValue(sheet, cell, h)
		_ = f.SetCellStyle(sheet, cell, cell, headerStyle)
	}
	return f.SetColWidth(sheet, "A", string('A'+rune(len(headers)-1)), 18)
}

// writeStaffTemplate is the shared shape for the 人员管理 子表 templates:
// header row + one styled example row. Sheet name shows up as the tab
// label in Excel, so it doubles as a hint about the template's purpose.
func writeStaffTemplate(w io.Writer, sheet string, headers []string, example []any) error {
	f := excelize.NewFile()
	defer f.Close()
	f.SetSheetName("Sheet1", sheet)
	if err := writeHeaders(f, sheet, headers); err != nil {
		return err
	}
	if err := writeExampleRow(f, sheet, example); err != nil {
		return err
	}
	_, err := f.WriteTo(w)
	return err
}

// writeExampleRow renders one example row at row 2 in italic grey text on a
// soft background, so users immediately see it's filler and can replace or
// delete it before re-uploading.
func writeExampleRow(f *excelize.File, sheet string, values []any) error {
	exampleStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Color: "#94A3B8", Italic: true},
		Fill: excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{"#F8FAFC"}},
	})
	if err != nil {
		return err
	}
	for i, v := range values {
		cell, _ := excelize.CoordinatesToCellName(i+1, 2)
		_ = f.SetCellValue(sheet, cell, v)
		_ = f.SetCellStyle(sheet, cell, cell, exampleStyle)
	}
	return nil
}

func deref(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func derefDate(p *model.Date) string {
	if p == nil {
		return ""
	}
	return time.Time(*p).Format("2006-01-02")
}

func optStr(cells []string, i int) *string {
	if i >= len(cells) {
		return nil
	}
	v := cells[i]
	if v == "" {
		return nil
	}
	return &v
}

func optDate(cells []string, i int) *model.Date {
	if i >= len(cells) {
		return nil
	}
	v := cells[i]
	if v == "" {
		return nil
	}
	formats := []string{"2006-01-02", "2006/1/2", "2006/01/02", time.RFC3339}
	for _, f := range formats {
		if t, err := time.Parse(f, v); err == nil {
			d := model.Date(t)
			return &d
		}
	}
	// excelize may return numeric serials (e.g. "44197"); leave nil for now.
	return nil
}

func isEmptyRow(cells []string) bool {
	for _, c := range cells {
		if c != "" {
			return false
		}
	}
	return true
}

// EnsureSheetSize is a convenience helper to validate basic invariants when
// large imports come in (currently unused but kept for future bulk paths).
func EnsureSheetSize(rows int, max int) error {
	if rows > max {
		return fmt.Errorf("rows %d exceeds max %d", rows, max)
	}
	return nil
}
