package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	pwd := os.Args[1]

	home := os.Getenv("HOME")
	srcin := home + "/Developer/src/"
	srcout := "~src/"
	confin := home + "/Developer/conf/"
	confout := "~conf/"

	if strings.HasPrefix(pwd, srcin) {
		fmt.Printf(strings.Replace(pwd, srcin, srcout, 1))
	} else if strings.HasPrefix(pwd, confin) {
		fmt.Printf(strings.Replace(pwd, confin, confout, 1))
	} else if strings.HasPrefix(pwd, home) {
		fmt.Printf(strings.Replace(pwd, home, "~", 1))
	} else {
		fmt.Printf(pwd)
	}
}
