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

	// List policies
	req := &morpheus.Request{}
	policyResponse, err := client.ListPolicies(req)
	if err != nil {
		log.Fatal(err)
	}
	result := policyResponse.Result.(*morpheus.ListPoliciesResult)
	for _, policy := range *result.Policies {
		fmt.Println(policy)
	}
}
