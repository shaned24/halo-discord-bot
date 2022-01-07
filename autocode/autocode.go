package autocode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"toughcrab.com/halo/autocode/generated"
)

type Client struct {
	URL        string
	APIKey     string
	HttpClient *http.Client
}

func NewAutoCodeClient(url string, apiKey string) *Client {
	return &Client{
		URL:        url,
		APIKey:     apiKey,
		HttpClient: &http.Client{},
	}
}

func (c *Client) ServiceRecordMultiplayer(gamertag string, filter string) (*generated.ServiceRecordMultiplayer, error) {
	bodyParams := map[string]string{"gamertag": gamertag}

	body, err := c.sendRequest(bodyParams, "POST")
	if err != nil {
		return nil, err
	}

	var serviceRecord *generated.ServiceRecordMultiplayer
	if err := json.Unmarshal(body, &serviceRecord); err != nil {
		panic(err)
	}

	return serviceRecord, nil
}

func (c *Client) sendRequest(params map[string]string, requestType string) ([]byte, error) {
	jsonData, _ := json.Marshal(params)
	request, err := http.NewRequest(requestType, c.URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
