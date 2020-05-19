package shopping

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

var _ Request = &GetUserProfileRequest{}

type GetUserProfileRequest struct {
	IncludeSelector string `xml:"IncludeSelector,omitempty"`
	UserId          string `xml:"UserID,omitempty"`
	MessageID       string `xml:"MessageID,omitempty"`
}

func (r GetUserProfileRequest) CallName() string {
	return "GetUserProfile"
}

func (r GetUserProfileRequest) UrlValues() url.Values {
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

func (s *Client) GetUserProfile(req GetUserProfileRequest, aff AffiliateParams) (GetUserProfileResponse, error) {
	var resp GetUserProfileResponse
	var httpResp *http.Response

	httpResp, err := s.doRequest(req, aff)
	if err != nil {
		return resp, errors.New("error making GetUserProfile request: " + err.Error())
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		return resp, errors.New("error reading GetUserProfile http response: " + err.Error())
	}

	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return resp, errors.New("error deserializing GetUserProfileResponse: " + err.Error())
	}

	return resp, nil
}
