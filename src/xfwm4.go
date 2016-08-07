package main

/*
#cgo pkg-config: x11 xext glib-2.0 gdk-2.0 atk libstartup-notification-1.0
#cgo pkg-config: libxfce4util-1.0 libxfconf-0 libxfce4kbd-private-2
#cgo LDFLAGS: -lm
#include <stdlib.h>
int real_main(int argc, char **argv);
*/
import "C"

import (
	"os"
	"unsafe"
)

const szptr = unsafe.Sizeof(uintptr(0))

func main() {
	argc := C.int(len(os.Args))

	argv := C.malloc(C.size_t(argc) * C.size_t(szptr))
	defer C.free(argv)

	for i := 0; i < len(os.Args); i++ {
		*(**C.char)(unsafe.Pointer(uintptr(argv) + uintptr(i)*szptr)) = C.CString(os.Args[i])
	}
	defer func() {
		for i := 0; i < len(os.Args); i++ {
			C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(uintptr(argv) + uintptr(i)*szptr))))
		}
	}()

	C.real_main(argc, argv)
}
