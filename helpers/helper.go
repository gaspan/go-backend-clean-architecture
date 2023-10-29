package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string, body []byte) ([]byte, error) {
	return doRequest("GET", url, body)
}

var client = &http.Client{}
var headers = map[string]string{
	"Content-Type": "application/json",
	"Accept":       "application/json",
}

func isStatusError(statusCode int) bool {
	return statusCode >= http.StatusBadRequest
}

func doRequest(method string, url string, body []byte) ([]byte, error) {

	payload := strings.NewReader(string(body))

	req, err := http.NewRequest(method, url, payload)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	fmt.Println("LOG ERR", err)

	if err != nil {
		return nil, err
	}

	fmt.Println("LOG REQ", req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	// fmt.Println("LOG RES", string(respBody))
	if isStatusError(resp.StatusCode) {
		return nil, fmt.Errorf("Status error: %v %v", resp.StatusCode, string(respBody))
	}

	defer resp.Body.Close()
	return respBody, err
}
