package taikungoclient

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
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
	mu        sync.Mutex
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

type TaikunError struct {
	HTTPStatusCode int
	Message        string
}

func (e *TaikunError) Error() string {
	return fmt.Sprintf("Taikun Error: %s (HTTP %d)", e.Message, e.HTTPStatusCode)
}
func CreateError(resp *http.Response, err error) error {
	if resp == nil {
		return err
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return err
	}
	return &TaikunError{
		HTTPStatusCode: resp.StatusCode,
		Message:        string(body),
	}
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
	if !customMode || taikunAuthMode == "" || taikunAuthMode == "default" {
		email := os.Getenv(TaikunEmailEnvVar)
		password := os.Getenv(TaikunPasswordEnvVar)
		if email == "" || password == "" {
			panic(fmt.Errorf("Please set your Taikun credentials."))
		}
		return NewClientFromCredentials(email, password, "", "", "", apiHost)
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
			c.mu.Lock()
			defer c.mu.Unlock()

			if c.Client.token == "" {
				var loginCmd *taikuncore.LoginCommand
				// keycloak
				// default
				// token
				// autoscaling
				if c.Client.authMode != "" || c.Client.authMode == "default" {
					loginCmd = &taikuncore.LoginCommand{
						SecretKey: *taikuncore.NewNullableString(&c.Client.secretKey),
						AccessKey: *taikuncore.NewNullableString(&c.Client.accessKey),
						Mode:      *taikuncore.NewNullableString(&c.Client.authMode),
					}
				} else {
					loginCmd = &taikuncore.LoginCommand{
						Email:    *taikuncore.NewNullableString(&c.Client.email),
						Password: *taikuncore.NewNullableString(&c.Client.password),
					}
				}
				result, body, err := c.Client.Client.AuthManagementApi.AuthLogin(req.Context()).LoginCommand(*loginCmd).Execute()
				if err != nil {
					return nil, CreateError(body, err)
				}
				c.Client.token = *result.Token.Get()
				c.Client.refreshToken = *result.RefreshToken.Get()
			}
		} else if c.Client.token != "" && c.hasTokenExpired() {
			result, body, err := c.Client.Client.AuthManagementApi.AuthRefresh(req.Context()).RefreshTokenCommand(taikuncore.RefreshTokenCommand{
				RefreshToken: *taikuncore.NewNullableString(&c.Client.refreshToken),
				Token:        *taikuncore.NewNullableString(&c.Client.token),
			}).Execute()
			if err != nil {
				return nil, CreateError(body, err)
			}
			c.Client.token = *result.Token.Get()
			c.Client.refreshToken = *result.RefreshToken.Get()
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
