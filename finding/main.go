package finding

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	findingServiceUrl = "https://svcs.ebay.com/services/search/FindingService/v1"
)

type Api interface {
	FindItemsAdvanced(req *FindItemsAdvancedRequest) (FindItemsAdvancedResponse, error)
}

type Client struct {
	applicationId string
	httpClient    *http.Client
}

var _ Api = &Client{}

func New(application_id string) *Client {
	e := new(Client)
	e.applicationId = application_id
	e.httpClient = http.DefaultClient
	return e
}

func (f *Client) SetHttpClient(httpClient *http.Client) {
	f.httpClient = httpClient
}

func (f *Client) FindItemsAdvanced(req *FindItemsAdvancedRequest) (FindItemsAdvancedResponse, error) {
	var response FindItemsAdvancedResponse

	b, err := xml.Marshal(req)
	if err != nil {
		return response, err
	}

	x := xml.Header + string(b)

	resp, err := f.doFindingServiceRequest([]byte(x), "findItemsAdvanced")
	if err != nil {
		return response, errors.New("error making findItemsAdvanced request: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, errors.New("error reading findItemsAdvanced response: " + err.Error())
	}

	err = xml.Unmarshal(body, &response)
	if err != nil {
		return response, errors.New(fmt.Sprintf("error deserializing response: %v, %s", err, string(body)))
	}

	return response, nil
}

func (f *Client) doFindingServiceRequest(b []byte, callName string) (*http.Response, error) {
	var response http.Response

	request, err := http.NewRequest("POST", findingServiceUrl, bytes.NewBuffer(b))
	if err != nil {
		return &response, errors.New("Error creating HTTP request: " + err.Error())
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
