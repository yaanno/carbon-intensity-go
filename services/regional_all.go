package services

import (
	e "carbon-intensity/entities"
	req "carbon-intensity/net"
	"encoding/json"
	"fmt"
)

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
	res, err := req.DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return res, nil
}

func (r *IntensityAllRegionsRequest) Validate(response []byte) bool {
	return req.ValidateResponse("regional", response)
}

func (r *IntensityAllRegionsRequest) UnMarshal(response []byte) error {
	err := json.Unmarshal(response, &r.Response)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
