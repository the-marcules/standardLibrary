package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	wd "github.com/emersion/go-webdav"
	"golang.org/x/crypto/pkcs12"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"time"
)

const (
	certPath = "testdata/keystore.p12"
	certPin  = "1234"
)

type WebDavResponse struct {
	Err error  `json:"error"`
	Msg string `json:"msg"`
}

type WebDavCommunicator struct {
	client *wd.Client
	url    string
	ctx    context.Context
	cancel context.CancelFunc
}

func NewWebDavCommunicator(url string, timeOut ...time.Duration) *WebDavCommunicator {

	ctx, cancel := context.WithCancel(context.Background())
	if timeOut != nil {
		ctx, cancel = context.WithTimeout(context.Background(), timeOut[0])
	}

	transport, err := createHttpTransportWithClientCert()
	if err != nil {
		panic(err)
	}

	httpClient := &http.Client{
		Transport:     transport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	webDavClient, err := wd.NewClient(httpClient, url)
	if err != nil {
		panic(err)
	}

	return &WebDavCommunicator{
		client: webDavClient,
		url:    url,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (w *WebDavCommunicator) GetStat(fileName string) *WebDavResponse {

	stat, err := w.client.Stat(w.ctx, fileName)
	if err != nil {
		stat, err = w.client.Stat(w.ctx, fileName)
	}

	msgString, _ := json.Marshal(stat)
	return &WebDavResponse{Err: err, Msg: string(msgString)}
}

func (w *WebDavCommunicator) PutFile(fileName string, remoteFilename string) *WebDavResponse {
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		return &WebDavResponse{Err: err}
	}
	writer, createErr := w.client.Create(w.ctx, remoteFilename)
	if createErr != nil {
		return &WebDavResponse{Err: createErr}
	}

	writtenBytes, writeErr := writer.Write(fileContents)
	_ = writer.Close()

	return &WebDavResponse{
		Err: writeErr,
		Msg: fmt.Sprintf("written %d bytes", writtenBytes),
	}
}

func createHttpTransportWithClientCert() (*http.Transport, error) {
	data, err := os.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	pemBlocks, err := pkcs12.ToPEM(data, certPin)
	if err != nil {
		return nil, err
	}

	var pemData []byte
	for _, b := range pemBlocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		return nil, err
	}

	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	caCertPool.AppendCertsFromPEM(pemData)

	tlsConfig := &tls.Config{
		Certificates:  []tls.Certificate{cert},
		RootCAs:       caCertPool,
		Renegotiation: tls.RenegotiateOnceAsClient,
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}

	return transport, nil
}
