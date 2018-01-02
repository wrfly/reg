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
	"gopkg.in/urfave/cli.v2"
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

func CmdBefore(c *cli.Context) error {
	// set log-level
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}
	return nil
}

// a="1",b="2" -> map["a":"1","b":"2"]
func String2Map(str string) map[string]string {
	pairs := strings.Split(str, ",")
	maps := make(map[string]string, 0)
	for _, pair := range pairs {
		p := strings.Split(pair, "=")
		if len(p) != 2 {
			continue
		}
		k := p[0]
		v := strings.Replace(p[1], "\"", "", -1)
		maps[k] = v
	}

	return maps
}

func Base64Encode(in string) string {
	n := base64.StdEncoding.EncodedLen(len([]byte(in)))
	dst := make([]byte, n)
	base64.StdEncoding.Encode(dst, []byte(in))
	return fmt.Sprintf("%s", dst)
}
