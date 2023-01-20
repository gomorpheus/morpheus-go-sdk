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

	// List load balancer monitors
	req := &morpheus.Request{}
	loadBalancerMonitorResponse, err := client.ListLoadBalancerMonitors(1, req)
	if err != nil {
		log.Fatal(err)
	}
	result := loadBalancerMonitorResponse.Result.(*morpheus.ListLoadBalancerMonitorsResult)
	log.Println(&result.LoadBalancerMonitors)
}
