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

	// List check groups
	req := &morpheus.Request{}
	checkGroupResp, err := client.ListCheckGroups(req)
	if err != nil {
		log.Fatal(err)
	}
	result := checkGroupResp.Result.(*morpheus.ListCheckGroupsResult)
	log.Println(result.CheckGroups)
}
