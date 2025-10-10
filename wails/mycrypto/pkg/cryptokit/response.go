package cryptokit

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type CryptoKitError struct {
	Code        string `json:"code,omitempty" `
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	Parameter   string `json:"parameter,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
}

type SignResponseContent struct {
	Format  string `json:"format,omitempty"`
	Content string `json:"content,omitempty"`
}

type SignResponse struct {
	Signature SignResponseContent `json:"signature,omitempty"`
}

type ResponseError struct {
	IsError bool   `json:"isError"`
	Message string `json:"message,omitempty"`
}

type VerificationResponseMessage struct {
	Result string `json:"result,omitempty"`
}

type VerificationResponseContent struct {
	Result   string                                 `json:"result,omitempty"`
	Messages map[string]VerificationResponseMessage `json:"messages,omitempty"`
}

type VerifyResponse struct {
	Verification VerificationResponseContent `json:"verification,omitempty"`
}

type EncryptResponse struct {
	Content string `json:"content,omitempty"`
	Format  string `json:"format,omitempty"`
}

type DecryptResponse struct {
	Encoding  string `json:"encoding,omitempty"`
	Plaintext string `json:"plaintext,omitempty"`
}

type Response struct {
	CryptoKitError
	SignResponse
	VerifyResponse
	Encryption EncryptResponse `json:"encryption,omitempty"`
	Decryption DecryptResponse `json:"decryption,omitempty"`
	Error      ResponseError   `json:"error"`
	TraceID    string          `json:"traceId,omitempty"`
}

func ResponseBuilder(apiResponse *http.Response, responseError error) string {

	var response []byte

	if responseError != nil {
		response, _ = json.Marshal(Response{
			Error: *NewResponseError(true, responseError.Error()),
		})
		return string(response)
	}

	traceId := apiResponse.Header.Get("x-b3-traceid")
	body, err := io.ReadAll(apiResponse.Body)

	if err != nil {
		slog.Error("Could not read apiResponse.body: ", "error", err)
		response, _ = json.Marshal(Response{
			Error: *NewResponseError(true, err.Error()),
		})
		return string(response)
	}

	var responseObject Response
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		slog.Error("Error unmarshalling the response: ", "error", err, "body", string(body))
		response, _ = json.Marshal(Response{
			Error: *NewResponseError(true, err.Error()),
		})
		return string(response)
	}

	if responseObject.Code != "" {
		responseObject.Error.IsError = true
	}
	responseObject.TraceID = traceId

	response, _ = json.Marshal(responseObject)
	return string(response)
}

func NewResponseError(isError bool, message string) *ResponseError {
	return &ResponseError{
		IsError: isError,
		Message: message,
	}
}
