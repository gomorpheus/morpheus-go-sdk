package morpheus

import (
	"fmt"
)

var (
	// TenantsPath is the API endpoint for tenants
	TenantsPath = "/api/accounts"
)

// Tenant structures for use in request and response payloads
type Tenant struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Subdomain   string `json:"subdomain"`
	Role        struct {
		ID          int64  `json:"id"`
		Authority   string `json:"authority"`
		Description string `json:"description"`
	} `json:"role"`
	Active         bool   `json:"active"`
	CustomerNumber string `json:"customerNumber"`
	AccountNumber  string `json:"accountNumber"`
	Currency       string `json:"currency"`
	AccountName    string `json:"accountName"`
	Stats          struct {
		InstanceCount int64 `json:"instanceCount"`
		UserCount     int64 `json:"userCount"`
	} `json:"stats"`
	DateCreated string `json:"dateCreated"`
	LastUpdated string `json:"lastUpdated"`
}

// ListTenantsResult structure parses the list tenants response payload
type ListTenantsResult struct {
	Accounts *[]Tenant   `json:"accounts"`
	Meta     *MetaResult `json:"meta"`
}

// GetTenantResult structure parses the get tenant response payload
type GetTenantResult struct {
	Tenant *Tenant `json:"account"`
}

// CreateTenantResult structure parses the create tenant response payload
type CreateTenantResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Tenant  *Tenant           `json:"account"`
}

// UpdateTenantResult structure parses the update tenant response payload
type UpdateTenantResult struct {
	CreateTenantResult
}

// DeleteTenantResult structure parses the delete tenant response payload
type DeleteTenantResult struct {
	DeleteResult
}

// Client request methods

// ListTenants lists all tenants
func (client *Client) ListTenants(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        TenantsPath,
		QueryParams: req.QueryParams,
		Result:      &ListTenantsResult{},
	})
}

// GetTenant gets a single tenant by id
func (client *Client) GetTenant(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", TenantsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetTenantResult{},
	})
}

// CreateTenant creates a new Morpheus tenant
func (client *Client) CreateTenant(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        TenantsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateTenantResult{},
	})
}

// UpdateTenant updates an existing Morpheus tenant
func (client *Client) UpdateTenant(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", TenantsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateTenantResult{},
	})
}

// DeleteTenant deletes an existing Morpheus tenant
func (client *Client) DeleteTenant(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", TenantsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteTenantResult{},
	})
}

// FindTenantByName gets an existing tenant by name
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
		return resp, fmt.Errorf("found %d Tenants for %v", tenantsCount, name)
	}
	firstRecord := (*listResult.Accounts)[0]
	tenantID := firstRecord.ID
	return client.GetTenant(tenantID, &Request{})
}
