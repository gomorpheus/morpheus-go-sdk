package morpheus

import (
	"fmt"
)

var (
	// CloudsPath is the API endpoint for clouds (zones)
	CloudsPath = "/api/zones"
)

// Cloud structures for use in request and response payloads
type Cloud struct {
	ID       int64  `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Location string `json:"location"`
	Owner    struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"owner"`
	AccountID int64 `json:"accountId"`
	Account   struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"account"`
	Visibility        string    `json:"visibility"`
	Enabled           bool      `json:"enabled"`
	Status            string    `json:"status"`
	StatusMessage     string    `json:"statusMessage"`
	StatusDate        string    `json:"statusDate"`
	CostStatus        string    `json:"costStatus"`
	CostStatusMessage string    `json:"costStatusMessage"`
	CostStatusDate    string    `json:"costStatusDate"`
	CloudType         CloudType `json:"zoneType"`
	CloudTypeID       int64     `json:"zoneTypeId"`
	GuidanceMode      string    `json:"guidanceMode"`
	StorageMode       string    `json:"storageMode"`
	AgentMode         string    `json:"agentMode"`
	ServiceVersion    string    `json:"serviceVersion"`
	Config            struct {
		APIUrl          string `json:"apiUrl"`
		APIVersion      string `json:"apiVersion"`
		Datacenter      string `json:"datacenter"`
		Cluster         string `json:"cluster"`
		DiskStorageType string `json:"diskStorageType"`
		DatacenterID    string `json:"datacenterId"`
		EnableVNC       string `json:"enableVnc"`
		AccessKey       string `json:"accessKey"`
		SecretKey       string `json:"secretKey"`
		VPC             string `json:"vpc"`
		DiskEncryption  string `json:"diskEncryption"`
		ApplianceUrl    string `json:"applianceUrl"`
		DatacenterName  string `json:"datacenterName"`
		SecretKeyHash   string `json:"secretKeyHash"`
		SubscriberID    string `json:"subscriberId"`
		TenantID        string `json:"tenantId"`
		ClientID        string `json:"clientId"`
		ClientSecret    string `json:"clientSecret"`
		ResourceGroup   string `json:"resourceGroup"`
		ImportExisting  string `json:"importExisting"`
		InventoryLevel  string `json:"inventoryLevel"`
		NetworkServerID string `json:"networkServer.id"`
		NetworkServer   struct {
			ID string `json:"id"`
		} `json:"networkServer"`
		SecurityMode          string      `json:"securityMode"`
		CertificateProvider   string      `json:"certificateProvider"`
		BackupMode            string      `json:"backupMode"`
		ReplicationMode       string      `json:"replicationMode"`
		DnsIntegrationID      string      `json:"dnsIntegrationId"`
		ServiceRegistryID     string      `json:"serviceRegistryId"`
		ConfigManagementID    string      `json:"configManagementId"`
		ConfigCmdbID          interface{} `json:"configCmdbId"`
		SecurityServer        string      `json:"securityServer"`
		CloudType             string      `json:"cloudType"`
		AccountType           string      `json:"accountType"`
		CSPTenantID           string      `json:"cspTenantId"`
		CSPClientID           string      `json:"cspClientId"`
		CSPClientSecret       string      `json:"cspClientSecret"`
		RPCMode               string      `json:"rpcMode"`
		EncryptionSet         string      `json:"encryptionSet"`
		ConfigCmID            string      `json:"configCmId"`
		ClientSecretHash      string      `json:"clientSecretHash"`
		CSPClientSecretHash   string      `json:"cspClientSecretHash"`
		ProjectID             string      `json:"projectId"`
		PrivateKey            string      `json:"privateKey"`
		ClientEmail           string      `json:"clientEmail"`
		GoogleRegionID        string      `json:"googleRegionId"`
		GoogleBucket          string      `json:"googleBucket"`
		KubeURL               string      `json:"kubeUrl"`
		CostingProjectID      string      `json:"costingProjectId"`
		ConfigCMDBDiscovery   bool        `json:"configCmdbDiscovery"`
		CostingDatasetID      string      `json:"costingDatasetId"`
		PrivateKeyHash        string      `json:"privateKeyHash"`
		Provider              string      `json:"provider"`
		Host                  string      `json:"host"`
		WorkingPath           string      `json:"workingPath"`
		VMPath                string      `json:"vmPath"`
		DiskPath              string      `json:"diskPath"`
		Username              string      `json:"username"`
		Password              string      `json:"password"`
		DistributedWorkerId   string      `json:"distributedWorkerId"`
		VCDVersion            string      `json:"vcdVersion"`
		DefaultStorageProfile string      `json:"defaultStorageProfile"`
		Passwordhash          string      `json:"passwordHash"`
		OrgID                 string      `json:"orgId"`
		VDCID                 string      `json:"vdcId"`
	} `json:"config"`
	DateCreated string `json:"dateCreated"`
	LastUpdated string `json:"lastUpdated"`
	Groups      []struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		AccountID int64  `json:"accountId"`
	} `json:"groups"`
}

type CloudType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// ListCloudsResult structure parses the list clouds response payload
type ListCloudsResult struct {
	Clouds *[]Cloud    `json:"zones"`
	Meta   *MetaResult `json:"meta"`
}

type GetCloudResult struct {
	Cloud *Cloud `json:"zone"`
}

type CreateCloudResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Cloud   *Cloud            `json:"zone"`
}

type UpdateCloudResult struct {
	CreateCloudResult
}

type DeleteCloudResult struct {
	DeleteResult
}

// API endpoints

func (client *Client) ListClouds(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        CloudsPath,
		QueryParams: req.QueryParams,
		Result:      &ListCloudsResult{},
	})
}

func (client *Client) GetCloud(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", CloudsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetCloudResult{},
	})
}

// CreateCloud creates a new cloud
func (client *Client) CreateCloud(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        CloudsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateCloudResult{},
	})
}

// UpdateCloud updates an existing cloud
func (client *Client) UpdateCloud(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", CloudsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateCloudResult{},
	})
}

// DeleteCloud deletes an existing cloud
func (client *Client) DeleteCloud(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", CloudsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteCloudResult{},
	})
}

func (client *Client) FindCloudByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListClouds(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListCloudsResult)
	cloudsCount := len(*listResult.Clouds)
	if cloudsCount != 1 {
		return resp, fmt.Errorf("found %d Clouds for %v", cloudsCount, name)
	}
	firstRecord := (*listResult.Clouds)[0]
	cloudId := firstRecord.ID
	return client.GetCloud(cloudId, &Request{})
}
