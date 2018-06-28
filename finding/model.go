package finding

import "github.com/datainq/xml-date-time"

type ImageSize string

const (
	LargeImage  ImageSize = "Large"
	MediumImage ImageSize = "Medium"
	SmallImage  ImageSize = "Small"
)

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
	ItemId          string                 `xml:"itemId"`
	Title           string                 `xml:"title"`
	PrimaryCategory CategorySummary        `xml:"primaryCategory"`
	ImageUrl        string                 `xml:"galleryURL"`
	GalleryUrls     []GalleryUrl           `xml:"galleryInfoContainer>galleryURL"`
	ListingUrl      string                 `xml:"viewItemURL"`
	Location        string                 `xml:"location"`
	Seller          Seller                 `xml:"sellerInfo"`
	CurrentPrice    float64                `xml:"sellingStatus>currentPrice"`
	ShippingPrice   float64                `xml:"shippingInfo>shippingServiceCost"`
	BinPrice        float64                `xml:"listingInfo>buyItNowPrice"`
	ShipsTo         []string               `xml:"shippingInfo>shipToLocations"`
	Site            string                 `xml:"globalId"`
	StartTime       xmldatetime.CustomTime `xml:"listingInfo>startTime"`
	EndTime         xmldatetime.CustomTime `xml:"listingInfo>endTime"`
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
	CustomId   string `xml:"customId"`
	NetworkId  string `xml:"networkId"`
	TrackingId string `xml:"trackingId"`
}
