package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func main() {
	// Validate that a Morpheus URL has been set
	if os.Getenv("MORPHEUS_URL") == "" {
		log.Fatal("MORPHEUS_URL variable not set")
	}
	client := morpheus.NewClient(os.Getenv("MORPHEUS_URL"))

	// Use Morpheus access token for authentication
	if os.Getenv("MORPHEUS_TOKEN") != "" {
		token := os.Getenv("MORPHEUS_TOKEN")
		client.SetAccessToken(token, "", 0, "write")
	}

	// Use Morpheus username and password for authentication
	username := os.Getenv("MORPHEUS_USERNAME")
	password := os.Getenv("MORPHEUS_PASSWORD")
	client.SetUsernameAndPassword(username, password)

	// Log into Morpheus
	_, err := client.Login()
	if err != nil {
		fmt.Println("LOGIN ERROR: ", err)
	}

	// Define request payload
	name := "goexample"
	code := "goexample"
	description := "golang example"
	visibility := "private"
	req := &morpheus.Request{
		Body: map[string]interface{}{
			"environment": map[string]interface{}{
				"name":        name,
				"code":        code,
				"description": description,
				"visibility":  visibility,
			},
		},
	}

	// JSON Request
	jsonRequest, _ := json.Marshal(req.Body)
	fmt.Println("JSON Request: ", string(jsonRequest))

	// JSON Response
	output, _ := client.CreateEnvironment(req)
	fmt.Println(output)
}
