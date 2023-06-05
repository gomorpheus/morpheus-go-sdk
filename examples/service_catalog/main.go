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

	// List inventory
	inventoryReq := &morpheus.Request{}
	inventoryResponse, err := client.ListCatalogInventoryItems(inventoryReq)
	if err != nil {
		log.Fatal(err)
	}
	ouptut := inventoryResponse.Result.(*morpheus.ListCatalogInventoryItemsResult)
	items := ouptut.CatalogInventoryItems
	fmt.Println(items)
}
