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
	var resultArr []string
	var resultPlainArr []string
	var returnString string
	var returnPlainString string

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

	if options.Split != "" {
		resultArr = strings.Split(result, options.Split)
		resultPlainArr = strings.Split(resultPlain, options.Split)
	} else {
		resultArr = append(resultArr, result)
		resultPlainArr = append(resultPlainArr, resultPlain)
	}

	for _, v := range resultArr {
		returnString = returnString + toFormat(v, options)
	}
	for _, v := range resultPlainArr {
		returnPlainString = returnPlainString + toFormat(v, options)
	}
	switch options.Format {
	case "json":
		returnString = "{\"line\":\"[" + returnString + "\"\"]\"}\n"
		returnPlainString = "{\"line\":\"[" + returnPlainString + "\"\"]\"}\n"
	case "md-table":
		returnString = "|" + returnString + "\n"
		returnPlainString = "|" + returnPlainString + "\n"
	case "html-table":
		returnString = "<tr>" + returnString + "</tr>\n"
		returnPlainString = "<tr>" + returnPlainString + "</tr>\n"
	}
	return returnString, returnPlainString
}

func toFormat(str string, options model.Options) string {
	var returnString string
	switch options.Format {
	case "json":
		returnString = str + ","
	case "md-table":
		returnString = str + "|"
	case "html-table":
		returnString = "<td>" + str + "</td>"
	default:
		returnString = strings.Replace(str, "\n", "", -1) + "\n"
	}
	return returnString
}

func toJSON() {

}

func toMarkdownTable() {

}

func toHTMLTable() {

}
