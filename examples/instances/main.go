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

	// List instances
	req := &morpheus.Request{}
	instancesResponse, err := client.ListInstances(req)
	if err != nil {
		log.Fatal(err)
	}
	result := instancesResponse.Result.(*morpheus.ListInstancesResult)
	for _, instance := range *result.Instances {
		fmt.Println(instance)
	}
}
