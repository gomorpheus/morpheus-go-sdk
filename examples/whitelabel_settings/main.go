package main

import (
	"fmt"
	"log"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func main() {
	client := morpheus.NewClient("https://yourmorpheus.com")
	client.SetUsernameAndPassword("username", "password")
	resp, err := client.Login()
	if err != nil {
		fmt.Println("LOGIN ERROR: ", err)
	}
	fmt.Println("LOGIN RESPONSE:", resp)

	// Get whitelabel settings
	req := &morpheus.Request{}
	whitelabelSettingsResponse, err := client.GetWhitelabelSettings(req)
	if err != nil {
		log.Fatal(err)
	}
	result := whitelabelSettingsResponse.Result.(*morpheus.GetWhitelabelSettingsResult)
	log.Println(result)
}
