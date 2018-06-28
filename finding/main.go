package finding

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	findingServiceUrl = "http://svcs.ebay.com/services/search/FindingService/v1"
)

type FindingApi struct {
	applicationId string
	httpClient    *http.Client
}

func New(application_id string) *FindingApi {
	e := FindingApi{}
	e.applicationId = application_id
	e.httpClient = http.DefaultClient
	return &e
}

func (f *FindingApi) SetHttpClient(httpClient *http.Client) {
	f.httpClient = httpClient
}

func (f *FindingApi) FindItemsAdvanced(req FindItemsAdvancedRequest) (FindItemsAdvancedResponse, error) {
	var response FindItemsAdvancedResponse

	b, err := xml.Marshal(req)
	if err != nil {
		return response, err
	}

	resp, err := f.doFindingServiceRequest(b, "findItemsAdvanced")
	if err != nil {
		return response, errors.New("error making findItemsAdvanced request: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, errors.New("error reading findItemsAdvanced response: " + err.Error())
	}

	err = xml.Unmarshal(body, response)
	if err != nil {
		return response, errors.New("error deserializing response: " + err.Error())
	}

	return response, nil
}

func (f *FindingApi) doFindingServiceRequest(b []byte, callName string) (*http.Response, error) {
	var response http.Response

	request, err := http.NewRequest("POST", findingServiceUrl, bytes.NewBuffer(b))
	if err != nil {
		return &response, errors.New("Error creating HTTP request: " + err.Error())
	}

	//Set standard call parameters
	request.URL.Query().Set("OPERATION-NAME", callName)
	request.URL.Query().Set("SECURITY-APPNAME", f.applicationId)
	request.URL.Query().Set("RESPONSE-DATA-FORMAT", "text/xml")
	request.URL.Query().Set("GLOBAL-ID", "EBAY-US")

	return f.httpClient.Do(request)
}
