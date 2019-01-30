package iis

import (
	"encoding/json"
	"fmt"
)

func (client Client) UpdateAppPool(id, name string) (*ApplicationPool, error) {
	reqBody := struct {
		Name string `json:"name"`
	}{name}
	url := fmt.Sprintf("/api/webserver/application-pools/%s", id)
	res, err := httpPatch(client, url, reqBody)
	if err != nil {
		return nil, err
	}
	var pool ApplicationPool
	err = json.Unmarshal(res, &pool)
	if err != nil {
		return nil, err
	}
	return &pool, nil
}
