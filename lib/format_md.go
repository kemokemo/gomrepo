package gomrepo

import (
	"fmt"
	"sort"
	"strings"
)

type md struct{}

const formatMdRow = "|%s|%s|%s|"

var (
	headerMd   = fmt.Sprintf(formatMdRow, "ID", "Version", "License")
	splitterMd = fmt.Sprintf(formatMdRow, ":---", ":---", ":---")
)

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
		rows = append(rows, fmt.Sprintf(formatMdRow, pkg.id, pkg.ver, pkg.lic))
	}
	sort.Strings(rows)
	rows = append([]string{splitterMd}, rows...)
	rows = append([]string{headerMd}, rows...)

	return strings.Join(rows, "\n")
}
