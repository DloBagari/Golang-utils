package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func HttpClient() (client *http.Client) {
	uckey := "/home/dlo/fuid_certs/client.fuid.key"
	ucert := "/home/dlo/fuid_certs/client.fuid.crt"
	x509cert, err := tls.LoadX509KeyPair(ucert, uckey)
	if err != nil {
		panic(err.Error())
	}
	certs := []tls.Certificate{x509cert}
	if len(certs) == 0 {
		client = &http.Client{}
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{Certificates: certs,
			InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	return
}

func main() {
	rurl := "https://51.91.237.250:5000/api/uid/v1.0/users?domain=example.com"
	client := HttpClient()
	req, err := http.NewRequest("GET", rurl, nil)
	if err != nil {
		log.Println("Unable to make GET request", err)
		os.Exit(1)
	}
	//username := "dlo"
	//password := "Forcepoint1"
	//login := fmt.Sprintf("%s:%s", username, password)
	//basic := base64.StdEncoding.EncodeToString([]byte(login))
	//basic = fmt.Sprintf("Basic %s", basic)
	//req.Header.Set("Authorization", basic)
	//req.Header.Add("Accept", "*/*")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	log.Println(string(data))
}
