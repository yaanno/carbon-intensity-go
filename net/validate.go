package net

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateResponse(endpoint string, body []byte) bool {
	file := fmt.Sprintf("file://./scheme/%v.json", endpoint)
	fmt.Println(file)
	schema := gojsonschema.NewReferenceLoader(file)
	doc := gojsonschema.NewStringLoader(string(body))
	result, err := gojsonschema.Validate(schema, doc)
	if err != nil {
		fmt.Println("Validator error: ", err)
		return false
	}
	if !result.Valid() {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		return false
	}
	return true
}
