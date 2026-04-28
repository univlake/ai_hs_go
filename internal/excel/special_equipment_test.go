package excel

import (
	"bytes"
	"testing"
	"time"

	"ai_hs_go/internal/model"
)

func TestSpecialEquipmentTemplateAndExport(t *testing.T) {
	var buf bytes.Buffer
	if err := SpecialEquipmentTemplate(&buf); err != nil {
		t.Fatalf("template: %v", err)
	}
	if buf.Len() == 0 {
		t.Fatal("empty template")
	}

	name := "球磨机"
	loc := "1号车间"
	useTime := model.Date(time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC))
	rows := []model.SpecialEquipment{
		{EquipmentName: &name, UtilizedLocation: &loc, UseTime: &useTime},
	}
	var out bytes.Buffer
	if err := SpecialEquipmentExport(&out, rows); err != nil {
		t.Fatalf("export: %v", err)
	}
	if out.Len() == 0 {
		t.Fatal("empty export")
	}
}

func TestSpecialEquipmentImportRoundTrip(t *testing.T) {
	name := "桥式起重机"
	loc := "1号车间"
	useTime := model.Date(time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC))
	rows := []model.SpecialEquipment{
		{EquipmentName: &name, UtilizedLocation: &loc, UseTime: &useTime},
	}
	var out bytes.Buffer
	if err := SpecialEquipmentExport(&out, rows); err != nil {
		t.Fatalf("export: %v", err)
	}

	parsed, err := SpecialEquipmentImport(bytes.NewReader(out.Bytes()))
	if err != nil {
		t.Fatalf("import: %v", err)
	}
	if len(parsed) != 1 {
		t.Fatalf("expected 1 row, got %d", len(parsed))
	}
	got := parsed[0]
	if got.EquipmentName == nil || *got.EquipmentName != name {
		t.Fatalf("equipment name mismatch: %+v", got.EquipmentName)
	}
	if got.UtilizedLocation == nil || *got.UtilizedLocation != loc {
		t.Fatalf("location mismatch: %+v", got.UtilizedLocation)
	}
}
