package main

import (
	"fmt"

	"github.com/mertcandav/dinogo/ansiescape"
	"github.com/mertcandav/dinogo/input"
)

func main() {
	fmt.Println(ansiescape.Bold, "TEST", ansiescape.ResetFont)
	_, _ = input.ReadLine()
}
