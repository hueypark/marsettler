package behavior_tree

import (
	"fmt"
	"strings"
)

// Indent creates an indented string.
func Indent(format string, a ...interface{}) string {
	strs := strings.Split(fmt.Sprintf(format, a...), "\n")
	var indentStr string
	for _, str := range strs {
		if str == "" {
			continue
		}
		indentStr += "\t" + str + "\n"
	}

	return indentStr
}
