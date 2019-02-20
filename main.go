package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	out,_:=exec.Command("/usr/local/go/bin/go","env").Output()
	lines:=strings.Split(string(out),"\n")
	for _,line:=range lines {
		if line != "" {
			fmt.Println("export " + line)
		}
	}
}
