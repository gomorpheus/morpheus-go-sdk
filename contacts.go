// Morpheus API types and Client methods for Option Types
package morpheus

import (
	"fmt"
)

// globals

var (
	ContactsPath = "/api/monitoring/contacts"
)

// Contact structures for use in request and response payloads

type Contact struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	SmsAddress   string `json:"smsAddress"`
	SlackHook    string `json:"slackHook"`
}

type ListContactsResult struct {
	Contacts *[]Contact  `json:"contacts"`
	Meta     *MetaResult `json:"meta"`
}

type GetContactResult struct {
	Contact *Contact `json:"contact"`
}

type CreateContactResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Contact *Contact          `json:"contact"`
}

type UpdateContactResult struct {
	CreateContactResult
}

type DeleteContactResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListContacts(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        ContactsPath,
		QueryParams: req.QueryParams,
		Result:      &ListContactsResult{},
	})
}

func (client *Client) GetContact(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", ContactsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetContactResult{},
	})
}

func (client *Client) CreateContact(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        ContactsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateContactResult{},
	})
}

func (client *Client) UpdateContact(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", ContactsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateContactResult{},
	})
}

func (client *Client) DeleteContact(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", ContactsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteContactResult{},
	})
}

// helper functions
func (client *Client) FindContactByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListContacts(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListContactsResult)
	contactCount := len(*listResult.Contacts)
	if contactCount != 1 {
		return resp, fmt.Errorf("found %d Contacts for %v", contactCount, name)
	}
	firstRecord := (*listResult.Contacts)[0]
	contactID := firstRecord.ID
	return client.GetContact(contactID, &Request{})
}
