package morpheusapi_test

import (
	"testing"
	"github.com/gomorpheus/morpheus-go-sdk"
)

func TestListClouds(t *testing.T) {
	client := getTestClient()
	resp, err := client.ListClouds(&morpheusapi.Request{})
	assertResponse(t, resp, err)
}

func TestGetCloud(t *testing.T) {
	client := getTestClient()
	resp, err := client.ListClouds(&morpheusapi.Request{})
	assertResponse(t, resp, err)
	// parse JSON and fetch the first one by ID
	listCloudsResult := resp.Result.(*morpheusapi.ListCloudsResult)
	cloudsCount := listCloudsResult.Meta.Total
	t.Logf("Found %d Clouds.", cloudsCount)
	if cloudsCount != 0 {
		firstCloud := (*listCloudsResult.Clouds)[0]	
		// fmt.Println(fmt.Sprintf("First Cloud: [%d] %v: ", firstCloud.ID, firstCloud.Name))
		// fmt.Println("resp.Result: ", resp.Result)
		resp, err = client.GetCloud(firstCloud.ID, &morpheusapi.Request{})
		assertResponse(t, resp, err)
	}
	
}