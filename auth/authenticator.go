package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	ScopeApiScope     = "https://api.ebay.com/oauth/api_scope"
	ScopeBuyItemFeed  = "https://api.ebay.com/oauth/api_scope/buy.item.feed"
	ScopeBuyMarketing = "https://api.ebay.com/oauth/api_scope/buy.marketing"
	SandboxBaseUrl    = "https://api.sandbox.ebay.com"
	ProductionBaseUrl = "https://api.ebay.com"
)

var ErrInvalidCredentials = errors.New("invalid ebay credentials")

type Authenticator interface {
	//AuthToken returns the current token.  The Authenticator may attempt to keep this token fresh, but it will not
	//explicitly check that the token is still valid
	AuthToken() (Token, error)

	//RefreshToken will force the authenticator to request a new token, and then return the new value.  The token will
	//be persisted for the next AuthToken() call.  Use this method when a token returned by AuthToken() is no longer valid
	RefreshToken() (Token, error)
}

type ApplicationAuthenticator struct {
	HttpClient   *http.Client
	BaseUrl      string
	Credentials  Credentials
	scope        string
	currentToken Token
	refreshTimer *time.Timer
}

var _ Authenticator = &ApplicationAuthenticator{}

type AuthenticatorConfig struct {
	Credentials Credentials
	BaseUrl     string
	Scope       string
}

func NewAuthenticator(config *AuthenticatorConfig) (*ApplicationAuthenticator, error) {
	auth := &ApplicationAuthenticator{
		Credentials: config.Credentials,
		scope:       config.Scope,
		BaseUrl:     config.BaseUrl,
	}

	if auth.HttpClient == nil {
		auth.HttpClient = http.DefaultClient
	}

	if auth.Credentials.Username == "" || auth.Credentials.Password == "" {
		return nil, ErrInvalidCredentials
	}

	if auth.BaseUrl == "" {
		auth.BaseUrl = ProductionBaseUrl
	}

	if auth.scope == "" {
		auth.scope = ScopeApiScope
	}

	return auth, nil
}

func MustNewAuthenticator(config *AuthenticatorConfig) *ApplicationAuthenticator {
	auth, err := NewAuthenticator(config)
	if err != nil {
		panic(err)
	}

	return auth
}

func (a *ApplicationAuthenticator) AuthToken() (Token, error) {
	if a.currentToken.AccessToken != "" {
		return a.currentToken, nil
	} else {
		return a.RefreshToken()
	}
}

func (a *ApplicationAuthenticator) RefreshToken() (Token, error) {
	var err error
	a.currentToken, err = a.fetchToken()
	if err != nil {
		return a.currentToken, err
	}

	if a.refreshTimer != nil {
		a.refreshTimer.Stop()
	}
	a.refreshTimer = time.AfterFunc(7140*time.Second, func() {
		_, _ = a.RefreshToken()
	})

	return a.currentToken, nil
}

func (a *ApplicationAuthenticator) fetchToken() (Token, error) {
	token := Token{}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("redirect_uri", a.Credentials.Username)
	data.Set("scope", a.scope)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/identity/v1/oauth2/token", a.BaseUrl), strings.NewReader(data.Encode()))
	req.SetBasicAuth(a.Credentials.Username, a.Credentials.Password)
	req.Header.Add("User-Agent", "UAPL-EbayClient/1.0") //TODO make User-Agent configurable
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := a.HttpClient.Do(req)
	if err != nil {
		return token, err
	}

	if resp.StatusCode == 401 {
		return token, ErrInvalidCredentials
	}
	if resp.StatusCode != 200 {
		return token, errors.New("expected http status 2xx during auth but got " + string(resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, fmt.Errorf("error reading body from http response:  %w", err)
	}

	err = json.Unmarshal(body, &token)
	if err != nil {
		return token, fmt.Errorf("could not unmarshall body:  %w", err)
	}

	return token, nil
}
