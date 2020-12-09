package gomrepo

import "io"

// MD is the table formatter for markdown.
var MD Formatter = (*md)(nil)

// HTML is the table formatter for HTML.
var HTML Formatter = (*html)(nil)

// Formatter is the formats to output go module report.
type Formatter interface {
	table(io.Writer, []pkginfo) error
}
