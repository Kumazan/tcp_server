package main

type respBody struct {
	Success string       `json:"success"`
	Result  resultStruct `json:"result"`
	Records records      `json:"records"`
}

type resultStruct struct {
	ResourceID string  `json:"resource_id"`
	Fields     []field `json:"fields"`
}

type field struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type records struct {
	Locations []locationCity `json:"locations"`
}

type locationCity struct {
	DatasetDescription string             `json:"datasetDescription"`
	LocationsName      string             `json:"locationsName"`
	DataID             string             `json:"dataid"`
	Locations          []locationDistrict `json:"location"`
}

type locationDistrict struct {
	LocationName    string           `json:"locationName"`
	Geocode         string           `json:"geocode"`
	Lat             string           `json:"lat"`
	Lon             string           `json:"lon"`
	WeatherElements []weatherElement `json:"weatherElement"`
}

type weatherElement struct {
	ElementName string     `json:"elementName"`
	Description string     `json:"description"`
	TimeData    []timeData `json:"time"`
}

type timeData struct {
	StartTime     string         `json:"startTime"`
	EndTime       string         `json:"endTime"`
	ElementValues []elementValue `json:"elementValue"`
}

type elementValue struct {
	Value    string `json:"value"`
	Measures string `json:"measures"`
}
