package shopping

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
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

func (r GetItemStatusRequest) UrlValues() url.Values {
	v := url.Values{}
	v.Set("ItemID", strings.Join(r.ItemIds, ","))
	v.Set("MessageID", r.MessageID)

	return v
}

func (g GetItemStatusRequest) CallName() string {
	panic("GetItemStatus")
}

func (s *Client) GetItemStatus(req GetItemStatusRequest, aff AffiliateParams) (GetItemStatusResponse, error) {
	var resp GetItemStatusResponse
	var httpResp *http.Response

	httpResp, err := s.doRequest(req, aff)
	if err != nil {
		return resp, fmt.Errorf("error making GetItemStatus request:  %w", err)
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		return resp, fmt.Errorf("error reading GetItemStatus http response:  %w", err)
	}

	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return resp, fmt.Errorf("error deserializing GetItemStatusResponse:  %w", err)
	}

	return resp, nil
}
