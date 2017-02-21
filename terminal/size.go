package terminal

import (
	"syscall"
	"unsafe"
)

type Size struct {
	Row, Col, Xpixel, Ypixel uint16
}

var (
	ws Size
)

func init() {
	Refresh()
}

func Refresh() {
	syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(0), uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)))
}

func Width() uint16 {
	return ws.Col
}

func Height() uint16 {
	return ws.Row
}
