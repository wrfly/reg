package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/wrfly/reg/types"
)

func ParseDockerCondig() (types.DockerConfig, error) {
	dc := types.DockerConfig{}

	HOME := os.Getenv("HOME")
	config := path.Join(HOME, ".docker/config.json")
	f, err := os.Open(config)
	if err != nil {
		if os.IsNotExist(err) {
			return dc, nil
		}
		return dc, err
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return dc, err
	}

	if err := json.Unmarshal(bs, &dc); err != nil {
		return dc, err
	}

	return dc, nil
}

func ParseAuth(auth string) (types.Credential, error) {
	logrus.Debug(auth)
	c := types.Credential{}
	bs, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return c, err
	}
	str := fmt.Sprintf("%s", bs)
	uAp := strings.Split(str, ":")
	if len(uAp) != 2 {
		return c, fmt.Errorf("bad credential!")
	}
	c.UserName = uAp[0]
	c.PassWord = uAp[1]

	return c, nil
}
