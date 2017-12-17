package cmds

import (
	"github.com/wrfly/reg/registry"
	"github.com/wrfly/reg/types"
)

var r registry.Registry

func SetRegistry(registryAddr string, credential types.Credential) {
	r = registry.Registry{
		RegistryAddr: registryAddr,
		Credential:   credential,
	}
}
