package shopping

import (
	"encoding/xml"
	"errors"
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