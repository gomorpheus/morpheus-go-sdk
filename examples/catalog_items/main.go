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

	// List catalog items
	req := &morpheus.Request{}
	catalogItemsResp, err := client.ListCatalogItems(req)
	if err != nil {
		log.Fatal(err)
	}
	result := catalogItemsResp.Result.(*morpheus.ListCatalogItemsResult)
	log.Println(result.CatalogItems)
}
