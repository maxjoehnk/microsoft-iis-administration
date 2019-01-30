package iis

import (
	"encoding/json"
	"fmt"
)

type AnonymousAuthentication struct {
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
	User    string `json:"user"`
}

func (anonymous AnonymousAuthentication) ToMap() map[string]interface{} {
	anonymousMap := make(map[string]interface{}, 2)
	anonymousMap["enabled"] = anonymous.Enabled
	anonymousMap["user"] = anonymous.User

	return anonymousMap
}

func (client Client) ReadAnonymousAuthenticationFromId(id string) (*AnonymousAuthentication, error) {
	url := fmt.Sprintf("/api/webserver/authentication/anonymous-authentication/%s", id)
	var auth AnonymousAuthentication
	err := getJson(client, url, &auth)
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

func (client Client) UpdateAnonymousAuthentication(auth *AnonymousAuthentication) (AnonymousAuthentication, error) {
	var anonymous AnonymousAuthentication
	url := fmt.Sprintf("/api/webserver/authentication/anonymous-authentication/%s", auth.ID)
	res, err := httpPatch(client, url, &auth)
	if err != nil {
		return anonymous, err
	}
	err = json.Unmarshal(res, &anonymous)
	if err != nil {
		return anonymous, err
	}
	return anonymous, nil
}
