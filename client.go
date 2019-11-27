// Client is the driver for interfacing with the Morpheus API
package morpheus

import (
    "fmt"
    "errors"
    "time"
	"encoding/json"
	"github.com/go-resty/resty"
)

type Client struct {
    Url string
    Username string
    Password string
    AccessToken string // todo: make internal
    RefreshToken string // todo: make internal
    AuthenticatedAt time.Time
    ExpiresIn int64
    Scope string
    UserAgent string
    //Headers map[string]string
 	//BaseURL   *url.URL
    //RestyClient *http.Client
    // LastLoginDate time
    // requests []*Request
    lastRequest *Request
    lastResponse *Response
    requestCount int64
    successCount int64
    errorCount int64
}

// func (client * Client) String() string {
//         return fmt.Sprintf("Client Url: %s Username: %s Logged In: %b", client.Url, client.Username, client.IsLoggedIn())
// }

func (client * Client) IsLoggedIn() bool {
	return client.AccessToken != ""
}

func (client * Client) RequestCount() int64 {
	return client.requestCount
}

func (client * Client) SuccessCount() int64 {
	return client.successCount
}

func (client * Client) ErrorCount() int64 {
	return client.errorCount
}

func (client * Client) incrementRequests(req * Request, resp * Response) {
	client.lastRequest = req
	client.lastResponse = resp
	client.requestCount++
	if resp.Success {
		client.successCount++
	} else {
		client.errorCount++ 
	}
}

func (client * Client) LastRequest() (*Request) {
	return client.lastRequest
}

func (client * Client) LastResponse() (*Response) {
	return client.lastResponse
}

// parseJsonToResult parses json into the given output (struct).
// The type of the ouput determines how it is parsed.
func parseJsonToResult(data []byte, output interface{}) (error) {
	var err error
	if (data != nil) {
		err = json.Unmarshal(data, &output)
	}
	return err
}

func NewClient(url string) (client * Client) {
 	var userAgent = "morpheus-terraform-plugin v0.1"
	//client := &Client{}
	// client.Url = url
	// client.UserAgent = userAgent
	// return client
    return &Client{
    	Url: url,
    	UserAgent: userAgent,
    }
}

func (client * Client) SetUsername(username string) (*Client) {
	// clear access token if switching users
	if (client.Username != username) {
		client.ClearAccessToken()
		//client.AccessToken = ""
	}
	client.Username = username
	return client
}

func (client * Client) SetPassword(password string) (*Client) {
	client.Password = password
	return client
}

func (client * Client) SetUsernameAndPassword(username string, password string) (*Client) {
	client.SetUsername(username)
	client.SetPassword(password)
	return client
}

func (client * Client) SetAccessToken(accessToken string, refreshToken string, expiresIn int64, scope string) (*Client) {
	client.AccessToken = accessToken
	client.RefreshToken = refreshToken
	client.ExpiresIn = expiresIn
	client.Scope = scope
	return client
}

func (client * Client) ClearAccessToken() (*Client) {
	client.AccessToken = ""
	client.RefreshToken = ""
	client.ExpiresIn = 0
	client.Scope = ""
	return client
}


