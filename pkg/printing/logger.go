package printing

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// DebugMsg is msg for debug
func DebugMsg(t string,m interface{}){
	fmt.Printf("%s%s%s\n",aurora.BrightYellow("[DEBUG]").String(),aurora.BrightBlue("["+t+"] ").String(),m)
}

// ErrPanic is panic error
func ErrPanic(e error) {

}

// ErrPrint is print error
func ErrPrint(e error) {

}
