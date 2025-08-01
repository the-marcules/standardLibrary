package cryptokit

import (
	"bytes"
	"encoding/base64"
	"mycrypto/pkg/config"
	"net/http"
)

type Client struct {
	config *config.Config
}

type AdditionalHeader struct {
	Key   string
	Value string
}

type JWK struct {
	Kty string `json:"kty,omitempty"`
	Crv string `json:"crv,omitempty"`
	X   string `json:"x,omitempty"`
	Y   string `json:"y,omitempty"`
	Kid string `json:"kid,omitempty"`
	Alg string `json:"alg,omitempty"`
	Use string `json:"use,omitempty"`
}

type Key struct {
	Kid string `json:"kid,omitempty"`
}

func NewCryptoKitApi(config *config.Config) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) requestHandler(method, path string, data []byte, additionalHeaders ...AdditionalHeader) (response *http.Response, err error) {
	reader := bytes.NewReader(data)

	httpClient := &http.Client{}
	req, err := http.NewRequest(method, c.config.BaseUrl+path, reader)
	if err != nil {
		return
	}

	auth := base64.StdEncoding.EncodeToString([]byte(c.config.AccessKey + ":" + c.config.SecretKey))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Content-Type", "application/json")

	for _, header := range additionalHeaders {
		req.Header.Add(header.Key, header.Value)
	}

	response, err = httpClient.Do(req)
	if err != nil {
		return
	}

	return
}
