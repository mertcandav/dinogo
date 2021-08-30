//go:build !windows && !linux
// +build !windows,!linux

// Copyright 2021.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package keyboard

import (
	"golang.org/x/sys/unix"
)

const (
	ioctl_GETATTR = unix.TIOCGETA
	ioctl_SETATTR = unix.TIOCSETA
)
