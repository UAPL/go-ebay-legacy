package shopping

import (
	"encoding/xml"
	"net/url"
	"strings"
)

var _ Request = &GetMultipleItemsRequest{}

type GetMultipleItemsRequest struct {
	IncludeSelector string   `xml:"IncludeSelector,omitempty"`
	ItemIds         []string `xml:"ItemID,omitempty"`
	MessageID       string   `xml:"MessageID,omitempty"`
}

func (r *GetMultipleItemsRequest) CallName() string {
	return "GetMultipleItems"
}

func (r *GetMultipleItemsRequest) UrlValues() url.Values {
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
