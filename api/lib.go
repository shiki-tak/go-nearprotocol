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

func RunStandalone(context, contextFile, input, methodName,
	state, config, configFile, wasmFile, vmKind string) {
	cContext := C.CString(context)
	cContextFile := C.CString(contextFile)
	cInput := C.CString(input)
	cMethodName := C.CString(methodName)
	cState := C.CString(state)
	cConfig := C.CString(config)
	cConfigFile := C.CString(configFile)
	cWasmFile := C.CString(wasmFile)
	cVmKind := C.CString(vmKind)

	defer C.free(unsafe.Pointer(cContext))
	defer C.free(unsafe.Pointer(cContextFile))
	defer C.free(unsafe.Pointer(cInput))
	defer C.free(unsafe.Pointer(cMethodName))
	defer C.free(unsafe.Pointer(cState))
	defer C.free(unsafe.Pointer(cConfig))
	defer C.free(unsafe.Pointer(cConfigFile))
	defer C.free(unsafe.Pointer(cWasmFile))
	defer C.free(unsafe.Pointer(cVmKind))

	C.run_with_standalone(cVmKind, cContext, cContextFile, cInput, cState, cMethodName, cConfig, cConfigFile, cWasmFile)
}
