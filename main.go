package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Response types

type ErrorResponse struct {
	ResponseError `json:"error"`
}

type IntensityFactorsResponse struct {
	Data []Factor `json:"data"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity
type IntensityRecentResponse struct {
	Data []struct {
		DateTime
		Intensity `json:"intensity"`
	} `json:"data"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity-date
type IntensityTodayResponse struct {
	Data []struct {
		DateTime
		Intensity `json:"intensity"`
	} `json:"data"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-intensity-stats-from-to
type IntensityByIntervalResponse struct {
	Data []struct {
		DateTime
		Intensity `json:"intensity"`
	} `json:"data"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional
type IntensityByAllRegionsResponse struct {
	DateTime
	Regions []RegionWithGenerationAndIntensity `json:"regions"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-england
type IntensityByMainRegionResponse struct {
	Data []struct {
		Region
		DateTime
		Data []struct {
			RegionWithGenerationAndIntensity
		} `json:"data"`
	} `json:"data"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-regionid-regionid
type IntensityByRegionIdResponse struct {
	Data []struct {
		Region
		Data []struct {
			DateTime
			RegionWithGenerationAndIntensity
		} `json:"data"`
	} `json:"data"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-regional-postcode-postcode
type IntensityByRegionPostCodeResponse struct {
	Data []struct {
		Region
		Data []struct {
			DateTime
			RegionWithGenerationAndIntensity
		} `json:"data"`
	} `json:"data"`
}

type IntensityByDatetimeAndRegionResponse struct {
	Data []struct {
		Region
		Data []struct {
			DateTime
			RegionWithGenerationAndIntensity
		} `json:"data"`
	} `json:"data"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#get-generation
type GenetrationMixRecentResponse struct {
	Data []struct {
		DateTime
		Generationmix []Generationmix `json:"generationmix"`
	} `json:"data"`
}

// Entities

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type RegionWithGenerationAndIntensity struct {
	Region
	Generationmix []Generationmix `json:"generationmix"`
	Intensity     Intensity
}
type DateTime struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type Factor struct {
	Biomass          int `json:"Biomass"`
	Coal             int `json:"Coal"`
	DutchImports     int `json:"Dutch Imports"`
	FrenchImports    int `json:"French Imports"`
	GasCombinedCycle int `json:"Gas (Combined Cycle)"`
	GasOpenCycle     int `json:"Gas (Open Cycle)"`
	Hydro            int `json:"Hydro"`
	IrishImports     int `json:"Irish Imports"`
	Nuclear          int `json:"Nuclear"`
	Oil              int `json:"Oil"`
	Other            int `json:"Other"`
	PumpedStorage    int `json:"Pumped Storage"`
	Solar            int `json:"Solar"`
	Wind             int `json:"Wind"`
}
type Intensity struct {
	Forecast int    `json:"forecast"`
	Max      int    `json:"max"`
	Average  int    `json:"average"`
	Min      int    `json:"min"`
	Index    string `json:"index"`
}
type Generationmix struct {
	Fuel string  `json:"fuel"`
	Perc float64 `json:"perc"`
}
type Region struct {
	Regionid  int    `json:"regionid"`
	Dnoregion string `json:"dnoregion"`
	Shortname string `json:"shortname"`
	Postcode  string `json:"postcode"`
}

const api = "https://api.carbonintensity.org.uk"

func main() {
	api := api + "/intensity"
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	fmt.Println(response.StatusCode)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var recent IntensityRecentResponse

	err = json.Unmarshal(body, &recent)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response Body:", recent.Data)
}
