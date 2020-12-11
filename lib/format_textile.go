package gomrepo

import (
	"fmt"
	"io"
	"sort"
	"text/template"
)

type textile struct{}

const formatTextile = `|_. ID |_. Version |_. License |
{{range .}}{{if .Error}}{{else}}| {{.ID}} | {{.Version}} | {{.License}} |
{{end}}{{end}}`

func (a *textile) table(w io.Writer, pkgs []pkginfo) error {
	if len(pkgs) == 0 {
		return fmt.Errorf("there is no data to be formatted")
	}
	tpl, err := template.New("").Parse(formatTextile)
	if err != nil {
		return fmt.Errorf("failed to parse template for Textile: %v", err)
	}
	sort.SliceStable(pkgs, func(i, j int) bool {
		return pkgs[i].ID < pkgs[j].ID
	})
	err = tpl.Execute(w, pkgs)
	if err != nil {
		return fmt.Errorf("failed to execute template for Textile: %v", err)
	}
	return nil
}
