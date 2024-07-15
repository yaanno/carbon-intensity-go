package net

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	e "carbon-intensity/entities"

	"github.com/xeipuuv/gojsonschema"
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

func DoRequest(endpoint string, flags map[string]any) {
	api := api + endpoint
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	refFile := fmt.Sprintf("file://./scheme/%v.json", endpoint)

	schema := gojsonschema.NewReferenceLoader(refFile)
	doc := gojsonschema.NewStringLoader(string(body))
	result, err := gojsonschema.Validate(schema, doc)
	if err != nil {
		fmt.Println("Validator error: ", err)
		return
	}
	if !result.Valid() {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		return
	}

	recent := IntensityByAllRegionsResponse{}
	err = json.Unmarshal(body, &recent)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response Body:", recent)
}
