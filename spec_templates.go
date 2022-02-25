package morpheus

import (
	"fmt"
)

var (
	// SpecTemplatesPath is the API endpoint for spec templates
	SpecTemplatesPath = "/api/library/spec-templates"
)

// SpecTemplate structures for use in request and response payloads
type SpecTemplate struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Type   Type   `json:"type"`
	Config Config `json:"config"`
	File   File   `json:"file"`
}

type Type struct {
	Code string `json:"code"`
}
type Cloudformation struct {
	Iam                  string `json:"IAM"`
	CapabilityNamedIam   string `json:"CAPABILITY_NAMED_IAM"`
	CapabilityAutoExpand string `json:"CAPABILITY_AUTO_EXPAND"`
}
type Config struct {
	Cloudformation Cloudformation `json:"cloudformation"`
}
type File struct {
	Sourcetype  string     `json:"sourceType"`
	Content     string     `json:"content"`
	ContentPath string     `json:"contentPath"`
	ContentRef  string     `json:"contentRef"`
	Repository  Repository `json:"repository"`
}

type Repository struct {
	ID int64 `json:"id"`
}

type ListSpecTemplatesResult struct {
	SpecTemplates *[]SpecTemplate `json:"specTemplates"`
	Meta          *MetaResult     `json:"meta"`
}

type GetSpecTemplateResult struct {
	SpecTemplate *SpecTemplate `json:"specTemplate"`
}

type CreateSpecTemplateResult struct {
	Success      bool              `json:"success"`
	Message      string            `json:"msg"`
	Errors       map[string]string `json:"errors"`
	SpecTemplate *SpecTemplate     `json:"specTemplate"`
}

type UpdateSpecTemplateResult struct {
	CreateSpecTemplateResult
}

type DeleteSpecTemplateResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListSpecTemplates(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        SpecTemplatesPath,
		QueryParams: req.QueryParams,
		Result:      &ListSpecTemplatesResult{},
	})
}

func (client *Client) GetSpecTemplate(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", SpecTemplatesPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetSpecTemplateResult{},
	})
}

func (client *Client) CreateSpecTemplate(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        SpecTemplatesPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateSpecTemplateResult{},
	})
}

func (client *Client) UpdateSpecTemplate(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", SpecTemplatesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateSpecTemplateResult{},
	})
}

func (client *Client) DeleteSpecTemplate(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", SpecTemplatesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteSpecTemplateResult{},
	})
}

// FindSpecTemplateByName gets an existing spec template by name
func (client *Client) FindSpecTemplateByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListSpecTemplates(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListSpecTemplatesResult)
	specTemplateCount := len(*listResult.SpecTemplates)
	if specTemplateCount != 1 {
		return resp, fmt.Errorf("found %d SpecTemplates for %v", specTemplateCount, name)
	}
	firstRecord := (*listResult.SpecTemplates)[0]
	specTemplateID := firstRecord.ID
	return client.GetSpecTemplate(specTemplateID, &Request{})
}
