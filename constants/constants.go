package constants

import "strings"

const (
	SEPARATOR            = "↝"
	STR_WALL             = "| "
	STR_TITLE_RIGHT      = "Rule's Right (↩ confirm/epsilon)\n| "
	STR_TITLE_LEFT_FIRST = "Rule's Left (↩ continue)\n| "
	STR_TITLE_LEFT       = "Rule's Left (↩ continue/cancel)\n| "

	// used for separating non-terminal symbols from terminal symbols
	// in the grammar input
	NT_SUFFIX        string = "*"
	REAL_NT_SUFFIX   string = "\a"
	ESCAPE_CHARACTER string = "\\"
)

func StringReal(str string) string {
	return strings.ReplaceAll(str, REAL_NT_SUFFIX, NT_SUFFIX)
}
