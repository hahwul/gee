package gee

import (
	"fmt"
	"strings"

	model "github.com/hahwul/gee/pkg/model"
)

// StdPrint is printing stdout
func StdPrint(line string, options model.Options) {
	if options.RemoveNewLine {
		fmt.Print(strings.Replace(line, "\n", "", -1))
	} else {
		fmt.Print(line)
	}
}
