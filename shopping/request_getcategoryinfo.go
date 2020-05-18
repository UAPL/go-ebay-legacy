package shopping

import (
	"encoding/xml"
	"net/url"
)

var _ Request = &GetCategoryInfoRequest{}

type GetCategoryInfoRequest struct {
	XMLName         xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetCategoryInfoRequest"`
	CategoryID      string   `xml:"CategoryID,omitempty"`
	IncludeSelector string   `xml:"IncludeSelector,omitempty"`
	MessageID       string   `xml:"MessageID,omitempty"`
}

func (r *GetCategoryInfoRequest) CallName() string {
	return "GetCategoryInfo"
}

func (r *GetCategoryInfoRequest) UrlValues() url.Values {
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