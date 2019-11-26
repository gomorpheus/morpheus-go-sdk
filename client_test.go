// This provides common functions for all morpheusapi.Client tests.
package morpheusapi_test

import (
	"testing"
	"fmt"
	"os"
	_ "strconv"
	_ "encoding/json"
	"github.com/gomorpheus/morpheusapi"
)

var (
	// DEBUG = (os.Getenv("MORPHEUS_API_DEBUG") == "true")
	testUrl = os.Getenv("MORPHEUS_API_TEST_URL")
	testUsername = os.Getenv("MORPHEUS_API_TEST_USERNAME")
	testPassword = os.Getenv("MORPHEUS_API_TEST_PASSWORD")
	// testUrl                  string
	// testUsername             string
	// testPassword             string
	sharedTestClient         *morpheusapi.Client
)


// // needed to hook up setup/teardown methods
// // this probably belongs in its own file..
// func TestMain(m *testing.M) {
//     setup()
//     code := m.Run() 
//     teardown()
//     os.Exit(code)
// }

// func setup() {
// 	//fmt.Println(fmt.Sprintf("Setup test suite..."))
// }

// func teardown() {
// 	//fmt.Println(fmt.Sprintf("Teardown test suite..."))
// 	if sharedTestClient != nil {
// 		fmt.Println(fmt.Sprintf("Client Request Summary | Success: %d, Error: %d", sharedTestClient.RequestCount(), sharedTestClient.ErrorCount()))
// 	}
// }

func getNewClient() (*morpheusapi.Client) {
	if testUrl == "" {
		panic("MORPHEUS_API_TEST_URL must be set to to run tests")
	}
	// testUsername = os.Getenv("MORPHEUS_API_TEST_USERNAME")
	// if testUsername == "" {
	// 	panic("MORPHEUS_API_TEST_USERNAME must be set to to run tests")
	// }
	// testPassword = os.Getenv("MORPHEUS_API_TEST_PASSWORD")
	// if testUsername == "" {
	// 	panic("MORPHEUS_API_TEST_PASSWORD must be set to to run tests")
	// }
	//fmt.Println(fmt.Sprintf("Initializing new client for %v @ %v", testUrl, testUsername))
	client := morpheusapi.NewClient(testUrl)
	return client
}

// this does not work for some reason...
// a client to be shared between requests.
func getTestClient() (*morpheusapi.Client) {
	if sharedTestClient == nil {
		if testUrl == "" {
			panic("MORPHEUS_API_TEST_URL must be set to to run tests.")
		}
		if testUsername == "" {
			panic("MORPHEUS_API_TEST_USERNAME must be set to to run tests")
		}
		if testPassword == "" {
			panic("MORPHEUS_API_TEST_PASSWORD must be set to to run tests")
		}
		fmt.Println(fmt.Sprintf("Initializing test client for %v @ %v", testUsername, testUrl))
		sharedTestClient = morpheusapi.NewClient(testUrl)
		sharedTestClient.SetUsernameAndPassword(testUsername, testPassword)
	}
	return sharedTestClient
}

// func parseInt64(s string) (int64, error) {
// 	n, err := strconv.ParseInt(s, 10, 64)
// 	return n, err
// }

func TestGet(t *testing.T) {
	client := getNewClient()
	testRequest := &morpheusapi.Request{
		Method: "GET",
		Path: "/api/setup/check",
		QueryParams:map[string]string{
	          "foo": "bar",
	          "gotest": "true",
	      },
	}
	resp, err := client.Execute(testRequest)
	assertResponse(t, resp, err)
}

func TestLogin(t *testing.T) {
	client := getNewClient()
	client.SetUsernameAndPassword(testUsername, testPassword)
	resp, err := client.Login()
	assertResponse(t, resp, err)
}

func TestLogout(t *testing.T) {
	client := getNewClient()
	client.SetUsernameAndPassword(testUsername, testPassword)
	resp, err := client.Login()
	assertResponse(t, resp, err)
	if err == nil {
		resp, err = client.Logout()
		assertError(t, err)
		// assertResponse(t, resp, err)
	}
}

func TestLoginRepeated(t *testing.T) {
	client := getNewClient()
	client.SetUsernameAndPassword(testUsername, testPassword)
	resp, err := client.Login()
	assertResponse(t, resp, err)
	if err == nil {
		resp, err = client.Login()
		resp, err = client.Login()
		resp, err = client.Login()
		resp, err = client.Login()
		// assertResponse(t, resp, err)
		assertError(t, err)
	}
}
