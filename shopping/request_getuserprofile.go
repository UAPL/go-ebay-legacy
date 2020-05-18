package shopping

import (
	"encoding/xml"
	"net/url"
)

var _ Request = &GetUserProfileRequest{}

type GetUserProfileRequest struct {
	IncludeSelector string `xml:"IncludeSelector,omitempty"`
	UserId          string `xml:"UserID,omitempty"`
	MessageID       string `xml:"MessageID,omitempty"`
}

func (r *GetUserProfileRequest) CallName() string {
	return "GetUserProfile"
}

func (r *GetUserProfileRequest) UrlValues() url.Values {
	v := url.Values{}
	v.Set("UserID", r.UserId)
	v.Set("IncludeSelector", r.IncludeSelector)
	v.Set("MessageID", r.MessageID)

	return v
}

type GetUserProfileResponse struct {
	*BaseShoppingResponse
	XmlName         xml.Name        `xml:"urn:ebay:apis:eBLBaseComponents GetUserProfileResponse"`
	FeedbackDetails FeedbackDetail  `xml:"FeedbackDetails"`
	FeedbackHistory FeedbackHistory `xml:"FeedbackHistory"`
	User            SimpleUser      `xml:"User"`
}
