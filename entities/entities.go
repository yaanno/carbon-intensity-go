package entities

// Entities

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GenerationAndIntensity struct {
	Generationmix []Generationmix `json:"generationmix"`
	Intensity     Intensity
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
type IntensityWithDate struct {
	Data []struct {
		DateTime
		Intensity `json:"intensity"`
	} `json:"data"`
}
type IntensityWithDateAndRegionWithGenerationAndIntensity struct {
	Data []struct {
		DateTime
		Intensity `json:"intensity"`
	} `json:"data"`
}

// https://carbon-intensity.github.io/api-definitions/?shell#factors
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

// https://carbon-intensity.github.io/api-definitions/?shell#intensity-1
type Intensity struct {
	Forecast int    `json:"forecast"`
	Actual   int    `json:"actual"`
	Index    string `json:"index"`
	Max      int    `json:"max"`
	Average  int    `json:"average"`
	Min      int    `json:"min"`
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
