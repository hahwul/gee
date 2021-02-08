package main

import (
	"flag"
	"fmt"

	app "github.com/hahwul/gee/pkg/gee"
	model "github.com/hahwul/gee/pkg/model"
	printing "github.com/hahwul/gee/pkg/printing"
)

func main() {
	// Commandline parse
	version := flag.Bool("-version", false, "version of gee")
	flag.Parse()
	options := model.Options{
		Files: flag.Args(),
	}

	// Show version
	if *version {
		fmt.Println(printing.VERSION)
	}
	// Running gee app
	app.Gee(options)
}
