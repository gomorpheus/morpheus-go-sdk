// Morpheus API types and Client methods for Option Types
package morpheus

import (
	"fmt"
)

// globals

var (
	CheckGroupsPath = "/api/monitoring/groups"
)

// CheckGroup structures for use in request and response payloads

type CheckGroup struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	CheckInterval string `json:"checkInterval"`
	CheckType     struct {
		ID   int64  `json:"id"`
		Code string `json:"code"`
	} `json:"checkType"`
	InUptime bool        `json:"inUptime"`
	Active   bool        `json:"active"`
	Severity string      `json:"severity"`
	Config   interface{} `json:"config"`
}

type ListCheckGroupsResult struct {
	CheckGroups *[]CheckGroup `json:"checkGroups"`
	Meta        *MetaResult   `json:"meta"`
}

type GetCheckGroupResult struct {
	CheckGroup *CheckGroup `json:"checkGroup"`
}

type CreateCheckGroupResult struct {
	Success    bool              `json:"success"`
	Message    string            `json:"msg"`
	Errors     map[string]string `json:"errors"`
	CheckGroup *CheckGroup       `json:"checkGroup"`
}

type UpdateCheckGroupResult struct {
	CreateCheckGroupResult
}

type DeleteCheckGroupResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListCheckGroups(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        CheckGroupsPath,
		QueryParams: req.QueryParams,
		Result:      &ListCheckGroupsResult{},
	})
}

func (client *Client) GetCheckGroup(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", CheckGroupsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetCheckGroupResult{},
	})
}

func (client *Client) CreateCheckGroup(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        CheckGroupsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateCheckGroupResult{},
	})
}

func (client *Client) UpdateCheckGroup(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", CheckGroupsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateCheckGroupResult{},
	})
}

func (client *Client) DeleteCheckGroup(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", CheckGroupsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteCheckGroupResult{},
	})
}

// helper functions
func (client *Client) FindCheckGroupByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListCheckGroups(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListCheckGroupsResult)
	checkGroupCount := len(*listResult.CheckGroups)
	if checkGroupCount != 1 {
		return resp, fmt.Errorf("found %d Check Groups for %v", checkGroupCount, name)
	}
	firstRecord := (*listResult.CheckGroups)[0]
	checkGroupID := firstRecord.ID
	return client.GetCheckGroup(checkGroupID, &Request{})
}
