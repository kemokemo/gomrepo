package gomrepo

import (
	"fmt"
	"io"
	"sort"
	"text/template"
)

type md struct{}

const formatMd = `|ID|Version|License|
|:---|:---|:---|
{{range .}}{{if .Error}}{{else}}|{{.ID}}|{{.Version}}|{{.License}}|
{{end}}{{end}}`

// Format returns the table format string for markdown.
func (m *md) table(w io.Writer, pkgs []pkginfo) error {
	if len(pkgs) == 0 {
		return fmt.Errorf("there is no data to be formatted")
	}
	tpl, err := template.New("").Parse(formatMd)
	if err != nil {
		return fmt.Errorf("failed to parse template for markdown: %v", err)
	}
	sort.SliceStable(pkgs, func(i, j int) bool {
		return pkgs[i].ID < pkgs[j].ID
	})
	err = tpl.Execute(w, pkgs)
	if err != nil {
		return fmt.Errorf("failed to execute template for markdown: %v", err)
	}
	return nil
}
