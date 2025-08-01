package cryptokit

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetKeyMetaJWK(keyId string) (*JWK, error) {

	url := fmt.Sprintf("/v2/accounts/%s/keys/%s/public", c.config.AccountName, keyId)

	header := []AdditionalHeader{
		{"accept", "application/jwk+json"},
	}

	response, err := c.requestHandler("GET", url, nil, header...)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Response: %v\n", response)

	parsedJWK := &JWK{}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, parsedJWK)
	if err != nil {
		return nil, err
	}

	return parsedJWK, nil
}

func (c *Client) GetKeys() ([]Key, error) {

	url := fmt.Sprintf("/v2/accounts/%s/keys", c.config.AccountName)

	response, err := c.requestHandler("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching keys: %v", response.Status)
	}

	parsedResponse := []Key{}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &parsedResponse)
	if err != nil {
		return nil, err
	}

	return parsedResponse, nil
}

func (c *Client) GetKeysString() string {

	keys, err := c.GetKeys()
	if err != nil {
		fmt.Printf("error getting keys: %v", err)
		return ResponseBuilder(nil, err)
	} else {
		fmt.Printf("Keys: %v\n", keys)
	}

	reponse, err := json.Marshal(keys)
	if err != nil {
		return ""
	}

	return string(reponse)
}
