package finding

import (
	"encoding/xml"
	"time"
)

type ImageSize string

const (
	LargeImage  ImageSize = "Large"
	MediumImage ImageSize = "Medium"
	SmallImage  ImageSize = "Small"
)

type EbayTime struct {
	*time.Time
}

func (c EbayTime) Parse(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

func (c *EbayTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	c.Time = &t
	return nil
}

func (c *EbayTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(c.Time.Format(time.RFC3339Nano), start)
}

type Seller struct {
	Username                string  `xml:"sellerUserName"`
	FeedbackScore           uint32  `xml:"feedbackScore"`
	PositiveFeedbackPercent float32 `xml:"positiveFeedbackPercent"`
	TopRated                bool    `xml:"topRatedSeller"`
}

type GalleryUrl struct {
	Size ImageSize `xml:"gallerySize,attr"`
	Url  string    `xml:",chardata"`
}

type Item struct {
	ItemId          string          `xml:"itemId"`
	Title           string          `xml:"title"`
	PrimaryCategory CategorySummary `xml:"primaryCategory"`
	ImageUrl        string          `xml:"galleryURL"`
	GalleryUrls     []GalleryUrl    `xml:"galleryInfoContainer>galleryURL"`
	ListingUrl      string          `xml:"viewItemURL"`
	Location        string          `xml:"location"`
	Seller          Seller          `xml:"sellerInfo"`
	CurrentPrice    float64         `xml:"sellingStatus>currentPrice"`
	ShippingPrice   float64         `xml:"shippingInfo>shippingServiceCost"`
	BinPrice        float64         `xml:"listingInfo>buyItNowPrice"`
	ShipsTo         []string        `xml:"shippingInfo>shipToLocations"`
	Site            string          `xml:"globalId"`
	StartTime       EbayTime        `xml:"listingInfo>startTime"`
	EndTime         EbayTime        `xml:"listingInfo>endTime"`
}

type CategorySummary struct {
	Name string `xml:"categoryName"`
	Id   int    `xml:"categoryId"`
}

type CategoryHistogram struct {
	CategoryId   int                 `xml:"categoryId"`
	CategoryName string              `xml:"categoryName"`
	Count        int                 `xml:"count"`
	Children     []CategoryHistogram `xml:"childCategoryHistogram"`
}

type ConditionHistogram struct {
	ConditionName string `xml:"condition>conditionDisplayName"`
	ConditionId   string `xml:"condition>conditionId"`
	Count         int    `xml:"count"`
}

type Aspect struct {
	Name   string        `xml:"name,attr"`
	Values []AspectValue `xml:"valueHistogram"`
}

type AspectValue struct {
	Value string `xml:"valueName,attr"`
	Count int    `xml:"xml:count"`
}

type AspectFilterInput struct {
	Name   string   `xml:"aspectName"`
	Values []string `xml:"aspectValueName"`
}

type ItemFilterInput struct {
	Name   string   `xml:"name"`
	Values []string `xml:"value"`
}

type Affiliate struct {
	CustomId   string `xml:"customId,omitempty"`
	NetworkId  string `xml:"networkId,omitempty"`
	TrackingId string `xml:"trackingId,omitempty"`
}
