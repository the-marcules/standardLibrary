package cryptokit

import (
	"io"
	"log/slog"
)

func (c *Client) Decrypt(payload, keyId string) string {

	data := c.BuildDecryptBody(payload, keyId)

	response, err := c.requestHandler("POST", "/v2/decrypt", data)
	if err != nil {
		slog.Error("Error signing the payload: ", "error", err)
		return ResponseBuilder(nil, err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ResponseBuilder(nil, err)
	}

	return ResponseBuilder(body, nil)
}

func (c *Client) BuildDecryptBody(payload, keyId string) []byte {

	return []byte(`{
		"authorization": {
			"account": "` + c.config.AccountName + `",
			"cryptoScenario": "` + c.config.Scenarios.Decrypt.ScenarioName + `"
		},
		"input": {
			"message": {
				"format": "application/mose",
				"content": "` + payload + `"
			}
		},
		"decrypt": {
			"key": {
				"static": {
					"kid": "` + keyId + `"
				}
			}
		}
	}`)

}
