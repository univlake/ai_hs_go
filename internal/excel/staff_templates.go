package excel

import (
	"bytes"
	"io"
	"strconv"

	"github.com/xuri/excelize/v2"

	"ai_hs_go/internal/model"
)

// 人员管理三个子表的导入模板。每张表只有几列，复用 writeHeaders +
// writeExampleRow（在 special_equipment.go 里），格式跟特种设备一致。

var safetyManagementStaffHeaders = []string{
	"姓名", "人员类型ID", "联系电话", "证书名称", "证书编号", "有效期",
}

var safetyManagementStaffExample = []any{
	"张三", 1, "13800000000", "安全管理人员证", "AQ-2024-001", "2026-12-31",
}

func SafetyManagementStaffTemplate(w io.Writer) error {
	return writeStaffTemplate(w, "安全管理人员",
		safetyManagementStaffHeaders, safetyManagementStaffExample)
}

var specialEquipmentStaffHeaders = []string{
	"姓名", "证书名称", "证书编号", "有效期",
}

var specialEquipmentStaffExample = []any{
	"李四", "电梯操作证", "TS-2024-002", "2025-12-31",
}

func SpecialEquipmentStaffTemplate(w io.Writer) error {
	return writeStaffTemplate(w, "特种作业人员",
		specialEquipmentStaffHeaders, specialEquipmentStaffExample)
}

var eliminationControlStaffHeaders = []string{
	"姓名", "证书名称", "证书编号", "有效期",
}

var eliminationControlStaffExample = []any{
	"王五", "建(构)筑物消防员证", "XF-2024-003", "2027-06-30",
}

func EliminationControlStaffTemplate(w io.Writer) error {
	return writeStaffTemplate(w, "消防控制室人员",
		eliminationControlStaffHeaders, eliminationControlStaffExample)
}

// ── Importers ──────────────────────────────────────────────────
//
// Caller is responsible for setting company_id on each returned row.
// First row is the header (skipped), and rows whose every cell is empty
// are dropped — that includes the example row when the user blanks it
// out before re-uploading.

// SafetyManagementStaffImport: name | safety_management_type_id | tel | cert_title | cert_number | valid_period
func SafetyManagementStaffImport(r io.Reader) ([]model.SafetyManagementStaff, error) {
	rowsIter, closeFn, err := openStaffSheet(r)
	if err != nil {
		return nil, err
	}
	defer closeFn()

	var rows []model.SafetyManagementStaff
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
		typeID := int64(1)
		if v := optStr(cells, 1); v != nil {
			if n, err := strconv.ParseInt(*v, 10, 64); err == nil {
				typeID = n
			}
		}
		rows = append(rows, model.SafetyManagementStaff{
			Name:                   optStr(cells, 0),
			SafetyManagementTypeID: typeID,
			Tel:                    optStr(cells, 2),
			CertTitle:              optStr(cells, 3),
			CertNumber:             optStr(cells, 4),
			ValidPeriod:            optDate(cells, 5),
		})
	}
	return rows, nil
}

// SpecialEquipmentStaffImport: name | cert_title | cert_number | valid_period
func SpecialEquipmentStaffImport(r io.Reader) ([]model.SpecialEquipmentStaff, error) {
	rowsIter, closeFn, err := openStaffSheet(r)
	if err != nil {
		return nil, err
	}
	defer closeFn()

	var rows []model.SpecialEquipmentStaff
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
		rows = append(rows, model.SpecialEquipmentStaff{
			Name:        optStr(cells, 0),
			CertTitle:   optStr(cells, 1),
			CertNumber:  optStr(cells, 2),
			ValidPeriod: optDate(cells, 3),
		})
	}
	return rows, nil
}

// EliminationControlStaffImport: name | cert_title | cert_number | valid_period
func EliminationControlStaffImport(r io.Reader) ([]model.EliminationControlStaff, error) {
	rowsIter, closeFn, err := openStaffSheet(r)
	if err != nil {
		return nil, err
	}
	defer closeFn()

	var rows []model.EliminationControlStaff
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
		rows = append(rows, model.EliminationControlStaff{
			Name:        optStr(cells, 0),
			CertTitle:   optStr(cells, 1),
			CertNumber:  optStr(cells, 2),
			ValidPeriod: optDate(cells, 3),
		})
	}
	return rows, nil
}

// openStaffSheet wraps reading + opening the first sheet of the
// uploaded xlsx. Returns the row iterator + a close func; caller defers
// the close — keeps the 3 importers DRY.
func openStaffSheet(r io.Reader) (*excelize.Rows, func(), error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}
	f, err := excelize.OpenReader(bytes.NewReader(buf))
	if err != nil {
		return nil, nil, err
	}
	sheet := f.GetSheetName(0)
	rowsIter, err := f.Rows(sheet)
	if err != nil {
		_ = f.Close()
		return nil, nil, err
	}
	closeFn := func() {
		rowsIter.Close()
		_ = f.Close()
	}
	return rowsIter, closeFn, nil
}
