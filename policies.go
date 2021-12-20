package morpheus

// Morpheus API types and Client methods for Policies

import (
	"fmt"
)

// globals
var (
	PoliciesPath = "/api/policies"
)

// Policy structures for use in request and response payloads
type Policy struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	PolicyType struct {
		ID   int64  `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"type"`
}

type ListPoliciesResult struct {
	Policies *[]Policy   `json:"policies"`
	Meta     *MetaResult `json:"meta"`
}

type GetPolicyResult struct {
	Policy *Policy `json:"policy"`
}

type CreatePolicyResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Policy  *Policy           `json:"policy"`
}

type UpdatePolicyResult struct {
	CreatePolicyResult
}

type DeletePolicyResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListPolicies(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        PoliciesPath,
		QueryParams: req.QueryParams,
		Result:      &ListPoliciesResult{},
	})
}

func (client *Client) GetPolicy(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", PoliciesPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetPolicyResult{},
	})
}

func (client *Client) CreatePolicy(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        PoliciesPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreatePolicyResult{},
	})
}

func (client *Client) UpdatePolicy(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", PoliciesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdatePolicyResult{},
	})
}

func (client *Client) DeletePolicy(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", PoliciesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeletePolicyResult{},
	})
}

// helper functions

func (client *Client) FindPolicyByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListPolicies(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListPoliciesResult)
	policyCount := len(*listResult.Policies)
	if policyCount != 1 {
		return resp, fmt.Errorf("found %d policies named %v", policyCount, name)
	}
	firstRecord := (*listResult.Policies)[0]
	policyID := firstRecord.ID
	return client.GetPolicy(policyID, &Request{})
}
