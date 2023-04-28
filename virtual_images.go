package morpheus

import (
	"fmt"
	"time"
)

var (
	// VirtualImagesPath is the API endpoint for virtual images
	VirtualImagesPath = "/api/virtual-images"
)

// VirtualImage structures for use in request and response payloads
type VirtualImage struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	OwnerID     int64       `json:"ownerId"`
	Tenant      struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"tenant"`
	ImageType       string      `json:"imageType"`
	UserUploaded    bool        `json:"userUploaded"`
	UserDefined     bool        `json:"userDefined"`
	SystemImage     bool        `json:"systemImage"`
	IsCloudInit     bool        `json:"isCloudInit"`
	SshUsername     interface{} `json:"sshUsername"`
	SshPassword     interface{} `json:"sshPassword"`
	SshPasswordHash interface{} `json:"sshPasswordHash"`
	SshKey          interface{} `json:"sshKey"`
	OsType          struct {
		ID          int64       `json:"id"`
		Code        string      `json:"code"`
		Name        string      `json:"name"`
		Description interface{} `json:"description"`
		Vendor      string      `json:"vendor"`
		Category    string      `json:"category"`
		OsFamily    string      `json:"osFamily"`
		OsVersion   string      `json:"osVersion"`
		BitCount    int64       `json:"bitCount"`
		Platform    string      `json:"platform"`
	} `json:"osType"`
	MinRam               int64       `json:"minRam"`
	MinRamGB             float32     `json:"minRamGB"`
	MinDisk              int64       `json:"minDisk"`
	MinDiskGB            float32     `json:"minDiskGB"`
	RawSize              interface{} `json:"rawSize"`
	RawSizeGB            interface{} `json:"rawSizeGB"`
	TrialVersion         bool        `json:"trialVersion"`
	VirtioSupported      bool        `json:"virtioSupported"`
	Uefi                 bool        `json:"uefi"`
	IsAutoJoinDomain     bool        `json:"isAutoJoinDomain"`
	VmtoolsInstalled     bool        `json:"vmToolsInstalled"`
	InstallAgent         bool        `json:"installAgent"`
	IsForceCustomization bool        `json:"isForceCustomization"`
	IsSysprep            bool        `json:"isSysprep"`
	FipsEnabled          bool        `json:"fipsEnabled"`
	UserData             interface{} `json:"userData"`
	ConsoleKeymap        interface{} `json:"consoleKeymap"`
	StorageProvider      interface{} `json:"storageProvider"`
	ExternalID           string      `json:"externalId"`
	Visibility           string      `json:"visibility"`
	Accounts             []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"accounts"`
	Config struct {
	} `json:"config"`
	Volumes []struct {
		Name       string `json:"name"`
		MaxStorage int64  `json:"maxStorage"`
		RawSize    int64  `json:"rawSize"`
		Size       int64  `json:"size"`
		RootVolume bool   `json:"rootVolume"`
		Resizeable bool   `json:"resizeable"`
	} `json:"volumes"`
	StorageControllers []struct {
		Name string `json:"name"`
		Type struct {
			ID   int64  `json:"id"`
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"type"`
		MaxDevices         int64 `json:"maxDevices"`
		ReservedUnitNumber int64 `json:"reservedUnitNumber"`
	} `json:"storageControllers"`
	NetworkInterfaces []interface{} `json:"networkInterfaces"`
	Tags              []interface{} `json:"tags"`
	Locations         []struct {
		ID    int64 `json:"id"`
		Cloud struct {
			ID   int64  `json:"id"`
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"cloud"`
		Code           string      `json:"code"`
		InternalID     string      `json:"internalId"`
		ExternalID     string      `json:"externalId"`
		ExternalDiskID interface{} `json:"externalDiskId"`
		RemotePath     interface{} `json:"remotePath"`
		ImagePath      interface{} `json:"imagePath"`
		ImageName      string      `json:"imageName"`
		ImageRegion    string      `json:"imageRegion"`
		ImageFolder    interface{} `json:"imageFolder"`
		RefType        string      `json:"refType"`
		RefID          int64       `json:"refId"`
		NodeRefType    interface{} `json:"nodeRefType"`
		NodeRefID      interface{} `json:"nodeRefId"`
		SubRefType     interface{} `json:"subRefType"`
		SubRefID       interface{} `json:"subRefId"`
		IsPublic       bool        `json:"isPublic"`
		SystemImage    bool        `json:"systemImage"`
		DiskIndex      int64       `json:"diskIndex"`
	} `json:"locations"`
	DateCreated time.Time `json:"dateCreated"`
	LastUpdated time.Time `json:"lastUpdated"`
	Status      string    `json:"status"`
}

type ListVirtualImagesResult struct {
	VirtualImages *[]VirtualImage `json:"virtualImages"`
	Meta          *MetaResult     `json:"meta"`
}

type GetVirtualImageResult struct {
	VirtualImage *VirtualImage `json:"virtualImage"`
}

type CreateVirtualImageResult struct {
	Success      bool              `json:"success"`
	Message      string            `json:"msg"`
	Errors       map[string]string `json:"errors"`
	VirtualImage *VirtualImage     `json:"virtualImage"`
}

type UpdateVirtualImageResult struct {
	CreateVirtualImageResult
}

type DeleteVirtualImageResult struct {
	DeleteResult
}

// ListVirtualImages lists all virtual images
func (client *Client) ListVirtualImages(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        VirtualImagesPath,
		QueryParams: req.QueryParams,
		Result:      &ListVirtualImagesResult{},
	})
}

// GetVirtualImage gets an existing virtualimage
func (client *Client) GetVirtualImage(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", VirtualImagesPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetVirtualImageResult{},
	})
}

// CreateVirtualImage creates a new virtual image
func (client *Client) CreateVirtualImage(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        VirtualImagesPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateVirtualImageResult{},
	})
}

// UpdateVirtualImage updates an existing virtual image
func (client *Client) UpdateVirtualImage(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", VirtualImagesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateVirtualImageResult{},
	})
}

// DeleteVirtualImage deletes an existing virtual image
func (client *Client) DeleteVirtualImage(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", VirtualImagesPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteVirtualImageResult{},
	})
}

// FindVirtualImageByName gets an existing virtual image by name
func (client *Client) FindVirtualImageByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListVirtualImages(&Request{
		QueryParams: map[string]string{
			"name":       name,
			"filterType": "All",
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListVirtualImagesResult)
	virtualImagesCount := len(*listResult.VirtualImages)
	if virtualImagesCount != 1 {
		return resp, fmt.Errorf("found %d Virtual Images for %v", virtualImagesCount, name)
	}
	firstRecord := (*listResult.VirtualImages)[0]
	virtualImageID := firstRecord.ID
	return client.GetVirtualImage(virtualImageID, &Request{})
}
func (client *Client) FindVirtualImageByNameAndType(name string, imagetype string) (*Response, error) {
        // Find by name, then get by ID
        resp, err := client.ListVirtualImages(&Request{
                QueryParams: map[string]string{
                        "name":       name,
                        "filterType": "All",
                        "imageType": imagetype,
                },
        })
        if err != nil {
                return resp, err
        }
        listResult := resp.Result.(*ListVirtualImagesResult)
        virtualImagesCount := len(*listResult.VirtualImages)
        if virtualImagesCount != 1 {
                return resp, fmt.Errorf("found %d Virtual Images for %v", virtualImagesCount, name)
        }
        firstRecord := (*listResult.VirtualImages)[0]
        virtualImageID := firstRecord.ID
        return client.GetVirtualImage(virtualImageID, &Request{})
}

