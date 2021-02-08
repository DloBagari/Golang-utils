package web_client

import (
	"crypto/tls"
	"net/http"
)

type NopTransport struct {
}

// when nop os true this method will be called for each request.
//we can modify this as we like,
// this method can be used for testing, we can mock it as we wish
func (n *NopTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}
func Setup(isSecure, nop bool) *http.Client {
	c := http.DefaultClient
	//if not SSL
	if !isSecure {
		c.Transport = &http.Transport{
			Proxy:          nil,
			DialContext:    nil,
			Dial:           nil,
			DialTLSContext: nil,
			DialTLS:        nil,
			TLSClientConfig: &tls.Config{
				Rand:                  nil,
				Time:                  nil,
				Certificates:          nil,
				NameToCertificate:     nil,
				GetCertificate:        nil,
				GetClientCertificate:  nil,
				GetConfigForClient:    nil,
				VerifyPeerCertificate: nil,
				RootCAs:               nil,
				NextProtos:            nil,
				ServerName:            "",
				ClientAuth:            0,
				ClientCAs:             nil,
				// set to false to skip ssl
				InsecureSkipVerify:          false,
				CipherSuites:                nil,
				PreferServerCipherSuites:    false,
				SessionTicketsDisabled:      false,
				SessionTicketKey:            [32]byte{},
				ClientSessionCache:          nil,
				MinVersion:                  0,
				MaxVersion:                  0,
				CurvePreferences:            nil,
				DynamicRecordSizingDisabled: false,
				Renegotiation:               0,
				KeyLogWriter:                nil,
			},
			TLSHandshakeTimeout:    0,
			DisableKeepAlives:      false,
			DisableCompression:     false,
			MaxIdleConns:           0,
			MaxIdleConnsPerHost:    0,
			MaxConnsPerHost:        0,
			IdleConnTimeout:        0,
			ResponseHeaderTimeout:  0,
			ExpectContinueTimeout:  0,
			TLSNextProto:           nil,
			ProxyConnectHeader:     nil,
			MaxResponseHeaderBytes: 0,
			WriteBufferSize:        0,
			ReadBufferSize:         0,
			ForceAttemptHTTP2:      false,
		}
	}
	if nop {
		c.Transport = &NopTransport{}
	}
	http.DefaultClient = c
	return c

}
