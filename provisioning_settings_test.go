package morpheus_test

import (
	"testing"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func TestGetProvisioningSettings(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.GetProvisioningSettings(req)
	assertResponse(t, resp, err)
}
