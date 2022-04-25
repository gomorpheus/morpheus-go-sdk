package morpheus

import (
	"fmt"
)

var (
	ProvisionTypesPath = "/api/provision-types"
)

// Provision Type structures for use in request and response payloads
type ProvisionType struct {
	ID                   int         `json:"id"`
	Name                 string      `json:"name"`
	Description          interface{} `json:"description"`
	Code                 string      `json:"code"`
	Aclenabled           bool        `json:"aclEnabled"`
	Multitenant          bool        `json:"multiTenant"`
	Managed              bool        `json:"managed"`
	Hostnetwork          bool        `json:"hostNetwork"`
	Customsupported      bool        `json:"customSupported"`
	Mapports             bool        `json:"mapPorts"`
	Exportserver         interface{} `json:"exportServer"`
	Viewset              string      `json:"viewSet"`
	Servertype           string      `json:"serverType"`
	Hosttype             string      `json:"hostType"`
	Addvolumes           bool        `json:"addVolumes"`
	Hasdatastore         bool        `json:"hasDatastore"`
	Hasnetworks          interface{} `json:"hasNetworks"`
	Maxnetworks          interface{} `json:"maxNetworks"`
	Customizevolume      bool        `json:"customizeVolume"`
	Rootdiskcustomizable bool        `json:"rootDiskCustomizable"`
	Lvmsupported         bool        `json:"lvmSupported"`
	Hostdiskmode         string      `json:"hostDiskMode"`
	Mindisk              int         `json:"minDisk"`
	Maxdisk              interface{} `json:"maxDisk"`
	Resizecopiesvolumes  bool        `json:"resizeCopiesVolumes"`
	Optiontypes          []struct {
		Name         string      `json:"name"`
		Description  interface{} `json:"description"`
		Fieldname    string      `json:"fieldName"`
		Fieldlabel   string      `json:"fieldLabel"`
		Fieldcontext string      `json:"fieldContext"`
		Fieldaddon   interface{} `json:"fieldAddOn"`
		Placeholder  interface{} `json:"placeHolder"`
		Helpblock    string      `json:"helpBlock"`
		Defaultvalue interface{} `json:"defaultValue"`
		Optionsource string      `json:"optionSource"`
		Type         string      `json:"type"`
		Advanced     bool        `json:"advanced"`
		Required     bool        `json:"required"`
		Editable     bool        `json:"editable"`
		Displayorder int         `json:"displayOrder"`
	} `json:"optionTypes"`
	Customoptiontypes []interface{} `json:"customOptionTypes"`
	Networktypes      []interface{} `json:"networkTypes"`
	Storagetypes      []struct {
		ID                int         `json:"id"`
		Code              string      `json:"code"`
		Name              string      `json:"name"`
		Displayorder      int         `json:"displayOrder"`
		Defaulttype       bool        `json:"defaultType"`
		Customlabel       bool        `json:"customLabel"`
		Customsize        bool        `json:"customSize"`
		Customsizeoptions interface{} `json:"customSizeOptions"`
	} `json:"storageTypes"`
	Rootstoragetypes []struct {
		ID                int         `json:"id"`
		Code              string      `json:"code"`
		Name              string      `json:"name"`
		Displayorder      int         `json:"displayOrder"`
		Defaulttype       bool        `json:"defaultType"`
		Customlabel       bool        `json:"customLabel"`
		Customsize        bool        `json:"customSize"`
		Customsizeoptions interface{} `json:"customSizeOptions"`
	} `json:"rootStorageTypes"`
	Controllertypes []interface{} `json:"controllerTypes"`
}

// ListPricesResult structure parses the list prices response payload
type ListProvisionTypeResult struct {
	ProvisionTypes *[]ProvisionType `json:"provisionTypes"`
	Meta           *MetaResult      `json:"meta"`
}

// GetPriceResult structure parses the get price response payload
type GetProvisionTypeResult struct {
	ProvisionType *ProvisionType `json:"provisionType"`
}

// API endpoints
// ListPrices lists all prices
// https://apidocs.morpheusdata.com/#get-all-prices
func (client *Client) ListProvisionTypes(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        ProvisionTypesPath,
		QueryParams: req.QueryParams,
		Result:      &ListProvisionTypeResult{},
	})
}

// GetPrice gets an existing price
// https://apidocs.morpheusdata.com/#get-a-specific-price
func (client *Client) GetProvisionType(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", ProvisionTypesPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetProvisionTypeResult{},
	})
}
