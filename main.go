package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"wg-manager/internal/api"
	"wg-manager/internal/vpn"
)

func main() {
	apiKey := ""
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Failed to load .env; protected endpoints will remain unavailable: %v\n", err)
	} else {
		apiKey = os.Getenv("API_KEY")
		if apiKey == "" {
			fmt.Println("API_KEY is empty; protected endpoints will remain unavailable")
		}
	}

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
	mux.Handle("/profile", api.RequireAPIKey(apiKey, api.MakeProfileHandler(&vpnApp)))

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
