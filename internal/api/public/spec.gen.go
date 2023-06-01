// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package public

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"gg1PAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
