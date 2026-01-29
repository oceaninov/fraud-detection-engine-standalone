package extSSOLoginMSF

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"gitlab.com/fds22/detection-sys/pkg/environments"
)

type (
	GetAccessTokenRequest struct {
		Code        string `json:"code"`
		RedirectUri string `json:"redirect_uri"`
	}
	GetAccessTokenResponse struct {
		ProcessTime string             `json:"process_time"`
		Status      string             `json:"status"`
		Code        string             `json:"code"`
		Message     string             `json:"message"`
		Data        GetAccessTokenData `json:"data"`
	}
	GetAccessTokenData struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		Scope        string `json:"scope"`
		IDToken      string `json:"id_token"`
		RefreshToken string `json:"refresh_token"`
		Error        string `json:"error,omitempty"`
		ErrorDesc    string `json:"error_description,omitempty"`
	}
)

type IProxyAdapterClient interface {
	GetAccessToken(code string) (resp GetAccessTokenResponse, err error)
}

type proxyAdapterClient struct {
	env    *environments.Envs
	client *resty.Client
}

func NewProxyAdapterClient(env *environments.Envs) IProxyAdapterClient {
	client := resty.New().
		SetTimeout(30 * time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(20 * time.Second)

	return &proxyAdapterClient{env: env, client: client}
}

func (a *proxyAdapterClient) GetAccessToken(code string) (resp GetAccessTokenResponse, err error) {
	level := a.env.LogLevel
	var debug bool
	if level == "DEBUG" {
		debug = true
	} else {
		debug = false
	}
	response, err := a.client.R().SetDebug(debug).
		SetHeader("Content-Type", "application/json").
		SetResult(&resp).
		SetBody(GetAccessTokenRequest{
			Code:        code,
			RedirectUri: a.env.ApplicationAzureSSORedirectUri,
		}).
		Post(fmt.Sprintf("%s/azure-ad/auth", a.env.ProxyAdapterHost))

	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(response.Body(), &resp)
	if err != nil {
		err = errors.New(fmt.Sprintf("[%d] Failed to parse response body", response.StatusCode()))
		return
	}
	if response.StatusCode() != http.StatusOK {
		if resp.Data.Error != "" {
			err = errors.New("[" + resp.Data.Error + "] " + resp.Data.ErrorDesc)
		} else {
			err = errors.New(fmt.Sprintf("[%d] Failed to get access token", response.StatusCode()))
		}
		return
	}

	return resp, nil
}
