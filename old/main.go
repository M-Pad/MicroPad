package main

import (
	"MicroPad/tui"
	"MicroPad/tui/common"
	"MicroPad/tui/theme"
	"fmt"
	"os"
)

func main() {
	t, err := theme.LoadFromPath("./themes/base.mptheme.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	common.Theme = t

	tui.TuiMain()
}
