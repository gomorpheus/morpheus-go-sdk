package morpheus_test

import (
	"testing"

	"github.com/gomorpheus/morpheus-go-sdk"
)

var (
	testNetworkName = "golangtest-network"
)

func TestListNetworks(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.ListNetworks(req)
	assertResponse(t, resp, err)
}

func TestGetNetwork(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.ListNetworks(req)
	assertResponse(t, resp, err)

	// parse JSON and fetch the first one by ID

	result := resp.Result.(*morpheus.ListNetworksResult)
	recordCount := result.Meta.Total
	t.Logf("Found %d Networks.", recordCount)
	if recordCount != 0 {
		// Get by ID
		record := (*result.Networks)[0]
		resp, err = client.GetNetwork(record.ID, &morpheus.Request{})
		assertResponse(t, resp, err)
	}
}

func _Busted_TestNetworksCRUD(t *testing.T) {
	client := getTestClient(t)
	//create
	// this has no uniqueness check on name, it probably should..
	req := &morpheus.Request{
		Body: map[string]interface{}{
			"network": map[string]interface{}{
				"name":        testNetworkName,
				"description": "a test network",
				"zone": map[string]interface{}{
					"id": 1,
				},
				// what else? varies by type...heh
			},
		},
	}
	resp, err := client.CreateNetwork(req)
	result := resp.Result.(*morpheus.CreateNetworkResult)
	assertResponse(t, resp, err)
	assertNotNil(t, result)
	assertEqual(t, result.Success, true)
	assertNotNil(t, result.Network)
	assertNotEqual(t, result.Network.ID, 0)
	assertEqual(t, result.Network.Name, testNetworkName)

	// update
	updateReq := &morpheus.Request{
		Body: map[string]interface{}{
			"network": map[string]interface{}{
				"description": "my new description",
			},
		},
	}
	updateResp, updateErr := client.UpdateNetwork(result.Network.ID, updateReq)
	updateResult := updateResp.Result.(*morpheus.UpdateNetworkResult)
	assertResponse(t, updateResp, updateErr)
	assertNotNil(t, updateResult)
	assertEqual(t, updateResult.Success, true)
	assertNotNil(t, updateResult.Network)
	assertNotEqual(t, updateResult.Network.ID, 0)
	assertEqual(t, updateResult.Network.Description, "my new description")

	// delete
	deleteReq := &morpheus.Request{}
	deleteResp, deleteErr := client.DeleteNetwork(result.Network.ID, deleteReq)
	deleteResult := deleteResp.Result.(*morpheus.DeleteNetworkResult)
	assertResponse(t, deleteResp, deleteErr)
	assertNotNil(t, deleteResult)
	assertEqual(t, deleteResult.Success, true)

}
