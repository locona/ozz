package hydra

import (
	"context"
	"log"
	"net/http"

	"github.com/ory/hydra/sdk/go/hydra"
	"github.com/ory/hydra/sdk/go/hydra/swagger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	clientID     = "subjects:hydra:clients:oathkeeper-client"
	clientSecret = "dummy-oathkeeper-secret"
)

const (
	userClientID     = "client-01"
	userClientSecret = "secret"
)

type transporter struct {
	*http.Transport
	FakeTLSTermination bool
}

type Client struct {
	sdk *hydra.CodeGenSDK
}

func New() (*Client, error) {
	sdk, err := hydra.NewSDK(&hydra.Configuration{
		AdminURL:     "http://localhost:24445",
		PublicURL:    "http://localhost:24444",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"openid", "offline", "healthcheck"},
	})

	if err != nil {
		return nil, err
	}

	return &Client{sdk: sdk}, nil
}

func (cli *Client) CreateClient() error {
	_, _, err := cli.sdk.CreateOAuth2Client(swagger.OAuth2Client{
		Owner:         "reckoner",
		ClientId:      userClientID,
		ClientName:    "Sample Client 01",
		ClientSecret:  userClientSecret,
		GrantTypes:    []string{"client_credentials"},
		Scope:         "openid,offline,healthcheck",
		ResponseTypes: []string{"token"},
	})

	if err != nil {
		log.Println("ERROR: ", err)
	}

	return err
}

func (cli *Client) Token() (*oauth2.Token, error) {
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
		Transport: &transporter{
			FakeTLSTermination: false,
			Transport:          &http.Transport{},
		},
	})
	oauthConfig := clientcredentials.Config{
		ClientID:     userClientID,
		ClientSecret: userClientSecret,
		TokenURL:     "http://localhost:24444/oauth2/token",
	}
	return oauthConfig.Token(ctx)
}
