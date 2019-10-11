package shopping

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	apiEndpoint               = "http://open.api.ebay.com/shopping"
	DefaultShoppingApiVersion = "1063"
)

type Api interface {
	GetCategoryInfo(req *GetCategoryInfoRequest) (GetCategoryInfoResponse, error)
	GetSingleItem(req *GetSingleItemRequest, aff AffiliateParams) (GetSingleItemResponse, error)
	GetMultipleItems(req *GetMultipleItemsRequest, aff AffiliateParams) (GetMultipleItemsResponse, error)
}

type Client struct {
	ApplicationId string
	Version       string
	HttpClient    *http.Client
}

var _ Api = &Client{}


func New(appId string, version string) *Client {
	s := Client{}
	s.ApplicationId = appId
	s.Version = version
	s.HttpClient = http.DefaultClient
	return &s
}

func (s *Client) SetHttpClient(httpClient *http.Client) {
	s.HttpClient = httpClient
}

func (s *Client) GetCategoryInfo(req *GetCategoryInfoRequest) (GetCategoryInfoResponse, error) {
	var resp GetCategoryInfoResponse

	b, err := xml.MarshalIndent(req, "", "\t")
	if err != nil {
		return resp, errors.New("Error serializing GetCategoryInfoRequest: " + err.Error())
	}

	x := xml.Header + string(b)

	httpResp, err := s.doShoppingRequest([]byte(x), "GetCategoryInfo", AffiliateParams{})
	if err != nil {
		return resp, err
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		return resp, errors.New("error reading GetCategoryInfoResponse: " + err.Error())
	}

	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return resp, errors.New("error deserializing response: " + err.Error())
	}

	return resp, nil
}

func (s *Client) GetSingleItem(req *GetSingleItemRequest, aff AffiliateParams) (GetSingleItemResponse, error) {
	var resp GetSingleItemResponse

	b, err := xml.Marshal(req)
	if err != nil {
		return resp, errors.New("Error serializing GetSingleItemRequest: " + err.Error())
	}

	b = bytes.Replace(b, []byte("<GetSingleItemRequest>"), []byte(xml.Header+"<GetSingleItemRequest xmlns=\"urn:ebay:apis:eBLBaseComponents\">"), 1)

	httpResp, err := s.doShoppingRequest(b, "GetSingleItem", aff)
	if err != nil {
		return resp, err
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return resp, errors.New("error reading GetSingleItemResponse: " + err.Error())
	}

	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return resp, errors.New("error deserializing response: " + err.Error())
	}

	return resp, nil
}

func (s *Client) GetMultipleItems(req *GetMultipleItemsRequest, aff AffiliateParams) (GetMultipleItemsResponse, error) {
	var resp GetMultipleItemsResponse

	b, err := xml.Marshal(req)
	if err != nil {
		return resp, err
	}

	b = bytes.Replace(b, []byte("<GetMultipleItemsRequest>"), []byte(xml.Header+"<GetMultipleItemsRequest xmlns=\"urn:ebay:apis:eBLBaseComponents\">"), 1)

	httpResp, err := s.doShoppingRequest(b, "GetMultipleItems", aff)
	if err != nil {
		return resp, errors.New("error making GetMultipleItems request: " + err.Error())
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return resp, errors.New("error reading GetMultipleItemsResponse: " + err.Error())
	}

	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return resp, errors.New("error deserializing response: " + err.Error())
	}

	return resp, nil
}

func (s *Client) doShoppingRequest(b []byte, callName string, aff AffiliateParams) (*http.Response, error) {
	var response http.Response

	request, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(b))
	if err != nil {
		return &response, errors.New("Error creating HTTP request: " + err.Error())
	}

	request.Header.Set("X-EBAY-API-REQUEST-ENCODING", "XML")

	q := url.Values{}

	//Set standard call parameters

	q.Add("appid", s.ApplicationId)
	q.Add("callname", callName)
	q.Add("version", s.Version)

	//Set affiliate call parameters
	if aff.TrackingId != "" {
		q.Add("trackingid", aff.TrackingId)
	}

	if aff.PartnerCode != "" {
		q.Add("trackingpartnercode", aff.PartnerCode)
	}

	if aff.AffiliateUserId != "" {
		q.Add("affiliateuserid", aff.AffiliateUserId)
	}

	request.URL.RawQuery = q.Encode()

	return s.HttpClient.Do(request)
}
