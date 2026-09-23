package main

import (
	"fmt"
	"net/http"
	"wg-manager/internal/api"
	"wg-manager/internal/vpn"
)

func main() {

	// vpn service object initialization
	vpnApp := vpn.VPNService{}

	// multiplexer to register and handle communications
	mux := http.ServeMux{}

	// Creating a server just to be explicit, we will configure it later
	server := http.Server{
		Handler: &mux,
		Addr:    ":8080",
	}

	// handlers

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/profile", api.MakeProfileHandler(&vpnApp))

	fmt.Println("Starting a server on port 8080")
	server.ListenAndServe() //here I should put a port somewhere or in the struct

}

// basic handler check below

func hello(rw http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello")
	response := "hello"
	_, err := rw.Write([]byte(response))
	if err != nil {
		fmt.Printf("Something is up we have %v", err)
	}
}
