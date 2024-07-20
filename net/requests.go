package net

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	e "carbon-intensity/entities"
)

// Response types

type ErrorResponse struct {
	e.ResponseError `json:"error"`
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

type IntensityMainRegionsRequest struct {
	Endpoint string
	Response IntensityByMainRegionResponse
}

func (r *IntensityMainRegionsRequest) GetEndpoint(args []string, flags map[string]string) {
	if len(args) > 0 {
		r.Endpoint = fmt.Sprintf("%v/%v", r.Endpoint, args[0])
	}
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-regionid-regionid
type IntensityByRegionIdResponse = e.IntensityWithDateAndRegionWithGenerationAndIntensity

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-postcode-postcode
type IntensityByRegionPostCodeResponse = e.IntensityWithDateAndRegionWithGenerationAndIntensity

type IntensityByDatetimeAndRegionResponse = e.IntensityWithDateAndRegionWithGenerationAndIntensity

const api = "https://api.carbonintensity.org.uk/"

var client = &http.Client{}

func DoRequest(endpoint string) ([]byte, error) {
	api := api + endpoint
	fmt.Println(api)
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
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%w", errors.New(string(body)))
	}

	return body, nil
}
