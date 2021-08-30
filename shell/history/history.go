// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package history

// History is command history.
type History struct {
	commands []string
	position int

	Duplicate bool // Let duplicates.
}

// Start sets position to begin of history.
func (h *History) Start() {
	h.position = 0
}

// End sets position to end of history.
func (h *History) End() {
	h.position = len(h.commands) - 1
	if h.position < 0 {
		h.position = 0
	}
}

// Next sets position to next history.
// Returns set state.
func (h *History) Next() bool {
	if h.position+1 >= len(h.commands) {
		return false
	}
	h.position++
	return true
}

// Next sets position to previous history.
// Returns set state.
func (h *History) Prev() bool {
	if h.position-1 < 0 {
		return false
	}
	h.position--
	return true
}

// Add adds new command to history.
// Returns add state.
func (h *History) Add(cmd string) bool {
	if !h.Duplicate && len(h.commands) > 0 {
		if h.commands[len(h.commands)-1] == cmd {
			return false
		}
	}
	h.commands = append(h.commands, cmd)
	return true
}

// Get returns command from history by position.
// Returns true if exist command, false if not.
func (h *History) Get() (string, bool) {
	if len(h.commands) == 0 {
		return "", false
	}
	return h.commands[h.position], true
}

// Position returns position.
func (h *History) Position() int {
	return h.position
}

// History returns mutable value of history list.
func (h *History) History() []string {
	return h.commands
}
