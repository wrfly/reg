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

func (r *Registry) getJson(URI string) (*http.Response, error) {
	// use https first
	URL := fmt.Sprintf("https://%s%s", r.RegistryAddr, URI)
	c := http.Client{
		Timeout: time.Second * 5,
	}
	logrus.Debugf("get %s", URL)

	// http stuff
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("Docker-Distribution-API-Version", "registry/2.0")

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
		req.Header.Set("Authorization", fmt.Sprintf("%s %s", scheme, token))
		return resp, fmt.Errorf("authentication failed")
	}
	return resp, err
}

func getToken(authStr string, credential types.Credential) (string, error) {
	m := utils.String2Map(authStr)
	logrus.Debugf("realm: %v", m["realm"])

	req, _ := http.NewRequest("GET", m["realm"], nil)
	q := req.URL.Query()
	q.Set("service", m["service"])
	q.Set("scope", m["scope"])
	req.URL.RawQuery = q.Encode()
	logrus.Debug(req.URL.String())

	req.SetBasicAuth(credential.UserName, credential.PassWord)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return "", nil
}
