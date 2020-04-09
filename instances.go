// Morpheus API types and Client methods for Instances (Sites)
package morpheus

import (
    "fmt"
    "time"
)

// globals

var (
	InstancesPath = "/api/instances"
)

// types

type Instance struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	InstanceType map[string]interface{} `json:"instanceType"`
	Layout map[string]interface{} `json:"layout"`
	Group map[string]interface{} `json:"group"`
	Cloud map[string]interface{} `json:"cloud"`
	Environment string `json:"instanceContext"`
	Plan InstancePlan `json:"plan"`
	Config map[string]interface{} `json:"config"`

	// might want to define types for these too
	Volumes *[]map[string]interface{} `json:"volumes"`
	Interfaces *[]map[string]interface{} `json:"interfaces"`
	Controllers *[]map[string]interface{} `json:"controllers"`
	Tags *[]string `json:"tags"`
	Metadata *[]map[string]interface{} `json:"metadata"`
	EnvironmentVaribles *[]map[string]interface{} `json:"evars"`

}

type InstancePlan struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type ListInstancesResult struct {
    Instances *[]Instance `json:"instances"`
    Meta *MetaResult `json:"meta"`
}

type GetInstanceResult struct {
    Instance *Instance `json:"instance"`
}

type CreateInstanceResult struct {
	Success bool `json:"success"`
	Message string `json:"msg"`
	Errors map[string]string `json:"errors"`
	Instance *Instance `json:"instance"`
}

type UpdateInstanceResult struct {
	CreateInstanceResult
}

type DeleteInstanceResult struct {
	DeleteResult
}

type ListInstancePlansResult struct {
    Plans *[]InstancePlan `json:"plans"`
    Meta *MetaResult `json:"meta"`
}

type GetInstancePlanResult struct {
    Plan *InstancePlan `json:"plan"`
}

// API endpoints

func (client * Client) ListInstances(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path: InstancesPath,
		QueryParams: req.QueryParams,
		Result: &ListInstancesResult{},
	})
}

func (client * Client) GetInstance(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path: fmt.Sprintf("%s/%d", InstancesPath, id),
		QueryParams: req.QueryParams,
		Result: &GetInstanceResult{},
	})
}

func (client * Client) CreateInstance(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "POST",
		Path: InstancesPath,
		QueryParams: req.QueryParams,
		Body: req.Body,
		Result: &CreateInstanceResult{},
	})
}

func (client * Client) UpdateInstance(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "PUT",
		Path: fmt.Sprintf("%s/%d", InstancesPath, id),
		QueryParams: req.QueryParams,
		Body: req.Body,
		Result: &UpdateInstanceResult{},
	})
}


func (client * Client) DeleteInstance(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "DELETE",
		Path: fmt.Sprintf("%s/%d", InstancesPath, id),
		QueryParams: req.QueryParams,
		Body: req.Body,
		Result: &DeleteInstanceResult{},
	})
}

// helper functions

func (client * Client) FindInstanceByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListInstances(&Request{
		QueryParams:map[string]string{
			"name": name,
      	},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListInstancesResult)
	instanceCount := len(*listResult.Instances)
	if instanceCount != 1 {
		return resp, fmt.Errorf("Found %d Instances for %v", instanceCount, name)
	}
	firstRecord := (*listResult.Instances)[0]
	instanceId := firstRecord.ID
	return client.GetInstance(instanceId, &Request{})
}

// Plan fetching
// todo: this needs to be refactored soon

func (client * Client) ListInstancePlans(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path: fmt.Sprintf("%s/service-plans", InstancesPath),
		QueryParams: req.QueryParams,
		Result: &ListInstancePlansResult{},
	})
}

//todo: need this api endpoint still, and consolidate to /api/plans perhaps
func (client * Client) GetInstancePlan(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path: fmt.Sprintf("%s/service-plans/%d", InstancesPath, id),
		QueryParams: req.QueryParams,
		Result: &GetInstancePlanResult{},
	})
}

func (client * Client) FindInstancePlanByName(name string, req *Request) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListInstancePlans(&Request{
		QueryParams:map[string]string{
			//"name": name, // this is not even supported..
			"zoneId": req.QueryParams["zoneId"], // todo: use cloudId
			"layoutId": req.QueryParams["layoutId"],
			"siteId": req.QueryParams["siteId"], // todo: use groupId
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
		if plan.Name == name ||  plan.Code == name || string(plan.ID) == name {
			matchingPlans = append(matchingPlans, plan)	
		}
	}
	matchingPlanCount := len(matchingPlans)
	if matchingPlanCount != 1 {
		return resp, fmt.Errorf("Found %d Plans for '%v'", matchingPlanCount, name)
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
    	Success: true,
    	StatusCode: 200,
    	Status: "200 OK",
    	ReceivedAt: time.Now(),
    	// Size: restyResponse.Size(),
    	// Body: restyResponse.Body(), // byte[]
    	Result: result,
    }
    return mockResp, nil
}

// this should work by code or name
// it also requires zoneId AND layoutId??
func (client * Client) FindInstancePlanByCode(code string, req *Request) (*Response, error) {
	return client.FindInstancePlanByName(code, req)
}

