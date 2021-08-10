package shopping

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/uapl/go-ebay-legacy/auth"
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
	httpClient    *http.Client
	authenticator auth.Authenticator
}

func New(a auth.Authenticator) *Client {
	s := Client{}
	s.httpClient = http.DefaultClient
	s.authenticator = a
	return &s
}

func (s *Client) SetHttpClient(httpClient *http.Client) {
	s.httpClient = httpClient
}

func (s *Client) SetAuthenticator(auther auth.Authenticator) {
	s.authenticator = auther
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
			return nil, fmt.Errorf("error marshalling request (%s) to XML: %w", req.CallName(), err)
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
		return &response, fmt.Errorf("Error creating HTTP request:  %w", err)
	}

	if err = s.prepareRequestHeaders(request); err != nil {
		return nil, fmt.Errorf("error setting request headers: %w", err)
	}

	//Set standard call parameters

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
	return s.httpClient.Do(request)
}

func (s *Client) prepareRequestHeaders(req *http.Request) error {
	token, err := s.authenticator.AuthToken()
	if err != nil {
		return err
	}

	req.Header.Set("X-EBAY-API-REQUEST-ENCODING", "XML")

	req.Header.Set("Accept-Charset", "utf-8")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("Content-Language", "en-US")
	req.Header.Set("X-EBAY-C-MARKETPLACE-ID", "EBAY_US") //TODO make marketplace configurable

	if token.AccessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
		//for Traditional APIs, we actually need to use X-EBAY-API-IAF-TOKEN
		req.Header.Set("X-EBAY-API-IAF-TOKEN", token.AccessToken)
	}

	return nil
}
