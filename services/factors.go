package services

import (
	e "carbon-intensity/entities"
	req "carbon-intensity/net"
	"encoding/json"
	"fmt"
)

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity-stats-from-to
type FactorsResponse = e.Factors

type FactorsRequest struct {
	Schema   string
	Endpoint string
	Response FactorsResponse
}

func NewFactorsRequest(endpoint string) FactorsRequest {
	return FactorsRequest{
		Schema:   "factors",
		Endpoint: endpoint,
		Response: FactorsResponse{},
	}
}

func (r *FactorsRequest) GetEndpoint() {
	r.Endpoint = fmt.Sprintf("%v/factors", r.Endpoint)
}

func (r *FactorsRequest) Get() ([]byte, error) {
	res, err := req.DoRequest(r.Endpoint)
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

func (r *FactorsRequest) Validate(response *[]byte) bool {
	return req.ValidateResponse(r.Schema, *response)
}

func (r *FactorsRequest) UnMarshal(response *[]byte) error {
	err := json.Unmarshal(*response, &r.Response)
	if err != nil {
		fmt.Println("Error:", &err)
		return err
	}
	return nil
}
