package gee

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//	"sync"

	model "github.com/hahwul/gee/pkg/model"
	printing "github.com/hahwul/gee/pkg/printing"
)

// Gee is running gee
func Gee(options model.Options) {
	sc := bufio.NewScanner(os.Stdin)
	mode := os.O_CREATE | os.O_WRONLY
	var files = []*os.File{}
	stdLine := 1
	stdPointer := 1
	distributePointer := 0

	if options.Append {
		mode = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	}

	for _, filename := range options.Files {
		f, err := os.OpenFile(filename, mode, 0644)
		if err != nil {

		} else {
			files = append(files, f)
		}
	}
	//	var wg sync.WaitGroup
	//	wg.Add(1 + len(options.Files))
	for sc.Scan() {
		l := sc.Text()
		line := ""
		if options.WithLine {
			l = "[" + strconv.Itoa(stdLine) + "] " + l
		}
		if options.WithTimestamp {
			l = "[" + GetNowTime() + "] " + l
		}

		// Prefix and Suffix
		line = options.Prefix + l + options.Suffix

		if options.RemoveNewLine {
			fmt.Print(line)
		} else {
			fmt.Println(line)
		}
		if (stdLine > options.ChunkedLine) && (options.ChunkedLine != 0) {
			ClosedFiles(files)
			for _, filename := range options.Files {
				f, err := os.OpenFile(filename+"_"+strconv.Itoa(stdPointer), mode, 0644)
				if err != nil {
					printing.ErrPrint(err)
				} else {
					files = append(files, f)
				}
			}
			stdLine = 1
			stdPointer = stdPointer + 1
		}
		if options.Distribute && (len(files) > 0) {
			if distributePointer < len(files) {
				WriteFile(files[distributePointer], line, options)
				distributePointer = distributePointer + 1
			} else {
				distributePointer = 0
			}
		} else {
			for _, k := range files {
				WriteFile(k, line, options)
			}
		}
		stdLine = stdLine + 1
	}
	//	wg.Wait()

	// Graceful shutdown
	ClosedFiles(files)
}
