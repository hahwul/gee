package gee

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//	"sync"

	model "github.com/hahwul/gee/pkg/model"
)

// Gee is running gee
func Gee(options model.Options) {
	sc := bufio.NewScanner(os.Stdin)
	mode := os.O_CREATE | os.O_WRONLY
	var files = []*os.File{}
	stdLine := 1
	stdPointer := 1

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

		line = l

		fmt.Println(line)
		if (stdLine > options.ChunkedLine) && (options.ChunkedLine != 0) {
			ClosedFiles(files)
			for _, filename := range options.Files {
				f, err := os.OpenFile(filename+"_"+strconv.Itoa(stdPointer), mode, 0644)
				if err != nil {

				} else {
					files = append(files, f)
				}
			}
			stdLine = 1
			stdPointer = stdPointer + 1
		}
		for _, k := range files {
			_, err := k.WriteString(line)
			if err != nil {

			}
		}
		stdLine = stdLine + 1
	}
	//	wg.Wait()

	// Graceful shutdown
	ClosedFiles(files)
}
