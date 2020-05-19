package shopping

import (
	"encoding/xml"
	"errors"
	"net/url"
)

type AckCode string

const (
	Failure        AckCode = "Failure"
	PartialFailure AckCode = "PartialFailure"
	Success        AckCode = "Success"
	Warning        AckCode = "Warning"
	CustomCode     AckCode = "CustomCode"
)

type AffiliateParams struct {
	TrackingId      string
	PartnerCode     string
	AffiliateUserId string
}

type Request interface {
	UrlValues() url.Values
	CallName() string
}

func (a AffiliateParams) UrlValues() url.Values {
	v := url.Values{}
	v.Set("trackingid", a.TrackingId)
	v.Set("trackingpartnercode", a.PartnerCode)
	v.Set("affiliateuserid", a.AffiliateUserId)

	return v
}

type Ack struct {
	Value AckCode `xml:",chardata"`
}

func (c *Ack) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "Failure":
		c.Value = Failure
	case "PartialFailure":
		c.Value = PartialFailure
	case "Success":
		c.Value = Success
	case "Warning":
		c.Value = Warning
	case "CustomCode":
		c.Value = CustomCode
	default:
		return errors.New("invalid Ack code received")
	}

	return nil
}

type ErrorParameter struct {
	ParamID string `xml:"ParamID,attr"`
	Value   string `xml:",chardata"`
}

type Error struct {
	ErrorCode           string           `xml:"ErrorCode,omitempty"`
	ShortMessage        string           `xml:"ShortMessage,omitempty"`
	LongMessage         string           `xml:"LongMessage,omitempty"`
	SeverityCode        string           `xml:"SeverityCode,omitempty"`
	ErrorParameters     []ErrorParameter `xml:"ErrorParameters,omitempty"`
	ErrorClassification string           `xml:"ErrorClassification,omitempty"`
}

type BaseShoppingResponse struct {
	Timestamp     string `xml:"Timestamp"`
	Ack           Ack    `xml:"Ack"`
	Build         string `xml:"Build"`
	Errors        Error  `xml:"Errors"`
	Version       string `xml:"Version"`
	CorrelationID string `xml:"CorrelationID"`
}
