package gee

import (
	"strconv"
	"strings"

	"github.com/hahwul/gee/pkg/model"
	"github.com/logrusorgru/aurora"
)

// StringProc is string
func StringProc(l string, stdLine int, options model.Options) (string, string) {
	result := l
	resultPlain := l
	au := aurora.NewAurora(options.Color)
	if options.Find != "" {
		if options.Replace != "" {
			result = strings.Replace(result, options.Find, au.BrightRed(options.Replace).String(), -1)
			resultPlain = strings.Replace(resultPlain, options.Find, options.Replace, -1)
		} else {
			result = strings.Replace(result, options.Find, au.BrightRed(options.Find).String(), -1)
		}
	}
	if options.WithLine {
		result = au.BrightBlue("["+strconv.Itoa(stdLine)+"] ").String() + result
		resultPlain = "[" + strconv.Itoa(stdLine) + "] " + resultPlain
	}
	if options.WithTimestamp {
		result = au.BrightBlue("["+GetNowTime()+"] ").String() + result
		resultPlain = "[" + GetNowTime() + "] " + resultPlain
	}
	return result, resultPlain
}
