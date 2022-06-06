package http

import (
	"crypto/tls"
	"github.com/parnurzeal/gorequest"
	"gitlab-webhook/constants"
	"net/http"
	"time"
)

var (
	debugClient bool
	timeOut     string
	retryCount  int
)

func init() {
	debugClient = true
	timeOut = "60s"
	retryCount = 1
}

// HTTPGet ..
func HTTPGet(url string, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClient)
	timeout, _ := time.ParseDuration(timeOut)
	reqAgent := request.Get(url)
	reqAgent.Header = header
	_, body, errs := reqAgent.
		Timeout(timeout).
		Retry(retryCount, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPost ..
func HTTPPost(url string, jsonData interface{}) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClient)
	timeout, _ := time.ParseDuration(timeOut)
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqAgent := request.Post(url)
	reqAgent.Header.Set("Content-Type", "application/json")
	_, body, errs := reqAgent.
		Send(jsonData).
		Timeout(timeout).
		Retry(retryCount, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPostWithHeader ..
func HTTPPostWithHeader(url string, jsonData interface{}, header http.Header) (gorequest.Response, []byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClient)
	timeout, _ := time.ParseDuration(timeOut)
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqAgent := request.Post(url)
	reqAgent.Header = header
	response, body, errs := reqAgent.
		Send(jsonData).
		Timeout(timeout).
		Retry(retryCount, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return nil, []byte(body), errs[0]
	}
	return response, []byte(body), nil
}

// HTTPPostFile ..
func HTTPPostFile(url string, jsonData interface{}, attachment []byte) (gorequest.Response, []byte, error) {
	var a gorequest.Response
	if len(attachment) == 0 {
		result, err := HTTPPost(url, jsonData)
		return a, result, err
	}
	request := gorequest.New()
	request.SetDebug(debugClient)
	timeout, _ := time.ParseDuration(timeOut)
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqAgent := request.Post(url)
	response, body, errs := reqAgent.
		Type("multipart").
		SendFile(attachment, "output.txt").
		Send(jsonData).
		Timeout(timeout).
		Retry(retryCount, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return nil, []byte(body), errs[0]
	}
	return response, []byte(body), nil
}

// HTTPPutWithHeader ..
func HTTPPutWithHeader(url string, jsonData interface{}, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClient)
	timeout, _ := time.ParseDuration(timeOut)
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqAgent := request.Put(url)
	reqAgent.Header = header
	_, body, errs := reqAgent.
		Send(jsonData).
		Timeout(timeout).
		Retry(retryCount, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPDeleteWithHeader ..
func HTTPDeleteWithHeader(url string, jsonData interface{}, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClient)
	timeout, _ := time.ParseDuration(timeOut)
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqAgent := request.Delete(url)
	reqAgent.Header = header
	_, body, errs := reqAgent.
		Send(jsonData).
		Timeout(timeout).
		Retry(retryCount, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// SendHttpRequest ..
func SendHttpRequest(method string, url string, header http.Header, body interface{}) (gorequest.Response, []byte, error) {
	var response gorequest.Response
	var data []byte
	var err error
	switch method {
	case constants.HttpMethodGet:
		data, err = HTTPGet(url, header)
	case constants.HttpMethodPost:
		response, data, err = HTTPPostWithHeader(url, body, header)
	case constants.HttpMethodPut:
		data, err = HTTPPutWithHeader(url, body, header)
	case constants.HttpMethodDelete:
		data, err = HTTPDeleteWithHeader(url, body, header)
	}
	return response, data, err
}
