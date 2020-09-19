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
