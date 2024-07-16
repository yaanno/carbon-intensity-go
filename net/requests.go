package net

import (
	"encoding/json"
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

type IntensityRecentRequest struct {
	Endpoint string
	Response IntensityRecentResponse
}

func NewIntensityRecentRequest(endpoint string) IntensityRecentRequest {
	return IntensityRecentRequest{
		Endpoint: endpoint,
		Response: IntensityRecentResponse{},
	}
}

func (r *IntensityRecentRequest) GetEndpoint() string {
	return ""
}

func (r *IntensityRecentRequest) Get() ([]byte, error) {
	res, err := DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return res, nil
}

func (r *IntensityRecentRequest) Validate(response []byte) bool {
	return ValidateResponse(r.Endpoint, response)
}

func (r *IntensityRecentRequest) UnMarshal(response []byte) error {
	err := json.Unmarshal(response, &r.Response)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}

type IntensityAllRegionsRequest struct {
	Endpoint string
	Response IntensityByAllRegionsResponse
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional
type IntensityByAllRegionsResponse struct {
	Data []struct {
		e.DateTime
		Regions []e.RegionWithGenerationAndIntensity `json:"regions"`
	} `json:"data"`
}

func NewIntensityAllRegionsRequest(endpoint string) IntensityAllRegionsRequest {
	return IntensityAllRegionsRequest{
		Endpoint: endpoint,
		Response: IntensityByAllRegionsResponse{},
	}
}

func (r *IntensityAllRegionsRequest) GetEndpoint(args []string, flags map[string]string) {
	// if len(args) > 0 {
	// 	r.Endpoint = fmt.Sprintf("%v/%v", r.Endpoint, args[0])
	// }
}

func (r *IntensityAllRegionsRequest) Get() ([]byte, error) {
	res, err := DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return res, nil
}

func (r *IntensityAllRegionsRequest) Validate(response []byte) bool {
	return ValidateResponse("regional", response)
}

func (r *IntensityAllRegionsRequest) UnMarshal(response []byte) error {
	err := json.Unmarshal(response, &r.Response)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
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

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity-date
type IntensityTodayResponse = e.IntensityWithDate

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity-stats-from-to
type IntensityByIntervalResponse = e.IntensityWithDate

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

	return body, nil
}
