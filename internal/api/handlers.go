package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wg-manager/internal/vpn"
)

func MakeProfileHandler(vpn *vpn.VPNService) http.HandlerFunc {

	return func(rw http.ResponseWriter, req *http.Request) {
		profile := (*vpn).CreateProfile()
		profileJson, err := json.Marshal(profile)
		if err != nil {
			fmt.Printf("Danger danger! We have an error of type: %s", err)
			return
		}

		_, err = rw.Write(profileJson)
		if err != nil {
			fmt.Println("Failed to write a response")
			return
		}
	}
}
