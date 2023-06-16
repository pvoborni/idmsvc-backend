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
	"DgXH5YTFRr++Wf8jAAD//wXIIJ7XVgAA",
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
