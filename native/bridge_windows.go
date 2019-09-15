// +build windows

package native

import (
	"sync"

	"golang.org/x/sys/windows"
)

type windowsProc struct {
	sync.Mutex
	*windows.Proc
}

func newWindowsProc(dll *windows.DLL, procName string) windowsProc {
	return windowsProc{
		Proc: dll.MustFindProc(procName),
	}
}

// TODO - combine architecture-specific implementations?
