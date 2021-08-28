package main

import (
	"fmt"

	"github.com/mertcandav/dinogo/input"
	"github.com/mertcandav/dinogo/terminal"
)

func main() {
	terminal.ForegroundByRGB(255, 0, 0)
	fmt.Println("TEST")
	_, _ = input.ReadLine()
}
