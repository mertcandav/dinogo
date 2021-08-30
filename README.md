# dinogo
Dinogo is an CLI framework for build terminal and shell applications in Go.

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mertcandav/dinogo/blob/main/LICENSE)

<h2 id="features">Features</h2>

+ Cross Platform
+ Fast and efficient
+ Keyboard API
+ Enable/Disable Windows Virtual Terminal Processing
+ Default actions (help, clear, exit...) for your shell
+ Shell infrastructure
+ Command history infrastructure
+ Input infrastructure (custom actions, key combos...)
+ ANSI Escape Code processor and parser
+ CLI functions (move left, move right, clear screen...)
+ By default supports cursor move on input via command line

<h2 id="example">Example</h2>

```go
package main

import (
	"github.com/mertcandav/dinogo/message"
	"github.com/mertcandav/dinogo/shell"
	"github.com/mertcandav/dinogo/shell/actions"
)

func about(shell.CommandActionInfo) {
	println("This shell is example for Dinogo framework.")
}

func main() {
	sh := shell.Init()
	sh.Commands = []shell.Command{
		{Name: "clear", Help: "Clear entire screen.", Action: actions.Clear},
		{Name: "about", Help: "Show about this shell.", Action: about},
		{Name: "history", Help: "List command history.", Action: actions.History},
		{Name: "help", Help: "Show help.", Action: actions.Help},
		{Name: "exit", Help: "Exit shell.", Action: actions.Exit},
	}
	for {
		err := sh.Prompt("Example Shell")
		if err != nil {
			message.Errorln(err)
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
