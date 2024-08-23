package gitlab

import (
	"encoding/json"
	"fmt"
	"github.com/ekristen/distillery/pkg/common"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

const baseURL = "https://gitlab.com/api/v4"

func NewClient(client *http.Client) *Client {
	return &Client{
		client: client,
	}
}

type Client struct {
	client *http.Client
	token  string
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) ListReleases(slug string) ([]*Release, error) {
	url := fmt.Sprintf("%s/projects/%s/releases", baseURL, url.QueryEscape(slug))
	logrus.Tracef("GET %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", fmt.Sprintf("%s/%s", common.NAME, common.AppVersion))

	if c.token != "" {
		//req.Header.Set("PRIVATE-TOKEN", c.token)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var releases []*Release
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return nil, err
	}

	return releases, nil
}

func (c *Client) GetLatestRelease(slug string) (*Release, error) {
	url := fmt.Sprintf("%s/projects/%s/releases?per_page=1", baseURL, url.QueryEscape(slug))
	logrus.Tracef("GET %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", fmt.Sprintf("%s/%s", common.NAME, common.AppVersion))

	if c.token != "" {
		//req.Header.Set("PRIVATE-TOKEN", c.token)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var releases []*Release
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return nil, err
	}

	return releases[0], nil

}

func (c *Client) GetRelease(slug, version string) (*Release, error) {
	url := fmt.Sprintf("%s/projects/%s/releases/%s", baseURL, url.QueryEscape(slug), url.QueryEscape(version))
	logrus.Tracef("GET %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", fmt.Sprintf("%s/%s", common.NAME, common.AppVersion))

	if c.token != "" {
		//req.Header.Set("PRIVATE-TOKEN", c.token)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var release *Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return release, nil
}