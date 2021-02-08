package main

import (
	"flag"
	"fmt"
	"strings"

	app "github.com/hahwul/gee/pkg/gee"
	model "github.com/hahwul/gee/pkg/model"
	printing "github.com/hahwul/gee/pkg/printing"
)

func main() {
	// Commandline parse
	versionOption := flag.Bool("version", false, "Version of gee")
	appendOption := flag.Bool("append", false, "Append mode for files")
	flag.Parse()

	// Show version
	if *versionOption {
		fmt.Println(printing.VERSION)
		return
	}

	// Finding file value
	var files []string
	args := flag.Args()
	for _, v := range args {
		if !strings.Contains(v, "-") {
			files = append(files, v)
		}
	}

	// Set Options
	options := model.Options{
		Files:  files,
		Append: *appendOption,
	}

	// Running gee app
	app.Gee(options)
}
