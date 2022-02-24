// Morpheus API types and Client methods for Option Types
package morpheus

import (
	"fmt"
)

// globals

var (
	ChecksPath = "/api/monitoring/checks"
)

// Check structures for use in request and response payloads

type Check struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	CheckInterval int64  `json:"checkInterval"`
	CheckType     struct {
		ID   int64  `json:"id"`
		Code string `json:"code"`
	} `json:"checkType"`
	InUptime bool        `json:"inUptime"`
	Active   bool        `json:"active"`
	Severity string      `json:"severity"`
	Config   interface{} `json:"config"`
}

type ListChecksResult struct {
	Checks *[]Check    `json:"checks"`
	Meta   *MetaResult `json:"meta"`
}

type GetCheckResult struct {
	Check *Check `json:"check"`
}

type CreateCheckResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Check   *Check            `json:"check"`
}

type UpdateCheckResult struct {
	CreateCheckResult
}

type DeleteCheckResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListChecks(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        ChecksPath,
		QueryParams: req.QueryParams,
		Result:      &ListChecksResult{},
	})
}

func (client *Client) GetCheck(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", ChecksPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetCheckResult{},
	})
}

func (client *Client) CreateCheck(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        ChecksPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateCheckResult{},
	})
}

func (client *Client) UpdateCheck(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", ChecksPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateCheckResult{},
	})
}

func (client *Client) DeleteCheck(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", ChecksPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteCheckResult{},
	})
}

// helper functions
func (client *Client) FindCheckByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListChecks(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListChecksResult)
	checkCount := len(*listResult.Checks)
	if checkCount != 1 {
		return resp, fmt.Errorf("found %d Checks for %v", checkCount, name)
	}
	firstRecord := (*listResult.Checks)[0]
	checkID := firstRecord.ID
	return client.GetCheck(checkID, &Request{})
}
