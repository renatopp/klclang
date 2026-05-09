package utils

import (
	"strings"
	"unicode/utf8"

	"github.com/renatopp/x/strx"
)

type TableWriter struct {
	title   string
	Headers []string
	Lengths []int
	Rows    [][]string
}

func NewTableWriter(headers ...string) *TableWriter {
	lengths := make([]int, len(headers))
	for i, header := range headers {
		lengths[i] = utf8.RuneCountInString(header)
	}

	return &TableWriter{
		Headers: headers,
		Lengths: lengths,
		Rows:    [][]string{},
	}
}

func (t *TableWriter) Title(title string) {
	t.title = title
}

func (t *TableWriter) Write(values ...string) {
	for i, value := range values {
		curLength := t.Lengths[i]
		valLength := utf8.RuneCountInString(value)
		if valLength > curLength {
			t.Lengths[i] = valLength
		}
	}
	t.Rows = append(t.Rows, values)
}

func (t *TableWriter) Render() string {
	var sb strings.Builder
	size := t.rowLength()
	if t.title != "" {
		sb.WriteString(strx.Format("+-%s-+\n", strx.Repeat("-", size)))
		sb.WriteString(strx.Format("| %s |\n", strx.PadCenter(t.title, size)))
	}
	sb.WriteString(strx.Format("+-%s-+\n", strx.Repeat("-", size)))
	sb.WriteString(strx.Format("| %s |\n", t.fprintHeader(t.Headers...)))
	sb.WriteString(strx.Format("+-%s-+\n", t.fprintSeparators()))
	for _, row := range t.Rows {
		sb.WriteString(strx.Format("| %s |\n", t.fprintRow(row...)))
	}
	sb.WriteString(strx.Format("+-%s-+\n", strx.Repeat("-", size)))
	return sb.String()
}

func (t *TableWriter) fprintHeader(values ...string) string {
	i := -1
	return strx.JoinFunc(values, func(value string) string {
		i++
		return strx.PadCenter(value, t.Lengths[i])
	}, " | ")
}

func (t *TableWriter) fprintRow(values ...string) string {
	i := -1
	return strx.JoinFunc(values, func(value string) string {
		i++
		return strx.PadRight(value, t.Lengths[i])
	}, " | ")
}

func (t *TableWriter) fprintSeparators() string {
	return strx.JoinFunc(t.Lengths, func(length int) string {
		return strx.Repeat("-", length)
	}, "-+-")
}

func (t *TableWriter) rowLength() int {
	size := 0
	for _, length := range t.Lengths {
		size += length
	}
	return size + 3*(len(t.Lengths)-1)
}
