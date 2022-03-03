package fhclient

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"net/http"

	client "github.com/firehydrant/api-client-go/client"
	transport "github.com/go-openapi/runtime/client"
)

type Config struct {
	ApiHost   string
	ApiKey    string
	Debug     bool
	Transport http.RoundTripper
}

type ApiClient struct {
	transport *transport.Runtime
	Client    *client.FireHydrantAPI
	Auth      runtime.ClientAuthInfoWriter
}

func NewApiClient(c Config) ApiClient {
	fhApiClient := ApiClient{}

	fhApiClient.transport = transport.New(c.ApiHost, "", []string{"https"})
	fhApiClient.transport.Debug = c.Debug
	if c.Transport != nil {
		fhApiClient.transport.Transport = c.Transport
	}

	fhApiClient.Client = client.New(fhApiClient.transport, strfmt.Default)
	fhApiClient.Auth = transport.BearerToken(c.ApiKey)

	return fhApiClient
}
