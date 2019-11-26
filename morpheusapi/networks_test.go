package morpheusapi_test

import (
	"testing"
	"github.com/gomorpheus/morpheus-go/morpheusapi"
)

var (
	testNetworkName = "morpheusapi-test-network"
)

func TestListNetworks(t *testing.T) {
	client := getTestClient()
	req := &morpheusapi.Request{}
	resp, err := client.ListNetworks(req)
	assertResponse(t, resp, err)
}

func TestGetNetwork(t *testing.T) {
	client := getTestClient()
	req := &morpheusapi.Request{}
	resp, err := client.ListNetworks(req)
	assertResponse(t, resp, err)

	// parse JSON and fetch the first one by ID

	result := resp.Result.(*morpheusapi.ListNetworksResult)
	recordCount := result.Meta.Total
	t.Logf("Found %d Network Domains.", recordCount)
	if recordCount != 0 {
		// Get by ID
		record := (*result.Networks)[0]
		resp, err = client.GetNetwork(record.ID, &morpheusapi.Request{})
		assertResponse(t, resp, err)

		// List by name

	} else {
		
	}
}

func TestNetworksCRUD(t *testing.T) {
	client := getTestClient()
	//create
	// this has no uniqueness check on name, it probably should..
	req := &morpheusapi.Request{
		Body: map[string]interface{}{
			"network": map[string]interface{}{
				"name": testNetworkName,
				"description": "a test network",
				"zone": map[string]interface{}{
					"id": 1,
				},
				// what else? varies by type...heh
			},
		},
	}
	resp, err := client.CreateNetwork(req)
	result := resp.Result.(*morpheusapi.CreateNetworkResult)
	assertResponse(t, resp, err)
	assertNotNil(t, result)
	assertEqual(t, result.Success, true)
	assertNotNil(t, result.Network)
	assertNotEqual(t, result.Network.ID, 0)
	assertEqual(t, result.Network.Name, testNetworkName)

	// update
	updateReq := &morpheusapi.Request{
		Body: map[string]interface{}{
			"network": map[string]interface{}{
				"description": "my new description",
			},
		},
	}
	updateResp, updateErr := client.UpdateNetwork(result.Network.ID, updateReq)
	updateResult := updateResp.Result.(*morpheusapi.UpdateNetworkResult)
	assertResponse(t, updateResp, updateErr)
	assertNotNil(t, updateResult)
	assertEqual(t, updateResult.Success, true)
	assertNotNil(t, updateResult.Network)
	assertNotEqual(t, updateResult.Network.ID, 0)
	assertEqual(t, updateResult.Network.Description, "my new description")
	
	// delete
	deleteReq := &morpheusapi.Request{}
	deleteResp, deleteErr := client.DeleteNetwork(result.Network.ID, deleteReq)
	deleteResult := deleteResp.Result.(*morpheusapi.DeleteNetworkResult)
	assertResponse(t, deleteResp, deleteErr)
	assertNotNil(t, deleteResult)
	assertEqual(t, deleteResult.Success, true)

}