package shopping

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var (
	ApiEndpoint   = "http://open.api.ebay.com/shopping"
	ApiVersion    = "1063"
	RequestMethod = "POST"
)

var _ Api = &Client{}

type Api interface {
	GetCategoryInfo(req GetCategoryInfoRequest, aff AffiliateParams) (GetCategoryInfoResponse, error)
	GetSingleItem(req GetSingleItemRequest, aff AffiliateParams) (GetSingleItemResponse, error)
	GetMultipleItems(req GetMultipleItemsRequest, aff AffiliateParams) (GetMultipleItemsResponse, error)
	GetUserProfile(req GetUserProfileRequest, aff AffiliateParams) (GetUserProfileResponse, error)
	GetItemStatus(req GetItemStatusRequest, aff AffiliateParams) (GetItemStatusResponse, error)
}

type Client struct {
	ApplicationId string
	HttpClient    *http.Client
}

func New(appId string) *Client {
	s := Client{}
	s.ApplicationId = appId
	s.HttpClient = http.DefaultClient
	return &s
}

func (s *Client) SetHttpClient(httpClient *http.Client) {
	s.HttpClient = httpClient
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

		//fix the XML header and namespacing errors
		b = bytes.Replace(b, []byte(fmt.Sprintf("<%sRequest>", req.CallName())), []byte(fmt.Sprintf("<%sRequest xmlns=\"urn:ebay:apis:eBLBaseComponents\">", req.CallName())), 1)
		b = append([]byte(xml.Header), b...)
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
	q.Set("version", ApiVersion)

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
