package morpheus

import (
	"fmt"
)

var (
	// NetworksPath is the API endpoint for networks
	NetworksPath = "/api/networks"
)

// Network structures for use in request and response payloads
type Network struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Visibility  string `json:"visibility"`
}

type ListNetworksResult struct {
	Networks *[]Network  `json:"networks"`
	Meta     *MetaResult `json:"meta"`
}

type GetNetworkResult struct {
	Network *Network `json:"network"`
}

type CreateNetworkResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Network *Network          `json:"network"`
}

type UpdateNetworkResult struct {
	CreateNetworkResult
}

type DeleteNetworkResult struct {
	DeleteResult
}

type NetworkPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Visibility  string `json:"visibility"`
}

type CreateNetworkPayload struct {
	NetworkPayload *NetworkPayload `json:"network"`
}

type UpdateNetworkPayload struct {
	CreateNetworkPayload
}

func (client *Client) ListNetworks(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        NetworksPath,
		QueryParams: req.QueryParams,
		Result:      &ListNetworksResult{},
	})
}

func (client *Client) GetNetwork(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", NetworksPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetNetworkResult{},
	})
}

func (client *Client) CreateNetwork(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        NetworksPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateNetworkResult{},
	})
}

func (client *Client) UpdateNetwork(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", NetworksPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateNetworkResult{},
	})
}

func (client *Client) DeleteNetwork(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", NetworksPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteNetworkResult{},
	})
}

func (client *Client) FindNetworkByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListNetworks(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListNetworksResult)
	networkCount := len(*listResult.Networks)
	if networkCount != 1 {
		return resp, fmt.Errorf("found %d networks for %v", networkCount, name)
	}
	firstRecord := (*listResult.Networks)[0]
	networkID := firstRecord.ID
	return client.GetNetwork(networkID, &Request{})
}
