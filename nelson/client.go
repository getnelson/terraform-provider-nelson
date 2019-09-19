package nelson

import (
	"net/http"
	"net/url"
	"sync"
)

var clientErrorPrefix string = "client: "
var DefaultConfig = &NelsonConfig{
	Endpoint: "",
	Session: &NelsonSession{
		Token:     "",
		ExpiresAt: 0,
	},
}

// Nelson is a data struct that contains an http client. This
// client will be used to make secure calls to the Nelson server.
type Nelson struct {
	Client       *http.Client
	apiVersion   string
	configPath   string
	NelsonConfig *NelsonConfig
}

type nelsonCookieJar struct {
	lock    sync.Mutex
	cookies map[string][]*http.Cookie
}

func (ncj *nelsonCookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	ncj.lock.Lock()
	ncj.cookies[u.Host] = cookies
	ncj.lock.Unlock()
}

func (ncj *nelsonCookieJar) Cookies(u *url.URL) []*http.Cookie {
	return ncj.cookies[u.Host]
}

type NelsonConfig struct {
	Endpoint string         `yaml:"endpoint"`
	Session  *NelsonSession `yaml:"session"`
}

type NelsonSession struct {
	Token     string `yaml:"token" json:"session_token"`
	ExpiresAt int64  `yaml:"expires_at" json:"expires_at"`
}

// CreateNelson creates a client from a given config. In the config is the
// path to the .nelson/config.yml. We need that file to grab the token.
func CreateNelson(address string, version string, configPath string) (*Nelson, error) {
	// Error handle the TLS thing
	_, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	nelson := &Nelson{
		Client:     http.DefaultClient,
		apiVersion: version,
		configPath: configPath,
		NelsonConfig: &NelsonConfig{
			Endpoint: address,
			Session: &NelsonSession{
				Token:     "",
				ExpiresAt: 0,
			},
		},
	}
	return nelson, nil
}
