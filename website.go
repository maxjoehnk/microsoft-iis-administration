package iis

type Website struct {
	Name         string           `json:"name"`
	ID           string           `json:"id"`
	PhysicalPath string           `json:"physical_path"`
	Bindings     []WebsiteBinding `json:"bindings"`
}

type WebsiteBinding struct {
	Protocol  string `json:"protocol"`
	Port      int64  `json:"port"`
	IPAddress string `json:"ip_address"`
	Hostname  string `json:"hostname"`
}
