package morpheus

import (
	"fmt"
	"time"
)

var (
	// AppsPath is the API endpoint for apps
	AppsPath = "/api/apps"
)

// App structures for use in request and response payloads
type App struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
	Environment string   `json:"environment"`
	AccountId   int64    `json:"accountId"`
	Account     struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"account"`
	Owner struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
	} `json:"owner"`
	SiteId int64 `json:"siteId"`
	Group  struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"group"`
	Blueprint struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"blueprint"`
	Type           string    `json:"type"`
	DateCreated    time.Time `json:"dateCreated"`
	LastUpdated    time.Time `json:"lastUpdated"`
	AppContext     string    `json:"appContext"`
	Status         string    `json:"status"`
	AppStatus      string    `json:"appStatus"`
	InstanceCount  int64     `json:"instanceCount"`
	ContainerCount int64     `json:"containerCount"`
	AppTiers       []struct {
		Tier struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"tier"`
		AppInstances []struct {
			Config   interface{} `json:"config"`
			Instance Instance    `json:"instance"`
		} `json:"appInstances"`
		BootSequence int64 `json:"bootSequence"`
	} `json:"appTiers"`
	Instances []struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"instances"`
	Stats struct {
		UsedMemory            int64   `json:"usedMemory"`
		MaxMemory             int64   `json:"maxMemory"`
		UsedStorage           int64   `json:"usedStorage"`
		MaxStorage            int64   `json:"maxStorage"`
		Running               int64   `json:"running"`
		Total                 int64   `json:"total"`
		CpuUsage              float64 `json:"cpuUsage"`
		InstanceCount         int64   `json:"instanceCount"`
		InstanceDayCount      []int64 `json:"instanceDayCount"`
		InstanceDayCountTotal int64   `json:"instanceDayCountTotal"`
	} `json:"stats"`
}

// ListAppsResult structure parses the list apps response payload
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
