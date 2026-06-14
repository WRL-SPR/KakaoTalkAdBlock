//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const errorAlreadyExists syscall.Errno = 183

var (
	kernel32     = windows.NewLazySystemDLL("kernel32.dll")
	createMutexW = kernel32.NewProc("CreateMutexW")
)

// AcquireSingleInstance keeps a named mutex open for the lifetime of the app.
func AcquireSingleInstance(name string) (release func(), acquired bool, err error) {
	mutexName, err := windows.UTF16PtrFromString(`Local\` + name)
	if err != nil {
		return nil, false, err
	}

	handle, _, callErr := createMutexW.Call(
		0,
		1,
		uintptr(unsafe.Pointer(mutexName)),
	)
	if handle == 0 {
		return nil, false, callErr
	}
	if callErr == errorAlreadyExists {
		_ = windows.CloseHandle(windows.Handle(handle))
		return func() {}, false, nil
	}

	return func() {
		_ = windows.CloseHandle(windows.Handle(handle))
	}, true, nil
}
