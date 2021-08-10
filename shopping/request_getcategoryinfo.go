package shopping

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

var _ Request = &GetCategoryInfoRequest{}

type GetCategoryInfoRequest struct {
	XMLName         xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetCategoryInfoRequest"`
	CategoryID      string   `xml:"CategoryID,omitempty"`
	IncludeSelector string   `xml:"IncludeSelector,omitempty"`
	MessageID       string   `xml:"MessageID,omitempty"`
}

func (r GetCategoryInfoRequest) CallName() string {
	return "GetCategoryInfo"
}

func (r GetCategoryInfoRequest) UrlValues() url.Values {
	v := url.Values{}
	v.Set("CategoryID", r.CategoryID)
	v.Set("IncludeSelector", r.IncludeSelector)
	v.Set("MessageID", r.MessageID)

	return v
}

type GetCategoryInfoResponse struct {
	*BaseShoppingResponse
	XmlName         xml.Name   `xml:"urn:ebay:apis:eBLBaseComponents GetCategoryInfoResponse"`
	Categories      []Category `xml:"CategoryArray>Category,omitempty"`
	CategoryCount   int        `xml:"CategoryCount,omitempty"`
	CategoryVersion string     `xml:"CategoryVersion,omitempty"`
	UpdateTime      string     `xml:"UpdateTime,omitempty"`
}

func (s *Client) GetCategoryInfo(req GetCategoryInfoRequest, aff AffiliateParams) (GetCategoryInfoResponse, error) {
	var resp GetCategoryInfoResponse
	var httpResp *http.Response

	httpResp, err := s.doRequest(req, aff)
	if err != nil {
		return resp, fmt.Errorf("error making GetMultipleItems request:  %w", err)
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		return resp, fmt.Errorf("error reading GetCategoryInfoResponse:  %w", err)
	}

	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return resp, fmt.Errorf("error deserializing response:  %w", err)
	}

	return resp, nil
}