// Morpheus API types and Client methods for Apps
package morpheus

import (
	"fmt"
)

// globals
var (
	AppsPath = "/api/apps"
)

// App structures for use in request and response payloads

type App struct {
	ID          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Group       map[string]interface{} `json:"group"`
	Environment string                 `json:"environment"`
	BlueprintID string                 `json:"blueprintId"`
}

type ListAppsResult struct {
	Apps *[]App      `json:"apps"`
	Meta *MetaResult `json:"meta"`
}

type GetAppResult struct {
	App *App `json:"app"`
}

type CreateAppResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	App     *App              `json:"app"`
}

type UpdateAppResult struct {
	CreateAppResult
}

type DeleteAppResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListApps(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        AppsPath,
		QueryParams: req.QueryParams,
		Result:      &ListAppsResult{},
	})
}

func (client *Client) GetApp(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", AppsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetAppResult{},
	})
}

func (client *Client) CreateApp(req *Request) (*Response, error) {
	fmt.Println(req.Body)
	return client.Execute(&Request{
		Method:      "POST",
		Path:        AppsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateAppResult{},
	})
}

func (client *Client) UpdateApp(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", AppsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateAppResult{},
	})
}

func (client *Client) DeleteApp(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", AppsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteAppResult{},
	})
}

// helper functions
func (client *Client) FindAppByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListApps(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListAppsResult)
	appsCount := len(*listResult.Apps)
	if appsCount != 1 {
		return resp, fmt.Errorf("found %d Apps for %v", appsCount, name)
	}
	firstRecord := (*listResult.Apps)[0]
	appID := firstRecord.ID
	return client.GetApp(appID, &Request{})
}
