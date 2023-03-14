package morpheus_test

import (
	"testing"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func TestGetGuidanceSettings(t *testing.T) {
	client := getTestClient(t)
	req := &morpheus.Request{}
	resp, err := client.GetGuidanceSettings(req)
	assertResponse(t, resp, err)
}
