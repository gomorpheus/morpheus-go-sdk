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

	// List network subnets
	req := &morpheus.Request{}
	networkSubnetsResponse, err := client.ListNetworkSubnets(req)
	if err != nil {
		log.Fatal(err)
	}
	result := networkSubnetsResponse.Result.(*morpheus.ListNetworkSubnetsResult)
	subnets := result.NetworkSubnets
	log.Println(subnets)
}
