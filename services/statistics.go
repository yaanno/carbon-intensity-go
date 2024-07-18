package services

import (
	e "carbon-intensity/entities"
	req "carbon-intensity/net"
	"encoding/json"
	"fmt"
)

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity-stats-from-to
type IntensityByIntervalResponse = e.IntensityWithDate

type IntensityIntervalRequest struct {
	Endpoint string
	Response IntensityByIntervalResponse
}

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
	res, err := req.DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return res, nil
}

func (r *IntensityIntervalRequest) Validate(response []byte) bool {
	return req.ValidateResponse("statistics", response)
}

func (r *IntensityIntervalRequest) UnMarshal(response []byte) error {
	err := json.Unmarshal(response, &r.Response)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
