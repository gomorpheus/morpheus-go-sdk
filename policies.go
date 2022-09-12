package morpheus

import (
	"fmt"
	"time"
)

var (
	// PoliciesPath is the API endpoint for policies
	PoliciesPath = "/api/policies"
)

// Policy structures for use in request and response payloads
type Policy struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	EachUser   bool   `json:"eachUser"`
	PolicyType struct {
		ID   int64  `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"type"`
	Config struct {
		ShutdownType                     string    `json:"shutdownType"`
		ShutdownAge                      int64     `json:"shutdownAge"`
		ShutdownRenewal                  int64     `json:"shutdownRenewal"`
		ShutdownNotify                   int64     `json:"shutdownNotify"`
		ShutdownMessage                  string    `json:"shutdownMessage"`
		ShutdownAutoRenew                string    `json:"shutdownAutoRenew"`
		ShutdownExtensionsBeforeApproval int64     `json:"shutdownExtensionsBeforeApproval"`
		ShutdownHideFixed                bool      `json:"shutdownHideFixed"`
		Strict                           string    `json:"strict"`
		Key                              string    `json:"key"`
		ValueListId                      int64     `json:"valueListId"`
		Value                            string    `json:"value"`
		PowerSchedule                    string    `json:"powerSchedule"`
		PowerScheduleType                string    `json:"powerScheduleType"`
		AccountIntegrationId             string    `json:"accountIntegrationId"`
		WorkflowID                       string    `json:"workflowId"`
		CreateUser                       string    `json:"createUser"`
		CreateUserType                   string    `json:"createUserType"`
		MaxRouters                       string    `json:"maxRouters"`
		MaxNetworks                      string    `json:"maxNetworks"`
		MaxVms                           string    `json:"maxVms"`
		MaxStorage                       string    `json:"maxStorage"`
		MaxPools                         int64     `json:"maxPools"`
		MaxPoolMembers                   int64     `json:"maxPoolMembers"`
		MaxMemory                        string    `json:"maxMemory"`
		MaxHosts                         int64     `json:"maxHosts"`
		MaxCores                         int64     `json:"maxCores"`
		MaxContainers                    int64     `json:"maxContainers"`
		MaxVirtualServers                int64     `json:"maxVirtualServers"`
		NamingType                       string    `json:"namingType"`
		NamingPattern                    string    `json:"namingPattern"`
		NamingConflict                   string    `json:"namingConflict"`
		HostNamingType                   string    `json:"hostNamingType"`
		HostNamingPattern                string    `json:"hostNamingPattern"`
		MaxPrice                         string    `json:"maxPrice"`
		MaxPriceCurrency                 string    `json:"maxPriceCurrency"`
		MaxPriceUnit                     string    `json:"maxPriceUnit"`
		RemovalAge                       int64     `json:"removalAge"`
		MotdTitle                        string    `json:"motd.title"`
		MotdMessage                      string    `json:"motd.message"`
		MotdType                         string    `json:"motd.type"`
		MotdFullPage                     string    `json:"motd.fullPage"`
		MotdDate                         time.Time `json:"motd.date"`
		Motd                             struct {
			Title    string `json:"title"`
			Message  string `json:"message"`
			Type     string `json:"type"`
			FullPage string `json:"fullPage"`
		} `json:"motd"`
		KeyPattern                        string `json:"keyPattern"`
		Read                              string `json:"read"`
		Write                             string `json:"write"`
		Update                            string `json:"update"`
		Delete                            string `json:"delete"`
		List                              string `json:"list"`
		UserGroup                         int64  `json:"userGroup"`
		ServerNamingType                  string `json:"serverNamingType"`
		ServerNamingPattern               string `json:"serverNamingPattern"`
		ServerNamingConflict              string `json:"serverNamingConflict"`
		CreateBackupType                  string `json:"createBackupType"`
		LifecycleType                     string `json:"lifecycleType"`
		LifecycleAge                      int64  `json:"lifecycleAge"`
		LifecycleRenewal                  int64  `json:"lifecycleRenewal"`
		LifecycleNotify                   int64  `json:"lifecycleNotify"`
		LifecycleMessage                  string `json:"lifecycleMessage"`
		LifecycleExtensionsBeforeApproval int64  `json:"lifecycleExtensionsBeforeApproval"`
		LifecycleAutoRenew                string `json:"lifecycleAutoRenew"`
		LifecycleHideFixed                bool   `json:"lifecycleHideFixed"`
	} `json:"config"`
	Owner struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"owner"`
	RefID   int64  `json:"refId"`
	RefType string `json:"refType"`
	Role    struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"role"`
	Site struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"site"`
	User struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	Zone struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"zone"`
	Accounts []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
}

type ListPoliciesResult struct {
	Policies *[]Policy   `json:"policies"`
	Meta     *MetaResult `json:"meta"`
}

type GetPolicyResult struct {
	Policy *Policy `json:"policy"`
}

type CreatePolicyResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Policy  *Policy           `json:"policy"`
}

type UpdatePolicyResult struct {
	CreatePolicyResult
}

type DeletePolicyResult struct {
	DeleteResult
}

// ListPolicies list all policies
func (client *Client) ListPolicies(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        PoliciesPath,
		QueryParams: req.QueryParams,
		Result:      &ListPoliciesResult{},
	})
}

// GetPolicy gets a policy
func (client *Client) GetPolicy(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", PoliciesPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetPolicyResult{},
	})
}

// CreatePolicy creates a new policy
func (client *Client) CreatePolicy(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        PoliciesPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreatePolicyResult{},
	})
}

// UpdatePolicy updates an existing policy
func (client *Client) UpdatePolicy(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", PoliciesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdatePolicyResult{},
	})
}

// DeletePolicy deletes an existing policy
func (client *Client) DeletePolicy(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", PoliciesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeletePolicyResult{},
	})
}

// FindPolicyByName gets an existing policy by name
func (client *Client) FindPolicyByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListPolicies(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListPoliciesResult)
	policyCount := len(*listResult.Policies)
	if policyCount != 1 {
		return resp, fmt.Errorf("found %d policies named %v", policyCount, name)
	}
	firstRecord := (*listResult.Policies)[0]
	policyID := firstRecord.ID
	return client.GetPolicy(policyID, &Request{})
}
