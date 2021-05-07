// Morpheus API types and Client methods for Tasks
package morpheus

import (
	"fmt"
)

// globals

var (
	TasksPath = "/api/tasks"
)

// Task structures for use in request and response payloads

type Task struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	ResultType string `json:"resultType"`
}

type ListTasksResult struct {
	Tasks *[]Task     `json:"tasks"`
	Meta  *MetaResult `json:"meta"`
}

type GetTaskResult struct {
	Task *Task `json:"task"`
}

type CreateTaskResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Task    *Task             `json:"task"`
}

type UpdateTaskResult struct {
	CreateTaskResult
}

type DeleteTaskResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListTasks(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        TasksPath,
		QueryParams: req.QueryParams,
		Result:      &ListTasksResult{},
	})
}

func (client *Client) GetTask(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", TasksPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetTaskResult{},
	})
}

func (client *Client) CreateTask(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        TasksPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateTaskResult{},
	})
}

func (client *Client) UpdateTask(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", TasksPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateTaskResult{},
	})
}

func (client *Client) DeleteTask(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", TasksPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteTaskResult{},
	})
}

// helper functions

func (client *Client) FindTaskByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListTasks(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListTasksResult)
	tenantsCount := len(*listResult.Tasks)
	if tenantsCount != 1 {
		return resp, fmt.Errorf("found %d Tasks for %v", tenantsCount, name)
	}
	firstRecord := (*listResult.Tasks)[0]
	tenantID := firstRecord.ID
	return client.GetTask(tenantID, &Request{})
}
