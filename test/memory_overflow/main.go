package main

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func main(){
	loader := openapi3.NewSwaggerLoader()
	loader.IsExternalRefsAllowed = true
	_, err := loader.LoadSwaggerFromFile("spec0.yaml")
	if err != nil {
		panic(err)
	}
}
