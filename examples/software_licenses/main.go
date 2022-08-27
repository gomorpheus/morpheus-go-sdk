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

	// List software licenses
	req := &morpheus.Request{}
	licensesResponse, err := client.ListSoftwareLicenses(req)
	if err != nil {
		log.Fatal(err)
	}
	result := licensesResponse.Result.(*morpheus.ListSoftwareLicensesResult)
	log.Println(result)
}
