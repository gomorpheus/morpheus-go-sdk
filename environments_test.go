package morpheus_test

import (
	"github.com/gomorpheus/morpheus-go-sdk"
	"testing"
)

func TestListEnvironments(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.ListEnvironments(req)
	assertResponse(t, resp, err)
}
