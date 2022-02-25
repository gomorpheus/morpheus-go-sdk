package morpheus

import (
	"fmt"
)

var (
	// CloudsPath is the API endpoint for clouds (zones)
	CloudsPath = "/api/zones"
)

// Cloud structures for use in request and response payloads
type Cloud struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Location  string    `json:"location"`
	CloudType CloudType `json:"zoneType"`
	//Active bool `json:"active"`
	Enabled    bool                   `json:"enabled"`
	Visibility string                 `json:"visibility"`
	Config     map[string]interface{} `json:"config"`
	// Clouds []string `json:"clouds"` //todo

}

type CloudType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// ListCloudsResult structure parses the list clouds response payload
type ListCloudsResult struct {
	Clouds *[]Cloud    `json:"zones"`
	Meta   *MetaResult `json:"meta"`
}

type GetCloudResult struct {
	Cloud *Cloud `json:"zone"`
}

type CreateCloudResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Cloud   *Cloud            `json:"zone"`
}

type UpdateCloudResult struct {
	CreateCloudResult
}

type DeleteCloudResult struct {
	DeleteResult
}

// API endpoints

func (client *Client) ListClouds(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        CloudsPath,
		QueryParams: req.QueryParams,
		Result:      &ListCloudsResult{},
	})
}

func (client *Client) GetCloud(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", CloudsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetCloudResult{},
	})
}

// CreateCloud creates a new cloud
func (client *Client) CreateCloud(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        CloudsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateCloudResult{},
	})
}

// UpdateCloud updates an existing cloud
func (client *Client) UpdateCloud(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", CloudsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateCloudResult{},
	})
}

// DeleteCloud deletes an existing cloud
func (client *Client) DeleteCloud(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", CloudsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteCloudResult{},
	})
}

func (client *Client) FindCloudByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListClouds(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListCloudsResult)
	cloudsCount := len(*listResult.Clouds)
	if cloudsCount != 1 {
		return resp, fmt.Errorf("found %d Clouds for %v", cloudsCount, name)
	}
	firstRecord := (*listResult.Clouds)[0]
	cloudId := firstRecord.ID
	return client.GetCloud(cloudId, &Request{})
}
