package morpheusapi_test

import (
	"testing"
	"github.com/gomorpheus/morpheus-go/morpheusapi"
)

func TestWhoami(t *testing.T) {
	client := getTestClient()
	resp, err := client.Whoami()
	assertResponse(t, resp, err)
	result := resp.Result.(*morpheusapi.WhoamiResult)
	assertNotNil(t, result.User)
	assertNotNil(t, result.User.ID)
	assertEqual(t, result.User.Username, testUsername)
}
