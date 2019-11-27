// Morpheus API types and Client methods for Networks
package morpheus

import (
    "fmt"
)

var (
	NetworksPath = "/api/networks"
)

// Network structures for use in request and response payloads

type Network struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Active bool `json:"active"`
	Visibility string `json:"visibility"`
	// what  else?
	// Account *TenantAbbrev `json:"account"`
	// Owner *TenantAbbrev `json:"owner"`
	// RefSource string `json:"refSource"`
	// RefType string `json:"refType"`
	// RefId int64 `json:"refId"`
}

type ListNetworksResult struct {
    Networks *[]Network `json:"networks"`
    Meta *MetaResult `json:"meta"`
}

type GetNetworkResult struct {
    Network *Network `json:"network"`
}

type CreateNetworkResult struct {
	Success bool `json:"success"`
	Message string `json:"msg"`
	Errors map[string]string `json:"errors"`
	Network *Network `json:"network"`
}

type UpdateNetworkResult struct {
	CreateNetworkResult
}

type DeleteNetworkResult struct {
	DeleteResult
}

type NetworkPayload struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Active bool `json:"active"`
	Visibility string `json:"visibility"`
	// what  else?
	// Account *TenantAbbrev `json:"account"`
	// Owner *TenantAbbrev `json:"owner"`
}

type CreateNetworkPayload struct {
	NetworkPayload *NetworkPayload `json:"network"`
}

type UpdateNetworkPayload struct {
	CreateNetworkPayload
}

// Request types

// type ListNetworksRequest struct {
// 	Request
// }

// type GetNetworkRequest struct {
// 	Request
// 	ID int64
// }

// type CreateNetworkRequest struct {
// 	Request
// 	Payload *CreateNetworkPayload
// }

// type UpdateNetworkRequest struct {
// 	Request
// 	ID int64
// 	Payload *UpdateNetworkPayload
// }

// type DeleteNetworkRequest struct {
// 	Request
// 	ID int64
// }

// // Response types

// type ListNetworksResponse struct {
// 	Response
// 	Result *ListNetworksResult
// }

// type GetNetworkResponse struct {
// 	Response
// 	Result *GetNetworkResult
// }

// type CreateNetworkResponse struct {
// 	Response
// 	Result *CreateNetworkResult
// }

// type UpdateNetworkResponse struct {
// 	Response
// 	Result *UpdateNetworkResult
// }

// type DeleteNetworkResponse struct {
// 	Request
// 	Result *DeleteNetworkResult
// }

// Client request methods

func (client * Client) ListNetworks(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path: NetworksPath,
		QueryParams: req.QueryParams,
		Result: &ListNetworksResult{},
	})
}

func (client * Client) GetNetwork(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path: fmt.Sprintf("%s/%d", NetworksPath, id),
		QueryParams: req.QueryParams,
		Result: &GetNetworkResult{},
	})
}

func (client * Client) CreateNetwork(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "POST",
		Path: NetworksPath,
		QueryParams: req.QueryParams,
		Body: req.Body,
		Result: &CreateNetworkResult{},
	})
}

func (client * Client) UpdateNetwork(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "PUT",
		Path: fmt.Sprintf("%s/%d", NetworksPath, id),
		QueryParams: req.QueryParams,
		Body: req.Body,
		Result: &UpdateNetworkResult{},
	})
}


func (client * Client) DeleteNetwork(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "DELETE",
		Path: fmt.Sprintf("%s/%d", NetworksPath, id),
		QueryParams: req.QueryParams,
		Body: req.Body,
		Result: &DeleteNetworkResult{},
	})
}
