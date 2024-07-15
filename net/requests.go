package net

import (
	"fmt"
	"io"
	"net/http"

	e "carbon-intensity/entities"
)

// Response types

type ErrorResponse struct {
	e.ResponseError `json:"error"`
}

type IntensityFactorsResponse struct {
	Data []e.Factor `json:"data"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity
type IntensityRecentResponse = e.IntensityWithDate

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity-date
type IntensityTodayResponse = e.IntensityWithDate

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity-stats-from-to
type IntensityByIntervalResponse = e.IntensityWithDate

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional
type IntensityByAllRegionsResponse struct {
	Data []struct {
		e.DateTime
		Regions []e.RegionWithGenerationAndIntensity `json:"regions"`
	} `json:"data"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-england
type IntensityByMainRegionResponse struct {
	Data []struct {
		e.Region
		Data []struct {
			e.DateTime
			e.GenerationAndIntensity
		} `json:"data"`
	} `json:"data"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-regionid-regionid
type IntensityByRegionIdResponse = e.IntensityWithDateAndRegionWithGenerationAndIntensity

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-postcode-postcode
type IntensityByRegionPostCodeResponse = e.IntensityWithDateAndRegionWithGenerationAndIntensity

type IntensityByDatetimeAndRegionResponse = e.IntensityWithDateAndRegionWithGenerationAndIntensity

// https://carbon-intensity.github.io/api-definitions/?shell#get-generation
type GenetrationMixRecentResponse struct {
	Data []struct {
		e.DateTime
		Generationmix []e.Generationmix `json:"generationmix"`
	} `json:"data"`
}

const api = "https://api.carbonintensity.org.uk/"

var client = &http.Client{}

func DoRequest(endpoint string) ([]byte, error) {
	api := api + endpoint
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return body, nil

	// isValid := validateResponse(endpoint, body)

	// if !isValid {
	// 	return
	// }

	// recent := getResponseTypeByEndpoint(endpoint)
	// var recent

	// recent := []byte{}
	// err = json.Unmarshal(body, &recent)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Response Body:", recent)
}

// func validateResponse(endpoint string, body []byte) bool {
// 	file := fmt.Sprintf("file://./scheme/%v.json", endpoint)
// 	schema := gojsonschema.NewReferenceLoader(file)
// 	doc := gojsonschema.NewStringLoader(string(body))
// 	result, err := gojsonschema.Validate(schema, doc)
// 	if err != nil {
// 		fmt.Println("Validator error: ", err)
// 		return false
// 	}
// 	if !result.Valid() {
// 		fmt.Printf("The document is not valid. see errors :\n")
// 		for _, desc := range result.Errors() {
// 			fmt.Printf("- %s\n", desc)
// 		}
// 		return false
// 	}
// 	return true
// }

func GetEndpoint(endpoint string, args []string, flags map[string]any) string {
	fmt.Println(args, flags)
	if len(args) > 0 {
		endpoint = fmt.Sprintf("%v/%v", endpoint, args[0])
		return endpoint
	}
	if flags["id"] != "" {
		endpoint = fmt.Sprintf("%v/regionid/%v", endpoint, flags["id"])
	}
	return endpoint
}
