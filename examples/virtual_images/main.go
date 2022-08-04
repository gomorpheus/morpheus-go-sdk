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

	// List virtual images
	req := &morpheus.Request{
		QueryParams: map[string]string{"filterType": "All"},
	}
	response, err := client.ListVirtualImages(req)
	if err != nil {
		log.Fatal(err)
	}
	result := response.Result.(*morpheus.ListVirtualImagesResult)
	log.Println(result.VirtualImages)
}
