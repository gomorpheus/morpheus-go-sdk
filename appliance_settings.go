package morpheus

import "fmt"

var (
	// ApplianceSettingsPath is the API endpoint for appliance settings
	ApplianceSettingsPath = "/api/appliance-settings"
)

// ApplianceSettings structures for use in request and response payloads
type ApplianceSettings struct {
	ApplianceURL             string `json:"applianceUrl"`
	InternalApplianceURL     string `json:"internalApplianceUrl"`
	CorsAllowed              string `json:"corsAllowed"`
	RegistrationEnabled      bool   `json:"registrationEnabled"`
	DefaultRoleID            string `json:"defaultRoleId"`
	DefaultUserRoleID        string `json:"defaultUserRoleId"`
	DockerPrivilegedMode     bool   `json:"dockerPrivilegedMode"`
	ExpirePwdDays            string `json:"expirePwdDays"`
	DisableAfterAttempts     string `json:"disableAfterAttempts"`
	DisableAfterDaysInactive string `json:"disableAfterDaysInactive"`
	WarnUserDaysBefore       string `json:"warnUserDaysBefore"`
	SMTPMailFrom             string `json:"smtpMailFrom"`
	SMTPServer               string `json:"smtpServer"`
	SMTPPort                 string `json:"smtpPort"`
	SMTPSSL                  bool   `json:"smtpSSL"`
	SMTPTLS                  bool   `json:"smtpTLS"`
	SMTPUser                 string `json:"smtpUser"`
	SMTPPassword             string `json:"smtpPassword"`
	SMTPPasswordHash         string `json:"smtpPasswordHash"`
	ProxyHost                string `json:"proxyHost"`
	ProxyPort                string `json:"proxyPort"`
	ProxyUser                string `json:"proxyUser"`
	ProxyPassword            string `json:"proxyPassword"`
	ProxyPasswordHash        string `json:"proxyPasswordHash"`
	ProxyDomain              string `json:"proxyDomain"`
	ProxyWorkstation         string `json:"proxyWorkstation"`
	CurrencyProvider         string `json:"currencyProvider"`
	CurrencyKey              string `json:"currencyKey"`
	EnabledZoneTypes         []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"enabledZoneTypes"`
	StatsRetainmentPeriod int64 `json:"statsRetainmentPeriod"`
}

type GetApplianceSettingsResult struct {
	ApplianceSettings *ApplianceSettings `json:"applianceSettings"`
}

type UpdateApplianceSettingsResult struct {
	Success           bool               `json:"success"`
	Message           string             `json:"msg"`
	Errors            map[string]string  `json:"errors"`
	ApplianceSettings *ApplianceSettings `json:"applianceSettings"`
}

type ToggleMaintenanceResult struct {
	StandardResult
}

// Client request methods
func (client *Client) GetApplianceSettings(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        ApplianceSettingsPath,
		QueryParams: req.QueryParams,
		Result:      &GetApplianceSettingsResult{},
	})
}

func (client *Client) UpdateApplianceSettings(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        ApplianceSettingsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateApplianceSettingsResult{},
	})
}

func (client *Client) ToggleMaintenanceMode(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        fmt.Sprintf("%s/maintenance", ApplianceSettingsPath),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &ToggleMaintenanceResult{},
	})
}
