package models

type Customer struct {
	Name string `json:"full_name" xml:"full_name"`
	City string `json:"city" xml:"city"`
	Zip  int    `json:"zip" xml:"zip"`
}
