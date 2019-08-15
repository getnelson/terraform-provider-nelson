package nelson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
)

type CreateSessionRequest struct {
	AccessToken string `json:"access_token"`
}

func (n *Nelson) Login(githubToken string) error {

	// If not logged in, log in and set the various session tokens and keys in the reciever.
	currentEpoch := time.Now()
	// Should we read the config or use whats already in n.NelsonConfig???
	if n.NelsonConfig == DefaultConfig {
		if err := n.readConfig(); err != nil {
			return multierror.Prefix(err, "reading config: ")
		}
	}
	expireTime := time.Unix(0, n.NelsonConfig.Session.ExpiresAt)
	if currentEpoch.After(expireTime) {
		if err := n.login(githubToken); err != nil {
			return multierror.Prefix(err, "login: ")
		}
	}
	return nil
}

func (n *Nelson) login(githubToken string) error {
	request, err := json.Marshal(CreateSessionRequest{AccessToken: githubToken})
	if err != nil {
		return err
	}

	endpoint := []string{n.NelsonConfig.Endpoint, "auth", "github"}
	req, _ := http.NewRequest(http.MethodPost, strings.Join(endpoint, "/"), bytes.NewBuffer(request))
	req.Header.Set("User-Agent", LibraryVersion())
	req.Header.Set("Content-Type", "application/json")

	response, err := n.Client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf(fmt.Sprintf("recieved %d response from server", response.StatusCode))
	}

	if err := json.NewDecoder(response.Body).Decode(n.NelsonConfig.Session); err != nil {
		return multierror.Prefix(err, "login")
	}

	if len(n.NelsonConfig.Session.Token) < 20 {
		return fmt.Errorf("couldn't create session token")
	}

	// Set up Cookiejar for all future requests
	ncj := &nelsonCookieJar{
		cookies: make(map[string][]*http.Cookie),
	}

	nelsonSessionCookie := &http.Cookie{
		Name:  "nelson.session",
		Value: n.NelsonConfig.Session.Token,
	}

	nelsonUrl, err := url.Parse(n.NelsonConfig.Endpoint)
	if err != nil {
		return multierror.Prefix(err, "could not parse endpoint")
	}

	ncj.SetCookies(nelsonUrl, []*http.Cookie{nelsonSessionCookie})
	n.Client = &http.Client{
		Jar: ncj,
	}

	return nil
}
