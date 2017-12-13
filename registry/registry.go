package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (r *Registry) ListRepos(num int, last string) ([]string, error) {
	// generate url
	URI := fmt.Sprintf("%s", catalogURI)
	if num != 0 {
		URI = fmt.Sprintf("%s?n=%d&last=%s", catalogURI, num, last)
	}

	resp, err := r.getJson(URI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Server Error: %s", bs)
	}

	repos := Catalog{}
	err = json.Unmarshal(bs, &repos)
	if err != nil {
		return nil, err
	}

	return repos.Repositories, nil
}
