// Morpheus API types and Client methods for Tenants
package morpheus

import (
	"fmt"
)

// globals

var (
	TenantsPath = "/api/accounts"
)

// Tenant structures for use in request and response payloads

type Tenant struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Active         bool   `json:"active"`
	CustomerNumber int64  `json:"customerNumber"`
	AccountNumber  int64  `json:"accountNumber"`
	Currency       string `json:"currency"`
}

type ListTenantsResult struct {
	Accounts *[]Tenant   `json:"accounts"`
	Meta     *MetaResult `json:"meta"`
}

type GetTenantResult struct {
	Tenant *Tenant `json:"account"`
}

type CreateTenantResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Tenant  *Tenant           `json:"account"`
}

type UpdateTenantResult struct {
	CreateTenantResult
}

type DeleteTenantResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListTenants(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        TenantsPath,
		QueryParams: req.QueryParams,
		Result:      &ListTenantsResult{},
	})
}

func (client *Client) GetTenant(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", TenantsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetTenantResult{},
	})
}

func (client *Client) CreateTenant(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        TenantsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateTenantResult{},
	})
}

func (client *Client) UpdateTenant(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", TenantsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateTenantResult{},
	})
}

func (client *Client) DeleteTenant(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", TenantsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteTenantResult{},
	})
}

// helper functions

func (client *Client) FindTenantByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListTenants(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListTenantsResult)
	tenantsCount := len(*listResult.Accounts)
	if tenantsCount != 1 {
		return resp, fmt.Errorf("Found %d Tenants for %v", tenantsCount, name)
	}
	firstRecord := (*listResult.Accounts)[0]
	tenantID := firstRecord.ID
	return client.GetTenant(tenantID, &Request{})
}
