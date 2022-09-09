package morpheus_test

import (
	"testing"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func TestClusterTypes(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.ListClusterTypes(req)
	assertResponse(t, resp, err)
}
