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

	// List network pools
	req := &morpheus.Request{}
	networkPoolResponse, err := client.ListNetworkPools(req)
	if err != nil {
		log.Fatal(err)
	}
	result := networkPoolResponse.Result.(*morpheus.ListNetworkPoolsResult)
	pools := result.NetworkPools
	log.Println(pools)
}
