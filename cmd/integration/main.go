package main

import (
	"crypto/tls"
	"net/http"
)

func main() {

	// ignore certificate validity
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec

}
