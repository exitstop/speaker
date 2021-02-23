package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input := `LaunchAll a browser with URL http://127.0.0.1:6080/#/?video, where video means to start with video mode. Now you can start Chromium in start menu (Internet -> Chromium Web Browser Sound) and try to play some video.

Following is the screen capture of these operations. Turn on your sound at the end of video!

`
	out := RegexWork(input)
	fmt.Println(out)
}

func RegexWork(tt string) string {
	reg, _ := regexp.Compile(`[\_\]\[\@\#\/]+`)
	reg2, _ := regexp.Compile(`([\p{L}])\.([\p{L}])`)
	reg3, _ := regexp.Compile(`([[:lower:]])([[:upper:]])`)
	reg4, _ := regexp.Compile(`(\b(\p{L}+)\b)`)
	tt = reg.ReplaceAllString(tt, " ")
	tt = reg2.ReplaceAllString(tt, "$1. $2")
	tt = reg3.ReplaceAllString(tt, "$1 $2")
	tt = reg4.ReplaceAllString(tt, " $1 ")

	tt = strings.TrimSpace(tt)
	return tt
}
