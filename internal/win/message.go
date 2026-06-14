//go:build windows

package win

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	messageBoxOK       = 0x00000000
	messageBoxIconInfo = 0x00000040
	messageBoxIconStop = 0x00000010
)

var (
	user32      = windows.NewLazySystemDLL("user32.dll")
	messageBoxW = user32.NewProc("MessageBoxW")
)

func ShowInfo(title, message string) {
	showMessage(title, message, messageBoxOK|messageBoxIconInfo)
}

func ShowError(title, message string) {
	showMessage(title, message, messageBoxOK|messageBoxIconStop)
}

func showMessage(title, message string, flags uintptr) {
	titlePtr, titleErr := windows.UTF16PtrFromString(title)
	messagePtr, messageErr := windows.UTF16PtrFromString(message)
	if titleErr != nil || messageErr != nil {
		return
	}
	_, _, _ = messageBoxW.Call(
		0,
		uintptr(unsafe.Pointer(messagePtr)),
		uintptr(unsafe.Pointer(titlePtr)),
		flags,
	)
}
