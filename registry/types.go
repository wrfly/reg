package registry

import "github.com/wrfly/reg/types"

const (
	apiVersion = 2
	catalogURI = "/v2/_catalog"
)

type Registry struct {
	RegistryAddr string
	Credential   types.Credential
}

type Catalog struct {
	Repositories []string `json:"repositories"`
}
