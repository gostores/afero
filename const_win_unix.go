// +build !darwin
// +build !openbsd
// +build !freebsd
// +build !dragonfly
// +build !netbsd

package fsintra

import (
	"syscall"
)

const BADFD = syscall.EBADFD
