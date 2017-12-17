package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/wrfly/reg/types"
)

func (r *Registry) ListRepos(num int, last string) ([]string, error) {
	// generate url
	URI := catalogURI
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

func (r *Registry) ListTags(repo string, filter types.TagsFilter) ([]string, error) {
	logrus.Debugf("list repo tags: %s", repo)
	// generate url
	URI := fmt.Sprintf(tagsURI, repo)

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
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("Repo not found")
		}
		return nil, fmt.Errorf("Server Error: %s", bs)
	}

	tags := Tags{}
	err = json.Unmarshal(bs, &tags)
	if err != nil {
		return nil, err
	}

	return tags.Tags, nil
}

func (r *Registry) DeleteImages(repo, tag string) error {
	logrus.Debugf("delete image %s:%s", repo, tag)
	// get digest
	URI := fmt.Sprintf(manifestURI, repo, tag)
	resp, err := r.getJson(URI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// ot found means deleted
		if resp.StatusCode == http.StatusNotFound {
			logrus.Debugf("image not found")
			return nil
		}
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Server Error: %s", bs)
	}

	digest := resp.Header.Get("Docker-Content-Digest")
	logrus.Debugf("delete manifest: digest: %s", digest)

	// generate url
	URI = fmt.Sprintf(deleteURI, repo, digest)

	resp, err = r.getJson(URI)
	if err != nil {
		return err
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusAccepted {
		// ot found means deleted
		if resp.StatusCode == http.StatusNotFound {
			return nil
		}
		return fmt.Errorf("Server Error: %s", bs)
	}

	logrus.Debugf("image %s:%s deleted", repo, tag)
	return nil
}
