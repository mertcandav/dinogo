package main

import (
	"fmt"

	"github.com/mertcandav/dinogo/input"
	"github.com/mertcandav/dinogo/shell/utils"
)

func main() {
	i := input.Init()
	fmt.Print("Input: ")
	i.Get()
	name, y := utils.SplitCommandName(string(i.Runes), " ")
	fmt.Println(name)
	fmt.Println("'" + y + "'")
}
