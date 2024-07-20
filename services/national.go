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

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity-date
type IntensityTodayResponse = e.IntensityWithDate

type IntensityTodayRequest struct {
	Schema   string
	Endpoint string
	Response IntensityRecentResponse
}

func NewIntensityTodayRequest(endpoint string) IntensityTodayRequest {
	return IntensityTodayRequest{
		Schema:   "intensity-today",
		Endpoint: endpoint,
		Response: IntensityTodayResponse{},
	}
}

func (r *IntensityTodayRequest) GetEndpoint() {
	r.Endpoint = fmt.Sprintf("%v/date", r.Endpoint)
}

func (r *IntensityTodayRequest) Get() ([]byte, error) {
	res, err := req.DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return res, nil
}

func (r *IntensityTodayRequest) Validate(response []byte) bool {
	return req.ValidateResponse(r.Schema, response)
}

func (r *IntensityTodayRequest) UnMarshal(response []byte) error {
	err := json.Unmarshal(response, &r.Response)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}

type IntensityDateAndPeriodResponse = e.IntensityWithDate

type IntensityDateAndPeriodRequest struct {
	Schema   string
	Endpoint string
	Response IntensityDateAndPeriodResponse
}

func NewIntensityDateAndPeriodRequest(endpoint string) IntensityDateAndPeriodRequest {
	return IntensityDateAndPeriodRequest{
		Schema:   "intensity-date",
		Endpoint: endpoint,
		Response: IntensityDateAndPeriodResponse{},
	}
}

func (r *IntensityDateAndPeriodRequest) GetEndpoint(flags map[string]string) {
	if len(flags) > 0 {
		r.Endpoint = fmt.Sprintf("%v/date/%v", r.Endpoint, flags["date"])
		if flags["period"] != "" {
			r.Endpoint = fmt.Sprintf("%v/%v", r.Endpoint, flags["period"])
		}
	}
}

func (r *IntensityDateAndPeriodRequest) Get() ([]byte, error) {
	res, err := req.DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return res, nil
}

func (r *IntensityDateAndPeriodRequest) Validate(response []byte) bool {
	return req.ValidateResponse(r.Schema, response)
}

func (r *IntensityDateAndPeriodRequest) UnMarshal(response []byte) error {
	err := json.Unmarshal(response, &r.Response)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}

type IntensityPeriodResponse = e.IntensityWithDate

type IntensityPeriodRequest struct {
	Schema   string
	Endpoint string
	Response IntensityPeriodResponse
}

func NewIntensityPeriodRequest(endpoint string) IntensityPeriodRequest {
	return IntensityPeriodRequest{
		Schema:   "intensity-period",
		Endpoint: endpoint,
		Response: IntensityPeriodResponse{},
	}
}

func (r *IntensityPeriodRequest) GetEndpoint(flags map[string]string) {
	if len(flags) > 0 {
		r.Endpoint = fmt.Sprintf("%v/%v", r.Endpoint, flags["from"])

		if flags["to"] != "" {
			r.Endpoint = fmt.Sprintf("%v/%v", r.Endpoint, flags["to"])
		}
		if flags["past"] == "true" {
			r.Endpoint = fmt.Sprintf("%v/pt24", r.Endpoint)
		}
		if flags["future"] == "true" {
			r.Endpoint = fmt.Sprintf("%v/fw%v", r.Endpoint, flags["hours"])
		}
	}
}

func (r *IntensityPeriodRequest) Get() ([]byte, error) {
	res, err := req.DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return res, nil
}

func (r *IntensityPeriodRequest) Validate(response []byte) bool {
	return req.ValidateResponse(r.Schema, response)
}

func (r *IntensityPeriodRequest) UnMarshal(response []byte) error {
	err := json.Unmarshal(response, &r.Response)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
