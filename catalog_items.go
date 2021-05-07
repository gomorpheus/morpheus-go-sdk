// Morpheus API types and Client methods for Catalog Items
package morpheus

import (
	"fmt"
)

// globals

var (
	CatalogItemsPath = "/api/catalog-item-types"
)

// CatalogItem structures for use in request and response payloads

type CatalogItem struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Featured    bool    `json:"featured"`
	Enabled     bool    `json:"enabled"`
	OptionTypes []int64 `json:"optionTypes"`
	Context     string  `json:"context"`
	Content     string  `json:"content"`
}

type ListCatalogItemsResult struct {
	CatalogItems *[]CatalogItem `json:"catalogItemTypes"`
	Meta         *MetaResult    `json:"meta"`
}

type GetCatalogItemResult struct {
	CatalogItem *CatalogItem `json:"catalogItemType"`
}

type CreateCatalogItemResult struct {
	Success     bool              `json:"success"`
	Message     string            `json:"msg"`
	Errors      map[string]string `json:"errors"`
	CatalogItem *CatalogItem      `json:"catalogItemType"`
}

type UpdateCatalogItemResult struct {
	CreateCatalogItemResult
}

type DeleteCatalogItemResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListCatalogItems(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        CatalogItemsPath,
		QueryParams: req.QueryParams,
		Result:      &ListCatalogItemsResult{},
	})
}

func (client *Client) GetCatalogItem(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", CatalogItemsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetCatalogItemResult{},
	})
}

func (client *Client) CreateCatalogItem(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        CatalogItemsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateCatalogItemResult{},
	})
}

func (client *Client) UpdateCatalogItem(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", CatalogItemsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateCatalogItemResult{},
	})
}

func (client *Client) DeleteCatalogItem(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", CatalogItemsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteCatalogItemResult{},
	})
}

// helper functions

func (client *Client) FindCatalogItemByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListCatalogItems(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListCatalogItemsResult)
	catalogItemCount := len(*listResult.CatalogItems)
	if catalogItemCount != 1 {
		return resp, fmt.Errorf("Found %d CatalogItems for %v", catalogItemCount, name)
	}
	firstRecord := (*listResult.CatalogItems)[0]
	optionTypeID := firstRecord.ID
	return client.GetCatalogItem(optionTypeID, &Request{})
}
