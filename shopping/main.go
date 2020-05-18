package shopping

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	ApiEndpoint   = "http://open.api.ebay.com/shopping"
	ApiVersion    = "1063"
	RequestMethod = "POST"
)

type Api interface {
	GetCategoryInfo(req *GetCategoryInfoRequest, aff AffiliateParams) (GetCategoryInfoResponse, error)
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

func (s *Client) GetCategoryInfo(req *GetCategoryInfoRequest, aff AffiliateParams) (GetCategoryInfoResponse, error) {
	var resp GetCategoryInfoResponse
	var httpResp *http.Response

	httpResp, err := s.doRequest(req, aff)
	if err != nil {
		return resp, errors.New("error making GetMultipleItems request: " + err.Error())
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

	httpResp, err := s.doRequest(req, aff)
	if err != nil {
		return resp, errors.New("error making GetMultipleItems request: " + err.Error())
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

	httpResp, err := s.doRequest(req, aff)
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

func (s *Client) doRequest(req Request, aff AffiliateParams) (*http.Response, error) {
	var response http.Response
	var b []byte
	var err error
	var q url.Values

	switch RequestMethod {

	case "POST":
		q = url.Values{}
		b, err = xml.Marshal(req)
		if err != nil {
			return nil, err
		}

		b = bytes.Replace(b, []byte(fmt.Sprintf("<%sRequest>", req.CallName())), []byte(fmt.Sprintf("%s<%sRequest xmlns=\"urn:ebay:apis:eBLBaseComponents\">", xml.Header, req.CallName())), 1)
	case "GET":
		q = req.UrlValues()

	default:
		return nil, fmt.Errorf("unsupported http request method: %s", RequestMethod)
	}

	request, err := http.NewRequest(RequestMethod, ApiEndpoint, bytes.NewBuffer(b))
	if err != nil {
		return &response, errors.New("Error creating HTTP request: " + err.Error())
	}

	request.Header.Set("X-EBAY-API-REQUEST-ENCODING", "XML")

	//Set standard call parameters

	q.Set("appid", s.ApplicationId)
	q.Set("callname", req.CallName())
	q.Set("version", s.Version)

	//Set affiliate call parameters
	if aff.TrackingId != "" {
		q.Set("trackingid", aff.TrackingId)
	}

	if aff.PartnerCode != "" {
		q.Set("trackingpartnercode", aff.PartnerCode)
	}

	if aff.AffiliateUserId != "" {
		q.Set("affiliateuserid", aff.AffiliateUserId)
	}

	request.URL.RawQuery = q.Encode()
	return s.HttpClient.Do(request)
}
