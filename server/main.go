package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Listening on https://localhost:8443")

	err := http.ListenAndServeTLS(
		":8443",
		"certs/server.crt",
		"certs/server.key",
		nil,
	)

	if err != nil {
		panic(err)
	}
}
