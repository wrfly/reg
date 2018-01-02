package registry

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/wrfly/reg/types"
	"github.com/wrfly/reg/utils"
)

func (r *Registry) getJson(what string) (*http.Response, error) {
	// get method and uri from `what`
	method := strings.Split(what, " ")[0]
	URI := strings.Split(what, " ")[1]
	// use https first
	URL := fmt.Sprintf("https://%s%s", r.RegistryAddr, URI)
	c := http.Client{
		Timeout: time.Second * 5,
	}
	logrus.Debugf("%s %s", method, URL)

	// http stuff
	req, _ := http.NewRequest(method, URL, nil)
	req.Header.Set("Docker-Distribution-API-Version", "registry/2.0")
	// need this header to delete image
	req.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	resp, err := c.Do(req)
	if err != nil {
		logrus.Debug(err)
		if strings.Contains(err.Error(), "Client.Timeout exceeded") {
			return nil, fmt.Errorf("timeout")
		} else if strings.Contains(err.Error(), "server gave HTTP response to HTTPS client") {
			// use http
			req.URL.Scheme = "http"
			resp, err = c.Do(req)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// 401
	if resp.StatusCode == http.StatusUnauthorized {
		authStr := resp.Header.Get("WWW-Authenticate")
		s := strings.Split(authStr, " ")
		scheme := s[0]
		details := s[1]
		token, err := getToken(details, r.Credential)
		if err != nil {
			return resp, err
		}
		logrus.Debugf("scheme: %s, token: %s", scheme, token)
		req.Header.Set("Authorization", fmt.Sprintf("%s %s", scheme, token))
		resp, err = c.Do(req)
		if err != nil {
			return resp, err
		}
		if resp.StatusCode == http.StatusUnauthorized {
			return resp, fmt.Errorf("authentication failed")
		}
	}
	return resp, err
}

func getToken(authStr string, credential types.Credential) (string, error) {
	m := utils.String2Map(authStr)

	logrus.Debugf("realm: %v", m["realm"])
	// htpasswd auth
	if m["realm"] == "Registry" {
		s := fmt.Sprintf("%s:%s", credential.UserName, credential.PassWord)
		return utils.Base64Encode(s), nil
	}

	req, _ := http.NewRequest("GET", m["realm"], nil)
	q := req.URL.Query()
	if m["service"] != "" {
		q.Set("service", m["service"])
	}
	if m["scope"] != "" {
		q.Set("scope", m["scope"])
	}

	req.URL.RawQuery = q.Encode()
	logrus.Debugf("get token url: %s", req.URL.String())

	req.SetBasicAuth(credential.UserName, credential.PassWord)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return "", nil
}
