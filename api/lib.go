package api

// #include <stdlib.h>
// #include "bindings.h"
import "C"

import (
	"unsafe"
)

func Greet(name string) {
	cStr := C.CString(name)
	C.greet(cStr)

	defer C.free(unsafe.Pointer(cStr))
}
