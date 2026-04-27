package service

import (
	"strings"
	"testing"
)

func TestMergeSplitPlaceholders(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "single-run placeholder",
			in:   `<w:t>${name}</w:t>`,
			want: `<w:t>${name}</w:t>`,
		},
		{
			name: "split across runs",
			in:   `<w:t>${company_</w:t></w:r><w:r><w:rPr><w:rFonts/></w:rPr><w:t>title</w:t></w:r><w:r><w:rPr/><w:t>}</w:t>`,
			want: `<w:t>${company_title}</w:t>`,
		},
		{
			name: "no placeholders",
			in:   `<w:t>plain text</w:t>`,
			want: `<w:t>plain text</w:t>`,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := mergeSplitPlaceholders(tc.in)
			if !strings.Contains(got, tc.want) {
				t.Errorf("mergeSplitPlaceholders(%q) = %q, want substring %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestCloneRowAndSetValues(t *testing.T) {
	// 简化的 docx 片段：一个表里只有一行，包含 ${sort} ${company_name} ${number} 三个占位符。
	template := `<w:tbl><w:tr><w:tc><w:p><w:r><w:t>${sort}</w:t></w:r></w:p></w:tc>` +
		`<w:tc><w:p><w:r><w:t>${company_name}</w:t></w:r></w:p></w:tc>` +
		`<w:tc><w:p><w:r><w:t>${number}</w:t></w:r></w:p></w:tc></w:tr></w:tbl>`

	values := []map[string]string{
		{"sort": "1", "company_name": "Foo Co", "number": "85"},
		{"sort": "2", "company_name": "Bar Co", "number": "92"},
	}

	got := cloneRowAndSetValues(template, "sort", values)

	// 期望两条 <w:tr>，分别替换好。
	if strings.Count(got, "<w:tr>") != 2 {
		t.Errorf("expected 2 rows, got %d in %q", strings.Count(got, "<w:tr>"), got)
	}
	if !strings.Contains(got, ">1<") || !strings.Contains(got, ">Foo Co<") || !strings.Contains(got, ">85<") {
		t.Errorf("missing first row substitutions in %q", got)
	}
	if !strings.Contains(got, ">2<") || !strings.Contains(got, ">Bar Co<") || !strings.Contains(got, ">92<") {
		t.Errorf("missing second row substitutions in %q", got)
	}
	// 占位符不应残留。
	if strings.Contains(got, "${sort}") || strings.Contains(got, "${number}") {
		t.Errorf("placeholders not substituted: %q", got)
	}
}

func TestCloneRowAndSetValues_NoTrigger(t *testing.T) {
	template := `<w:tbl><w:tr><w:tc>foo</w:tc></w:tr></w:tbl>`
	got := cloneRowAndSetValues(template, "missing", []map[string]string{{"missing": "v"}})
	if got != template {
		t.Errorf("expected unchanged when trigger missing, got %q", got)
	}
}

func TestFormatFloat(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{85.0, "85"},
		{85.5, "85.5"},
		{0, "0"},
		{0.1, "0.1"},
	}
	for _, tc := range cases {
		got := formatFloat(tc.in)
		if got != tc.want {
			t.Errorf("formatFloat(%v) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
