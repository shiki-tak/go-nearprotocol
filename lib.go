package nearprotocol

import api "github.com/shiki-tak/go-nearprotocol/api"

type Wasmer struct {
}

func NewWasmer() (*Wasmer, error) {
	return &Wasmer{}, nil
}

// Greet
func (w *Wasmer) Greet(name string) {
	api.Greet(name)
}

// DeployContract

// FunctionCall

// FunctionCallByStandalone
// 	near.RunWithStandalone(context, contextFile, input, methodName, state, promiseResults, config, configFile, wasmFile, vmKind, profileGas)

func (w *Wasmer) RunWithStandalone(context, contextFile, input, methodName,
	state, config, configFile, wasmFile, vmKind string) {
	api.RunStandalone(context, contextFile, input, methodName,
		state, config, configFile, wasmFile, vmKind)
}
