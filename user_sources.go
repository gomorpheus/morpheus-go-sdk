// Morpheus API types and Client methods for Tenants
package morpheus

import ()

// globals
var (
	UserSourcesPath     = "/api/accounts"
	UserSourcesPathList = "/api/user-sources"
)

type UserSource struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type CreateUserSourceResult struct {
	Success    bool              `json:"success"`
	Message    string            `json:"msg"`
	Errors     map[string]string `json:"errors"`
	UserSource *UserSource       `json:"userSource"`
}

type UpdateUserSourceResult struct {
	CreateUserSourceResult
}

type DeleteUserSourceResult struct {
	DeleteResult
}

type ListUserSourcesResult struct {
	UserSources *[]UserSource `json:"userSource"`
	Meta        *MetaResult   `json:"meta"`
}

// Client request methods

func (client *Client) ListUserSources(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        UserSourcesPathList,
		QueryParams: req.QueryParams,
		Result:      &ListUserSourcesResult{},
	})
}

/*
func (client *Client) GetNetworkDomain(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", NetworkDomainsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetNetworkDomainResult{},
	})
}

func (client *Client) CreateNetworkDomain(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        NetworkDomainsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateNetworkDomainResult{},
	})
}

func (client *Client) UpdateNetworkDomain(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", NetworkDomainsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateNetworkDomainResult{},
	})
}

func (client *Client) DeleteNetworkDomain(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", NetworkDomainsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteNetworkDomainResult{},
	})
}
*/
