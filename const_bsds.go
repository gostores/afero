// +build darwin openbsd freebsd netbsd dragonfly

package afero

import (
	"syscall"
)

const BADFD = syscall.EBADF
