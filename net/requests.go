package net

import (
	"encoding/json"
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

type IntensityIntervalRequest struct {
	Endpoint string
	Response IntensityByIntervalResponse
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity-stats-from-to
type IntensityByIntervalResponse = e.IntensityWithDate

func NewIntensityIntervalRequest(endpoint string) IntensityIntervalRequest {
	return IntensityIntervalRequest{
		Endpoint: endpoint,
		Response: IntensityByIntervalResponse{},
	}
}

func (r *IntensityIntervalRequest) GetEndpoint(args []string, flags map[string]string) {
	if len(flags) > 0 {
		r.Endpoint = fmt.Sprintf("%v/stats/%v/%v", r.Endpoint, flags["start-date"], flags["end-date"])
	}
}

func (r *IntensityIntervalRequest) Get() ([]byte, error) {
	res, err := DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return res, nil
}

func (r *IntensityIntervalRequest) Validate(response []byte) bool {
	return ValidateResponse("statistics", response)
}

func (r *IntensityIntervalRequest) UnMarshal(response []byte) error {
	err := json.Unmarshal(response, &r.Response)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-regionid-regionid
type IntensityByRegionIdResponse = e.IntensityWithDateAndRegionWithGenerationAndIntensity

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-postcode-postcode
type IntensityByRegionPostCodeResponse = e.IntensityWithDateAndRegionWithGenerationAndIntensity

type IntensityByDatetimeAndRegionResponse = e.IntensityWithDateAndRegionWithGenerationAndIntensity

// https://carbon-intensity.github.io/api-definitions/?shell#get-generation
type GenerationMixRecentResponse struct {
	Data []struct {
		e.DateTime
		Generationmix []e.Generationmix `json:"generationmix"`
	} `json:"data"`
}

type GenerationMixRecentRequest struct {
	Endpoint string
	Response GenerationMixRecentResponse
}

func NewGenerationMixRequest(endpoint string) GenerationMixRecentRequest {
	return GenerationMixRecentRequest{
		Endpoint: endpoint,
		Response: GenerationMixRecentResponse{},
	}
}

func (r *GenerationMixRecentRequest) GetEndpoint(args []string, flags map[string]string) {
	if len(flags) > 0 {
		if flags["past"] != "false" {
			r.Endpoint = fmt.Sprintf("%v/%v/%v", r.Endpoint, flags["start-date"], "pt24h")
		} else if flags["start-date"] != "" {
			r.Endpoint = fmt.Sprintf("%v/%v/%v", r.Endpoint, flags["start-date"], flags["end-date"])
		}
	}
}

func (r *GenerationMixRecentRequest) Get() ([]byte, error) {
	res, err := DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return res, nil
}

func (r *GenerationMixRecentRequest) Validate(response []byte) bool {
	return ValidateResponse("generation", response)
}

func (r *GenerationMixRecentRequest) UnMarshal(response []byte) error {
	err := json.Unmarshal(response, &r.Response)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
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
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%w", errors.New(string(body)))
	}

	return body, nil
}
