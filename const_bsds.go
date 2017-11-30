// +build darwin openbsd freebsd netbsd dragonfly

package fsintra

import (
	"syscall"
)

const BADFD = syscall.EBADF
