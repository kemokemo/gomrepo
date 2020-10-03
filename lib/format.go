package gomrepo

import (
	"fmt"
	"sort"
	"strings"
)

// MD is the table formatter for markdown.
var MD tableFormatter = (*md)(nil)

// tableFormatter is the formats to output go module report.
type tableFormatter interface {
	table([]pkginfo) string
}

type md struct{}

// Format returns the table format string for markdown.
func (m *md) table(pkgs []pkginfo) string {
	if len(pkgs) == 0 {
		return ""
	}

	var rows []string
	for _, pkg := range pkgs {
		if pkg.err != nil {
			continue
		}
		rows = append(rows, fmt.Sprintf("|%s|%s|%s|", pkg.id, pkg.ver, pkg.lic))
	}
	sort.Strings(rows)

	return strings.Join(rows, "\n")
}
