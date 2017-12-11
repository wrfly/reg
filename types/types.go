package types

type Credential struct {
	UserName string
	PassWord string
}

// {
// 	"auths": {
// 		"registry.kfd.me": {
// 			"auth": "YWRtaW46cGFzcw=="
// 		}
// 	}
// }

// DockerConfig is ~/.docker/config.json
type DockerConfig struct {
	Auths map[string]struct {
		Auth string `json:"auth"`
	} `json:"auths"`
}
