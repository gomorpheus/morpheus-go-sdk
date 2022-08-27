package morpheus_test

import (
	"testing"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func TestListIdentitySources(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.ListIdentitySources(req)
	assertResponse(t, resp, err)
}

func TestGetIdentitySource(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.ListScaleThresholds(req)
	assertResponse(t, resp, err)

	// parse JSON and fetch the first one by ID
	result := resp.Result.(*morpheus.ListIdentitySourcesResult)
	recordCount := result.Meta.Total
	t.Logf("Found %d Identity Sources.", recordCount)
	if recordCount != 0 {
		// Get by ID
		record := (*result.IdentitySources)[0]
		resp, err = client.GetIdentitySource(record.ID, &morpheus.Request{})
		assertResponse(t, resp, err)
	}
}
