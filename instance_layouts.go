package morpheus

import (
	"fmt"
)

var (
	// InstanceLayoutsPath is the API endpoint for instance layouts
	InstanceLayoutsPath = "/api/library/layouts"
)

// InstanceLayout structures for use in request and response payloads
type InstanceLayout struct {
	ID      int64 `json:"id"`
	Account struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"account"`
	Name                     string `json:"name"`
	Description              string `json:"description"`
	Code                     string `json:"code"`
	ContainerVersion         string `json:"instanceVersion"`
	Creatable                bool   `json:"creatable"`
	MemoryRequirement        int64  `json:"memoryRequirement"`
	SupportsConvertToManaged bool   `json:"supportsConvertToManaged"`
	ProvisionType            struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"provisionType"`
	TaskSets []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"taskSets"`
	ContainerTypes []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"containerTypes"`
	SpecTemplates []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"specTemplates"`
	OptionTypes []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"optionTypes"`
	EnvironmentVariables []struct {
		EvarName         string `json:"evarName"`
		Name             string `json:"name"`
		DefaultValue     string `json:"defaultValue"`
		DefaultValueHash string `json:"defaultValueHash"`
		ValueType        string `json:"valueType"`
		Export           bool   `json:"export"`
		Masked           bool   `json:"masked"`
	} `json:"environmentVariables"`
}

// ListInstanceLayoutsResult structure parses the list instance layouts response payload
type ListInstanceLayoutsResult struct {
	InstanceLayouts *[]InstanceLayout `json:"instanceTypeLayouts"`
	Meta            *MetaResult       `json:"meta"`
}

type GetInstanceLayoutResult struct {
	InstanceLayout *InstanceLayout `json:"instanceTypeLayout"`
}

type CreateInstanceLayoutResult struct {
	Success        bool              `json:"success"`
	Message        string            `json:"msg"`
	Errors         map[string]string `json:"errors"`
	InstanceLayout *InstanceLayout   `json:"instanceTypeLayout"`
}

type UpdateInstanceLayoutResult struct {
	CreateInstanceLayoutResult
}

type DeleteInstanceLayoutResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListInstanceLayouts(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        InstanceLayoutsPath,
		QueryParams: req.QueryParams,
		Result:      &ListInstanceLayoutsResult{},
	})
}

func (client *Client) GetInstanceLayout(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", InstanceLayoutsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetInstanceLayoutResult{},
	})
}

func (client *Client) CreateInstanceLayout(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        fmt.Sprintf("/api/library/instance-types/%d/layouts", id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateInstanceLayoutResult{},
	})
}

func (client *Client) UpdateInstanceLayout(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", InstanceLayoutsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateInstanceLayoutResult{},
	})
}

func (client *Client) DeleteInstanceLayout(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", InstanceLayoutsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteInstanceLayoutResult{},
	})
}

// helper functions
func (client *Client) FindInstanceLayoutByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListInstanceLayouts(&Request{
		QueryParams: map[string]string{
			"name": name,
			"max":  "5000",
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListInstanceLayoutsResult)
	instanceLayoutsCount := len(*listResult.InstanceLayouts)
	if instanceLayoutsCount != 1 {
		return resp, fmt.Errorf("found %d InstanceLayouts for %v", instanceLayoutsCount, name)
	}
	firstRecord := (*listResult.InstanceLayouts)[0]
	instanceLayoutId := firstRecord.ID
	return client.GetInstanceLayout(instanceLayoutId, &Request{})
}
