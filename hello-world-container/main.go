package main

import (
	"log"
	"net/http"
	"os"
)

var Version = "4.0.1"

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		hostname, err := os.Hostname()

		if err != nil {
			log.Fatal(err)
		}

		rw.WriteHeader(200)
		rw.Write([]byte("Welcome to Kubernetes!\n"))
		rw.Write([]byte("Hostname: " + hostname + "\n"))
		rw.Write([]byte("Version: " + Version))
	})

	http.ListenAndServe(":3000", nil)
}
