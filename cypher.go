package morpheus

import (
	"fmt"
	"time"
)

var (
	// CypherPath is the API endpoint for cypher
	CypherPath = "/api/cypher"
)

// Cypher structures for use in request and response payloads
type Cypher struct {
	ID           int64     `json:"id"`
	ItemKey      string    `json:"itemKey"`
	LeaseTimeout int64     `json:"leaseTimeout"`
	ExpireDate   time.Time `json:"expireDate"`
	DateCreated  time.Time `json:"dateCreated"`
	LastUpdated  time.Time `json:"lastUpdated"`
	LastAccessed time.Time `json:"lastAccessed"`
	CreatedBy    string    `json:"createdBy"`
}

type CypherData struct {
	Keys []string `json:"keys"`
}

type ListCypherResult struct {
	Success bool        `json:"success"`
	Data    CypherData  `json:"data"`
	Cyphers *[]Cypher   `json:"cyphers"`
	Meta    *MetaResult `json:"meta"`
}

type GetCypherResult struct {
	Cypher *Cypher `json:"cypher"`
}

type CreateCypherResult struct {
	Success bool              `json:"success"`
	Data    CypherData        `json:"data"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Cypher  *Cypher           `json:"cypher"`
}

type DeleteCypherResult struct {
	DeleteResult
}

// Client request methods
func (client *Client) ListCyphers(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        CypherPath,
		QueryParams: req.QueryParams,
		Result:      &ListCypherResult{},
	})
}

func (client *Client) GetCypher(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", CypherPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetCypherResult{},
	})
}

func (client *Client) CreateCypher(path string, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        fmt.Sprintf("%s/%s", CypherPath, path),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateCypherResult{},
	})
}

func (client *Client) DeleteCypher(path string, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%s", CypherPath, path),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteCypherResult{},
	})
}

// FindCredentialByName gets an existing credential by name
func (client *Client) FindCypherByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListCyphers(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListCypherResult)
	cypherCount := len(*listResult.Cyphers)
	if cypherCount != 1 {
		return resp, fmt.Errorf("found %d cyphers for %v", cypherCount, name)
	}
	firstRecord := (*listResult.Cyphers)[0]
	cypherID := firstRecord.ID
	return client.GetCypher(cypherID, &Request{})
}
