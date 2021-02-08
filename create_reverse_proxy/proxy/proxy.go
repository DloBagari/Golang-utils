package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Proxy struct {
	Client  *http.Client
	BaseUrl string
}

//implement the handler interface, which manipulate the request and forwards it to BaseUrl, then return the response
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := p.ProcessRequest(r); err != nil {
		logrus.Errorf("error occured during processing the request: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := p.Client.Do(r)
	if err != nil {
		logrus.Errorf("error occur during client operation: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	copyResponse(w, resp)
}

func copyResponse(writer http.ResponseWriter, response *http.Response) {
	var out bytes.Buffer
	_, _ = out.ReadFrom(response.Body)
	for key, values := range response.Header {
		for _, v := range values {
			writer.Header().Add(key, v)
		}
	}
	writer.WriteHeader(response.StatusCode)
	_, _ = writer.Write(out.Bytes())
}

func (p *Proxy) ProcessRequest(r *http.Request) error {
	logrus.Infof(" r.URL= %s", r.URL.String())
	proxyURLRaw := p.BaseUrl + r.URL.String()
	logrus.Infof("proxyURLRaw= %s", proxyURLRaw)
	proxyURL, err := url.Parse(proxyURLRaw)
	logrus.Infof("proxyURL= %s", proxyURL)

	if err != nil {
		return err
	}
	r.URL = proxyURL
	r.Host = proxyURL.Host
	r.RequestURI = ""
	return nil
}
