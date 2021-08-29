package main

import (
	"fmt"

	"github.com/mertcandav/dinogo/input"
)

func main() {
	i := input.Init()
	fmt.Print("Input: ")
	i.Get()
	fmt.Println(string(i.Runes))
}
