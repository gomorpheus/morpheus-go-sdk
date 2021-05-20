// Morpheus API types and Client methods for Plans
package morpheus

import (
	"fmt"
)

// globals
var (
	PlansPath = "/api/service-plans"
)

// Plan structures for use in request and response payloads

type Plan struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Visibility  string `json:"visibility"`
}

type ListPlansResult struct {
	Plans *[]Plan     `json:"servicePlans"`
	Meta  *MetaResult `json:"meta"`
}

type GetPlanResult struct {
	Plan *Plan `json:"servicePlan"`
}

type CreatePlanResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Plan    *Plan             `json:"servicePlan"`
}

type UpdatePlanResult struct {
	CreatePlanResult
}

type DeletePlanResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListPlans(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        PlansPath,
		QueryParams: req.QueryParams,
		Result:      &ListPlansResult{},
	})
}

func (client *Client) GetPlan(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", PlansPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetPlanResult{},
	})
}

func (client *Client) CreatePlan(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        PlansPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreatePlanResult{},
	})
}

func (client *Client) UpdatePlan(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", PlansPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdatePlanResult{},
	})
}

func (client *Client) DeletePlan(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", PlansPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeletePlanResult{},
	})
}

// helper functions
func (client *Client) FindPlanByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListPlans(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListPlansResult)
	plansCount := len(*listResult.Plans)
	if plansCount != 1 {
		return resp, fmt.Errorf("found %d Plans for %v", plansCount, name)
	}
	firstRecord := (*listResult.Plans)[0]
	planID := firstRecord.ID
	return client.GetPlan(planID, &Request{})
}
