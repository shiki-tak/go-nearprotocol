package nearprotocol

type Wasmer struct {
}

func NewWasmer() (*Wasmer, error) {
	return &Wasmer{}, nil
}

// Greet
func (w *Wasmer) Greet(name string) string {
	return "Hello, " + name
}

// DeployContract

// FunctionCall
