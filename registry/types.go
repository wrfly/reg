package registry

import "github.com/wrfly/reg/types"

const (
	apiVersion = 2
	catalogURI = "/v2/_catalog"
	tagsURI    = "/v2/%s/tags/list"
)

type Registry struct {
	RegistryAddr string
	Credential   types.Credential
}

type Catalog struct {
	Repositories []string `json:"repositories"`
}

type Tags struct{
	Name string `json:"name"`
	Tags []string  `json:"tags"`
}