package main

type IndicatorLight struct {
	Label   string `json:"title"`
	Details string `json:"desc"`
	//0-gray/irrelevant, 1-blue/improving, 2-green/good, 3-yellow/danger, 4-red/crisis
	Color int `json:"color"`
}

type Alter struct {
	Label    string `json:"title"`
	Details  string `json:"desc"`
	Fronting bool   `json:"fronting"`
}
