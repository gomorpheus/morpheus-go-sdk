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

	// List vdi pools
	req := &morpheus.Request{}
	vdiPoolsResponse, err := client.ListVDIPools(req)
	if err != nil {
		log.Fatal(err)
	}
	result := vdiPoolsResponse.Result.(*morpheus.ListVDIPoolsResult)
	log.Println(result)
}
