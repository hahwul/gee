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
		result = au.BrightGreen("["+GetNowTime()+"] ").String() + result
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
		returnString = "{\"line\":[" + returnString + "\"\"]}\n"
		returnPlainString = "{\"line\":[" + returnPlainString + "\"\"]}\n"
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
		returnString = "\"" + setFix(str, options) + "\","
	case "md-table":
		returnString = setFix(str+"|", options)
	case "html-table":
		returnString = "<td>" + setFix(str, options) + "</td>"
	default:
		returnString = setFix(str, options) + "\n"
	}
	return returnString
}

func setFix(str string, options model.Options) string {
	pf := setPrefix(strings.Replace(str, "\n", "", -1), options)
	sf := setSuffix(pf, options)
	return sf
}

func setPrefix(str string, options model.Options) string {
	return options.Prefix + str
}
func setSuffix(str string, options model.Options) string {
	return str + options.Suffix
}

func setReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
