package main

import (
	"flag"

	near "github.com/shiki-tak/go-nearprotocol"
)

func main() {
	var (
		standalone  = flag.Bool("standalone", false, "execute vm mode")
		context     = flag.String("context", "", "Specifies the execution context in JSON format, see `VMContext`.")
		contextFile = flag.String("context-file", "", "Reads the context from the file.")
		input       = flag.String("input", "", "Overrides input field of the context with the given string.")
		methodName  = flag.String("method-name", "", "The name of the method to call on the smart contract.")
		state       = flag.String("state", "", "Key-value state in JSON base64 format for the smart contract as HashMap.")
		config      = flag.String("config", "", "Specifies the economics and Wasm config in JSON format, see `Config`.")
		configFile  = flag.String("config-file", "", "Reads the config from the file.")
		wasmFile    = flag.String("wasm-file", "", "File path that contains the Wasm code to run.")
		vmKind      = flag.String("vm-kind", "", "Select VM kind to run.")
	)

	flag.Parse()

	near, err := near.NewWasmer()
	if err != nil {
		panic(err)
	}

	if *standalone {
		runStandalone(near, *context, *contextFile,
			*input, *methodName, *state, *config, *configFile, *wasmFile, *vmKind)

	}

	near.Greet("Go-NearProtocol")
}

func runStandalone(near *near.Wasmer, context, contextFile,
	input, methodName, state, config, configFile, wasmFile, vmKind string) {
	near.RunWithStandalone(context, contextFile, input, methodName, state, config, configFile, wasmFile, vmKind)
}
