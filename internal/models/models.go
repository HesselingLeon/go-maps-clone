package models

type CoordinateInput struct {
	Zoom   float64  `path:"z"`
	Column float64  `path:"x"`
	Row    float64  `path:"y"`
	_      struct{} `query:"_" cookie:"_" additionalProperties:"false"`
}

type CoordinateOutput struct {
	Lon_deg_nw float64 `json:"latitude_nw"`
	Lat_deg_nw float64 `json:"longitude_nw"`
	Lon_deg_se float64 `json:"latitude_se"`
	Lat_deg_se float64 `json:"longitude_se"`
}
