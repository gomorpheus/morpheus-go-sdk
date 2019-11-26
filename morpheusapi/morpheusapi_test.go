// Test Package for the Morpheus API (morpheusapi).
// This file defines global assertion methods for use in your tests.

package morpheusapi_test

import (
	_ "fmt"
	_ "io"
	_ "os"
	"reflect"
	"testing"
	"github.com/gomorpheus/morpheus-go/morpheusapi"
)

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Testing Unexported methods
//___________________________________

func assertNil(t *testing.T, v interface{}) {
	if !isNil(v) {
		t.Errorf("[%v] was expected to be nil", v)
	}
}

func assertNotNil(t *testing.T, v interface{}) {
	if isNil(v) {
		t.Errorf("[%v] was expected to be non-nil", v)
	}
}

func assertType(t *testing.T, typ, v interface{}) {
	if reflect.DeepEqual(reflect.TypeOf(typ), reflect.TypeOf(v)) {
		t.Errorf("Expected type %t, got %t", typ, v)
	}
}

func assertError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
}

func assertErrorNotNil(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Error expected to be non-nil")
	}
}

func assertEqual(t *testing.T, e, g interface{}) (r bool) {
	if !equal(e, g) {
		t.Errorf("Expected [%v], got [%v]", e, g)
	}

	return
}

func assertNotEqual(t *testing.T, e, g interface{}) (r bool) {
	if equal(e, g) {
		t.Errorf("Expected [%v], got [%v]", e, g)
	} else {
		r = true
	}

	return
}

func equal(expected, got interface{}) bool {
	return reflect.DeepEqual(expected, got)
}

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)
	kind := rv.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && rv.IsNil() {
		return true
	}

	return false
}

func assertResponse(t *testing.T, resp *morpheusapi.Response, err error) {
	assertError(t, err)
	if resp == nil {
		t.Errorf("API Response was nil")
	} else {
		if resp.Success != true {
			t.Errorf("API Response Error [%v]", err)
			printApiFailure(t, resp, err)
		} else {
			// success
			// logResponse(t, resp)
		}
	}
	
}

func assertBadResponse(t *testing.T, resp *morpheusapi.Response, err error) {
	assertErrorNotNil(t, err)
	if resp == nil {
		t.Errorf("API Response was nil")
	} else {
		if resp.Success != false {
			t.Errorf("API Request should have failed and did not [%v]", err)
			logResponse(t, resp)
		} else {
			// failure
			// logResponse(t, resp)
		}
	}
}

func assertResponseStatusCode(t *testing.T, resp *morpheusapi.Response, statusCode int) {
	if resp.StatusCode != statusCode {
		t.Errorf("Expected API response status to be %d and got [%d]", statusCode, resp.StatusCode)
		logResponse(t, resp)
	}
}

//todo: make this pretty
// just use logResponse always probably

func printApiFailure(t *testing.T, resp *morpheusapi.Response, err error) {
	t.Logf("API FAILUREt: %v", resp)
	if err != nil {
		t.Logf("ERROR: %v", err)
	}
}


func logResponse(t *testing.T, resp *morpheusapi.Response) {
	t.Logf("API RESPONSE: %v", resp)
	// need resp.restyResponse for this
	// t.Logf("Response Status: %v", resp.Status())
	// t.Logf("Response Time: %v", resp.Time())
	// t.Logf("Response Headers: %v", resp.Header())
	// t.Logf("Response Cookies: %v", resp.Cookies())
	// t.Logf("Response Body: %v", resp)
}

