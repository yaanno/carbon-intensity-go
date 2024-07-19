package services

import (
	e "carbon-intensity/entities"
	req "carbon-intensity/net"
	"encoding/json"
	"fmt"
)

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
	return r.Endpoint
}

func (r *IntensityRecentRequest) Get() ([]byte, error) {
	res, err := req.DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return res, nil
}

func (r *IntensityRecentRequest) Validate(response []byte) bool {
	return req.ValidateResponse(r.Endpoint, response)
}

func (r *IntensityRecentRequest) UnMarshal(response []byte) error {
	err := json.Unmarshal(response, &r.Response)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
