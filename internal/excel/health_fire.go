package excel

import (
	"io"
	"strconv"

	"github.com/xuri/excelize/v2"

	"ai_hs_go/internal/model"
)

// 职业危害因素 / 健康档案的导入模板 / 导出 / 解析。复用 special_equipment.go
// 里的 writeHeaders / writeExampleRow / openStaffSheet / optStr / optDate / deref。

// ── 职业危害因素 ────────────────────────────────────────────────

var occupationalHazardHeaders = []string{
	"作业场所", "职业危害因素", "接触人数", "申报状态", "申报日期",
}

var occupationalHazardExample = []any{
	"配料车间", "粉尘类", 10, "已申报", "2024-01-15",
}

func OccupationalHazardTemplate(w io.Writer) error {
	f := excelize.NewFile()
	defer f.Close()
	sheet := "职业危害因素"
	f.SetSheetName("Sheet1", sheet)
	if err := writeHeaders(f, sheet, occupationalHazardHeaders); err != nil {
		return err
	}
	if err := writeExampleRow(f, sheet, occupationalHazardExample); err != nil {
		return err
	}
	_, err := f.WriteTo(w)
	return err
}

func OccupationalHazardExport(w io.Writer, rows []model.OccupationalHazard) error {
	f := excelize.NewFile()
	defer f.Close()
	sheet := "职业危害因素"
	f.SetSheetName("Sheet1", sheet)
	if err := writeHeaders(f, sheet, occupationalHazardHeaders); err != nil {
		return err
	}
	for i, r := range rows {
		row := i + 2
		number := ""
		if r.Number != nil {
			number = strconv.FormatInt(int64(*r.Number), 10)
		}
		values := []any{
			deref(r.Name), deref(r.Type), number,
			deref(r.IsDeclare), derefDate(r.DeclareDate),
		}
		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			_ = f.SetCellValue(sheet, cell, v)
		}
	}
	_, err := f.WriteTo(w)
	return err
}

func OccupationalHazardImport(r io.Reader) ([]model.OccupationalHazard, error) {
	rowsIter, closeFn, err := openStaffSheet(r)
	if err != nil {
		return nil, err
	}
	defer closeFn()

	var rows []model.OccupationalHazard
	idx := 0
	for rowsIter.Next() {
		idx++
		cells, err := rowsIter.Columns()
		if err != nil {
			return nil, err
		}
		if idx == 1 || isEmptyRow(cells) {
			continue
		}
		var num *int32
		if v := optStr(cells, 2); v != nil {
			if n, err := strconv.ParseInt(*v, 10, 32); err == nil {
				n32 := int32(n)
				num = &n32
			}
		}
		rows = append(rows, model.OccupationalHazard{
			Name:        optStr(cells, 0),
			Type:        optStr(cells, 1),
			Number:      num,
			IsDeclare:   optStr(cells, 3),
			DeclareDate: optDate(cells, 4),
		})
	}
	return rows, nil
}

// ── 健康档案 ───────────────────────────────────────────────────

var healthFileHeaders = []string{
	"姓名", "工号", "部门", "岗位", "体检日期", "体检结论", "复查情况", "处理意见", "处理结果",
}

var healthFileExample = []any{
	"张三", "EMP001", "生产部", "操作员", "2024-03-01", "合格", "", "", "",
}

func HealthFileTemplate(w io.Writer) error {
	f := excelize.NewFile()
	defer f.Close()
	sheet := "健康档案"
	f.SetSheetName("Sheet1", sheet)
	if err := writeHeaders(f, sheet, healthFileHeaders); err != nil {
		return err
	}
	if err := writeExampleRow(f, sheet, healthFileExample); err != nil {
		return err
	}
	_, err := f.WriteTo(w)
	return err
}

func HealthFileExport(w io.Writer, rows []model.HealthFile) error {
	f := excelize.NewFile()
	defer f.Close()
	sheet := "健康档案"
	f.SetSheetName("Sheet1", sheet)
	if err := writeHeaders(f, sheet, healthFileHeaders); err != nil {
		return err
	}
	for i, r := range rows {
		row := i + 2
		values := []any{
			deref(r.Name), deref(r.Number), deref(r.Department), deref(r.Post),
			derefDate(r.CheckupDate), deref(r.CheckupExpiry),
			deref(r.ReviewSituation), deref(r.HandlingOpinions), deref(r.ProcessResult),
		}
		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			_ = f.SetCellValue(sheet, cell, v)
		}
	}
	_, err := f.WriteTo(w)
	return err
}

func HealthFileImport(r io.Reader) ([]model.HealthFile, error) {
	rowsIter, closeFn, err := openStaffSheet(r)
	if err != nil {
		return nil, err
	}
	defer closeFn()

	var rows []model.HealthFile
	idx := 0
	for rowsIter.Next() {
		idx++
		cells, err := rowsIter.Columns()
		if err != nil {
			return nil, err
		}
		if idx == 1 || isEmptyRow(cells) {
			continue
		}
		rows = append(rows, model.HealthFile{
			Name:             optStr(cells, 0),
			Number:           optStr(cells, 1),
			Department:       optStr(cells, 2),
			Post:             optStr(cells, 3),
			CheckupDate:      optDate(cells, 4),
			CheckupExpiry:    optStr(cells, 5),
			ReviewSituation:  optStr(cells, 6),
			HandlingOpinions: optStr(cells, 7),
			ProcessResult:    optStr(cells, 8),
		})
	}
	return rows, nil
}
