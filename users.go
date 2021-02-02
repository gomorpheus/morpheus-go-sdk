// Morpheus API types and Client methods for Tenants
package morpheus

import (
	"fmt"
)

// globals

var (
	UsersPath = "/api/accounts"
)

// Tenant structures for use in request and response payloads

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type ListUsersResult struct {
	Accounts *[]Network  `json:"accounts"`
	Meta     *MetaResult `json:"meta"`
}

type GetUserResult struct {
	Tenant *Tenant `json:"tenant"`
}

type CreateUserResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Network *Network          `json:"network"`
}

type UpdateUserResult struct {
	CreateNetworkResult
}

type DeleteUserResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListUsers(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        TenantsPath,
		QueryParams: req.QueryParams,
		Result:      &ListUsersResult{},
	})
}

func (client *Client) GetUser(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", TenantsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetUserResult{},
	})
}

func (client *Client) CreateUser(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        TenantsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateUserResult{},
	})
}

func (client *Client) UpdateUser(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", TenantsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateUserResult{},
	})
}

func (client *Client) DeleteUser(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", TenantsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteUserResult{},
	})
}
