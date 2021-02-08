package gee

import (
	"bufio"
	"fmt"
	"os"
	//	"sync"

	model "github.com/hahwul/gee/pkg/model"
)

// Gee is running gee
func Gee(options model.Options) {
	sc := bufio.NewScanner(os.Stdin)
	mode := os.O_CREATE | os.O_WRONLY
	var files = []*os.File{}

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
		line := sc.Text()
		fmt.Println(line)
		for _, k := range files {
			_, err := k.WriteString(line)
			if err != nil {

			}
		}
	}
	//	wg.Wait()
}
