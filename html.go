package utils

import (
	"html"
)

// htmlentities
func HTMLEntityEncode(str string) string {
	return html.EscapeString(str)
}

// html_entity_decode
func HTMLEntityDecode(str string) string {
	return html.UnescapeString(str)
}
