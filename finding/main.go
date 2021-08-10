package finding

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	ServiceUrl = "https://svcs.ebay.com/services/search/FindingService/v1"
)

type Api interface {
	FindItemsAdvanced(req FindItemsAdvancedRequest) (FindItemsAdvancedResponse, error)
	GetHistograms(req GetHistogramsRequest) (GetHistogramsResponse, error)
}

type Client struct {
	applicationId string
	httpClient    *http.Client
}

var _ Api = &Client{}

func New(applicationId string) *Client {
	e := new(Client)
	e.applicationId = applicationId
	e.httpClient = http.DefaultClient
	return e
}

func (f *Client) SetHttpClient(httpClient *http.Client) {
	f.httpClient = httpClient
}

func (f *Client) FindItemsAdvanced(req FindItemsAdvancedRequest) (FindItemsAdvancedResponse, error) {
	var response FindItemsAdvancedResponse

	b, err := xml.Marshal(req)
	if err != nil {
		return response, err
	}

	x := xml.Header + string(b)

	resp, err := f.doFindingServiceRequest([]byte(x), "findItemsAdvanced")
	if err != nil {
		return response, fmt.Errorf("error making findItemsAdvanced request:  %w", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("error reading findItemsAdvanced response:  %w", err)
	}

	err = xml.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf("error deserializing FindItemsAdvanced response: %w, %s", err, string(body))
	}

	return response, nil
}

func (f *Client) GetHistograms(req GetHistogramsRequest) (GetHistogramsResponse, error) {
	var response GetHistogramsResponse

	b, err := xml.Marshal(req)
	if err != nil {
		return response, err
	}

	x := xml.Header + string(b)

	resp, err := f.doFindingServiceRequest([]byte(x), "getHistograms")
	if err != nil {
		return response, fmt.Errorf("error making getHistograms request: %w", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("error reading getHistograms response:  %w", err)
	}

	err = xml.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf("error deserializing GetHistograms response: %w, %s", err, string(body))
	}

	return response, nil
}

func (f *Client) doFindingServiceRequest(b []byte, callName string) (*http.Response, error) {
	var response http.Response

	request, err := http.NewRequest("POST", ServiceUrl, bytes.NewBuffer(b))
	if err != nil {
		return &response, fmt.Errorf("Error creating HTTP request:  %w", err)
	}

	q := url.Values{}

	//Set standard call parameters
	q.Add("OPERATION-NAME", callName)
	q.Add("SECURITY-APPNAME", f.applicationId)
	q.Add("RESPONSE-DATA-FORMAT", "xml")
	q.Add("GLOBAL-ID", "EBAY-US")

	request.URL.RawQuery = q.Encode()

	return f.httpClient.Do(request)
}
