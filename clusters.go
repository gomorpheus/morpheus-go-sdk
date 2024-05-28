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
	ID                  int64    `json:"id"`
	Name                string   `json:"name"`
	Code                string   `json:"code"`
	Category            string   `json:"category"`
	Visibility          string   `json:"visibility"`
	Description         string   `json:"description"`
	Location            string   `json:"location"`
	Enabled             bool     `json:"enabled"`
	ServiceUrl          string   `json:"serviceUrl"`
	ServiceHost         string   `json:"serviceHost"`
	ServicePath         string   `json:"servicePath"`
	ServiceHostname     string   `json:"serviceHostname"`
	ServicePort         int64    `json:"servicePort"`
	ServiceUsername     string   `json:"serviceUsername"`
	ServicePassword     string   `json:"servicePassword"`
	ServicePasswordHash string   `json:"servicePasswordHash"`
	ServiceToken        string   `json:"serviceToken"`
	ServiceTokenHash    string   `json:"serviceTokenHash"`
	ServiceAccess       string   `json:"serviceAccess"`
	ServiceAccessHash   string   `json:"serviceAccessHash"`
	ServiceCert         string   `json:"serviceCert"`
	ServiceCertHash     string   `json:"serviceCertHash"`
	ServiceVersion      string   `json:"serviceVersion"`
	SearchDomains       string   `json:"searchDomains"`
	EnableInternalDns   bool     `json:"enableInternalDns"`
	InternalId          string   `json:"internalId"`
	ExternalId          string   `json:"externalId"`
	DatacenterId        string   `json:"datacenterId"`
	StatusMessage       string   `json:"statusMessage"`
	InventoryLevel      string   `json:"inventoryLevel"`
	LastSyncDuration    int64    `json:"lastSyncDuration"`
	Labels              []string `json:"labels"`
	Type                struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
	Layout struct {
		Id                int64  `json:"id"`
		Name              string `json:"name"`
		ProvisionTypeCode string `json:"provisionTypeCode"`
	} `json:"layout"`
	Group map[string]interface{} `json:"group"`
	Site  struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"site"`
	Zone struct {
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		ZoneType struct {
			Id int64 `json:"id"`
		} `json:"zoneType"`
	} `json:"zone"`
	Servers      []Server `json:"servers"`
	Status       string   `json:"status"`
	Managed      bool     `json:"managed"`
	ServiceEntry string   `json:"serviceEntry"`
	CreatedBy    struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
	} `json:"createdBy"`
	UserGroup string `json:"userGroup"`
	Owner     struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"owner"`
	WorkerStats struct {
		UsedStorage  int64   `json:"usedStorage"`
		MaxStorage   int64   `json:"maxStorage"`
		UsedMemory   int64   `json:"usedMemory"`
		MaxMemory    int64   `json:"maxMemory"`
		UsedCpu      float64 `json:"usedCpu"`
		CpuUsage     float64 `json:"cpuUsage"`
		CpuUsagePeak float64 `json:"cpuUsagePeak"`
		CpuUsageAvg  float64 `json:"cpuUsageAvg"`
	}
	ContainersCount  int64 `json:"containersCount"`
	DeploymentsCount int64 `json:"deploymentsCount"`
	PodsCount        int64 `json:"podsCount"`
	JobsCount        int64 `json:"jobsCount"`
	VolumesCount     int64 `json:"volumesCount"`
	NamespacesCount  int64 `json:"namespacesCount"`
	WorkersCount     int64 `json:"workersCount"`
	ServicesCount    int64 `json:"servicesCount"`
}

type Server struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	TypeSet struct {
		Id   int64  `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"typeSet"`
	ComputeServerType struct {
		Id       int64  `json:"id"`
		Code     string `json:"code"`
		NodeType string `json:"nodeType"`
	} `json:"computeServerType"`
}

// ListClustersResult structure parses the list clusters response payload
type ListClustersResult struct {
	Clusters *[]Cluster  `json:"clusters"`
	Meta     *MetaResult `json:"meta"`
}

type GetClusterResult struct {
	Cluster *Cluster `json:"cluster"`
}

type GetClusterApiConfigResult struct {
	ServiceUrl          string `json:"serviceUrl"`
	ServiceHost         string `json:"serviceHost"`
	ServicePath         string `json:"servicePath"`
	ServiceHostname     string `json:"serviceHostname"`
	ServicePort         int64  `json:"servicePort"`
	ServiceUsername     string `json:"serviceUsername"`
	ServicePassword     string `json:"servicePassword"`
	ServicePasswordHash string `json:"servicePasswordHash"`
	ServiceToken        string `json:"serviceToken"`
	ServiceAccess       string `json:"serviceAccess"`
	ServiceCert         string `json:"serviceCert"`
	ServiceVersion      string `json:"serviceVersion"`
}

type ListClusterNamespacesResults struct {
	Namespaces []Namespaces `json:"namespaces"`
	Meta       *MetaResult  `json:"meta"`
}

type Namespaces struct {
	Id                 int64  `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	RegionCode         string `json:"regionCode"`
	ExternalId         string `json:"externalId"`
	Status             string `json:"status"`
	Visibility         string `json:"visibility"`
	Active             bool   `json:"active"`
	ResourcePermission struct {
		AllGroups            bool   `json:"allGroups"`
		DefaultStore         bool   `json:"defaultStore"`
		AllPlans             bool   `json:"allPlans"`
		DefaultTarget        bool   `json:"defaultTarget"`
		MorpheusResourceType string `json:"morpheusResourceType"`
		MorpheusResourceId   int64  `json:"morpheusResourceId"`
		CanManage            bool   `json:"canManage"`
		All                  bool   `json:"all"`
		Account              struct {
			ID int64 `json:"id"`
		} `json:"account"`
		Sites []struct {
			ID      int64  `json:"id"`
			Name    string `json:"name"`
			Default bool   `json:"default"`
		} `json:"sites"`
		Plans []struct {
			ID      int64  `json:"id"`
			Name    string `json:"name"`
			Default bool   `json:"default"`
		} `json:"plans"`
	} `json:"resourcePermission"`
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

func (client *Client) ListClusterNamespaces(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d/namespaces", ClustersPath, id),
		QueryParams: req.QueryParams,
		Result:      &ListClusterNamespacesResults{},
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

// UpdateCluster updates an existing cluster
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

// GetClusterApiConfig gets the api configuration for an existing cluster
func (client *Client) GetClusterApiConfig(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d/api-config", ClustersPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &GetClusterApiConfigResult{},
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
