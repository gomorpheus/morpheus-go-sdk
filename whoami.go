// Client is the driver for interfacing with the Morpheus API
package morpheus

import (
    "fmt"
    "errors"
	_ "encoding/json"
)

var (
	WhoamiPath = "/api/whoami"
)

type WhoamiResult struct {
    User *WhoamiUserResult `json:"user"`
    IsMasterAccount bool `json:"isMasterAccount"`
    Permissions map[string]string `json:"permissions"`
    Appliance WhoamiApplianceResult `json:"appliance"`
}

type WhoamiUserResult struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}
type WhoamiApplianceResult struct {
	BuildVersion string `json:"buildVersion"`
}

func (client * Client) Whoami() (*Response, error) {
	// login if needed
	//var resp Response
	// var resp *Response
	// var err error
	
	// if (client.IsLoggedIn() == false) {
	// 	resp, err := client.Login()
	// 	if err != nil {
	// 		return resp, err
	// 	}
	// }

	whoamiRequest := &Request{
		Method: "GET",
		Path: WhoamiPath,
		Timeout: 10,
		Result: &WhoamiResult{},
	}

	resp, err := client.Execute(whoamiRequest)
	
	// parse JSON response
	// var whoamiResult WhoamiResult
	// err = json.Unmarshal(resp.Body, &whoamiResult)

	whoamiResult := resp.Result.(*WhoamiResult)

	// if err != nil {
	//     fmt.Println(err)
	// }

	if (whoamiResult.User != nil) {
		if  whoamiResult.User.Username != client.Username {
			//currentAccessToken := client.AccessToken
			//client.SetUsername(whoamiResult.User.Username)
			//client.SetAccessToken(currentAccessToken, client.RefreshToken, client.ExpiresIn, client.Scope)
			//client.AccessToken = currentAccessToken
		}
		// fmt.Println("Whoami Success")
		//fmt.Println("Username: " + whoamiResult.User.Username)
		//fmt.Println("Name: " + whoamiResult.User.FirstName + " " + whoamiResult.User.LastName)
		var perms []string
		for k,v := range whoamiResult.Permissions {
			perms = append(perms, fmt.Sprintf("%v (%v)", k, v))
		}
		//fmt.Println("Permissions: " + strings.Join(perms, ", "))
		// fmt.Println("Access Token: " + client.AccessToken)
	} else {
		err = errors.New("Unable to parse whoami result from api response")
		fmt.Println(err)
	}
	
	// client.setLastWhoamiResult(whoamiResult)
	
	return resp, err
	

	//return resp, err
}

