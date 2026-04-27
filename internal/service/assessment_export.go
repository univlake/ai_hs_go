// Word 导出 — 考核汇总（assessment-summary/export）。
//
//   bloc:    public/check_bloc.docx        （动态行：${sort} ${company_name} ${number}）
//   company: public/check_company.docx     year ≤ 2025
//   company: public/check_company_new.docx year > 2025
//
// PhpWord 行为对齐：
//   - bloc 模板里 `${company_title}` 被拆成多个 <w:t> run，本地做一次合并预处理。
//   - bloc 模板里 ${sort} 所在 <w:tr> 是模板行；按 cloneRowAndSetValues 复制并替换。
//   - 普通占位符使用 expert_service.go 的同款 substitute 流程（zip → 替换 → 重新打包）。

package service

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"ai_hs_go/internal/errs"
)

// ExportSummary — 把汇总结果渲染到 docx 模板，落盘到 storageDir 后返回路径 + 文件名。
//
//   templateDir: 模板根目录，需包含 check_bloc.docx / check_company.docx / check_company_new.docx
//   storageDir : 文件落盘根目录；最终路径会落到 storageDir/check_word/YYYY/MM/<filename>
//
// kind == "bloc" / "company" 决定使用哪套模板。
func (s *AssessmentService) ExportSummary(
	ctx context.Context,
	companyID int64,
	year int,
	kind string,
	templateDir, storageDir string,
) (absPath, filename string, err error) {
	now := time.Now()
	dir := filepath.Join(storageDir, "check_word", now.Format("2006"), now.Format("01"))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", "", errs.Wrap(errs.CodeInternal, "创建导出目录失败", err)
	}

	switch kind {
	case "bloc":
		rows, err := s.SummaryBloc(ctx, companyID, year)
		if err != nil {
			return "", "", err
		}
		parentTitle := s.CompanyTitle(ctx, companyID)
		templatePath := filepath.Join(templateDir, "check_bloc.docx")
		if _, err := os.Stat(templatePath); err != nil {
			return "", "", errs.New(errs.CodeInternal, "导出模板缺失：check_bloc.docx")
		}
		filename = fmt.Sprintf("%s%d%d考核结果.docx", parentTitle, year, rand.Int())
		absPath = filepath.Join(dir, filename)

		topReplacements := map[string]string{
			"${company_title}": parentTitle,
			"${year}":          strconv.Itoa(year),
		}
		rowValues := make([]map[string]string, 0, len(rows))
		for i, r := range rows {
			rowValues = append(rowValues, map[string]string{
				"sort":         strconv.Itoa(i + 1),
				"company_name": r.CompanyName,
				"number":       formatFloat(r.Number),
			})
		}
		if err := renderTemplateWithRowClone(templatePath, absPath, topReplacements, "sort", rowValues); err != nil {
			return "", "", errs.Wrap(errs.CodeInternal, "生成文档失败", err)
		}
		return absPath, filename, nil

	case "company":
		row, err := s.SummaryCompany(ctx, companyID, year)
		if err != nil {
			return "", "", err
		}
		templateName := "check_company.docx"
		if year > 2025 {
			templateName = "check_company_new.docx"
		}
		templatePath := filepath.Join(templateDir, templateName)
		if _, err := os.Stat(templatePath); err != nil {
			return "", "", errs.New(errs.CodeInternal, "导出模板缺失："+templateName)
		}
		filename = fmt.Sprintf("%s%d%d考核结果.docx", row.CompanyName, year, rand.Int())
		absPath = filepath.Join(dir, filename)

		replacements := map[string]string{
			"${company_name}": row.CompanyName,
			"${year}":         strconv.Itoa(year),
			"${number}":       formatFloat(row.Number),
			"${number1}":      formatFloat(row.Number1),
			"${number2}":      formatFloat(row.Number2),
			"${number3}":      formatFloat(row.Number3),
			"${number4}":      formatFloat(row.Number4),
			"${number5}":      formatFloat(row.Number5),
			"${number6}":      formatFloat(row.Number6),
		}
		if err := renderTemplateWithRowClone(templatePath, absPath, replacements, "", nil); err != nil {
			return "", "", errs.Wrap(errs.CodeInternal, "生成文档失败", err)
		}
		return absPath, filename, nil

	default:
		return "", "", errs.New(errs.CodeUnprocessable, "未知 type: "+kind)
	}
}

