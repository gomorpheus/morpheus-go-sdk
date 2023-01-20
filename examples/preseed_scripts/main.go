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

	// List preseed scripts
	req := &morpheus.Request{}
	preseedScriptsResponse, err := client.ListPreseedScripts(req)
	if err != nil {
		log.Fatal(err)
	}
	result := preseedScriptsResponse.Result.(*morpheus.ListPreseedScriptsResult)
	log.Println(result)
}
