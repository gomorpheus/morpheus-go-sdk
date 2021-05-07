// Morpheus API types and Client methods for Groups (Sites)
package morpheus

import (
	"fmt"
)

// globals

var (
	GroupsPath = "/api/groups"
)

// types

type Group struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Location string `json:"location"`
	Clouds   []Zone `json:"zones"`
}

type Zone struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ListGroupsResult struct {
	Groups *[]Group    `json:"groups"`
	Meta   *MetaResult `json:"meta"`
}

type GetGroupResult struct {
	Group *Group `json:"group"`
}

type CreateGroupResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Group   *Group            `json:"group"`
}

type UpdateGroupResult struct {
	CreateGroupResult
}

type DeleteGroupResult struct {
	DeleteResult
}

// API endpoints

func (client *Client) ListGroups(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        GroupsPath,
		QueryParams: req.QueryParams,
		Result:      &ListGroupsResult{},
	})
}

func (client *Client) GetGroup(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", GroupsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetGroupResult{},
	})
}

func (client *Client) CreateGroup(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        GroupsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateGroupResult{},
	})
}

func (client *Client) UpdateGroup(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", GroupsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateGroupResult{},
	})
}

func (client *Client) UpdateGroupClouds(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d/update-zones", GroupsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &StandardResult{},
	})
}

func (client *Client) UpdateGroupZones(id int64, req *Request) (*Response, error) {
	return client.UpdateGroupClouds(id, req)
}

func (client *Client) DeleteGroup(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", GroupsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteGroupResult{},
	})
}

// helper functions

func (client *Client) FindGroupByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListGroups(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListGroupsResult)
	groupsCount := len(*listResult.Groups)
	if groupsCount != 1 {
		return resp, fmt.Errorf("Found %d Groups for %v", groupsCount, name)
	}
	firstRecord := (*listResult.Groups)[0]
	groupId := firstRecord.ID
	return client.GetGroup(groupId, &Request{})
}
