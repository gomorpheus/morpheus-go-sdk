package main

import (
	"encoding/json"
	"fmt"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func main() {
	client := morpheus.NewClient("")
	client.SetUsernameAndPassword("", "")
	_, err := client.Login()
	if err != nil {
		fmt.Println("LOGIN ERROR: ", err)
	}
	//fmt.Println("LOGIN RESPONSE:", resp)
	hostname := "kubedemo13"
	req := &morpheus.Request{
		Body: map[string]interface{}{
			"cluster": map[string]interface{}{
				"name":        hostname,
				"description": "testing terraform",
				"type":        "kubernetes-cluster",
				"group":       map[string]interface{}{"id": 1},
				"cloud":       map[string]interface{}{"id": 1},
				"layout":      map[string]interface{}{"id": 121},
				"server": map[string]interface{}{
					"hostname": hostname,
					"name":     hostname,
					"config": map[string]interface{}{
						"resourcePool": "resgroup-23",
					},
					"plan": map[string]interface{}{
						"id":   370,
						"code": "vm-8192",
					},
					"volumes": []map[string]interface{}{
						{
							"id":          -1,
							"rootVolume":  true,
							"name":        "root",
							"size":        20,
							"sizeId":      nil,
							"storageType": 1,
							"datastoreId": "auto",
						},
					},
					"networkInterfaces": []map[string]interface{}{
						{
							"network": map[string]interface{}{
								"id": "network-1",
							},
						},
					},
					"visibility":    "private",
					"nodeCount":     1,
					"networkDomain": nil,
				},
			},
		},
	}
	//fmt.Println(req)
	slcB, _ := json.Marshal(req.Body)
	fmt.Println(string(slcB))
	output, _ := client.CreateCluster(req)
	fmt.Println(output)

	//options, err := client.ListOptionTypes(&morpheus.Request{})
	//fmt.Println(options)
	/*
		req := &morpheus.Request{
			Body: map[string]interface{}{
				"optionTypeList": map[string]interface{}{
					"name":            "tfnew",
					"description":     "testing terraform",
					"type":            "rest",
					"sourceUrl":       "https://gist.githubusercontent.com/martezr/56a1c7d69e96a8dd89c0d47236a90249/raw/ca2cd8719b17a669d398d9b576dd084c12ef0237/projects.json",
					"sourceMethod":    "GET",
					"ignoreSSLErrors": true,
					"realTime":        true,
					"config": map[string]interface{}{
						"sourceHeaders": []map[string]interface{}{
							{"name": "content-type",
								"value":  "application/json",
								"masked": false},
						},
					},
				},
			},
		}
		fmt.Println(req)
		slcB, _ := json.Marshal(req.Body)
		fmt.Println(string(slcB))
		output, _ := client.CreateOptionList(req)
		fmt.Println(output)
	*/
	/*
		options := []int64{1523, 1524, 1525, 1526, 1527}
		req := &morpheus.Request{
			Body: map[string]interface{}{
				"taskSet": map[string]interface{}{
					"name":        "avi workflow",
					"description": "avi workflow",
					"optionTypes": options,
					"type":        "operation",
				},
			},
		}
		fmt.Println(req)
		resp, _ = client.CreateTaskSet(req)
		fmt.Println(resp)

		/*
			usersources, err := client.ListUserSources(&morpheus.Request{})
			fmt.Println(usersources)

			usersource, err := client.FindUserSourceByName("GRTAD")
			fmt.Println((usersource))

			req := &morpheus.Request{
				Body: map[string]interface{}{
					"userSource": map[string]interface{}{
						"name":        "TestAD",
						"description": "GO AD",
						"type":        "activeDirectory",
						"config": map[string]interface{}{
							"url":             "grtdc01.grt.local",
							"domain":          "grt.local",
							"bindingUsername": "administrator",
							"bindingPassword": "test",
						},
					},
				},
			}

			createdsource, err := client.CreateUserSource(req)
			fmt.Println(createdsource)

			//output, err := client.DeleteUserSource(1, &morpheus.Request{})
			//fmt.Println((output))
	*/
}
