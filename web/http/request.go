package http

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendUnsecureHTTPSRequest(ip string, method string, body string, headers map[string]string) (int, string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	var req *http.Request
	var err error

	if !strings.Contains(ip, "https") {
		ip = "https://" + ip
	}

	method = strings.ToUpper(method)

	if method == "DELETE" {
		req, err = http.NewRequest(method, ip, nil)
	} else {
		req, err = http.NewRequest(method, ip, strings.NewReader(body))
	}

	if err == nil {
		if headers != nil {
			for k, v := range headers {
				req.Header.Add(k, v)
			}
		}
		response, err := client.Do(req)
		if err == nil {
			defer response.Body.Close()
			if response.StatusCode == 200 {
				bodyBytes, err2 := ioutil.ReadAll(response.Body)
				bodyString := string(bodyBytes)
				if err2 == nil {
					return response.StatusCode, string(bodyString), nil
				}
				return response.StatusCode, "", err2
			}
		} else {
			return 0, "", err
		}
	}
	return 0, "", nil
}
