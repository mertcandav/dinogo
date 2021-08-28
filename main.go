package main

import (
	"fmt"

	"github.com/mertcandav/dinogo/input"
)

func main() {
	fmt.Print("Input: ")
	r, _ := input.ReadRune()
	fmt.Println(string(r))
}
