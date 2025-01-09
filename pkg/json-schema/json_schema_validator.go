package xjsonschema

import (
	"fmt"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

type JsonSchemaValidator struct {
	basePath string
}

func NewJsonSchemaValidator(basePath string) *JsonSchemaValidator {
	return &JsonSchemaValidator{
		basePath: basePath,
	}
}

func (jsv *JsonSchemaValidator) Validate(dataToValidate []byte, jsonSchemaLocation string) (*gojsonschema.Result, error) {
	absPath, _ := filepath.Abs(fmt.Sprintf("%s%s", jsv.basePath, jsonSchemaLocation))
	schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", absPath))
	documentLoader := gojsonschema.NewBytesLoader(dataToValidate)
	return gojsonschema.Validate(schemaLoader, documentLoader)
}
