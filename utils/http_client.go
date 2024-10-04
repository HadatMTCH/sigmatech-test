package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

func client() *http.Client {
	return &http.Client{
		// Timeout: time.Second * constants.HttpClientDefaultTimeout,
		Transport: &http.Transport{
			DisableCompression: true,
			DisableKeepAlives:  true,
			// Dial: (&net.Dialer{
			// 	Timeout: 60 * time.Second,
			// }).Dial,
			// TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}

func do(method string, baseUrl string, params map[string]string, headers map[string]string, body io.Reader) ([]byte, *string, error) {
	req, err := http.NewRequest(method, baseUrl, body)
	if err != nil {
		return nil, nil, err
	}
	req.Close = true

	q := req.URL.Query()
	for i, v := range params {
		q.Add(i, v)
	}
	req.URL.RawQuery = q.Encode()

	for i, v := range headers {
		req.Header.Add(i, v)
	}

	netClient := client()
	resp, err := netClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != 200 {
		err = ErrHttpClient(baseUrl)
		errResp := ConvertBytesToString(buf)
		log.Error(err, errResp)
		return nil, &errResp, err
	}

	return buf, nil, nil
}

func HttpClientDoJson(method string, baseUrl string, params map[string]string, headers map[string]string, body interface{}, result interface{}) (*string, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(jsonBody)
	buf, errResp, err := do(method, baseUrl, params, headers, reqBody)
	if err != nil {
		return errResp, err
	}

	err = json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func HttpClientDo(method string, baseUrl string, params map[string]string, headers map[string]string, result interface{}) (*string, error) {
	buf, errResp, err := do(method, baseUrl, params, headers, nil)
	if err != nil {
		return errResp, err
	}

	err = json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func HttpClientDoUrlEncoded(method string, baseUrl string, params map[string]string, headers map[string]string, body map[string]string, result interface{}) (*string, error) {
	data := url.Values{}
	for i, v := range body {
		data.Set(i, v)
	}

	reqBody := strings.NewReader(data.Encode())
	buf, errResp, err := do(method, baseUrl, params, headers, reqBody)
	if err != nil {
		return errResp, err
	}

	err = json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
