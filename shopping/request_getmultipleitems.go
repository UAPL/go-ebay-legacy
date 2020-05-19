package shopping

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/url"
	"strings"
)

var _ Request = &GetMultipleItemsRequest{}

type GetMultipleItemsRequest struct {
	IncludeSelector string   `xml:"IncludeSelector,omitempty"`
	ItemIds         []string `xml:"ItemID,omitempty"`
	MessageID       string   `xml:"MessageID,omitempty"`
}

func (r GetMultipleItemsRequest) CallName() string {
	return "GetMultipleItems"
}

func (r GetMultipleItemsRequest) UrlValues() url.Values {
	v := url.Values{}
	v.Set("ItemID", strings.Join(r.ItemIds, ","))
	v.Set("IncludeSelector", r.IncludeSelector)
	v.Set("MessageID", r.MessageID)

	return v
}

type GetMultipleItemsResponse struct {
	*BaseShoppingResponse
	XmlName xml.Name     `xml:"GetMultipleItemsResponse"`
	Items   []SimpleItem `xml:"Item"`
}

func (s *Client) GetMultipleItems(req GetMultipleItemsRequest, aff AffiliateParams) (GetMultipleItemsResponse, error) {
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