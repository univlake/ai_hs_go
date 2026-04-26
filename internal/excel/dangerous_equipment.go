package excel

import (
	"bytes"
	"io"
	"time"

	"github.com/xuri/excelize/v2"

	"ai_hs_go/internal/model"
)

var dangerousEquipmentHeaders = []string{
	"设备名称", "使用位置", "规格型号", "使用证编号", "投用时间",
	"自动化程度", "维护情况", "备注",
}

// Example data row — `自动化程度` must match one of the AUTOMATION_OPTIONS
// the form's <select> renders, otherwise the field shows as blank when the
// user opens the imported record for edit.
var dangerousEquipmentExample = []any{
	"管式反应釜", "2号车间-A区",
	"5m³ / 设计压力 1.6MPa", "DQ-2024-001",
	"2024-05-20", "全自动设备", "运行良好，月度巡检",
	"示例数据 — 导入前请删除本行",
}

func DangerousEquipmentTemplate(w io.Writer) error {
	f := excelize.NewFile()
	defer f.Close()
	sheet := "重大危险设备"
	f.SetSheetName("Sheet1", sheet)
	if err := writeHeaders(f, sheet, dangerousEquipmentHeaders); err != nil {
		return err
	}
	if err := writeExampleRow(f, sheet, dangerousEquipmentExample); err != nil {
		return err
	}
	_, err := f.WriteTo(w)
	return err
}

func DangerousEquipmentExport(w io.Writer, rows []model.DangerousEquipment) error {
	f := excelize.NewFile()
	defer f.Close()
	sheet := "重大危险设备"
	f.SetSheetName("Sheet1", sheet)
	if err := writeHeaders(f, sheet, dangerousEquipmentHeaders); err != nil {
		return err
	}
	for i, r := range rows {
		row := i + 2
		values := []any{
			deref(r.EquipmentName),
			deref(r.UtilizedLocation),
			deref(r.NormModel),
			deref(r.SerialNumber),
			derefDate(r.UseTime),
			deref(r.AutomationType),
			deref(r.MaintainInfo),
			deref(r.Comment),
		}
		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			_ = f.SetCellValue(sheet, cell, v)
		}
	}
	_, err := f.WriteTo(w)
	return err
}

func DangerousEquipmentImport(r io.Reader) ([]model.DangerousEquipment, error) {
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

	var rows []model.DangerousEquipment
	rowIdx := 0
	for rowsIter.Next() {
		rowIdx++
		cells, err := rowsIter.Columns()
		if err != nil {
			return nil, err
		}
		if rowIdx == 1 || isEmptyRow(cells) {
			continue
		}
		row := model.DangerousEquipment{
			EquipmentName:    optStr(cells, 0),
			UtilizedLocation: optStr(cells, 1),
			NormModel:        optStr(cells, 2),
			SerialNumber:     optStr(cells, 3),
			UseTime:          optDate(cells, 4),
			AutomationType:   optStr(cells, 5),
			MaintainInfo:     optStr(cells, 6),
			Comment:          optStr(cells, 7),
		}
		rows = append(rows, row)
	}
	return rows, nil
}

// Time helpers reuse the special_equipment.go pkg-locals (deref, derefDate,
// optStr, optDate, isEmptyRow, writeHeaders). The compiler enforces no-dupe.
var _ = time.Time{}
