package main

import (
	"fmt"

	"github.com/mertcandav/dinogo/shell/history"
)

func main() {
	h := history.History{Duplicate: true}
	fmt.Println(h.Get())
}
