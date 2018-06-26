package finding

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

type Pagination struct {
	Page         uint16 `xml:"pageNumber"`
	PerPage      uint16 `xml:"entriesPerPage"`
	TotalPages   uint16 `xml:"totalPages"`
	TotalEntries uint16 `xml:"totalEntries"`
}

type Error struct {
	ErrorId   string `xml:"errorId"`
	Domain    string `xml:"domain"`
	Severity  string `xml:"severity"`
	Category  string `xml:"category"`
	Message   string `xml:"message"`
	SubDomain string `xml:"subdomain"`
}

type BaseRequest struct {
	Affiliate  Affiliate  `xml:"affiliate"`
	Pagination Pagination `xml:"paginationInput"`
	SortOrder  string     `xml:"sortOrder"`
}

type BaseResponse struct {
	Ack        Ack                    `xml:"ack"`
	Errors     []Error                `xml:"errorMessage>error"`
	Pagination Pagination             `xml:"paginationOutput"`
	Timestamp  xmldatetime.CustomTime `xml:"timestamp"`
	Version    string                 `xml:"version"`
}

type FindItemsAdvancedRequest struct {
	*BaseRequest
	XmlName           xml.Name            `xml:"findItemsAdvancedRequest"`
	AspectFilters     []AspectFilterInput `xml:"aspectFilter"`
	Keywords          string              `xml:"keywords"`
	Categories        []int               `xml:"categoryId"`
	DescriptionSearch bool                `xml:"descriptionSearch"`
	BuyerPostalCode   string              `xml:"buyerPostalCode"`
	OutputSelectors   []string            `xml:"outputSelector"`
}

type FindItemsAdvancedResponse struct {
	*BaseResponse
	XmlName            xml.Name             `xml:"findItemsAdvancedResponse"`
	Items              []Item               `xml:"searchResult>item"`
	Aspects            []Aspect             `xml:"aspectHistogramContainer>aspect"`
	CategoryHistogram  []CategoryHistogram  `xml:"categoryHistogramContainer>categoryHistogram"`
	ConditionHistogram []ConditionHistogram `xml:"conditionHistogramContainer>conditionHistogram"`
}

type FindItemsByKeywordResponse struct {
	*BaseResponse
	XmlName xml.Name `xml:"findItemsByKeywordsResponse"`
	Items   []Item   `xml:"searchResult>item"`
}
