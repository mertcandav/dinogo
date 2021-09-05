# dinogo
Dinogo is an CLI framework for build shell, terminal and command line applications in Go.

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mertcandav/dinogo/blob/main/LICENSE)

## Table of Contents
<div class="toc">
  <ul>
    <li><a href="#features">Features</a></li>
    <li><a href="#example_shell_application">Example Shell Application</a></li>
    <li><a href="#packages">Packages</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ul>
</div>

<h2 id="features">Features</h2>

+ Cross Platform
+ Fast and efficient
+ Keyboard API
+ Enable/Disable Windows Virtual Terminal Processing
+ Input infrastructure (custom actions, key combos...)
+ Input profiles (classic, password...)
+ Drawing tools
+ Shell infrastructure
+ Default actions (help, clear, exit...) for your shell
+ Command history infrastructure
+ Auto command complete support for your shell
+ ANSI escape codes and parser
+ CLI functions (move left, move right, clear screen...)
+ By default supports cursor move on input via command line

<h2 id="example_shell_application">Example Shell Application</h2>

```go
package main

import (
	"github.com/mertcandav/dinogo/shell"
	"github.com/mertcandav/dinogo/shell/actions"
)

func about(_ shell.CommandActionInfo) {
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
	sh.Loop()
}
```

<h2 id="packages">Packages</h2>

+ ``ansiescape/`` ANSI escape code constants and parser.
+ ``drawing/`` Drawing tools.
+ ``input/`` Input infrastructure for command line inputs.
+ ``keyboard/`` Keyboard API.
+ ``message/`` Message templates for print.
+ ``shell/`` Shell infrastructure.
+ ``terminal/`` CLI functions and tools.

<h2 id="contributing">Contributing</h2>
Thanks for you want contributing to Dinogo!

The Dinogo project use issues for only bug reports and proposals. <br>
To contribute, please read the contribution guidelines from [here](https://github.com/mertcandav/dinogo/blob/main/CONTRIBUTING.md).

All contributions to Dinogo, no matter how small or large, are welcome. <br>
From a simple typo correction to a contribution to the code, all contributions are welcome and appreciated.

<h2 id="license">License</h2>

Dinogo is distributed under the terms of the MIT license. <br>
[See license details.](https://github.com/mertcandav/dinogo/blob/main/LICENSE)
