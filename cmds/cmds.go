package cmds

import "github.com/wrfly/reg/types"

var (
	registryAddr string
	credential   types.Credential
)

func SetRegistry(addr string, cre types.Credential) {
	registryAddr = addr
	credential = cre
}
