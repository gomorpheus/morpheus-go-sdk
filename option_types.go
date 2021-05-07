// Morpheus API types and Client methods for Option Types
package morpheus

import (
	"fmt"
)

// globals

var (
	OptionTypesPath = "/api/library/option-types"
)

// OptionType structures for use in request and response payloads

type OptionType struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	FieldName    string `json:"fieldName"`
	Type         string `json:"type"`
	FieldLabel   string `json:"fieldLabel"`
	PlaceHolder  string `json:"placeHolder"`
	DefaultValue string `json:"defaultValue"`
	Required     bool   `json:"required"`
	ExportMeta   bool   `json:"exportMeta"`
	OptionListId int64  `json:"optionList,omitEmpty"`
	HelpBlock    string `json:"helpBlock"`
}

type ListOptionTypesResult struct {
	OptionTypes *[]OptionType `json:"optionTypes"`
	Meta        *MetaResult   `json:"meta"`
}

type GetOptionTypeResult struct {
	OptionType *OptionType `json:"optionType"`
}

type CreateOptionTypeResult struct {
	Success    bool              `json:"success"`
	Message    string            `json:"msg"`
	Errors     map[string]string `json:"errors"`
	OptionType *OptionType       `json:"optionType"`
}

type UpdateOptionTypeResult struct {
	CreateOptionTypeResult
}

type DeleteOptionTypeResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListOptionTypes(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        OptionTypesPath,
		QueryParams: req.QueryParams,
		Result:      &ListOptionTypesResult{},
	})
}

func (client *Client) GetOptionType(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", OptionTypesPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetOptionTypeResult{},
	})
}

func (client *Client) CreateOptionType(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        OptionTypesPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateOptionTypeResult{},
	})
}

func (client *Client) UpdateOptionType(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", OptionTypesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateOptionTypeResult{},
	})
}

func (client *Client) DeleteOptionType(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", OptionTypesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteOptionTypeResult{},
	})
}

// helper functions

func (client *Client) FindOptionTypeByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListOptionTypes(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListOptionTypesResult)
	optionTypeCount := len(*listResult.OptionTypes)
	if optionTypeCount != 1 {
		return resp, fmt.Errorf("Found %d OptionTypes for %v", optionTypeCount, name)
	}
	firstRecord := (*listResult.OptionTypes)[0]
	optionTypeID := firstRecord.ID
	return client.GetOptionType(optionTypeID, &Request{})
}
