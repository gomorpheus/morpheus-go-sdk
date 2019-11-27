package morpheusapi_test

import (
	"testing"
	"github.com/gomorpheus/morpheus-go-sdk"
)

func TestListInstances(t *testing.T) {
	client := getTestClient()
	resp, err := client.ListInstances(&morpheusapi.Request{})
	assertResponse(t, resp, err)
}

func TestGetInstance(t *testing.T) {
	client := getTestClient()
	resp, err := client.ListInstances(&morpheusapi.Request{})
	assertResponse(t, resp, err)
	// parse JSON and fetch the first one by ID
	listInstancesResult := resp.Result.(*morpheusapi.ListInstancesResult)
	instancesCount := listInstancesResult.Meta.Total
	t.Logf("Found %d Instances.", instancesCount)
	// if instancesCount != 0 {
		firstInstance := (*listInstancesResult.Instances)[0]	
		// fmt.Println(fmt.Sprintf("First Instance: [%d] %v: ", firstInstance.ID, firstInstance.Name))
		resp, err = client.GetInstance(firstInstance.ID, &morpheusapi.Request{})
		assertResponse(t, resp, err)
	// }
	
}


// this requires params zoneId&layoutId&siteId, heh
// func TestListInstancePlans(t *testing.T) {
// 	client := getTestClient()
// 	resp, err := client.ListInstancePlans(&morpheusapi.Request{})
// 	assertResponse(t, resp, err)
// }

