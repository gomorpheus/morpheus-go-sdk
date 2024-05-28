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

	// List email templates
	req := &morpheus.Request{}
	emailTemplatesResponse, err := client.ListEmailTemplates(req)
	if err != nil {
		log.Fatal(err)
	}
	result := emailTemplatesResponse.Result.(*morpheus.ListEmailTemplatesResult)
	for _, emailTemplate := range *result.EmailTemplates {
		fmt.Println(emailTemplate)
	}
}
