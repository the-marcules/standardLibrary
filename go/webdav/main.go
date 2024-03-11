package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/studio-b12/gowebdav"
	"golang.org/x/crypto/pkcs12"
	"net/http"
	"os"
)

func main() {
	a := NewApp()
	noProxyTransport, err := a.initCom()
	if err != nil {
		println(err.Error())
	}
	a.doTest(noProxyTransport, "GUKMOHD")
}

type App struct {
	stage    string // todo: what about stage?
	dtsHost  string
	certPath string
	certPin  string
	proxyUrl string
	//disableUploadDT         bool
	//testConnectionDT        bool
	//enableAttemptIdInUpload bool
}

func NewApp() App {
	dtsHost := os.Getenv("DTS_HOST")
	certPath := os.Getenv("CERT_PATH")
	certPin := os.Getenv("CERT_PIN")
	proxyUrl := os.Getenv("MY_HTTPS_PROXY")

	return App{
		dtsHost:  dtsHost,
		certPath: certPath,
		certPin:  certPin,
		proxyUrl: proxyUrl,
	}
}

func (a App) doTest(transportWithClientCert *http.Transport, username string) {
	webdavClient := gowebdav.NewClient("https://"+a.dtsHost, "anonymous", "")
	webdavClient.SetTransport(transportWithClientCert)
	err := webdavClient.Connect()
	if err != nil {
		println(err)
	}
	dir, err := webdavClient.ReadDir("/")
	//dir, err := webdavClient.ReadDir("/"+username) ??
	if err != nil {
		println(err)
	}

	for _, info := range dir {
		fmt.Printf("+ %s", info.Name())

		if info.Name() == username && info.IsDir() {
			println("dir is present")
		}
	}
}

func (a App) initCom() (*http.Transport, error) {

	noProxyTransport, err := a.createHttpTransportWithClientCert()

	return noProxyTransport, err
}
func (a App) createHttpTransportWithClientCert() (*http.Transport, error) {
	data, err := os.ReadFile(a.certPath)
	if err != nil {
		return nil, err
	}

	pemBlocks, err := pkcs12.ToPEM(data, a.certPin)
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
