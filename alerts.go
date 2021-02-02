// Morpheus API types and Client methods for Alerts
package morpheus

import (
	"fmt"
)

// globals
var (
	AlertsPath = "/api/monitoring/alerts"
)

// Alert structures for use in request and response payloads

type Alert struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Visibility  string `json:"visibility"`
}

type ListAlertsResult struct {
	Alerts *[]Alert    `json:"alerts"`
	Meta   *MetaResult `json:"meta"`
}

type GetAlertResult struct {
	Alert *Alert `json:"alert"`
}

type CreateAlertResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Alert   *Alert            `json:"alert"`
}

type UpdateAlertResult struct {
	CreateAlertResult
}

type DeleteAlertResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListAlerts(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        AlertsPath,
		QueryParams: req.QueryParams,
		Result:      &ListAlertsResult{},
	})
}

func (client *Client) GetAlert(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", AlertsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetAlertResult{},
	})
}

func (client *Client) CreateAlert(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        AlertsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateAlertResult{},
	})
}

func (client *Client) UpdateAlert(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", AlertsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateAlertResult{},
	})
}

func (client *Client) DeleteAlert(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", AlertsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteAlertResult{},
	})
}

// helper functions
func (client *Client) FindAlertByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListAlerts(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListAlertsResult)
	alertsCount := len(*listResult.Alerts)
	if alertsCount != 1 {
		return resp, fmt.Errorf("Found %d Alerts for %v", alertsCount, name)
	}
	firstRecord := (*listResult.Alerts)[0]
	alertID := firstRecord.ID
	return client.GetAlert(alertID, &Request{})
}
