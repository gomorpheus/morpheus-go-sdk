package morpheus

import (
	"fmt"
)

var (
	// NetworkPoolsPath is the API endpoint for network pools
	NetworkPoolsPath = "/api/networks/pools"
)

// NetworkPool structures for use in request and response payloads
type NetworkPool struct {
	ID   int64 `json:"id"`
	Type struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"type"`
	Account struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"account"`
	Category      interface{} `json:"category"`
	Code          string      `json:"code"`
	Name          string      `json:"name"`
	DisplayName   string      `json:"displayName"`
	InternalId    interface{} `json:"internalId"`
	ExternalId    string      `json:"externalId"`
	DnsDomain     string      `json:"dnsDomain"`
	DnsSearchPath string      `json:"dnsSearchPath"`
	HostPrefix    interface{} `json:"hostPrefix"`
	HttpProxy     interface{} `json:"httpProxy"`
	DnsServers    []string    `json:"dnsServers"`
	DnsSuffixlist []string    `json:"dnsSuffixList"`
	DhcpServer    bool        `json:"dhcpServer"`
	DhcpIp        interface{} `json:"dhcpIp"`
	Gateway       string      `json:"gateway"`
	Netmask       string      `json:"netmask"`
	SubnetAddress string      `json:"subnetAddress"`
	IpCount       int64       `json:"ipCount"`
	FreeCount     int64       `json:"freeCount"`
	PoolEnabled   bool        `json:"poolEnabled"`
	TftpServer    interface{} `json:"tftpServer"`
	BootFile      interface{} `json:"bootFile"`
	RefType       string      `json:"refType"`
	RefId         string      `json:"refId"`
	ParentType    string      `json:"parentType"`
	ParentId      string      `json:"parentId"`
	PoolGroup     interface{} `json:"poolGroup"`
	IpRanges      []struct {
		ID           int64       `json:"id"`
		StartAddress string      `json:"startAddress"`
		EndAddress   string      `json:"endAddress"`
		InternalId   interface{} `json:"internalId"`
		ExternalId   interface{} `json:"externalId"`
		Description  interface{} `json:"description"`
		AddressCount int64       `json:"addressCount"`
		Active       bool        `json:"active"`
		DateCreated  string      `json:"dateCreated"`
		LastUpdated  string      `json:"lastUpdated"`
		Cidr         interface{} `json:"cidr"`
	} `json:"ipRanges"`
}

// ListNetworkPoolsResult structure parses the list network pools response payload
type ListNetworkPoolsResult struct {
	NetworkPools *[]NetworkPool `json:"networkPools"`
	Meta         *MetaResult    `json:"meta"`
}

type GetNetworkPoolResult struct {
	NetworkPool *NetworkPool `json:"networkPool"`
}

type CreateNetworkPoolResult struct {
	Success     bool              `json:"success"`
	Message     string            `json:"msg"`
	Errors      map[string]string `json:"errors"`
	NetworkPool *NetworkPool      `json:"networkPool"`
}

type UpdateNetworkPoolResult struct {
	CreateNetworkPoolResult
}

type DeleteNetworkPoolResult struct {
	DeleteResult
}

// ListNetworkPools lists all network pools
func (client *Client) ListNetworkPools(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        NetworkPoolsPath,
		QueryParams: req.QueryParams,
		Result:      &ListNetworkPoolsResult{},
	})
}

// GetNetworkPool gets an existing network pool
func (client *Client) GetNetworkPool(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", NetworkPoolsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetNetworkPoolResult{},
	})
}

// CreateNetworkPool creates a new network pool
func (client *Client) CreateNetworkPool(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        NetworkPoolsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateNetworkPoolResult{},
	})
}

// UpdateNetworkPool updates an existing network pool
func (client *Client) UpdateNetworkPool(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", NetworkPoolsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateNetworkPoolResult{},
	})
}

// DeleteNetworkPool deletes an existing network pool
func (client *Client) DeleteNetworkPool(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", NetworkPoolsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteNetworkPoolResult{},
	})
}

// FindNetworkPoolByName gets an existing network pool by name
func (client *Client) FindNetworkPoolByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListNetworkPools(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListNetworkPoolsResult)
	networkPoolsCount := len(*listResult.NetworkPools)
	if networkPoolsCount != 1 {
		return resp, fmt.Errorf("found %d Network Pools for %v", networkPoolsCount, name)
	}
	firstRecord := (*listResult.NetworkPools)[0]
	networkPoolID := firstRecord.ID
	return client.GetNetworkPool(networkPoolID, &Request{})
}
