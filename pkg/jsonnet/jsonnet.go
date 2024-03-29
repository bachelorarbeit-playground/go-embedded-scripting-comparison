package jsonnet

import (
	"fmt"

	jsonnet "github.com/google/go-jsonnet"
)

const initScript = "local stdin=std.extVar(\"stdin\");"

// Transformer is an object to use jsonnet
type Transformer struct {
	vm *jsonnet.VM
}

func WithJsonnet(script []byte, input []byte) (string, error) {
	// exec jsonnet
	vm := MakeTransformer()
	scriptstr := string(script)
	inputstr := string(input)
	output, err := vm.Transform(&scriptstr, &inputstr)
	if err != nil {
		return "", fmt.Errorf("Could not transform input using jsonnet script: %s", err.Error())
	}
	return *output, nil
}

// MakeTransformer creates a jsonnet virtual machine instance specific for our use case
func MakeTransformer() *Transformer {
	t := Transformer{}

	t.vm = jsonnet.MakeVM()

	return &t
}

// Transform uses a jsonnet script to transform stdin json data to the resulting json data
func (t *Transformer) Transform(script *string, input *string) (*string, error) {
	t.vm.ExtCode("stdin", *input)
	out, err := t.vm.EvaluateAnonymousSnippet("Error", initScript+*script)

	if err != nil {
		return nil, err
	}

	return &out, nil
}
