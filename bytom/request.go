package bytom

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type HTTPProvider struct {
	address string
	timeout int32
	client  *http.Client
}

var (
	//test
	port       = "0.0.0.0:9888"
	httpClient *http.Client
)

func NewHTTPProvider(timeout int32) *HTTPProvider {
	return NewHTTPProviderWithClient(timeout, httpClient)
}

func NewHTTPProviderWithClient(timeout int32, client *http.Client) *HTTPProvider {
	provider := new(HTTPProvider)
	provider.timeout = timeout
	provider.client = client

	return provider
}

func (provider HTTPProvider) SendRequest(v interface{}, params interface{}) error {

	bodyString := JSONRPCObject{Params: params}

	body := strings.NewReader(bodyString.AsJsonString())
	req, err := http.NewRequest("POST", port, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := provider.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	}

	return json.Unmarshal(bodyBytes, v)

}

func (provider HTTPProvider) Close() error { return nil }
