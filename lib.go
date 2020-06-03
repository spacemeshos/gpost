package gpost
import "C"

// #cgo CFLAGS: -I${SRCDIR}/include
// #cgo LDFLAGS: -llibgpu
// #cgo darwin,amd64 LDFLAGS:-L${SRCDIR}/lib/darwin/amd64
// #include "libgpu.h"
import "C"

func Echo(message string) string {
	i := C.stats()
	println(i)
	return message
}