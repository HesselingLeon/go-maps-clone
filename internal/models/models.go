package models

type CoordinateInput struct {
	Zoom   string   `path:"z"`
	Column string   `path:"x"`
	Row    string   `path:"y"`
	_      struct{} `query:"_" cookie:"_" additionalProperties:"false"`
}

type CoordinateOutput struct {
	Min       string `json:"min"`
	Max       string `json:"max"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}
