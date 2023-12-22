package app

import "net/http"

func Start() {

	/* http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", customers)
	http.ListenAndServe("localhost:4001", nil) */

	//CUstomer ROutes

	mux := http.NewServeMux()

	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", customers)
	http.ListenAndServe(":4001", mux)
}
