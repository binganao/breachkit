package lib

import (
	"crypto/tls"
	"github.com/binganao/breachkit/config"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type response struct {
	Status         bool
	StatusCode     string
	ResponseHeader map[string][]string
	ResponseBody   string
}

func requests(method string, url string, headers map[string]string, body string) response {
	var Response response
	var Status = false
	var StatusCode = ""
	var ResponseHeader = map[string][]string{}
	var ResponseBody = ""

	Response.Status = Status
	Response.StatusCode = StatusCode
	Response.ResponseHeader = ResponseHeader
	Response.ResponseBody = ResponseBody

	var client = &http.Client{
		Timeout: time.Duration(TimeOut) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	req, err := http.NewRequest(method, url, strings.NewReader(body))

	if err != nil {
		return Response
	}

	for key, value := range config.DefaultHeader {
		req.Header.Set(key, value)
	}

	if method == "POST" && headers["Content-Type"] == "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)

	if err != nil {
		return Response
	}

	defer resp.Body.Close()

	ByteResponse, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Response
	}

	Status = true
	StatusCode = strconv.Itoa(resp.StatusCode)
	for key, value := range resp.Header {
		ResponseHeader[key] = value
	}
	ResponseBody = string(ByteResponse)

	Response.Status = Status
	Response.StatusCode = StatusCode
	Response.ResponseHeader = ResponseHeader
	Response.ResponseBody = ResponseBody

	return Response
}
