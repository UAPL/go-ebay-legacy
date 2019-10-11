package finding

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
	Page         int `xml:"pageNumber,omitempty"`
	PerPage      int `xml:"entriesPerPage,omitempty"`
	TotalPages   int `xml:"totalPages,omitempty"`
	TotalEntries int `xml:"totalEntries,omitempty"`
}

type Error struct {
	ErrorId   string `xml:"errorId,omitempty"`
	Domain    string `xml:"domain,omitempty"`
	Severity  string `xml:"severity,omitempty"`
	Category  string `xml:"category,omitempty"`
	Message   string `xml:"message,omitempty"`
	SubDomain string `xml:"subdomain,omitempty"`
}

type BaseRequest struct {
	Affiliate  Affiliate  `xml:"affiliate,omitempty"`
	Pagination Pagination `xml:"paginationInput,omitempty"`
	SortOrder  string     `xml:"sortOrder,omitempty"`
}

type BaseResponse struct {
	Ack        Ack        `xml:"ack,omitempty"`
	Errors     []Error    `xml:"errorMessage>error,omitempty"`
	Pagination Pagination `xml:"paginationOutput,omitempty"`
	Timestamp  EbayTime   `xml:"timestamp,omitempty"`
	Version    string     `xml:"version,omitempty"`
}

type FindItemsAdvancedRequest struct {
	*BaseRequest

	XMLName           xml.Name            `xml:"http://www.ebay.com/marketplace/search/v1/services findItemsAdvancedRequest"`
	ItemFilters		  []ItemFilterInput   `xml:"itemFilter,omitempty"`
	AspectFilters     []AspectFilterInput `xml:"aspectFilter,omitempty"`
	Keywords          string              `xml:"keywords,omitempty"`
	Categories        []int               `xml:"categoryId,omitempty"`
	DescriptionSearch bool                `xml:"descriptionSearch,omitempty"`
	BuyerPostalCode   string              `xml:"buyerPostalCode,omitempty"`
	OutputSelectors   []string            `xml:"outputSelector,omitempty"`
}

type FindItemsAdvancedResponse struct {
	*BaseResponse

	XMLName            xml.Name             `xml:"findItemsAdvancedResponse" json:"ignore"`
	Items              []Item               `xml:"searchResult>item"`
	Aspects            []Aspect             `xml:"aspectHistogramContainer>aspect"`
	CategoryHistogram  []CategoryHistogram  `xml:"categoryHistogramContainer>categoryHistogram"`
	ConditionHistogram []ConditionHistogram `xml:"conditionHistogramContainer>conditionHistogram"`
}

type FindItemsByKeywordResponse struct {
	*BaseResponse
	XMLName xml.Name `xml:"findItemsByKeywordsResponse" json:"ignore"`
	Items   []Item   `xml:"searchResult>item"`
}
