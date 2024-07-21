package services

import (
	e "carbon-intensity/entities"
	req "carbon-intensity/net"
	"encoding/json"
	"fmt"
)

type IntensityAllRegionsRequest struct {
	Schema   string
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
		Schema:   "regional",
		Endpoint: endpoint,
		Response: IntensityByAllRegionsResponse{},
	}
}

func (r *IntensityAllRegionsRequest) GetEndpoint() {
}

func (r *IntensityAllRegionsRequest) Get() ([]byte, error) {
	res, err := req.DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	valid := r.Validate(&res)
	if valid {
		err = r.UnMarshal(&res)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, err
}

func (r *IntensityAllRegionsRequest) Validate(response *[]byte) bool {
	return req.ValidateResponse(r.Schema, *response)
}

func (r *IntensityAllRegionsRequest) UnMarshal(response *[]byte) error {
	err := json.Unmarshal(*response, &r.Response)
	if err != nil {
		fmt.Println("Error:", &err)
		return err
	}
	return nil
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-postcode-postcode
type IntensityByRegionPostCodeResponse = e.IntensityWithDateAndRegionWithGenerationAndIntensity

type IntensityRegionsPostcodeRequest struct {
	Scheme   string
	Endpoint string
	Response IntensityByRegionPostCodeResponse
}

func NewIntensityRegionsPostcodeRequest(endpoint string) IntensityRegionsPostcodeRequest {
	return IntensityRegionsPostcodeRequest{
		Scheme:   "regional-postcode",
		Endpoint: endpoint,
		Response: IntensityByRegionPostCodeResponse{},
	}
}

func (r *IntensityRegionsPostcodeRequest) GetEndpoint(postcode *string) {
	r.Endpoint = fmt.Sprintf("%v/postcode/%v", &r.Endpoint, &postcode)
}

func (r *IntensityRegionsPostcodeRequest) Get() (*[]byte, error) {
	res, err := req.DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	valid := r.Validate(&res)
	if valid {
		err = r.UnMarshal(&res)
		if err != nil {
			return nil, err
		}
		return &res, nil
	}
	return nil, err
}

func (r *IntensityRegionsPostcodeRequest) Validate(response *[]byte) bool {
	return req.ValidateResponse(r.Scheme, *response)
}

func (r *IntensityRegionsPostcodeRequest) UnMarshal(response *[]byte) error {
	err := json.Unmarshal(*response, &r.Response)
	if err != nil {
		fmt.Println("Error:", &err)
		return err
	}
	return nil
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-regionid-regionid
type IntensityByRegionIdResponse = e.IntensityWithDateAndRegionWithGenerationAndIntensity
