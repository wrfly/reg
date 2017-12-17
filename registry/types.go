package registry

import "github.com/wrfly/reg/types"

const (
	apiVersion = 2
	// GET /v2/_catalog
	catalogURI = "GET /v2/_catalog"
	// GET /v2/<name>/tags/list
	tagsURI = "GET /v2/%s/tags/list"
	// DELETE /v2/<name>/manifests/<reference>
	deleteURI = "DELETE /v2/%s/manifests/%s"
	// GET /v2/<name>/manifests/<reference>
	manifestURI = "GET /v2/%s/manifests/%s"
)

type Registry struct {
	RegistryAddr string
	Credential   types.Credential
}

type Catalog struct {
	Repositories []string `json:"repositories"`
}

type Tags struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}
