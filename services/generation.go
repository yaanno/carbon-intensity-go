package services

import (
	e "carbon-intensity/entities"
	req "carbon-intensity/net"
	"encoding/json"
	"fmt"
)

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
	res, err := req.DoRequest(r.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return res, nil
}

func (r *GenerationMixRecentRequest) Validate(response []byte) bool {
	return req.ValidateResponse("generation", response)
}

func (r *GenerationMixRecentRequest) UnMarshal(response []byte) error {
	err := json.Unmarshal(response, &r.Response)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
