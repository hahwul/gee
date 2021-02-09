package gee

import (
	"os"

	model "github.com/hahwul/gee/pkg/model"
	printing "github.com/hahwul/gee/pkg/printing"
)

// ClosedFiles is closing list of files
func ClosedFiles(files []*os.File) {
	for _, v := range files {
		v.Close()
	}
}

// WriteFile is write to file
func WriteFile(k *os.File, line string, options model.Options) {
	if options.RemoveNewLine {
		_, err := k.WriteString(line)
		printing.ErrPrint(err)
	} else {
		_, err := k.WriteString(line + "\r\n")
		printing.ErrPrint(err)
	}
}
