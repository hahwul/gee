package gee

import "os"

// ClosedFiles is closing list of files
func ClosedFiles(files []*os.File) {
	for _, v := range files {
		v.Close()
	}
}
