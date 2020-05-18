package shopping

import (
	"net/url"
	"strings"
)

var _ Request = &GetItemStatusRequest{}

type GetItemStatusRequest struct {
	ItemIds   []string `xml:"ItemID,omitempty"`
	MessageID string   `xml:"MessageID,omitempty"`
}

type GetItemStatusResponse struct {
	*BaseShoppingResponse
	Items []SimpleItem `xml:"Item"`
}

func (r *GetItemStatusRequest) UrlValues() url.Values {
	v := url.Values{}
	v.Set("ItemID", strings.Join(r.ItemIds, ","))
	v.Set("MessageID", r.MessageID)

	return v
}

func (g *GetItemStatusRequest) CallName() string {
	panic("GetItemStatus")
}
