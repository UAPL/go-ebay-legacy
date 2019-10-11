package shopping

import (
	"encoding/xml"
	"errors"
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

type Ack struct {
	Value AckCode `xml:",chardata"`
}

func (c *Ack) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)

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

type BaseRequest struct {
}

type BaseShoppingResponse struct {
	Timestamp     string `xml:"Timestamp"`
	Ack           Ack    `xml:"Ack"`
	Build         string `xml:"Build"`
	Errors        Error  `xml:"Errors"`
	Version       string `xml:"Version"`
	CorrelationID string `xml:"CorrelationID"`
}

type GetMultipleItemsRequest struct {
	IncludeSelector string   `xml:"IncludeSelector,omitempty"`
	ItemIds         []string `xml:"ItemID,omitempty"`
}

type GetMultipleItemsResponse struct {
	*BaseShoppingResponse
	XmlName xml.Name `xml:"GetMultipleItemsResponse"`
	Items   []Item   `xml:"Item"`
}

type GetSingleItemRequest struct {
	IncludeSelector string `xml:"IncludeSelector,omitempty"`
	ItemId          string `xml:"ItemID,omitempty"`
	MessageID       string `xml:"MessageID,omitempty"`
}

type GetSingleItemResponse struct {
	*BaseShoppingResponse
	XmlName xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetSingleItemResponse"`
	Item    Item     `xml:"Item"`
}

type GetCategoryInfoRequest struct {
	XMLName         xml.Name `xml:"urn:ebay:apis:eBLBaseComponents GetCategoryInfoRequest"`
	CategoryID      string   `xml:"CategoryID,omitempty"`
	IncludeSelector string   `xml:"IncludeSelector,omitempty"`
	MessageID       string   `xml:"MessageID,omitempty"`
}

type GetCategoryInfoResponse struct {
	*BaseShoppingResponse
	XmlName         xml.Name   `xml:"urn:ebay:apis:eBLBaseComponents GetCategoryInfoResponse"`
	Categories      []Category `xml:"CategoryArray>Category,omitempty"`
	CategoryCount   int        `xml:"CategoryCount,omitempty"`
	CategoryVersion string     `xml:"CategoryVersion,omitempty"`
	UpdateTime      string     `xml:"UpdateTime,omitempty"`
}