func (client * Client) Execute(req * Request) (*Response, error) {
	// first, login if needed
	if (req.SkipLogin != true) {
		if (client.IsLoggedIn() != true) {
			//fmt.Println("Autologin as " + client.Username)
			if client.Username == "" {
				// fmt.Println("Skipping login because Username is not set.")
			} else {
				loginResp, loginErr := client.Login()
				if loginErr != nil {
					return loginResp, loginErr
				}
			}
		} else {
			//fmt.Println("You are logged in as " + client.Username)
		}
	}

	// The transient resty response object
	var restyResponse *resty.Response


	// The response object to be returned
	var resp *Response

	// potential error to be returned
	var err error

	// construct the request
    var httpMethod = req.Method
    if (httpMethod == "") {
    	// httpMethod = "GET"
    	return nil, errors.New("Invalid Request: Method is required eg. GET,POST,PUT,DELETE")
    }

    var url string = client.Url + req.Path

	//var url string = client.Url + req.Path
	// construct resty.Client
    restyClient := resty.New()
    //set timeout
    if (req.Timeout > 0) {
    	restyClient.SetTimeout(time.Duration(req.Timeout) * time.Second)
    } else {
    	// restyClient.SetTimeout(time.Duration(30) * time.Second)
    }
    // construct resty.Request
	restyReq := restyClient.R()
	
	// set query params
	restyReq.SetQueryParams(req.QueryParams)
	
	// set Headers
	// Set default headers: application/json
	if req.Headers != nil {
		// restyReq.SetHeaders(req.Headers)
		for k,v := range req.Headers {
			restyReq.SetHeader(k, v)
		}
	}
	
	// add Authorization Header with our access token
	if req.SkipAuthorization != true {
		if restyReq.Header["Authorization"] == nil {
			if client.AccessToken != "" {
				restyReq.SetHeader("Authorization", "Bearer " + client.AccessToken)
			}
		}
	}

	// set body
	if (httpMethod == "POST" || httpMethod == "PUT" || httpMethod == "PATCH") {
    	// FormData means use application/x-www-form-urlencoded
    	if req.FormData != nil {
    		//fmt.Println("SETTING REQUEST FORM DATA:", req.FormData)
    		// var formData map[string]string
    		// for k,v := range req.FormData {
    		// 	formData[k] = fmt.Sprintf("%v", v)
    		// }
    		// restyReq.SetFormData(formData)
    		restyReq.SetFormData(req.FormData)
    		if restyReq.Header["Content-Type"] == nil {
				restyReq.SetHeader("Content-Type", "application/x-www-form-urlencoded")
			}
    	} 
    	if req.Body != nil {
    		//fmt.Println("SETTING REQUEST BODY:", req.Body)
    		// Aways json for now...
    		// todo: use encoder
    		restyReq.SetBody(req.Body)
    		if restyReq.Header["Content-Type"] == nil {
				restyReq.SetHeader("Content-Type", "application/json")
			}

    	}
    	
    	// Set default headers: application/json
		if restyReq.Header["Content-Type"] == nil {
			restyReq.SetHeader("Content-Type", "application/json")
		}
    }

    // Set default Accept header
	if restyReq.Header["Accept"] == nil {
		restyReq.SetHeader("Accept", "application/json")
	}

	// print for debugging
    logDebug(fmt.Sprintf("Request: %s %s", req.Method, url))

	// Make the request
    if httpMethod == "GET" {
    	restyResponse, err = restyReq.Get(url)
    } else if httpMethod == "POST" {
    	restyResponse, err = restyReq.Post(url)
    } else if httpMethod == "PUT" {
    	restyResponse, err = restyReq.Put(url)
    } else if httpMethod == "DELETE" {
    	restyResponse, err = restyReq.Delete(url)
   	} else if httpMethod == "PATCH" {
   		restyResponse, err = restyReq.Patch(url)
    } else if httpMethod == "HEAD" {
    	restyResponse, err = restyReq.Head(url)
    } else if httpMethod == "OPTIONS" {
    	restyResponse, err = restyReq.Options(url)
    // } else if httpMethod == "LIST" {
    // restyResponse, err = restyReq.List(url)
    } else {
    	return nil, errors.New(fmt.Sprintf("Invalid Request.  Unknown HTTP Method: %v", httpMethod))
    }

    // convert a resty response into our Response object

	//var err error
	
	
    resp = &Response{
    	//RestyResponse: restyResponse,
    	Success: restyResponse.IsSuccess(),
    	StatusCode: restyResponse.StatusCode(),
    	Status: restyResponse.Status(),
    	ReceivedAt: restyResponse.ReceivedAt(),
    	Size: restyResponse.Size(),
    	Body: restyResponse.Body(), // byte[]
    }

    // determine success and set err accordingly
    if resp.Success != true {
    	err = errors.New(fmt.Sprintf("API returned HTTP %d", resp.StatusCode))
    	// try to parse the result as a standard result to get success info
    	var standardResult StandardResult
		standardResultParseErr := json.Unmarshal(resp.Body, &standardResult)
		if standardResultParseErr != nil {
			// failed to parse body as standard result json
			// err = standardResultParseErr
		} else {
			if standardResult.Message != "" {
				err = errors.New(standardResult.Message)
			}
			if len(standardResult.Errors) != 0 {
				// err = errors.New(standardResult.Errors[0])
				// resp.Errors = standardResult.Errors
			}
		}
    }
    // resp.Error = err
    // RestyResponse is a the underlying resty object, 
    // This is handy for inspecting the complete request
    // The http response is available at RestyResponse.RawResponse
    resp.RestyResponse = restyResponse
    
    // attempt to parse as json, populates JsonData
    var parsedResult interface{} 
    jsonError := parseJsonToResult(resp.Body, &parsedResult)
    resp.JsonData = parsedResult
    resp.JsonParseError = jsonError

    // attempt to parse json into result type
        	//Error: restyResponse.Error()
	// arbitrary interface{} data is parsed and stored in here
    // if (req.Result != nil) {
    	resp.Result = req.Result
    	// could even just skip this and not return the error
    	// parseJsonToResult(resp.Body, &resp.Result)
    	jsonParseResultError := parseJsonToResult(resp.Body, &resp.Result)
    	
    	if jsonParseResultError != nil {
    		logError(fmt.Sprintf("Error parsing JSON result for type %T [%v]", resp.Result, jsonParseResultError))
    		// maybe this should be treated as a failure..
    		// err = jsonParseResultError
    		// resp.Success = false
    		// return resp, jsonParseResultError
    	} else {
    		//fmt.Println(fmt.Sprintf("Parsed JSON result for type %T", resp.Result))
    	}
    // }

    
    // print for debugging
    // avoid printing request body for now, it may have secrets.
    // if req.Body != nil {
    // 	fmt.Println(fmt.Sprintf("==> Request: %s %s JSON: %s", req.Method, url, req.Body))
    // } else if req.FormData != nil {
    // 	fmt.Println(fmt.Sprintf("==> Request: %s %s BODY: %s", req.Method, url, req.FormData))
    // } else {
    // 	fmt.Println(fmt.Sprintf("==> Request: %s %s", req.Method, url))
    // }
    
    // only print body for failures
    if err != nil {
    	logDebug(fmt.Sprintf("Response: %d %s", resp.StatusCode, resp.Body))
    } else {
    	//fmt.Println(fmt.Sprintf("Response: %d", resp.StatusCode))
    	logDebug(fmt.Sprintf("Response: %d %s", resp.StatusCode, resp.Body))
    }

    client.incrementRequests(req, resp)

    return resp, err
}

