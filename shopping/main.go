package shopping

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	apiEndpoint = "http://open.api.ebay.com/shopping"
)

type ShoppingApi struct {
	ApplicationId string
	Version       string
	HttpClient    *http.Client
}

func New(app_id string, version string) *ShoppingApi {
	s := ShoppingApi{}
	s.ApplicationId = app_id
	s.Version = version
	s.HttpClient = http.DefaultClient
	return &s
}

func (s *ShoppingApi) SetHttpClient(httpClient *http.Client) {
	s.HttpClient = httpClient
}

func (s *ShoppingApi) GetSingleItem(req *GetSingleItemRequest, aff *AffiliateParams) (GetSingleItemResponse, error) {
	var resp GetSingleItemResponse

	b, err := xml.Marshal(req)
	if err != nil {
		return resp, errors.New("Error serializing GetSingleItemRequest: " + err.Error())
	}

	httpResp, err := s.doShoppingRequest(b, "GetSingleItem", aff)
	if err != nil {
		return resp, err
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return resp, errors.New("error reading GitSingleItemResponse: " + err.Error())
	}

	err = xml.Unmarshal(body, resp)
	if err != nil {
		return resp, errors.New("error deserializing response: " + err.Error())
	}

	return resp, nil
}

func (s *ShoppingApi) GetMultipleItems(req *GetMultipleItemsRequest, aff *AffiliateParams) (GetMultipleItemsResponse, error) {
	var resp GetMultipleItemsResponse

	b, err := xml.Marshal(req)
	if err != nil {
		return resp, err
	}

	httpResp, err := s.doShoppingRequest(b, "GetMultipleItems", aff)
	if err != nil {
		return resp, errors.New("error making GetMultipleItems request: " + err.Error())
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return resp, errors.New("error reading GetMultipleItemsResponse: " + err.Error())
	}

	err = xml.Unmarshal(body, resp)
	if err != nil {
		return resp, errors.New("error deserializing response: " + err.Error())
	}

	return resp, nil
}

func (s *ShoppingApi) doShoppingRequest(b []byte, callName string, aff *AffiliateParams) (*http.Response, error) {
	var response http.Response

	request, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(b))
	if err != nil {
		return &response, errors.New("Error creating HTTP request: " + err.Error())
	}

	//Set standard call parameters
	request.URL.Query().Set("appid", s.ApplicationId)
	request.URL.Query().Set("callname", callName)
	request.URL.Query().Set("requestencoding", "XML")
	request.URL.Query().Set("responseencoding", "XML")
	request.URL.Query().Set("version", s.Version)

	//Set affiliate call parameters
	request.URL.Query().Set("trackingid", aff.TrackingId)
	request.URL.Query().Set("trackingpartnercode", aff.PartnerCode)
	request.URL.Query().Set("affiliateuserid", aff.AffiliateUserId)

	return s.HttpClient.Do(request)
}
