package cryptokit

import (
	"log/slog"
)

func (c *Client) Verify(format, signature string) string {

	data := []byte(`
	{	  
		"authorization": {
			"account": "` + c.config.AccountName + `",
			"cryptoScenario": "` + c.config.Scenarios.Verify.ScenarioName + `"
	  	},
		"input": {
			"messages": {
				"message-1": {
				"format": "` + format + `",
				"content" : "` + signature + `"
				}
			}
		}
	}`)

	response, err := c.requestHandler("POST", "/v2/verify", data)
	if err != nil {
		slog.Error("Error signing the payload: ", "error", err)
		return ResponseBuilder(nil, err)
	}

	return ResponseBuilder(response, nil)
}
