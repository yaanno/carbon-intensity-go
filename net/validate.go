package net

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateResponse(endpoint string, body []byte) bool {
	file := fmt.Sprintf("file://./scheme/%v.json", endpoint)
	schema := gojsonschema.NewReferenceLoader(file)
	// error400Schema := gojsonschema.NewReferenceLoader("file://./scheme/400.json")
	// error500Schema := gojsonschema.NewReferenceLoader("file://./scheme/500.json")
	doc := gojsonschema.NewStringLoader(string(body))
	// check for response errors
	// result, err := gojsonschema.Validate(error400Schema, doc)
	// if err != nil {
	// 	fmt.Println("Validator error: ", err)
	// 	return false
	// }
	// if !result.Valid() {
	// 	fmt.Println("Api responded with error: 400 Bad Request")
	// 	fmt.Println(string(body))
	// 	return false
	// }
	// result, err = gojsonschema.Validate(error500Schema, doc)
	// if err != nil {
	// 	fmt.Println("Validator error: ", err)
	// 	return false
	// }
	// if result.Valid() {
	// 	fmt.Println("Api responded with error: 500 Internal Server Error")
	// 	// fmt.Println(string(body))
	// 	return false
	// }
	// validate document
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
