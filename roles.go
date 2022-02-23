// Morpheus API types and Client methods for Tenants
package morpheus

import (
	"fmt"
)

// globals

var (
	RolesPath       = "/api/roles"
	TenantRolesPath = "/api/accounts/available-roles"
)

// Role structures for use in request and response payloads

type Role struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Authority   string      `json:"authority"`
	RoleType    string      `json:"roleType"`
	Owner       interface{} `json:"owner"`
}

type ListRolesResult struct {
	Roles *[]Role     `json:"roles"`
	Meta  *MetaResult `json:"meta"`
}

type GetRoleResult struct {
	Role *Role `json:"role"`
}

type CreateRoleResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Role    *Role             `json:"role"`
}

type UpdateRoleResult struct {
	CreateRoleResult
}

type DeleteRoleResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListRoles(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        RolesPath,
		QueryParams: req.QueryParams,
		Result:      &ListRolesResult{},
	})
}

func (client *Client) ListTenantRoles(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        TenantRolesPath,
		QueryParams: req.QueryParams,
		Result:      &ListRolesResult{},
	})
}

func (client *Client) GetRole(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", RolesPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetRoleResult{},
	})
}

func (client *Client) CreateRole(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        RolesPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateRoleResult{},
	})
}

func (client *Client) UpdateRole(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", RolesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateRoleResult{},
	})
}

func (client *Client) DeleteRole(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", RolesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteRoleResult{},
	})
}

// helper functions

func (client *Client) FindRoleByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListRoles(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListRolesResult)
	rolesCount := len(*listResult.Roles)
	if rolesCount != 1 {
		return resp, fmt.Errorf("found %d Roles for %v", rolesCount, name)
	}
	firstRecord := (*listResult.Roles)[0]
	roleID := firstRecord.ID
	return client.GetRole(roleID, &Request{})
}

func (client *Client) FindTenantRoleByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListTenantRoles(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListRolesResult)
	for _, role := range *listResult.Roles {
		if role.Authority == name {
			return client.GetRole(role.ID, &Request{})
		}
	}
	return resp, fmt.Errorf("not matching role found for %v", name)
}
