package shopping

import "time"

type NameValueList struct {
	Name  string   `xml:"Name"`
	Value []string `xml:"Value"`
}

type Seller struct {
	UserID                  string  `xml:"UserID"`
	FeedbackRatingStar      string  `xml:"FeedbackRatingStar"`
	FeedbackScore           int     `xml:"FeedbackScore"`
	PositiveFeedbackPercent float64 `xml:"PositiveFeedbackPercent"`
	TopRatedSeller          bool    `xml:"TopRatedSeller"`
}

type Item struct {
	ItemID                  string              `xml:"ItemID"`
	Title                   string              `xml:"Title"`
	Subtitle                string              `xml:"Subtitle"`
	Description             string              `xml:"Description"`
	ItemUrl                 string              `xml:"ViewItemURLForNaturalSearch"`
	GalleryURL              string              `xml:"GalleryURL"`
	PictureURLs             []string            `xml:"PictureURL"`
	PrimaryCategoryID       string              `xml:"PrimaryCategoryID"`
	PrimaryCategoryName     string              `xml:"PrimaryCategoryName"`
	PrimaryCategoryIDPath   string              `xml:"PrimaryCategoryIDPath"`
	SecondaryCategoryID     string              `xml:"SecondaryCategoryID"`
	SecondaryCategoryName   string              `xml:"SecondaryCategoryName"`
	SecondaryCategoryIDPath string              `xml:"SecondaryCategoryIDPath"`
	CurrentPrice            float64             `xml:"CurrentPrice"`
	BestOfferEnabled        bool                `xml:"BestOfferEnabled"`
	EndTime                 time.Time           `xml:"EndTime"`
	StartTime               time.Time           `xml:"StartTime"`
	ListingType             string              `xml:"ListingType"`
	Location                string              `xml:"Location"`
	PostalCode              string              `xml:"PostalCode"`
	Quantity                int                 `xml:"Quantity"`
	Seller                  Seller              `xml:"Seller"`
	ListingStatus           string              `xml:"ListingStatus"`
	HitCount                int                 `xml:"HitCount"`
	ConditionDisplayName    string              `xml:"ConditionDisplayName"`
	ConditionDescription    string              `xml:"ConditionDescription"`
	SKU                     string              `xml:"SKU"`
	ItemSpecifics           []NameValueList     `xml:"ItemSpecifics>NameValueList"`
	ItemCompatibilities     []ItemCompatibility `xml:"ItemCompatibilityList>Compatibility"`
}

type ItemCompatibility struct {
	Compatibility []NameValueList `xml:"NameValueList"`
	Notes         string          `xml:"CompatibilityNotes"`
}

type Category struct {
	Id           string `xml:"CategoryID"`
	IdPath       string `xml:"CategoryIDPath"`
	Name         string `xml:"CategoryName"`
	NamePath     string `xml:"CategoryNamePath"`
	Level        uint   `xml:"CategoryLevel"`
	LeafCategory bool   `xml:"LeafCategory"`
	ParentID     string `xml:"CategoryParentID"`
}
