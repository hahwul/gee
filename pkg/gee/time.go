package gee

import (
	"fmt"
	"time"
)

// GetNowTime is Get time string now
func GetNowTime() string {
	now := time.Now()
	sec := now.Unix()
	rtn := fmt.Sprintf("%v", time.Unix(sec, 0))
	return rtn
}
