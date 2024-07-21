package services

import (
	e "carbon-intensity/entities"
	req "carbon-intensity/net"
	"encoding/json"
	"fmt"
)

// https://carbon-intensity.github.io/api-definitions/?shell#get-generation
type GenerationMixRecentResponse struct {
	Data struct {
		Generationmix []e.Generationmix `json:"generationmix"`
	} `json:"data"`
}

type GenerationMixRecentRequest struct {
	Endpoint string
	Response GenerationMixRecentResponse
	Schema   string
}

func NewGenerationMixRecentRequest(endpoint string) GenerationMixRecentRequest {
	return GenerationMixRecentRequest{
		Schema:   "generation-current",
		Endpoint: endpoint,
		Response: GenerationMixRecentResponse{},
	}
}

func (r *GenerationMixRecentRequest) GetEndpoint() {
}

func (r *GenerationMixRecentRequest) Get() ([]byte, error) {
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

func (r *GenerationMixRecentRequest) Validate(response *[]byte) bool {
	return req.ValidateResponse(r.Schema, *response)
}

func (r *GenerationMixRecentRequest) UnMarshal(response *[]byte) error {
	err := json.Unmarshal(*response, &r.Response)
	if err != nil {
		fmt.Println("Error:", &err)
		return err
	}
	return nil
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-generation
type GenerationMixIntervalResponse struct {
	Data []struct {
		e.DateTime
		Generationmix []e.Generationmix `json:"generationmix"`
	} `json:"data"`
}

type GenerationMixIntervalRequest struct {
	Endpoint string
	Response GenerationMixIntervalResponse
	Schema   string
}

func NewGenerationMixRequest(endpoint string) GenerationMixIntervalRequest {
	return GenerationMixIntervalRequest{
		Endpoint: endpoint,
		Response: GenerationMixIntervalResponse{},
		Schema:   "generation",
	}
}

func (r *GenerationMixIntervalRequest) GetEndpoint(flags map[string]interface{}) {
	if len(flags) > 0 {
		if flags["past"] == true {
			r.Endpoint = fmt.Sprintf("%v/%v/%v", r.Endpoint, flags["start-date"], "pt24h")
		} else if flags["start-date"] != "" {
			r.Endpoint = fmt.Sprintf("%v/%v/%v", r.Endpoint, flags["start-date"], flags["end-date"])
		}
	}
}

func (r *GenerationMixIntervalRequest) Get() ([]byte, error) {
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

func (r *GenerationMixIntervalRequest) Validate(response *[]byte) bool {
	return req.ValidateResponse(r.Schema, *response)
}

func (r *GenerationMixIntervalRequest) UnMarshal(response *[]byte) error {
	err := json.Unmarshal(*response, &r.Response)
	if err != nil {
		fmt.Println("Error:", &err)
		return err
	}
	return nil
}
