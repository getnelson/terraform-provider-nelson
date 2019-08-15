package nelson

import (
	"bytes"
	"net/http"
	"strings"
)

// ReposHook is the small bit that says whether the thing is on.
type ReposHook struct {
	IsActive bool   `json:"is_active"`
	ID       uint32 `json:"id"`
}

// Repos list of repos that come back
type Repos struct {
	Repository string     `json:"repository"`
	Slug       string     `json:"slug"`
	ID         uint32     `json:"id"`
	Hook       *ReposHook `json:"hook"`
	Owner      string     `json:"owner"`
	Access     string     `json:"access"`
}

// GetRepos uses the http.Client in Nelson api struct to talk to
// the Nelson server
// https://nelson.local/v1/repos?owner=githubuser360
func (n *Nelson) GetRepos(owner string) {
	var b []byte
	response := bytes.NewBuffer(b)

	endpoint := []string{n.NelsonConfig.Endpoint, n.apiVersion, "repos"}
	req, _ := http.NewRequest(http.MethodGet, strings.Join(endpoint, "/"), response)
	req.URL.Query().Add("owner", owner)

	n.Client.Do(req)
}

// CreateWebhook puts a webhook on a repository to enable nelson on a given repository.
// https://nelson.local/v1/repos/getnelson/terraform-provider-nelson/hook
func (n *Nelson) CreateWebhook(owner string, repo string) {
	var b []byte
	response := bytes.NewBuffer(b)

	endpoint := []string{n.NelsonConfig.Endpoint, n.apiVersion, "repos", owner, repo, "hook"}
	req, _ := http.NewRequest(http.MethodPost, strings.Join(endpoint, "/"), response)

	n.Client.Do(req)
}

// DeleteWebhook tells Nelson to delete a webhook on the repository.
func (n *Nelson) DeleteWebhook(owner string, repo string) {
	var b []byte
	response := bytes.NewBuffer(b)

	endpoint := []string{n.NelsonConfig.Endpoint, n.apiVersion, "repos", owner, repo, "hook"}
	req, _ := http.NewRequest(http.MethodDelete, strings.Join(endpoint, "/"), response)

	n.Client.Do(req)
}
