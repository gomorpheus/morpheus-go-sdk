package morpheus

// Morpheus API types and Client methods for Node Types

import (
	"fmt"
)

// globals
var (
	NodeTypesPath = "/api/library/container-types"
)

// NodeType structures for use in request and response payloads
type NodeType struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Code      string `json:"code"`
}

type ListNodeTypesResult struct {
	NodeTypes *[]NodeType `json:"containerTypes"`
	Meta      *MetaResult `json:"meta"`
}

type GetNodeTypeResult struct {
	NodeType *NodeType `json:"containerType"`
}

type CreateNodeTypeResult struct {
	Success  bool              `json:"success"`
	Message  string            `json:"msg"`
	Errors   map[string]string `json:"errors"`
	NodeType *NodeType         `json:"nodeType"`
}

type UpdateNodeTypeResult struct {
	CreateNodeTypeResult
}

type DeleteNodeTypeResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListNodeTypes(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        NodeTypesPath,
		QueryParams: req.QueryParams,
		Result:      &ListNodeTypesResult{},
	})
}

func (client *Client) GetNodeType(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", NodeTypesPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetNodeTypeResult{},
	})
}

func (client *Client) CreateNodeType(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        NodeTypesPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateNodeTypeResult{},
	})
}

func (client *Client) UpdateNodeType(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", NodeTypesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateNodeTypeResult{},
	})
}

func (client *Client) DeleteNodeType(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", NodeTypesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteNodeTypeResult{},
	})
}

// helper functions

func (client *Client) FindNodeTypeByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListNodeTypes(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListNodeTypesResult)
	nodeTypeCount := len(*listResult.NodeTypes)
	if nodeTypeCount != 1 {
		return resp, fmt.Errorf("found %d node types named %v", nodeTypeCount, name)
	}
	firstRecord := (*listResult.NodeTypes)[0]
	nodeTypeID := firstRecord.ID
	return client.GetNodeType(nodeTypeID, &Request{})
}
