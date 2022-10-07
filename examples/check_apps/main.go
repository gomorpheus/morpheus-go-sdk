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

	// List check apps
	req := &morpheus.Request{}
	checkAppResp, err := client.ListCheckApps(req)
	if err != nil {
		log.Fatal(err)
	}
	result := checkAppResp.Result.(*morpheus.ListCheckAppsResult)
	log.Println(result.CheckApps)
}
