package gee

import (
	"fmt"

	model "github.com/hahwul/gee/pkg/model"
)

// StdPrint is printing stdout
func StdPrint(line string, options model.Options) {
	if options.RemoveNewLine {
		fmt.Print(line)
	} else {
		fmt.Println(line)
	}
}
