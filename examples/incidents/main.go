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

	// List incidents
	req := &morpheus.Request{}
	incidentResp, err := client.ListIncidents(req)
	if err != nil {
		log.Fatal(err)
	}
	result := incidentResp.Result.(*morpheus.ListIncidentsResult)
	log.Println(result.Incidents)
}
