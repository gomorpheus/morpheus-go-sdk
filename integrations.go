package morpheus

import (
	"fmt"
)

var (
	// IntegrationsPath is the API endpoint for integrations
	IntegrationsPath = "/api/integrations"
)

// Integration structures for use in request and response payloads
type Integration struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Enabled         bool   `json:"enabled"`
	Type            string `json:"type"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Version         string `json:"version"`
	IntegrationType struct {
		ID   int64  `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	}
	URL             string `json:"url"`
	ServiceUrl      string `json:"serviceUrl"`
	ServiceUsername string `json:"serviceUsername"`
	ServicePassword string `json:"servicePassword"`
	Token           string `json:"token"`
	TokenHash       string `json:"tokenHash"`
	Config          struct {
		DefaultBranch                string                            `json:"defaultBranch"`
		CacheEnabled                 bool                              `json:"cacheEnabled"`
		AnsiblePlaybooks             string                            `json:"ansiblePlaybooks"`
		AnsibleRoles                 string                            `json:"ansibleRoles"`
		AnsibleGroupVars             string                            `json:"ansibleGroupVars"`
		AnsibleHostVars              string                            `json:"ansibleHostVars"`
		AnsibleCommandBus            string                            `json:"ansibleCommandBus"`
		AnsibleVerbose               bool                              `json:"ansibleVerbose"`
		AnsibleGalaxyEnabled         string                            `json:"ansibleGalaxyEnabled"`
		AnsibleDefaultBranch         string                            `json:"ansibleDefaultBranch"`
		Plugin                       interface{}                       `json:"plugin"`
		IncidentAccess               bool                              `json:"incidentAccess"`
		RequestAccess                bool                              `json:"requestAccess"`
		ServiceNowCMDBBusinessObject string                            `json:"serviceNowCMDBBusinessObject"`
		ServiceNowCustomCmdbMapping  string                            `json:"serviceNowCustomCmdbMapping"`
		ServiceNowCmdbClassMapping   []serviceNowCmdbClassMappingEntry `json:"serviceNowCmdbClassMapping"`
		Databags                     []chefDatabagEntry                `json:"databags"`
		ApprovalUser                 string                            `json:"approvalUser"`
		Company                      string                            `json:"company"`
		AppID                        string                            `json:"appId"`
		InventoryExisting            string                            `json:"inventoryExisting"`
		ExtraAttributes              string                            `json:"extraAttributes"`
		EngineMount                  string                            `json:"engineMount"`
		SecretPath                   string                            `json:"secretPath"`
		SecretEngine                 string                            `json:"secretEngine"`
		SecretPathHash               string                            `json:"secretPathHash"`
		SecretEngineHash             string                            `json:"secretEngineHash"`
	}
	Status     string `json:"status"`
	StatusDate string `json:"statusDate"`
	IsPlugin   bool   `json:"isPlugin"`
	ServiceKey struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	ServiceMode string `json:"serviceMode"`
}

type serviceNowCmdbClassMappingEntry struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	NowClass string `json:"nowClass"`
}

type chefDatabagEntry struct {
	Path string `json:"path"`
	Key  string `json:"key"`
}

// ListIntegrationsResult structure parses the list integrations response payload
type ListIntegrationsResult struct {
	Integrations *[]Integration `json:"integrations"`
	Meta         *MetaResult    `json:"meta"`
}

type GetIntegrationResult struct {
	Integration *Integration `json:"integration"`
}

type CreateIntegrationResult struct {
	Success     bool              `json:"success"`
	Message     string            `json:"msg"`
	Errors      map[string]string `json:"errors"`
	Integration *Integration      `json:"integration"`
}

type UpdateIntegrationResult struct {
	CreateIntegrationResult
}

type DeleteIntegrationResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListIntegrations(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        IntegrationsPath,
		QueryParams: req.QueryParams,
		Result:      &ListIntegrationsResult{},
	})
}

func (client *Client) GetIntegration(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", IntegrationsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetIntegrationResult{},
	})
}

func (client *Client) CreateIntegration(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        IntegrationsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateIntegrationResult{},
	})
}

func (client *Client) UpdateIntegration(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", IntegrationsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateIntegrationResult{},
	})
}

func (client *Client) DeleteIntegration(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", IntegrationsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteIntegrationResult{},
	})
}

// helper functions

func (client *Client) FindIntegrationByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListIntegrations(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListIntegrationsResult)
	integrationCount := len(*listResult.Integrations)
	if integrationCount != 1 {
		return resp, fmt.Errorf("found %d integrations named %v", integrationCount, name)
	}
	firstRecord := (*listResult.Integrations)[0]
	integrationID := firstRecord.ID
	return client.GetIntegration(integrationID, &Request{})
}
