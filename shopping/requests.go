package shopping

import (
	"encoding/xml"
	"errors"
	"github.com/datainq/xml-date-time"
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
	ErrorCode           string           `xml:"ErrorCode"`
	ShortMessage        string           `xml:"ShortMessage"`
	LongMessage         string           `xml:"LongMessage"`
	SeverityCode        string           `xml:"SeverityCode"`
	ErrorParameters     []ErrorParameter `xml:"ErrorParameters"`
	ErrorClassification string           `xml:"ErrorClassification"`
}

type BaseShoppingResponse struct {
	Timestamp     xmldatetime.CustomTime `xml:"Timestamp"`
	Ack           Ack                    `xml:"Ack"`
	Build         string                 `xml:"Build"`
	CorrelationID string                 `xml:"CorrelationID"`
}

type GetMultipleItemsRequest struct {
	XmlName         xml.Name `xml:"GetMultipleItemsRequest"`
	IncludeSelector string   `xml:"IncludeSelector"`
	ItemIds         []string `xml:"ItemID"`
}

type GetMultipleItemsResponse struct {
	*BaseShoppingResponse
	XmlName xml.Name     `xml:"GetMultipleItemsResponse"`
	Items   []ItemDetail `xml:"Item"`
}

type GetSingleItemRequest struct {
	XmlName         xml.Name `xml:"GetSingleItemRequest"`
	IncludeSelector string   `xml:"IncludeSelector"`
	ItemId          string   `xml:"ItemID"`
	MessageID       string   `xml:"MessageID"`
}

type GetSingleItemResponse struct {
	*BaseShoppingResponse
	XmlName xml.Name `xml:"GetSingleItemResponse"`
	Items   []ItemDetail
}

type GetCategoryInfoRequest struct {
	XmlName         xml.Name `xml:"GetCategoryInfoRequest"`
	CategoryID      string   `xml:"CategoryID"`
	IncludeSelector string   `xml:"IncludeSelector"`
	MessageID       string   `xml:"MessageID"`
}

type GetCategoryInfoResponse struct {
	*BaseShoppingResponse
	XmlName         xml.Name               `xml:"GetCategoryInfoResponse"`
	Categories      []Category             `xml:"CategoryArray>Category"`
	CategoryCount   int                    `xml:"CategoryCount"`
	CategoryVersion string                 `xml:"CategoryVersion"`
	UpdateTime      xmldatetime.CustomTime `xml:"UpdateTime"`
}
