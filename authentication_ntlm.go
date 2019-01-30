package iis

import (
	"encoding/json"
	"fmt"
)

type WindowsAuthentication struct {
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
}

func (windows WindowsAuthentication) ToMap() map[string]interface{} {
	windowsMap := make(map[string]interface{}, 1)
	windowsMap["enabled"] = windows.Enabled

	return windowsMap
}

func (client Client) UpdateWindowsAuthentication(auth *WindowsAuthentication) (*WindowsAuthentication, error) {
	url := fmt.Sprintf("/api/webserver/authentication/windows-authentication/%s", auth.ID)
	res, err := httpPatch(client, url, &auth)
	if err != nil {
		return nil, err
	}
	var windows WindowsAuthentication
	err = json.Unmarshal(res, &windows)
	if err != nil {
		return nil, err
	}
	return &windows, nil
}
