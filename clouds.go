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
	ID         int64  `json:"id"`
	UUID       string `json:"uuid"`
	ExternalID string `json:"externalId"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Location   string `json:"location"`
	Owner      struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"owner"`
	AccountID int64 `json:"accountId"`
	Account   struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"account"`
	Visibility           string    `json:"visibility"`
	Enabled              bool      `json:"enabled"`
	Status               string    `json:"status"`
	StatusMessage        string    `json:"statusMessage"`
	StatusDate           string    `json:"statusDate"`
	LastSync             string    `json:"lastSync"`
	NextRunDate          string    `json:"nextRunDate"`
	LastSyncDuration     string    `json:"lastSyncDuration"`
	CostStatus           string    `json:"costStatus"`
	CostStatusMessage    string    `json:"costStatusMessage"`
	CostStatusDate       string    `json:"costStatusDate"`
	CostLastSyncDuration string    `json:"costLastSyncDuration"`
	CostLastSync         string    `json:"costLastSync"`
	CloudType            CloudType `json:"zoneType"`
	CloudTypeID          int64     `json:"zoneTypeId"`
	GuidanceMode         string    `json:"guidanceMode"`
	StorageMode          string    `json:"storageMode"`
	AgentMode            string    `json:"agentMode"`
	ConsoleKeymap        string    `json:"consoleKeymap"`
	ContainerMode        string    `json:"containerMode"`
	CostingMode          string    `json:"costingMode"`
	ServiceVersion       string    `json:"serviceVersion"`
	SecurityMode         string    `json:"securityMode"`
	InventoryLevel       string    `json:"inventoryLevel"`
	TimeZone             string    `json:"timezone"`
	NetworkDomain        struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"networkDomain"`
	DomainName            string `json:"domainName"`
	RegionCode            string `json:"regionCode"`
	AutoRecoverPowerState bool   `json:"autoRecoverPowerState"`
	ScalePriority         int64  `json:"scalePriority"`
	Config                struct {
		// AWS
		Endpoint             string `json:"endpoint"`
		IsVpc                string `json:"isVpc"`
		ImageStoreId         string `json:"imageStoreId"`
		EbsEncryption        string `json:"ebsEncryption"`
		CostingReport        string `json:"costingReport"`
		CostingRegion        string `json:"costingRegion"`
		CostingSecretKeyHash string `json:"costingSecretKeyHash"`
		SecretKeyHash        string `json:"secretKeyHash"`
		AccessKey            string `json:"accessKey"`
		SecretKey            string `json:"secretKey"`
		VPC                  string `json:"vpc"`
		StsAssumeRole        string `json:"stsAssumeRole"`
		UseHostCredentials   string `json:"useHostCredentials"`
		CostingAccessKey     string `json:"costingAccessKey"`
		CostingBucketName    string `json:"costingBucketName"`
		CostingFolder        string `json:"costingFolder"`
		CostingReportName    string `json:"costingReportName"`
		CostingSecretKey     string `json:"costingSecretKey"`

		// vSphere
		APIUrl                     string `json:"apiUrl"`
		APIVersion                 string `json:"apiVersion"`
		Datacenter                 string `json:"datacenter"`
		Cluster                    string `json:"cluster"`
		DiskStorageType            string `json:"diskStorageType"`
		DatacenterID               string `json:"datacenterId"`
		EnableVNC                  string `json:"enableVnc"`
		DiskEncryption             string `json:"diskEncryption"`
		EnableDiskTypeSelection    string `json:"enableDiskTypeSelection"`
		EnableStorageTypeSelection string `json:"enableStorageTypeSelection"`
		EnableNetworkTypeSelection string `json:"enableNetworkTypeSelection"`
		ResourcePool               string `json:"resourcePool"`
		ResourcePoolId             string `json:"resourcePoolId"`
		HideHostSelection          string `json:"hideHostSelection"`

		// Azure
		AzureCostingMode    string `json:"azureCostingMode"`
		SubscriberID        string `json:"subscriberId"`
		TenantID            string `json:"tenantId"`
		ClientID            string `json:"clientId"`
		ClientSecret        string `json:"clientSecret"`
		ClientSecretHash    string `json:"clientSecretHash"`
		ResourceGroup       string `json:"resourceGroup"`
		CSPCustomer         string `json:"cspCustomer"`
		CSPTenantID         string `json:"cspTenantId"`
		CSPClientID         string `json:"cspClientId"`
		CSPClientSecret     string `json:"cspClientSecret"`
		CSPClientSecretHash string `json:"cspClientSecretHash"`

		// GCP
		GoogleRegionID string `json:"googleRegionId"`
		GoogleBucket   string `json:"googleBucket"`
		PrivateKey     string `json:"privateKey"`
		PrivateKeyHash string `json:"privateKeyHash"`
		ClientEmail    string `json:"clientEmail"`

		// Hyperv
		Provider    string `json:"provider"`
		Host        string `json:"host"`
		WorkingPath string `json:"workingPath"`
		VMPath      string `json:"vmPath"`
		DiskPath    string `json:"diskPath"`

		// OpenStack
		IdentityApi                  string `json:"identityApi"`
		DomainId                     string `json:"domainId"`
		ProjectName                  string `json:"projectName"`
		OsRelease                    string `json:"osRelease"`
		DiskMode                     string `json:"diskMode"`
		IdentityVersion              string `json:"identityVersion"`
		ComputeApi                   string `json:"computeApi"`
		ComputeVersion               string `json:"computeVersion"`
		ImageApi                     string `json:"imageApi"`
		ImageVersion                 string `json:"imageVersion"`
		StorageApi                   string `json:"storageApi"`
		StorageVersion               string `json:"storageVersion"`
		NetworkApi                   string `json:"networkApi"`
		NetworkVersion               string `json:"networkVersion"`
		ApiProjectId                 string `json:"apiProjectId"`
		ApiTokenExpiresAt            string `json:"apiTokenExpiresAt"`
		LbaasType                    string `json:"lbaasType"`
		ApiDomainId                  string `json:"apiDomainId"`
		ApiUserId                    string `json:"apiUserId"`
		ProvisionMethod              string `json:"provisionMethod"`
		ComputeMicroVersion          string `json:"computeMicroVersion"`
		ImageMicroVersion            string `json:"imageMicroVersion"`
		StorageMicroVersion          string `json:"storageMicroVersion"`
		NetworkMicroVersion          string `json:"networkMicroVersion"`
		LoadBalancerApi              string `json:"loadBalancerApi"`
		LoadBalancerVersion          string `json:"loadBalancerVersion"`
		LoadBalancerMicroVersion     string `json:"loadBalancerMicroVersion"`
		LoadBalancerV1Api            string `json:"loadBalancerV1Api"`
		LoadBalancerV1Version        string `json:"loadBalancerV1Version"`
		LoadBalancerV1MicroVersion   string `json:"loadBalancerV1MicroVersion"`
		ObjectStorageApi             string `json:"objectStorageApi"`
		ObjectStorageVersion         string `json:"objectStorageVersion"`
		ObjectStorageMicroVersion    string `json:"objectStorageMicroVersion"`
		SharedFileSystemApi          string `json:"sharedFileSystemApi"`
		SharedFileSystemVersion      string `json:"sharedFileSystemVersion"`
		SharedFileSystemMicroVersion string `json:"sharedFileSystemMicroVersion"`

		// VCD
		OrgID                 string `json:"orgId"`
		VDCID                 string `json:"vdcId"`
		VCDVersion            string `json:"vcdVersion"`
		DefaultStorageProfile string `json:"defaultStorageProfile"`
		Catalog               string `json:"catalog"`

		// General
		ClusterRef      string `json:"clusterRef"`
		ProjectID       string `json:"projectId"`
		ApplianceUrl    string `json:"applianceUrl"`
		DatacenterName  string `json:"datacenterName"`
		ImportExisting  string `json:"importExisting"`
		InventoryLevel  string `json:"inventoryLevel"`
		NetworkServerID string `json:"networkServer.id"`
		NetworkServer   struct {
			ID string `json:"id"`
		} `json:"networkServer"`
		SecurityMode        string      `json:"securityMode"`
		CertificateProvider string      `json:"certificateProvider"`
		BackupMode          string      `json:"backupMode"`
		ReplicationMode     string      `json:"replicationMode"`
		DnsIntegrationID    string      `json:"dnsIntegrationId"`
		ServiceRegistryID   string      `json:"serviceRegistryId"`
		ConfigManagementID  string      `json:"configManagementId"`
		ConfigCmdbID        interface{} `json:"configCmdbId"`
		SecurityServer      string      `json:"securityServer"`
		CloudType           string      `json:"cloudType"`
		AccountType         string      `json:"accountType"`
		RPCMode             string      `json:"rpcMode"`
		EncryptionSet       string      `json:"encryptionSet"`
		ConfigCmID          string      `json:"configCmId"`
		KubeURL             string      `json:"kubeUrl"`
		CostingProjectID    string      `json:"costingProjectId"`
		ConfigCMDBDiscovery bool        `json:"configCmdbDiscovery"`
		CostingDatasetID    string      `json:"costingDatasetId"`
		Username            string      `json:"username"`
		Password            string      `json:"password"`
		DistributedWorkerId string      `json:"distributedWorkerId"`
		PasswordHash        string      `json:"passwordHash"`
	} `json:"config"`
	Credential struct {
		Type string `json:"type"`
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"credential"`
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
