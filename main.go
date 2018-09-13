package main

import (
	"flag"
	"fmt"
	"os/exec"
)

var (
	row = flag.Int("row", 8, "行数")
	col = flag.Int("col", 10, "列数")
)
var (
	cmd1 = "defaults write com.apple.dock springboard-rows -int %d"
	cmd2 = "defaults write com.apple.dock springboard-columns -int %d"
	cmd3 = "killall Dock"
)

func main() {
	flag.Parse()
	exec.Command("/bin/bash", "-c", fmt.Sprintf(cmd1, *row)).Run()
	exec.Command("/bin/bash", "-c", fmt.Sprintf(cmd2, *col)).Run()
	exec.Command("/bin/bash", "-c", fmt.Sprintf(cmd3)).Run()
}
