// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package terminal

// Position represents CLI cursor position.
type Position struct {
	Line   int
	Column int
}

// Set cursor position by values.
func (p *Position) Set() { SetPosition(p.Line, p.Column) }
