package morpheus

import (
	"fmt"
)

var (
	// ClustersPath is the API endpoint for clusters
	ClustersPath = "/api/clusters"
)

// Cluster structures for use in request and response payloads
type Cluster struct {
	ID          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Type        interface{}            `json:"type"`
	Layout      interface{}            `json:"layout"`
	Group       map[string]interface{} `json:"group"`
	Cloud       map[string]interface{} `json:"cloud"`
	Server      Server                 `json:"server"`
	Status      string                 `json:"status"`
}

type Server struct {
	Config            map[string]interface{}    `json:"config"`
	Name              string                    `json:"name"`
	HostName          string                    `json:"hostname"`
	Plan              map[string]interface{}    `json:"plan"`
	Volumes           *[]map[string]interface{} `json:"volumes"`
	NetworkInterfaces *[]map[string]interface{} `json:"networkInterfaces"`
	Visibility        string                    `json:"visibility"`
	NodeCount         int64                     `json:"nodeCount"`
}

// ListClustersResult structure parses the list clusters response payload
type ListClustersResult struct {
	Clusters *[]Cluster  `json:"clusters"`
	Meta     *MetaResult `json:"meta"`
}

type GetClusterResult struct {
	Cluster *Cluster `json:"cluster"`
}

type CreateClusterResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Cluster *Cluster          `json:"cluster"`
}

type UpdateClusterResult struct {
	CreateClusterResult
}

type DeleteClusterResult struct {
	DeleteResult
}

// API endpoints

func (client *Client) ListClusters(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        ClustersPath,
		QueryParams: req.QueryParams,
		Result:      &ListClustersResult{},
	})
}

func (client *Client) GetCluster(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", ClustersPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetClusterResult{},
	})
}

// CreateCluster creates a new cluster
func (client *Client) CreateCluster(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        ClustersPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateClusterResult{},
	})
}

// UpdateCluster udpates an existing cluster
func (client *Client) UpdateCluster(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", ClustersPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateClusterResult{},
	})
}

// DeleteCluster deletes an existing cluster
func (client *Client) DeleteCluster(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", ClustersPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteClusterResult{},
	})
}

func (client *Client) FindClusterByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListClusters(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListClustersResult)
	clustereCount := len(*listResult.Clusters)
	if clustereCount != 1 {
		return resp, fmt.Errorf("found %d Clusters for %v", clustereCount, name)
	}
	firstRecord := (*listResult.Clusters)[0]
	clustereId := firstRecord.ID
	return client.GetCluster(clustereId, &Request{})
}
