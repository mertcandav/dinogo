# dinogo
Dinogo is an CLI framework for build terminal and shell applications in Go.

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mertcandav/dinogo/blob/main/LICENSE)

<h2 id="example">Example</h2>

```go
package main

import (
	"fmt"

	"github.com/mertcandav/dinogo/input"
	"github.com/mertcandav/dinogo/shell"
	"github.com/mertcandav/dinogo/shell/actions"
	"github.com/mertcandav/dinogo/shell/history"
)

func about(info shell.CommandActionInfo) {
	println("This shell is example for Dinogo framework.")
}

func main() {
	sh := shell.Shell{
		Sep:     " ",
		Prefix:  "$ ",
		Input:   input.Init(),
		History: &history.History{},
		Commands: []shell.Command{
			{Name: "help", Help: "Show help.", Action: actions.Help},
			{Name: "history", Help: "List command history.", Action: actions.History},
			{Name: "exit", Help: "Exit shell.", Action: actions.Exit},
			{Name: "about", Help: "Show about this shell.", Action: about},
		},
	}
	for {
		err := sh.Prompt("Example Shell")
		if err != nil {
			fmt.Println("error:", err)
		}
	}
}
```

<h2 id="goals">Contributing</h2>
Thanks for you want contributing to Dinogo!

The Dinogo project use issues for only bug reports and proposals. <br>
To contribute, please read the contribution guidelines from [here](https://github.com/mertcandav/dinogo/blob/main/CONTRIBUTING.md).

All contributions to Dinogo, no matter how small or large, are welcome. <br>
From a simple typo correction to a contribution to the code, all contributions are welcome and appreciated.


<h2 id="license">License</h2>

Fract and standard library is distributed under the terms of the MIT license. <br>
[See license details.](https://github.com/mertcandav/dinogo/blob/main/LICENSE)
