package gpost

// #cgo CFLAGS: -I${SRCDIR}/include
// #cgo LDFLAGS: -l${SRCDIR}/lib/libgpu
// #include "libgpu.h"
import "C"

func Echo(message string) string {
	i := C.stats()
	println(i)
	return message
}