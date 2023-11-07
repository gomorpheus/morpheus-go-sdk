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

	// List clouds
	req := &morpheus.Request{}
	cloudsResp, err := client.ListClouds(req)
	if err != nil {
		log.Fatal(err)
	}
	result := cloudsResp.Result.(*morpheus.ListCloudsResult)
	log.Println(result.Clouds)

	// List cloud datastores
	req = &morpheus.Request{}
	cloudDatastoresResp, err := client.ListCloudDatastores((*result.Clouds)[0].ID, req)
	if err != nil {
		log.Fatal(err)
	}
	datastore := cloudDatastoresResp.Result.(*morpheus.ListCloudDatastoresResult)
	log.Println(datastore.Datastores)

	// List cloud resource folders
	req = &morpheus.Request{}
	cloudResourceFoldersResp, err := client.ListCloudResourceFolders((*result.Clouds)[0].ID, req)
	if err != nil {
		log.Fatal(err)
	}
	folder := cloudResourceFoldersResp.Result.(*morpheus.ListCloudResourceFoldersResult)
	log.Println(folder.Folders)
}
