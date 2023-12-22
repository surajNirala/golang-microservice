package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	/* http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", customers)
	http.ListenAndServe("localhost:4001", nil) */

	//CUstomer ROutes

	// router := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", customers)
	// router.HandleFunc("/customers/{customer_id}", getCustomer)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer) // accept only numeric

	http.ListenAndServe(":4001", router)
}
