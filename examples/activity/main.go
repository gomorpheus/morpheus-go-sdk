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

	// Get activity
	req := &morpheus.Request{}
	activityResponse, err := client.GetActivity(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(activityResponse.JsonData)
	result := activityResponse.Result.(*morpheus.GetActivityResult)
	log.Println(result.Activity)
}
