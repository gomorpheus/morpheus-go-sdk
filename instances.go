package morpheus

import (
	"fmt"
	"time"
)

var (
	// InstancesPath is the API endpoint for instances
	InstancesPath = "/api/instances"
)

// Instance structures for use in request and response payloads
type Instance struct {
	ID        int64  `json:"id"`
	UUID      string `json:"uuid"`
	AccountId int64  `json:"accountId"`
	Tenant    struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"tenant"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	InstanceType struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		Code     string `json:"code"`
		Category string `json:"category"`
		Image    string `json:"image"`
	} `json:"instanceType"`
	Layout struct {
		ID                int64  `json:"id"`
		Name              string `json:"name"`
		ProvisionTypeId   int64  `json:"provisionTypeId"`
		ProvisionTypeCode string `json:"provisionTypeCode"`
	} `json:"layout"`
	Group struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"group"`
	Cloud struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"cloud"`
	Containers []int64 `json:"containers"`
	Servers    []int64 `json:"servers"`
	Resources  []struct {
		ID              int64    `json:"id"`
		UUID            string   `json:"uuid"`
		Code            string   `json:"code"`
		Category        string   `json:"category"`
		Name            string   `json:"name"`
		DisplayName     string   `json:"displayName"`
		Labels          []string `json:"labels"`
		ResourceVersion string   `json:"resourceVersion"`
		ResourceContext string   `json:"resourceContext"`
		Owner           struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"owner"`
		ResourceType string `json:"resourceType"`
		ResourceIcon string `json:"resourceIcon"`
		Type         struct {
			Id   int64  `json:"id"`
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"type"`
		Status     string `json:"status"`
		Enabled    bool   `json:"enabled"`
		ExternalId string `json:"externalId"`
	} `json:"resources"`
	ConnectionInfo []struct {
		Ip   string `json:"ip"`
		Port int64  `json:"port"`
	} `json:"connectionInfo"`
	Environment string                   `json:"environment"`
	Plan        InstancePlan             `json:"plan"`
	Config      map[string]interface{}   `json:"config"`
	Labels      []string                 `json:"labels"`
	Version     string                   `json:"instanceVersion"`
	Status      string                   `json:"status"`
	Owner       Owner                    `json:"owner"`
	Volumes     []map[string]interface{} `json:"volumes"`
	Interfaces  []map[string]interface{} `json:"interfaces"`
	Controllers []map[string]interface{} `json:"controllers"`
	Tags        []struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"tags"`
	Metadata             []map[string]interface{} `json:"metadata"`
	EnvironmentVariables []struct {
		Name   string `json:"name"`
		Value  string `json:"value"`
		Export bool   `json:"export"`
		Masked bool   `json:"masked"`
	} `json:"evars"`
	CustomOptions  map[string]interface{} `json:"customOptions"`
	MaxMemory      int64                  `json:"maxMemory"`
	MaxStorage     int64                  `json:"maxStorage"`
	MaxCores       int64                  `json:"maxCores"`
	CoresPerSocket int64                  `json:"coresPerSocket"`
	MaxCpu         int64                  `json:"maxCpu"`
	HourlyCost     float64                `json:"hourlyCost"`
	HourlyPrice    float64                `json:"hourlyPrice"`
	InstancePrice  struct {
		Price    float64 `json:"price"`
		Cost     float64 `json:"cost"`
		Currency string  `json:"currency"`
		Unit     string  `json:"unit"`
	} `json:"instancePrice"`
	DateCreated       string `json:"dateCreated"`
	LastUpdated       string `json:"lastUpdated"`
	HostName          string `json:"hostName"`
	DomainName        string `json:"domainName"`
	EnvironmentPrefix string `json:"environmentPrefix"`
	FirewallEnabled   bool   `json:"firewallEnabled"`
	NetworkLevel      string `json:"networkLevel"`
	AutoScale         bool   `json:"autoScale"`
	InstanceContext   string `json:"instanceContext"`
	Locked            bool   `json:"locked"`
	IsScalable        bool   `json:"isScalable"`
	CreatedBy         struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
	} `json:"createdBy"`
}

type InstancePlan struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Owner struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type ContainerDetails struct {
	ID               int64  `json:"id"`
	UUID             string `json:"uuid"`
	Name             string `json:"name"`
	IP               string `json:"ip"`
	InternalIp       string `json:"internalIp"`
	InternalHostname string `json:"internalHostname"`
	ExternalHostname string `json:"externalHostname"`
	ExternalDomain   string `json:"externalDomain"`
	ExternalFqdn     string `json:"externalFqdn"`
	AccountId        int64  `json:"accountId"`
	Instance         struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"instance"`
	ContainerType struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		Code     string `json:"code"`
		Category string `json:"category"`
	} `json:"containerType"`
	Server struct {
		ID               int64  `json:"id"`
		UUID             string `json:"uuid"`
		ExternalId       string `json:"externalId"`
		InternalId       string `json:"internalId"`
		ExternalUniqueId string `json:"externalUniqueId"`
		Name             string `json:"name"`
		ExternalName     string `json:"externalName"`
		Hostname         string `json:"hostname"`
		AccountId        int64  `json:"accountId"`
		SshHost          string `json:"sshHost"`
		ExternalIp       string `json:"externalIp"`
		InternalIp       string `json:"internalIp"`
		Platform         string `json:"platform"`
		PlatformVersion  string `json:"platformVersion"`
		AgentInstalled   bool   `json:"agentInstalled"`
		AgentVersion     string `json:"agentVersion"`
		SourceImage      struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"sourceImage"`
	} `json:"server"`
}

// ListInstancesResult structure parses the list instances response payload
type ListInstancesResult struct {
	Instances *[]Instance `json:"instances"`
	Meta      *MetaResult `json:"meta"`
}

type GetInstanceResult struct {
	Instance *Instance `json:"instance"`
}

type CreateInstanceResult struct {
	Success  bool              `json:"success"`
	Message  string            `json:"msg"`
	Errors   map[string]string `json:"errors"`
	Instance *Instance         `json:"instance"`
}

type UpdateInstanceResult struct {
	CreateInstanceResult
}

type DeleteInstanceResult struct {
	DeleteResult
}

type ListInstancePlansResult struct {
	Plans *[]InstancePlan `json:"plans"`
	Meta  *MetaResult     `json:"meta"`
}

type GetInstancePlanResult struct {
	Plan *InstancePlan `json:"plan"`
}

// API endpoints

func (client *Client) ListInstances(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        InstancesPath,
		QueryParams: req.QueryParams,
		Result:      &ListInstancesResult{},
	})
}

func (client *Client) GetInstance(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", InstancesPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetInstanceResult{},
	})
}

func (client *Client) CreateInstance(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        InstancesPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateInstanceResult{},
	})
}

func (client *Client) UpdateInstance(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", InstancesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateInstanceResult{},
	})
}

func (client *Client) DeleteInstance(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", InstancesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteInstanceResult{},
	})
}

// helper functions

func (client *Client) FindInstanceByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListInstances(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListInstancesResult)
	instanceCount := len(*listResult.Instances)
	if instanceCount != 1 {
		return resp, fmt.Errorf("found %d Instances for %v", instanceCount, name)
	}
	firstRecord := (*listResult.Instances)[0]
	instanceId := firstRecord.ID
	return client.GetInstance(instanceId, &Request{})
}

// Plan fetching
// todo: this needs to be refactored soon

func (client *Client) ListInstancePlans(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/service-plans", InstancesPath),
		QueryParams: req.QueryParams,
		Result:      &ListInstancePlansResult{},
	})
}

// todo: need this api endpoint still, and consolidate to /api/plans perhaps
func (client *Client) GetInstancePlan(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/service-plans/%d", InstancesPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetInstancePlanResult{},
	})
}

func (client *Client) FindInstancePlanByName(name string, req *Request) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListInstancePlans(&Request{
		QueryParams: map[string]string{
			//"name": name, // this is not even supported..
			"zoneId":   req.QueryParams["zoneId"], // todo: use cloudId
			"layoutId": req.QueryParams["layoutId"],
			"siteId":   req.QueryParams["siteId"], // todo: use groupId
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListInstancePlansResult)
	planCount := len(*listResult.Plans)
	// need to filter the list ourselves for now..
	var matchingPlans []InstancePlan
	for i := 0; i < planCount; i++ {
		plan := (*listResult.Plans)[i] // .(InstancePlan)
		if plan.Name == name || plan.Code == name || string(rune(plan.ID)) == name {
			matchingPlans = append(matchingPlans, plan)
		}
	}
	matchingPlanCount := len(matchingPlans)
	if matchingPlanCount != 1 {
		return resp, fmt.Errorf("found %d Plans for '%v'", matchingPlanCount, name)
	}
	firstRecord := matchingPlans[0]

	// planId := firstRecord.ID
	// return client.GetInstancePlan(planId, &Request{})

	// for now just return a fake response until endpoint is ready
	var result = &GetInstancePlanResult{
		Plan: &firstRecord,
	}
	mockResp := &Response{
		//RestyResponse: restyResponse,
		Success:    true,
		StatusCode: 200,
		Status:     "200 OK",
		ReceivedAt: time.Now(),
		// Size: restyResponse.Size(),
		// Body: restyResponse.Body(), // byte[]
		Result: result,
	}
	return mockResp, nil
}

// this should work by code or name
// it also requires zoneId AND layoutId??
func (client *Client) FindInstancePlanByCode(code string, req *Request) (*Response, error) {
	return client.FindInstancePlanByName(code, req)
}
