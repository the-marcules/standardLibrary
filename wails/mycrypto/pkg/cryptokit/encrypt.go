package cryptokit

import (
	"fmt"
	"io"
	"log/slog"
)

func (c *Client) Encrypt(payload string, pubKeyId string) string {

	keys, err := c.GetKeys()
	if err != nil {
		fmt.Printf("error getting keys: %v", err)
	} else {
		fmt.Printf("Keys: %v\n", keys)
	}

	keyMeta, err := c.GetKeyMetaJWK(pubKeyId)
	if err != nil {
		slog.Error("Error getting key meta: ", "error", err)
		return ResponseBuilder(nil, err)
	}

	data := []byte(`{
		"authorization": {
			"account": "` + c.config.AccountName + `",
			"cryptoScenario": "` + c.config.Scenarios.Encrypt.ScenarioName + `"
		},
		"input": {
			"payload": {
				"data": "` + payload + `",
				"encoding": "PLAIN"
			}
		},
		"encrypt": {
			"payloadType": {
				"name": "text-plain"
			},
			"outputPolicy": {
				"name": "MOSE_plain"
			},
			"key": {
				"ephemeral": {}
			},
			"receiverPublicKey": {
      "supplied": {
        "jwk": {
            "kty": "EC",
            "crv": "P-256",
            "x": "` + keyMeta.X + `",
            "y": "` + keyMeta.Y + `"
        }
      }
    },
			"alg": "ECDH-ES",
			"enc": "A256GCM",
			"apu": "Q3J5cHRvS2l0",
			"apv": "UmVjZWl2ZXI"
		}
	}`)

	response, err := c.requestHandler("POST", "/v2/encrypt", data)
	if err != nil {
		slog.Error("Error encrypting the payload: ", "error", err)
		return ResponseBuilder(nil, err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ResponseBuilder(nil, err)
	}

	return ResponseBuilder(body, nil)
}
