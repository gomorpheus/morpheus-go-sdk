package morpheus

import (
	"fmt"
)

var (
	// ServicePlansPath is the API endpoint for servicePlans
	ServicePlansPath = "/api/service-plans"
)

// ServicePlans structures for use in request and response payloads
type ServicePlan struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Active bool   `json:"active"`
}

// ListServicePlansResult structure parses the list servicePlans response payload
type ListServicePlansResult struct {
	ServicePlans *[]ServicePlan `json:"servicePlans"`
	Meta         *MetaResult    `json:"meta"`
}

// GetServicePlanResult structure parses the get servicePlan response payload
type GetServicePlanResult struct {
	ServicePlan *ServicePlan `json:"servicePlan"`
}

// CreateServicePlanResult structure parses the create servicePlan response payload
type CreateServicePlanResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	ID      int64             `json:"id"`
}

// UpdateServicePlanResult structure parses the update servicePlan response payload
type UpdateServicePlanResult struct {
	CreateServicePlanResult
}

// DeleteServicePlanResult structure parses the delete servicePlan response payload
type DeleteServicePlanResult struct {
	DeleteResult
}

// ListServicePlans lists all service plans
// https://apidocs.morpheusdata.com/#get-all-service-plans
func (client *Client) ListServicePlans(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        ServicePlansPath,
		QueryParams: req.QueryParams,
		Result:      &ListServicePlansResult{},
	})
}

// GetServicePlan gets an existing service plan
// https://apidocs.morpheusdata.com/#get-a-specific-service-plan
func (client *Client) GetServicePlan(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", ServicePlansPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetServicePlanResult{},
	})
}

// CreateServicePlan creates a new servicePlan
// https://apidocs.morpheusdata.com/#create-a-service-plan
func (client *Client) CreateServicePlan(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        ServicePlansPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateServicePlanResult{},
	})
}

// UpdateServicePlan updates an existing servicePlan
// https://apidocs.morpheusdata.com/#update-a-servicePlan
func (client *Client) UpdateServicePlan(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", ServicePlansPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateServicePlanResult{},
	})
}

// DeleteServicePlan deletes an existing servicePlan
// https://apidocs.morpheusdata.com/#delete-a-service-plan
func (client *Client) DeleteServicePlan(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", ServicePlansPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteServicePlanResult{},
	})
}

// DeactivateServicePlan deactivates an existing service plan
// https://apidocs.morpheusdata.com/#deactivate-a-service-plan
func (client *Client) DeactivateServicePlan(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d/deactivate", ServicePlansPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteServicePlanResult{},
	})
}

// FindServicePlanByName gets an existing servicePlan by name
func (client *Client) FindServicePlanByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListServicePlans(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListServicePlansResult)
	servicePlanCount := len(*listResult.ServicePlans)
	if servicePlanCount != 1 {
		return resp, fmt.Errorf("found %d ServicePlans for %v", servicePlanCount, name)
	}
	firstRecord := (*listResult.ServicePlans)[0]
	servicePlanID := firstRecord.ID
	return client.GetServicePlan(servicePlanID, &Request{})
}
