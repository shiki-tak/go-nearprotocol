// +build linux,!muslc darwin

package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lgo_nearprotocol
import "C"