func (client * Client) Get(req * Request) (*Response, error) {
	req.Method = "GET"
	return client.Execute(req)
}

func (client * Client) Post(req * Request) (*Response, error) {
	req.Method = "POST"
	return client.Execute(req)
}

func (client * Client) Put(req * Request) (*Response, error) {
	req.Method = "PUT"
	return client.Execute(req)
}

func (client * Client) Delete(req * Request) (*Response, error) {
	req.Method = "DELETE"
	return client.Execute(req)
}

func (client * Client) Patch(req * Request) (*Response, error) {
	req.Method = "PATCH"
	return client.Execute(req)
}

func (client * Client) Head(req * Request) (*Response, error) {
	req.Method = "HEAD"
	return client.Execute(req)
}

func (client * Client) Options(req * Request) (*Response, error) {
	req.Method = "OPTIONS"
	return client.Execute(req)
}

// func (client * Client) List(req * Request) (*Response, error) {
// 	req.Method = "LIST"
// 	return client.Execute(req)
// }


type LoginResult struct {
    AccessToken string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    ExpiresIn int64 `json:"expires_in"`
    Scope string `json:"scope"`
}

func (client * Client) Login() (* Response, error) {
	// already logged in
	if (client.IsLoggedIn()) {
		logDebug("Login skipped. Already logged in as: " + client.Username)
		// todo: validate token at /api/whoami
		return nil, nil
	} else {
		//fmt.Println(fmt.Sprintf("Logging in as %s at %s", client.Username, client.Url))
		loginRequest := &Request{
			Method: "POST",
			Path: "/oauth/token",
			QueryParams:map[string]string{
		          "client_id": "morph-api",
		          "grant_type": "password",
		          "scope": "write",
		          "username": client.Username,
		    },
		    FormData:map[string]string{
		    	//"username": client.username,
		        "password": client.Password,
		    },
		    Timeout: 10,
		    SkipLogin: true,
		}
		resp, err := client.Execute(loginRequest)

		if (resp.Success) {
			var loginResult LoginResult
			jsonErr := json.Unmarshal(resp.Body, &loginResult)
			if jsonErr != nil {
				logError(fmt.Sprintf("Error parsing JSON result for type %T [%v]", loginResult, jsonErr))
			    return resp, jsonErr
			}
			// fmt.Println("LOGIN RESPONSE: ", resp, err)
			// fmt.Println("PARSED LOGIN RESULT: ", loginResult)
			
			if (loginResult.AccessToken != "") {
				client.SetAccessToken(loginResult.AccessToken, loginResult.RefreshToken, loginResult.ExpiresIn, loginResult.Scope)
				logDebug(fmt.Sprintf("Logged in as %v @ %v", client.Username, client.Url))
				// fmt.Println("Access Token: " + client.AccessToken)
			} else {
				err = errors.New("Login failed, unable to parse access token from login response")
				logError(err)
			}
			// client.setLastLoginResult(loginResult)
			return resp, err
		} else {
			//fmt.Println("Login Failure:", resp.RestyResponse)
			//err = errors.New("Login error: %v", resp.RestyResponse)
			// err = errors.New("Login error: %v", resp.Body)
			// need to parse with LoginErrorResult which has msg: "blah blah"
			return resp, err
		}
		
	}

	//return resp, err
}

func (client * Client) Logout() (*Response, error) {
	client.ClearAccessToken()
	// client.AccessToken = ""
	// client.RefreshToken = ""
	// client.ExpiresIn = 0
	// client.Scope = ""
	client.Username = ""
	client.Password = ""
	// there is no serverside endpoint for this right now
	// create mock response
	resp := &Response{Success:true,Status:"200 OK",StatusCode:200}
	return resp, nil
}
