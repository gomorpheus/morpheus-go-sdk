package morpheusapi_test

import (
	"testing"
	"github.com/gomorpheus/morpheusapi"
)

var (
	testNetworkDomainName = "morpheusapi.gomorpheus.com"
)

func TestListNetworkDomains(t *testing.T) {
	client := getTestClient()
	req := &morpheusapi.Request{}
	resp, err := client.ListNetworkDomains(req)
	assertResponse(t, resp, err)
}

func TestGetNetworkDomain(t *testing.T) {
	client := getTestClient()
	req := &morpheusapi.Request{}
	resp, err := client.ListNetworkDomains(req)
	assertResponse(t, resp, err)

	// parse JSON and fetch the first one by ID

	result := resp.Result.(*morpheusapi.ListNetworkDomainsResult)
	recordCount := result.Meta.Total
	t.Logf("Found %d Network Domains.", recordCount)
	if recordCount != 0 {
		// Get by ID
		record := (*result.NetworkDomains)[0]
		resp, err = client.GetNetworkDomain(record.ID, &morpheusapi.Request{})
		assertResponse(t, resp, err)

		// List by name

	} else {
		
	}
}

func TestNetworkDomainsCRUD(t *testing.T) {
	client := getTestClient()
	//create
	// this has no uniqueness check on name, it probably should..
	req := &morpheusapi.Request{
		Body: map[string]interface{}{
			"networkDomain": map[string]interface{}{
				"name": testNetworkDomainName,
				"description": "a test domain",
				"publicZone": false,
				"domainController": false,
				"visibility":"private",
			},
		},
	}
	resp, err := client.CreateNetworkDomain(req)
	result := resp.Result.(*morpheusapi.CreateNetworkDomainResult)
	assertResponse(t, resp, err)
	assertNotNil(t, result)
	assertEqual(t, result.Success, true)
	assertNotNil(t, result.NetworkDomain)
	assertNotEqual(t, result.NetworkDomain.ID, 0)
	assertEqual(t, result.NetworkDomain.Name, testNetworkDomainName)

	// update
	updateReq := &morpheusapi.Request{
		Body: map[string]interface{}{
			"networkDomain": map[string]interface{}{
				"description": "my new description",
			},
		},
	}
	updateResp, updateErr := client.UpdateNetworkDomain(result.NetworkDomain.ID, updateReq)
	updateResult := updateResp.Result.(*morpheusapi.UpdateNetworkDomainResult)
	assertResponse(t, updateResp, updateErr)
	assertNotNil(t, updateResult)
	assertEqual(t, updateResult.Success, true)
	assertNotNil(t, updateResult.NetworkDomain)
	assertNotEqual(t, updateResult.NetworkDomain.ID, 0)
	assertEqual(t, updateResult.NetworkDomain.Description, "my new description")
	
	// delete
	deleteReq := &morpheusapi.Request{}
	deleteResp, deleteErr := client.DeleteNetworkDomain(result.NetworkDomain.ID, deleteReq)
	deleteResult := deleteResp.Result.(*morpheusapi.DeleteNetworkDomainResult)
	assertResponse(t, deleteResp, deleteErr)
	assertNotNil(t, deleteResult)
	assertEqual(t, deleteResult.Success, true)

}