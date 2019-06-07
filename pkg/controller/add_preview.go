package controller

import (
	"github.com/okteto/preview-operator/pkg/controller/preview"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, preview.Add)
}
