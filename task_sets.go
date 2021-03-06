// Morpheus API types and Client methods for Task Sets
package morpheus

import (
	"fmt"
)

// globals

var (
	TaskSetsPath = "/api/task-sets"
)

// TaskSet structures for use in request and response payloads

type TaskSet struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Type        string        `json:"type"`
	OptionTypes []interface{} `json:"optionTypes"`
	Tasks       []int64       `json:"tasks"`
}

type TaskSetPayload struct {
	ID                int64         `json:"id"`
	Name              string        `json:"name"`
	Description       string        `json:"description"`
	Type              string        `json:"type"`
	Visibility        string        `json:"visibility"`
	Platform          string        `json:"platform"`
	AllowCustomConfig bool          `json:"allowCustomConfig"`
	OptionTypes       []interface{} `json:"optionTypes"`
	Tasks             []int64       `json:"tasks"`
}

type ListTaskSetsResult struct {
	TaskSets *[]TaskSet  `json:"taskSets"`
	Meta     *MetaResult `json:"meta"`
}

type GetTaskSetResult struct {
	TaskSet *TaskSetPayload `json:"taskSet"`
}

type CreateTaskSetResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	TaskSet *TaskSetPayload   `json:"taskSet"`
}

type UpdateTaskSetResult struct {
	CreateTaskSetResult
}

type DeleteTaskSetResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListTaskSets(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        TaskSetsPath,
		QueryParams: req.QueryParams,
		Result:      &ListTaskSetsResult{},
	})
}

func (client *Client) GetTaskSet(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", TaskSetsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetTaskSetResult{},
	})
}

func (client *Client) CreateTaskSet(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        TaskSetsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateTaskSetResult{},
	})
}

func (client *Client) UpdateTaskSet(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", TaskSetsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateTaskSetResult{},
	})
}

func (client *Client) DeleteTaskSet(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", TaskSetsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteTaskSetResult{},
	})
}

// helper functions

func (client *Client) FindTaskSetByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListTaskSets(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListTaskSetsResult)
	taskSetCount := len(*listResult.TaskSets)
	if taskSetCount != 1 {
		return resp, fmt.Errorf("found %d TaskSets for %v", taskSetCount, name)
	}
	firstRecord := (*listResult.TaskSets)[0]
	taskSetID := firstRecord.ID
	return client.GetTaskSet(taskSetID, &Request{})
}
