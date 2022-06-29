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

	// List log settings
	req := &morpheus.Request{}
	logSettingsResponse, err := client.GetLogSettings(req)
	if err != nil {
		log.Fatal(err)
	}
	result := logSettingsResponse.Result.(*morpheus.GetLogSettingsResult)
	log.Println(result.LogSettings)
}
