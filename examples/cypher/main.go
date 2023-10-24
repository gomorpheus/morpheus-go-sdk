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

	createReq := &morpheus.Request{
		Body: map[string]interface{}{
			"ttl":   "15m",
			"value": "test secret",
		},
	}
	createResp, err := client.CreateCypher("secret/gotest", createReq)
	if err != nil {
		fmt.Println(err)
	}
	result := createResp.Result.(*morpheus.CreateCypherResult)
	fmt.Println(result)

	// List cyphers
	req := &morpheus.Request{}
	cyphersResponse, err := client.ListCyphers(req)
	if err != nil {
		log.Fatal(err)
	}
	listResult := cyphersResponse.Result.(*morpheus.ListCypherResult)
	log.Println(listResult.Cyphers)

	// Delete cypher
	deleteReq := &morpheus.Request{}
	deleteResp, deleteErr := client.DeleteCypher("secret/gotest", deleteReq)
	if deleteErr != nil {
		fmt.Println(deleteErr)
	}
	deleteResult := deleteResp.Result.(*morpheus.DeleteCypherResult)
	fmt.Println(deleteResult)
}
