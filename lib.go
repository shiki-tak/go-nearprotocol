package nearprotocol

import (
	api "github.com/shiki-tak/go-nearprotocol/api"
)

type VM struct {
	cache api.Cache
}

func NewVM() (*VM, error) {
	cache, err := api.InitVM()
	if err != nil {
		return nil, err
	}

	return &VM{cache: cache}, nil
}

// Greet
func (vm *VM) Greet(name string) {
	api.Greet(name)
}

// DeployContract

// FunctionCall

// FunctionCallByStandalone
func (vm *VM) RunWithStandalone(context, contextFile, input, methodName,
	state, config, configFile, wasmFile, vmKind string) {
	api.RunStandalone(context, contextFile, input, methodName, state, config, configFile, wasmFile, vmKind)
}
