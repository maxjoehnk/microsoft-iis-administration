package iis

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Do a GET request and parse the body to json
func getJson(client Client, path string, r interface{}) error {
	data, err := httpGet(client, path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &r)
}

// Do a GET request
func httpGet(client Client, path string) ([]byte, error) {
	response, err := request(client, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	return fetchBody(response)
}

// Do a POST request
func httpPost(client Client, path string, body interface{}) ([]byte, error) {
	response, err := request(client, "POST", path, body)
	if err != nil {
		return nil, err
	}
	return fetchBody(response)
}

// Do a PATCH request
func httpPatch(client Client, path string, body interface{}) ([]byte, error) {
	response, err := request(client, "PATCH", path, body)
	if err != nil {
		return nil, err
	}
	return fetchBody(response)
}

// Do a DELETE request
func httpDelete(client Client, path string) error {
	if _, err := request(client, "DELETE", path, nil); err != nil {
		return err
	}
	return nil
}

func buildHttpClient() *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: transport}
}

func buildRequest(client Client, method, path string, body interface{}) (*http.Request, error) {
	log.Printf("%s %s", method, path)
	b := new(bytes.Buffer)
	if body != nil {
		if err := json.NewEncoder(b).Encode(body); err != nil {
			return nil, err
		}
	}

	url := fmt.Sprintf("%s%s", client.Host, path)
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Access-Token", fmt.Sprintf("Bearer %s", client.AccessKey))
	req.Header.Set("Accept", "application/hal+json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func executeRequest(req *http.Request) (*http.Response, error) {
	httpClient := buildHttpClient()
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if err := guardStatusCode(response); err != nil {
		return nil, err
	}
	return response, err
}

func request(client Client, method, path string, body interface{}) (*http.Response, error) {
	req, err := buildRequest(client, method, path, body)
	if err != nil {
		return nil, err
	}
	return executeRequest(req)
}

func fetchBody(res *http.Response) ([]byte, error) {
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err = res.Body.Close(); err != nil {
		return nil, err
	}
	return resBody, nil
}

func guardStatusCode(response *http.Response) error {
	if response.StatusCode < 200 || response.StatusCode > 400 {
		return fmt.Errorf("invalid status code: %d - %s", response.StatusCode, response.Status)
	}
	return nil
}
