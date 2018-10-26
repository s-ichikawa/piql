package client

import (
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Url        *url.URL
	HttpClient http.Client

	Token *string
}

func NewClient(urlStr, token string, insecureSkipVerify bool) (*Client, error) {
	url, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}

	if len(token) == 0 {
		return nil, errors.New("missing user token.")
	}

	tlsConfig := tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}
	transport := http.DefaultTransport.(*http.Transport)
	transport.TLSClientConfig = &tlsConfig
	httpClient := &http.Client{
		Transport: transport,
	}

	client := Client{
		Url:        url,
		Token:      &token,
		HttpClient: *httpClient,
	}

	return &client, nil
}

func (c *Client) NewRequest(method string, body io.Reader) (*http.Request, error) {
	fmt.Println(c.Url.Host + c.Url.RequestURI())
	req, err := http.NewRequest(
		method,
		c.Url.Host+c.Url.RequestURI(),
		body,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Token != nil && len(*c.Token) > 0 {
		req.Header.Set("X-USER-TOKEN", *c.Token)
	}

	return req, nil
}
