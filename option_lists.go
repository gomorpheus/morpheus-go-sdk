// Morpheus API types and Client methods for Option Types
package morpheus

import (
	"fmt"
)

// globals

var (
	OptionListsPath = "/api/library/option-type-lists"
)

// OptionLists structures for use in request and response payloads

type OptionList struct {
	ID                int64                  `json:"id"`
	Name              string                 `json:"name"`
	Description       string                 `json:"description"`
	Type              string                 `json:"type"`
	SourceURL         string                 `json:"sourceUrl"`
	Visibility        string                 `json:"visibility"`
	SourceMethod      string                 `json:"sourceMethod"`
	APIType           string                 `json:"apiType,omitempty"`
	IgnoreSSLErrors   bool                   `json:"ignoreSSLErrors"`
	RealTime          bool                   `json:"realTime"`
	InitialDataset    string                 `json:"initialDataset"`
	TranslationScript string                 `json:"translationScript"`
	Config            map[string]interface{} `json:"config"`
	//Config            struct {
	//	SourceHeaders []SourceHeader `json:"sourceHeaders"`
	//} `json:"config"`
}

type SourceHeader struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Masked bool   `json:"masked"`
}

type ListOptionListsResult struct {
	OptionLists *[]OptionList `json:"optionTypeLists"`
	Meta        *MetaResult   `json:"meta"`
}

type GetOptionListResult struct {
	OptionList *OptionList `json:"optionTypeList"`
}

type CreateOptionListResult struct {
	Success    bool              `json:"success"`
	Message    string            `json:"msg"`
	Errors     map[string]string `json:"errors"`
	OptionList *OptionList       `json:"optionTypeList"`
}

type UpdateOptionListResult struct {
	CreateOptionListResult
}

type DeleteOptionListResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListOptionLists(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        OptionListsPath,
		QueryParams: req.QueryParams,
		Result:      &ListOptionListsResult{},
	})
}

func (client *Client) GetOptionList(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", OptionListsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetOptionListResult{},
	})
}

func (client *Client) CreateOptionList(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        OptionListsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateOptionListResult{},
	})
}

func (client *Client) UpdateOptionList(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", OptionListsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateOptionListResult{},
	})
}

func (client *Client) DeleteOptionList(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", OptionListsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteOptionListResult{},
	})
}

// helper functions

func (client *Client) FindOptionListByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListOptionLists(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListOptionListsResult)
	optionListCount := len(*listResult.OptionLists)
	if optionListCount != 1 {
		return resp, fmt.Errorf("Found %d OptionLists for %v", optionListCount, name)
	}
	firstRecord := (*listResult.OptionLists)[0]
	optionListID := firstRecord.ID
	return client.GetOptionList(optionListID, &Request{})
}
