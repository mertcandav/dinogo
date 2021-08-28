package main

import (
	"fmt"

	"github.com/mertcandav/dinogo/ansiescape"
	"github.com/mertcandav/dinogo/input"
)

func main() {
	fmt.Print(ansiescape.PrevLine(1))
	_, _ = input.ReadRune()
}