// renderTemplateWithRowClone — 单 docx 渲染 helper：解压 → 预处理拆分占位符 → cloneRow（如 trigger 非空）
// → 占位符替换 → 重新打包。
//
// rowTrigger 为 "" 时跳过 cloneRow（适用于 company 模板）。
func renderTemplateWithRowClone(srcPath, dstPath string, replacements map[string]string, rowTrigger string, rowValues []map[string]string) error {
	r, err := zip.OpenReader(srcPath)
	if err != nil {
		return fmt.Errorf("open template: %w", err)
	}
	defer r.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("create output: %w", err)
	}
	defer dstFile.Close()
	w := zip.NewWriter(dstFile)
	defer w.Close()

	for _, f := range r.File {
		fw, err := w.CreateHeader(&zip.FileHeader{
			Name:     f.Name,
			Method:   f.Method,
			Modified: f.Modified,
		})
		if err != nil {
			return fmt.Errorf("zip header %s: %w", f.Name, err)
		}
		fr, err := f.Open()
		if err != nil {
			return fmt.Errorf("open entry %s: %w", f.Name, err)
		}
		if isDocxTextPart(f.Name) {
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, fr); err != nil {
				fr.Close()
				return fmt.Errorf("read entry %s: %w", f.Name, err)
			}
			fr.Close()
			content := buf.String()

			// 1) 合并跨 <w:t> run 的占位符（PhpWord 兼容）。
			content = mergeSplitPlaceholders(content)

			// 2) cloneRow（仅 document.xml 有意义；trigger 在哪个 part 就在哪个 part 处理）。
			if rowTrigger != "" && len(rowValues) > 0 {
				content = cloneRowAndSetValues(content, rowTrigger, rowValues)
			}

			// 3) 占位符替换。
			for k, v := range replacements {
				content = strings.ReplaceAll(content, k, escapeXML(v))
			}
			if _, err := fw.Write([]byte(content)); err != nil {
				return fmt.Errorf("write entry %s: %w", f.Name, err)
			}
			continue
		}
		// 非文本 part 原样拷贝。
		if _, err := io.Copy(fw, fr); err != nil {
			fr.Close()
			return fmt.Errorf("copy entry %s: %w", f.Name, err)
		}
		fr.Close()
	}
	return nil
}

// mergeSplitPlaceholdersRe — 占位符 `${...}` 内部允许出现 XML 标签：
//   匹配 `\$\{[^}]*\}` 后剥离 markup，重组为单 token。
//
// 示例：`<w:t>${company_</w:t></w:r><w:r>...<w:t>title</w:t></w:r><w:r>...<w:t>}</w:t>`
// 结果：`<w:t>${company_title}</w:t></w:r>`（后两个 <w:r> 被吞掉，但其包裹结构早已闭合）。
var mergeSplitPlaceholdersRe = regexp.MustCompile(`\$\{[^}]*\}`)
var stripXMLTagsRe = regexp.MustCompile(`<[^>]*>`)

func mergeSplitPlaceholders(content string) string {
	return mergeSplitPlaceholdersRe.ReplaceAllStringFunc(content, func(match string) string {
		// match 形如 ${...}；剥掉中间的 XML markup。
		inner := match[2 : len(match)-1]
		clean := stripXMLTagsRe.ReplaceAllString(inner, "")
		// 占位符名只能是字母 / 数字 / 下划线；做一次防御性过滤。
		var b strings.Builder
		for _, r := range clean {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
				b.WriteRune(r)
			}
		}
		if b.Len() == 0 {
			return match
		}
		return "${" + b.String() + "}"
	})
}

// cloneRowAndSetValues — 找到包含 `${trigger}` 的 <w:tr>...</w:tr>，按 rowValues 复制 N 份并代入。
//
// 形如 PhpWord TemplateProcessor::cloneRowAndSetValues(triggerName, rowValues)。
func cloneRowAndSetValues(content string, trigger string, rowValues []map[string]string) string {
	placeholder := "${" + trigger + "}"
	idx := strings.Index(content, placeholder)
	if idx < 0 {
		return content
	}
	rowStart := findRowStart(content, idx)
	if rowStart < 0 {
		return content
	}
	rowEndOff := strings.Index(content[idx:], "</w:tr>")
	if rowEndOff < 0 {
		return content
	}
	rowEnd := idx + rowEndOff + len("</w:tr>")
	rowXML := content[rowStart:rowEnd]

	var sb strings.Builder
	sb.Grow(rowEnd - rowStart)
	for _, vals := range rowValues {
		clone := rowXML
		for k, v := range vals {
			clone = strings.ReplaceAll(clone, "${"+k+"}", escapeXML(v))
		}
		sb.WriteString(clone)
	}
	return content[:rowStart] + sb.String() + content[rowEnd:]
}

// findRowStart — 从 idx 向前找 `<w:tr>` 或 `<w:tr ...>` 开标签的起始位置。
func findRowStart(content string, idx int) int {
	for j := idx; j > 0; j-- {
		if j+6 <= len(content) && content[j:j+6] == "<w:tr>" {
			return j
		}
		if j+6 <= len(content) && content[j:j+6] == "<w:tr " {
			return j
		}
	}
	return -1
}

// formatFloat — 跟前端展示对齐：去掉无意义的尾随 0，但保留至少 1 位小数。
func formatFloat(v float64) string {
	if v == float64(int64(v)) {
		return strconv.FormatInt(int64(v), 10)
	}
	return strconv.FormatFloat(v, 'f', -1, 64)
}
