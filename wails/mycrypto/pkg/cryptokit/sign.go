package cryptokit

import (
	"io"
	"log/slog"
)

func (c *Client) Sign(payload string) string {

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
	      "name": "MOSE_plain"
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ResponseBuilder(nil, err)
	}

	return ResponseBuilder(body, nil)
}
