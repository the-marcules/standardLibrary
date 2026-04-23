package cryptokit

import (
	"encoding/base64"
	"log/slog"
	"os"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (c *Client) Sign(payload string, codeSign bool) string {

	cryptoScenario := c.config.Scenarios.Sign.ScenarioName
	scenarioKeyName := c.config.Scenarios.Sign.ScenarioKeyName
	payloadType := "text-plain"
	outputPolicy := "JWS"
	encoding := "PLAIN"

	if codeSign {
		cryptoScenario = c.config.Scenarios.CodeSign.ScenarioName
		scenarioKeyName = c.config.Scenarios.CodeSign.ScenarioKeyName
		payloadType = "application-ps1"
		outputPolicy = "PS1"
		encoding = "BASE64"
	}

	data := []byte(`{
	  "authorization": {
	    "account": "` + c.config.AccountName + `",
	    "cryptoScenario": "` + cryptoScenario + `"
	  },
	  "input": {
	    "payload": {
	      "data": "` + payload + `",
	      "encoding": "` + encoding + `"
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

	startTime := time.Now()
	response, err := c.requestHandler("POST", "/v2/sign", data)
	elapsedTime := time.Since(startTime)
	slog.Info("Time taken to sign the payload: ", "elapsedTime", elapsedTime)

	if err != nil {
		slog.Error("Error signing the payload: ", "error", err)
		return ResponseBuilder(nil, err)
	}

	if codeSign {
		return CodeSignResponseBuilder(c.Ctx, response, nil)
	}

	return ResponseBuilder(response, nil)
}

func (c *Client) EngageCodeSignProcess() string {

	slog.Info("Engaging code signing process...")

	filename, err := runtime.OpenFileDialog(c.Ctx, runtime.OpenDialogOptions{
		Title: "Select a file to sign",
	})
	if err != nil {
		slog.Error("Error opening file dialog: ", "error", err)
		return ""
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		slog.Error("Error reading the file: ", "error", err)
		return ""
	}

	base64Payload := base64.StdEncoding.EncodeToString(file)

	return c.Sign(base64Payload, true)
}

func (c *Client) MakeFileFromResponse(response string) string {

	return ""
}
