package cryptokit

import (
	"log/slog"
)

func (c *Client) Sign(payload string, codeSign bool) string {

	cryptoScenario := c.config.Scenarios.Sign.ScenarioName
	scenarioKeyName := c.config.Scenarios.Sign.ScenarioKeyName
	payloadType := "text-plain"
	outputPolicy := "JWS"

	if codeSign {
		cryptoScenario = c.config.Scenarios.CodeSign.ScenarioName
		scenarioKeyName = c.config.Scenarios.CodeSign.ScenarioKeyName
		payloadType = "application-ps1"
		outputPolicy = "SIGNATURE_PS1"
	}

	data := []byte(`{
	  "authorization": {
	    "account": "` + c.config.AccountName + `",
	    "cryptoScenario": "` + cryptoScenario + `"
	  },
	  "input": {
	    "payload": {
	      "data": "` + payload + `",
	      "encoding": "PLAIN"
	    }
	  },
	  "sign": {
	    "payloadType": {
	      "name": "` + payloadType + `"
	    },
	    "outputPolicy": {
	      "name": "` + outputPolicy + `" 
	    },
	    "key": {
	      "static": {
	        "kid": "` + scenarioKeyName + `"
	      }
	    }
	  }
	}
	`)
	slog.Info("payload", "data", string(data))

	response, err := c.requestHandler("POST", "/v2/sign", data)

	if err != nil {
		slog.Error("Error signing the payload: ", "error", err)
		return ResponseBuilder(nil, err)
	}

	if codeSign {
		return CodeSignResponseBuilder(response, nil)
	}

	return ResponseBuilder(response, nil)
}
