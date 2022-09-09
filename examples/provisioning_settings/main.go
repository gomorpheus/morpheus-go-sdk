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

	// Get provisioning settings
	req := &morpheus.Request{}
	provisioningSettingsResponse, err := client.GetProvisioningSettings(req)
	if err != nil {
		log.Fatal(err)
	}
	result := provisioningSettingsResponse.Result.(*morpheus.GetProvisinoingSettingsResult)
	log.Println(result)
}
