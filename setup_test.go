package morpheusapi_test

import (
	"testing"
	"github.com/gomorpheus/morpheus-go-sdk"
	"net/http"
)

func TestSetupCheck(t *testing.T) {
	client := getTestClient()
	req := &morpheusapi.Request{}
	resp, err := client.SetupCheck(req)
	assertResponse(t, resp, err)
}


// func TestSetupInit(t *testing.T) {
// 	client := getTestClient()
// 	req := &morpheusapi.Request{
// 		Body: map[string]interface{}{
// 			"accountName": "root",
// 			"firstName": "Super",
// 			"lastName": "Admin",
// 			"password": "ChangeM3!",
// 		}
// 	}
// 	resp, err := client.SetupInit(req)
// 	assertResponse(t, resp, err)
// }


// It should expect to be already setup
func TestSetupInitShouldBeAlreadySetup(t *testing.T) {
	client := getTestClient()
	req := &morpheusapi.Request{
		Body: map[string]interface{}{
			"accountName": "root",
			"firstName": "Super",
			"lastName": "Admin",
			"password": "ChangeM3!",
		},
	}
	resp, err := client.SetupInit(req)
	assertResponseStatusCode(t, resp, http.StatusBadRequest)
	assertErrorNotNil(t, err)
	//assertBadResponse(t, resp, err)
	// could assert resp.Msg == "Already setup" or whatever
}

