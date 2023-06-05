package taikungoclient

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	taikuncore "github.com/chnyda/taikungoclient/client"
	taikunshowback "github.com/chnyda/taikungoclient/showbackclient"
)

const TaikunEmailEnvVar = "TAIKUN_EMAIL"
const TaikunPasswordEnvVar = "TAIKUN_PASSWORD"
const TaikunAuthMode = "TAIKUN_AUTH_MODE"
const TaikunAccessKey = "TAIKUN_ACCESS_KEY"
const TaikunSecretKey = "TAIKUN_SECRET_KEY"
const TaikunApiHostEnvVar = "TAIKUN_API_HOST"

type CustomTransport struct {
	Transport http.RoundTripper
	Client    *Client
}

type Client struct {
	Client          *taikuncore.APIClient
	ShowbackClient  *taikunshowback.APIClient
	CustomTransport *CustomTransport
	email           string
	password        string
	secretKey       string
	accessKey       string
	authMode        string
	token           string
	refreshToken    string
}

type jwtData struct {
	Nameid     string `json:"nameid"`
	Email      string `json:"email"`
	UniqueName string `json:"unique_name"`
	Role       string `json:"role"`
	Nbf        int    `json:"nbf"`
	Exp        int    `json:"exp"`
	Iat        int    `json:"iat"`
}

func NewClientFromCredentials(email string, password string, accessKey string, secretKey string, authMode string, apiHost string) *Client {
	cfg := taikuncore.NewConfiguration()
	cfg.Host = apiHost
	cfg.Scheme = "https"

	cfg2 := taikunshowback.NewConfiguration()
	cfg2.Host = apiHost
	cfg2.Scheme = "https"

	client := &Client{
		email:     email,
		password:  password,
		accessKey: accessKey,
		secretKey: secretKey,
		authMode:  authMode,
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	customTransport := &CustomTransport{
		Transport: tr,
		Client:    client,
	}
	cfg.HTTPClient = &http.Client{Transport: customTransport}
	cfg2.HTTPClient = &http.Client{Transport: customTransport}
	client.Client = taikuncore.NewAPIClient(cfg)
	client.ShowbackClient = taikunshowback.NewAPIClient(cfg2)

	return client
}

func NewClient() *Client {
	apiHost := os.Getenv(TaikunApiHostEnvVar)
	if apiHost == "" {
		apiHost = "api.taikun.coud"
	}

	taikunAuthMode, customMode := os.LookupEnv(TaikunAuthMode)
	if !customMode {
		taikunAuthMode = "default"
	}

	if taikunAuthMode == "default" {
		email := os.Getenv(TaikunEmailEnvVar)
		password := os.Getenv(TaikunPasswordEnvVar)
		if email == "" || password == "" {
			panic(fmt.Errorf("Please set your Taikun credentials."))
		}
		return NewClientFromCredentials(email, password, "", "", taikunAuthMode, apiHost)
	}

	accessKey := os.Getenv(TaikunAccessKey)
	secretKey := os.Getenv(TaikunSecretKey)
	if accessKey == "" || secretKey == "" {
		panic(fmt.Errorf("Please set your Taikun credentials."))
	}

	return NewClientFromCredentials("", "", accessKey, secretKey, taikunAuthMode, apiHost)
}

func (c *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Path != "/api/v1/auth/login" && req.URL.Path != "/api/v1/auth/refresh" {
		if c.Client.token == "" {
			var loginCmd *taikuncore.LoginCommand
			// keycloak
			// default
			// token
			// autoscaling
			if c.Client.authMode != "" || c.Client.authMode == "default" {
				loginCmd = &taikuncore.LoginCommand{
					SecretKey: &c.Client.secretKey,
					AccessKey: &c.Client.accessKey,
					Mode:      &c.Client.authMode,
				}
			} else {
				loginCmd = &taikuncore.LoginCommand{
					Email:    &c.Client.email,
					Password: &c.Client.password,
				}
			}

			result, _, err := c.Client.Client.AuthApi.AuthLogin(req.Context(), "1").Body(*loginCmd).Execute()
			if err != nil {
				return nil, err
			}
			c.Client.token = *result.Token
			c.Client.refreshToken = *result.RefreshToken
		} else if c.Client.token != "" && c.hasTokenExpired() {
			result, _, err := c.Client.Client.AuthApi.AuthRefreshToken(req.Context(), "1").Body(taikuncore.RefreshTokenCommand{
				RefreshToken: &c.Client.refreshToken,
				Token:        &c.Client.token,
			}).Execute()
			if err != nil {
				return nil, err
			}
			c.Client.token = *result.Token
			c.Client.refreshToken = *result.RefreshToken
		}
		req.Header.Set("Authorization", "Bearer "+c.Client.token)
	}
	return c.Transport.RoundTrip(req)

}

func (transport *CustomTransport) hasTokenExpired() bool {
	jwtSplit := strings.Split(transport.Client.token, ".")
	if len(jwtSplit) != 3 {
		return true
	}

	data, err := base64.RawURLEncoding.DecodeString(jwtSplit[1])
	if err != nil {
		return true
	}

	jwtData := jwtData{}

	err = json.Unmarshal(data, &jwtData)
	if err != nil {
		return true
	}

	tm := time.Unix(int64(jwtData.Exp), 0)

	return tm.Before(time.Now())
}
