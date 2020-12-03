package gomrepo

// MD is the table formatter for markdown.
var MD tableFormatter = (*md)(nil)

// tableFormatter is the formats to output go module report.
type tableFormatter interface {
	table([]pkginfo) string
}
