package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Listening on https://localhost:8443")

	// Check if Certificate Exists. if not new one is generated.

	if _, err := os.Stat("certs/server.crt"); os.IsNotExist(err) {
		if err := generateCert(); err != nil {
			panic(err)
		}
	}

	// Open The Http and Server TLS for encryption.

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
