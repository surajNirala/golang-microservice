package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Customer struct {
	Name string `json:"full_name" xml:"full_name"`
	City string `json:"city" xml:"city"`
	Zip  int    `json:"zip" xml:"zip"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World Response.")
}

func customers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Suraj", "Gurgaon", 202002},
		{"RIva", "Delhi", 212002},
	}
	if r.Header.Get("Content-type") == "application/xml" {
		w.Header().Add("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else if r.Header.Get("Content-type") == "application/json" {
		w.Header().Add("Content-type", "application/json")
		response := ApiResponse{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    customers,
		}
		json.NewEncoder(w).Encode(response)
	} else {
		res := []string{"message", "Data Not Found"}
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
