package cryptokit

import (
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
