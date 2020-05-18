package shopping

import (
	"encoding/xml"
	"net/url"
)

var _ Request = &GetSingleItemRequest{}

type GetSingleItemRequest struct {
	IncludeSelector string `xml:"IncludeSelector,omitempty"`
	ItemId          string `xml:"ItemID,omitempty"`
	MessageID       string `xml:"MessageID,omitempty"`
}

func (r *GetSingleItemRequest) CallName() string {
	return "GetSingleItem"
}

func (r *GetSingleItemRequest) UrlValues() url.Values {
	v := url.Values{}
	v.Set("ItemID", r.ItemId)
	v.Set("IncludeSelector", r.IncludeSelector)
	v.Set("MessageID", r.MessageID)

	return v
}

type GetSingleItemResponse struct {
	*BaseShoppingResponse
	XmlName xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetSingleItemResponse"`
	Item    Item     `xml:"Item"`
}
