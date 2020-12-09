package gomrepo

import "io"

// MD is the table formatter for markdown.
var MD tableFormatter = (*md)(nil)

// tableFormatter is the formats to output go module report.
type tableFormatter interface {
	table(io.Writer, []pkginfo) error
}
