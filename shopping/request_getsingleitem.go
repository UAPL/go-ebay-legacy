package shopping

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/url"
)

var _ Request = &GetSingleItemRequest{}

type GetSingleItemRequest struct {
	IncludeSelector string `xml:"IncludeSelector,omitempty"`
	ItemId          string `xml:"ItemID,omitempty"`
	MessageID       string `xml:"MessageID,omitempty"`
}

func (r GetSingleItemRequest) CallName() string {
	return "GetSingleItem"
}

func (r GetSingleItemRequest) UrlValues() url.Values {
	v := url.Values{}
	v.Set("ItemID", r.ItemId)
	v.Set("IncludeSelector", r.IncludeSelector)
	v.Set("MessageID", r.MessageID)

	return v
}

type GetSingleItemResponse struct {
	*BaseShoppingResponse
	XmlName xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetSingleItemResponse"`
	Item    SimpleItem     `xml:"Item"`
}

func (s *Client) GetSingleItem(req GetSingleItemRequest, aff AffiliateParams) (GetSingleItemResponse, error) {
	var resp GetSingleItemResponse

	httpResp, err := s.doRequest(req, aff)
	if err != nil {
		return resp, fmt.Errorf("error making GetMultipleItems request:  %w", err)
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return resp, fmt.Errorf("error reading GetSingleItemResponse:  %w", err)
	}

	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return resp, fmt.Errorf("error deserializing response: %w", err)
	}

	return resp, nil
}