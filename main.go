package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	app "github.com/hahwul/gee/pkg/gee"
	model "github.com/hahwul/gee/pkg/model"
	printing "github.com/hahwul/gee/pkg/printing"
	"github.com/logrusorgru/aurora"
)

func main() {
	// Commandline parse
	versionOption := flag.Bool("version", false, "Version of gee")
	appendOption := flag.Bool("append", false, "Append mode for files")
	chunkedLineOption := flag.Int("chunked", 0, "Chuked files from line (e.g output / output_1 / output_2)")
	withLineOption := flag.Bool("with-line", false, "With line number (colorize blue)")
	withTimeOption := flag.Bool("with-time", false, "With timestamp (colorize blue)")
	prefixOption := flag.String("prefix", "", "Prefix string")
	suffixOption := flag.String("suffix", "", "Suffix string")
	rmnlOption := flag.Bool("rmnl", false, "Remove newline(\\r\\n)")
	distributeOption := flag.Bool("distribute", false, "Distribution to files")
	colorOption := flag.Bool("uncolor", false, "Uncolorize stdout")
	regexOption := flag.String("regex", "", "Match with Regular Expression (like grep)")
	regexvOption := flag.String("regexv", "", "Unmatch with Regular Expression (like grep -v)")
	findOption := flag.String("find", "", "Find string in line (colorize red)")
	replaceOption := flag.String("replace", "", "Replace string in line with '-find' option")
	splitOption := flag.String("split", "", "Split string within line. (to line , to table, to md-table)")
	formatOption := flag.String("format", "line", "Change output format (json, md-table, html-table)")

	// Custom usage
	flag.Usage = func() {
		printing.Banner()
		fmt.Fprintf(os.Stderr, aurora.White("Usage: %s [flags] [file1] [file2] ...\n").String(), os.Args[0])
		fmt.Fprintf(os.Stderr, "(If you do not specify a file, only stdout is output)\n\n")
		fmt.Fprintf(os.Stderr, aurora.White("Flags:\n").String())
		flag.PrintDefaults()
	}

	// Flag parse
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
		Files:         files,
		Append:        *appendOption,
		ChunkedLine:   *chunkedLineOption,
		WithLine:      *withLineOption,
		WithTimestamp: *withTimeOption,
		Prefix:        *prefixOption,
		Suffix:        *suffixOption,
		RemoveNewLine: *rmnlOption,
		Distribute:    *distributeOption,
		Regex:         *regexOption,
		RegexV:        *regexvOption,
		Replace:       *replaceOption,
		Find:          *findOption,
		Color:         !*colorOption,
		Split:         *splitOption,
		Format:        *formatOption,
	}

	// Running gee app
	app.Gee(options)
}
