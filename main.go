package main

import (
	"flag"
	"fmt"

	model "github.com/hahwul/gee/pkg/model"
	printing "github.com/hahwul/gee/pkg/printing"
)

func main() {
	version := flag.Bool("version", false, "version of gee")
	flag.Parse()
	options := model.Options{}
	_ = options
	if *version {
		fmt.Println(printing.VERSION)
	}

}
