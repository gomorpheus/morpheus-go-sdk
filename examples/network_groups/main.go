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

	// List network groups
	req := &morpheus.Request{}
	networkGroupResponse, err := client.ListNetworkGroups(req)
	if err != nil {
		log.Fatal(err)
	}
	result := networkGroupResponse.Result.(*morpheus.ListNetworkGroupsResult)
	networkGroups := result.NetworkGroups
	log.Println(networkGroups)
}
