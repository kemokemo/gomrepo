package gomrepo

import (
	"fmt"
	"io"
	"sort"
	"text/template"
)

type html struct{}

const formatHTML = `<table>
	<thead>
		<tr>
		  <th>ID</th>
		  <th>Version</th>
		  <th>License</th>
		</tr>
	</thead>
	<tbody>{{range .}}{{if .Error}}{{else}}
		<tr>
		  <td>{{.ID}}</td>
		  <td>{{.Version}}</td>
		  <td>{{.License}}</td>
		</tr>{{end}}{{end}}
	</tbody>
</table>
`

func (h *html) table(w io.Writer, pkgs []pkginfo) error {
	if len(pkgs) == 0 {
		return fmt.Errorf("there is no data to be formatted")
	}
	tpl, err := template.New("").Parse(formatHTML)
	if err != nil {
		return fmt.Errorf("failed to parse template for HTML: %v", err)
	}
	sort.SliceStable(pkgs, func(i, j int) bool {
		return pkgs[i].ID < pkgs[j].ID
	})
	err = tpl.Execute(w, pkgs)
	if err != nil {
		return fmt.Errorf("failed to execute template for HTML: %v", err)
	}
	return nil
}
