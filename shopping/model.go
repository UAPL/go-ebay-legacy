package shopping

import "github.com/datainq/xml-date-time"

type NameValueList struct {
	Name  string   `xml:"Name"`
	Value []string `xml:"Value"`
}

type ItemDetail struct {
	ItemId               string                 `xml:"ItemId"`
	Title                string                 `xml:"Title"`
	Subtitle             string                 `xml:"Subtitle"`
	Description          string                 `xml:"Description"`
	ItemUrl              string                 `xml:"ViewItemURLForNaturalSearch"`
	StartTime            xmldatetime.CustomTime `xml:"StartTime"`
	EndTime              xmldatetime.CustomTime `xml:"EndTime"`
	Location             string                 `xml:"Location"`
	BestOffer            bool                   `xml:"BestOfferEnabled"`
	ListingType          string                 `xml:"ListingType"`
	PictureUrls          []string               `xml:"PictureURL"`
	PostalCode           string                 `xml:"PostalCode"`
	PrimaryCategory      string                 `xml:"PrimaryCategoryID"`
	SecondaryCategory    string                 `xml:"SecondaryCategoryID"`
	CategoryName         string                 `xml:"PrimaryCategoryName"`
	CurrentPrice         float32                `xml:"CurrentPrice"`
	Sku                  string                 `xml:"SKU"`
	ConditionDisplayName string                 `xml:"ConditionDisplayName"`
	ConditionId          string                 `xml:"ConditionID"`
	ItemSpecifics        []NameValueList        `xml:"ItemSpecifics>NameValueList"`
}

type Category struct {
	Id           string `xml:"CategoryID"`
	IdPath       string `xml:"CategoryIDPath"`
	Name         string `xml:"CategoryName"`
	NamePath     string `xml:"CategoryNamePath"`
	Level        int    `xml:"CategoryLevel"`
	LeafCategory bool   `xml:"LeafCategory"`
}
