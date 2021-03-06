// Morpheus API types and Client methods for Instance Types
package morpheus

import (
	"fmt"
)

// globals
var (
	InstanceTypesPath = "/api/library/instance-types"
)

// InstanceType structures for use in request and response payloads

type InstanceType struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Active      bool   `json:"active"`
	Visibility  string `json:"visibility"`
}

type ListInstanceTypesResult struct {
	InstanceTypes *[]InstanceType `json:"instanceTypes"`
	Meta          *MetaResult     `json:"meta"`
}

type GetInstanceTypeResult struct {
	InstanceType *InstanceType `json:"instanceType"`
}

type CreateInstanceTypeResult struct {
	Success      bool              `json:"success"`
	Message      string            `json:"msg"`
	Errors       map[string]string `json:"errors"`
	InstanceType *InstanceType     `json:"instanceType"`
}

type UpdateInstanceTypeResult struct {
	CreateInstanceTypeResult
}

type DeleteInstanceTypeResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListInstanceTypes(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        InstanceTypesPath,
		QueryParams: req.QueryParams,
		Result:      &ListInstanceTypesResult{},
	})
}

func (client *Client) GetInstanceType(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", InstanceTypesPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetInstanceTypeResult{},
	})
}

func (client *Client) CreateInstanceType(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        InstanceTypesPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateInstanceTypeResult{},
	})
}

func (client *Client) UpdateInstanceType(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", InstanceTypesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateInstanceTypeResult{},
	})
}

func (client *Client) DeleteInstanceType(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", InstanceTypesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteInstanceTypeResult{},
	})
}

// helper functions
func (client *Client) FindInstanceTypeByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListInstanceTypes(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListInstanceTypesResult)
	instanceTypesCount := len(*listResult.InstanceTypes)
	if instanceTypesCount != 1 {
		return resp, fmt.Errorf("found %d InstanceTypes for %v", instanceTypesCount, name)
	}
	firstRecord := (*listResult.InstanceTypes)[0]
	instanceTypeId := firstRecord.ID
	return client.GetInstanceType(instanceTypeId, &Request{})
}
