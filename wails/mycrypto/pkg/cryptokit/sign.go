package cryptokit

import (
	"encoding/base64"
	"log/slog"
	"os"
)

func (c *Client) Sign(payload string, codeSign bool) string {

	if codeSign {
		payload = getBase64PayloadForCodeSign(payload)
	}

	data := []byte(`{
	  "authorization": {
	    "account": "` + c.config.AccountName + `",
	    "cryptoScenario": "` + c.config.Scenarios.Sign.ScenarioName + `"
	  },
	  "input": {
	    "payload": {
	      "data": "` + payload + `",
	      "encoding": "PLAIN"
	    }
	  },
	  "sign": {
	    "payloadType": {
	      "name": "text-plain"
	    },
	    "outputPolicy": {
	      "name": "JWS" 
	    },
	    "key": {
	      "static": {
	        "kid": "` + c.config.Scenarios.Sign.ScenarioKeyName + `"
	      }
	    }
	  }
	}
	`)

	response, err := c.requestHandler("POST", "/v2/sign", data)

	if err != nil {
		slog.Error("Error signing the payload: ", "error", err)
		return ResponseBuilder(nil, err)
	}

	return ResponseBuilder(response, nil)
}

func getBase64PayloadForCodeSign(filePath string) string {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("Error reading the file: ", "error", err)
		return ResponseBuilder(nil, err)
	}
	return base64.StdEncoding.EncodeToString(fileContent)
}
