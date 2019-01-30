package iis

type Website struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type WebsiteListResponse struct {
	Websites []Website `json:"websites"`
}

func (client Client) ListWebsites() ([]Website, error) {
	var res WebsiteListResponse
	err := getJson(client, "/api/webserver/websites", &res)
	if err != nil {
		return nil, err
	}
	return res.Websites, nil
}
