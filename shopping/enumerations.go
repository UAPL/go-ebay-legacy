package shopping

import (
	"encoding/xml"
	"errors"
)

type Enum interface {
	String() string
	UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
}

var _ Enum = &ListingStatus{}
var _ Enum = &FeedbackRatingStar{}
var _ Enum = &SiteCode{}
var _ Enum = &SellerBusinessCode{}
var _ Enum = &SellerLevelCode{}
var _ Enum = &UserStatusCode{}
var _ Enum = &CharityStatusCode{}
var _ Enum = &MAPExposureCode{}
var _ Enum = &PricingTreatmentCode{}
var _ Enum = &ListingTypeCode{}
var _ Enum = &BuyerPaymentMethodCode{}
var _ Enum = &QuantityAvailableHintCode{}
var _ Enum = &ShippingTypeCode{}
var _ Enum = &CommentTypeCode{}
var _ Enum = &TradingRoleCode{}
var _ Enum = &FeedbackRatingDetailCode{}

type ListingStatusCode string

const (
	ListingStatusActive    ListingStatusCode = "Active"
	ListingStatusCompleted ListingStatusCode = "Completed"
	ListingStatusEnded     ListingStatusCode = "Ended"

	ListingStatusCustomCode ListingStatusCode = "CustomCode"
)

type ListingStatus struct {
	Value ListingStatusCode `xml:",chardata"`
}

func (c *ListingStatus) String() string {
	return string(c.Value)
}

func (c *ListingStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "Active":
		c.Value = ListingStatusActive
	case "Completed":
		c.Value = ListingStatusCompleted
	case "Ended":
		c.Value = ListingStatusEnded
	case "CustomCode":
		c.Value = ListingStatusCustomCode

	default:
		return errors.New("invalid ListingStatus code received")
	}

	return nil
}

type FeedbackRatingStarCode string

const (
	FeedbackRatingNone      FeedbackRatingStarCode = "None"
	FeedbackRatingYellow    FeedbackRatingStarCode = "Yellow"
	FeedbackRatingBlue      FeedbackRatingStarCode = "Blue"
	FeedbackRatingTurquoise FeedbackRatingStarCode = "Turquoise"
	FeedbackRatingPurple    FeedbackRatingStarCode = "Purple"
	FeedbackRatingRed       FeedbackRatingStarCode = "Red"
	FeedbackRatingGreen     FeedbackRatingStarCode = "Green"

	FeedbackRatingYellowShooting    FeedbackRatingStarCode = "YellowShooting"
	FeedbackRatingTurquoiseShooting FeedbackRatingStarCode = "TurquoiseShooting"
	FeedbackRatingPurpleShooting    FeedbackRatingStarCode = "PurpleShooting"
	FeedbackRatingRedShooting       FeedbackRatingStarCode = "RedShooting"
	FeedbackRatingGreenShooting     FeedbackRatingStarCode = "GreenShooting"
	FeedbackRatingSilverShooting    FeedbackRatingStarCode = "SilverShooting"

	FeedbackRatingCustomCode FeedbackRatingStarCode = "CustomCode"
)

type FeedbackRatingStar struct {
	Value FeedbackRatingStarCode `xml:",chardata"`
}

func (c *FeedbackRatingStar) String() string {
	return string(c.Value)
}

func (c *FeedbackRatingStar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "None":
		c.Value = FeedbackRatingNone
	case "Yellow":
		c.Value = FeedbackRatingYellow
	case "Blue":
		c.Value = FeedbackRatingBlue
	case "Turquoise":
		c.Value = FeedbackRatingTurquoise
	case "Purple":
		c.Value = FeedbackRatingPurple
	case "Red":
		c.Value = FeedbackRatingRed
	case "Green":
		c.Value = FeedbackRatingGreen
	case "YellowShooting":
		c.Value = FeedbackRatingYellowShooting
	case "TurquoiseShooting":
		c.Value = FeedbackRatingTurquoiseShooting
	case "PurpleShooting":
		c.Value = FeedbackRatingPurpleShooting
	case "RedShooting":
		c.Value = FeedbackRatingRedShooting
	case "GreenShooting":
		c.Value = FeedbackRatingGreenShooting
	case "SilverShooting":
		c.Value = FeedbackRatingSilverShooting
	case "CustomCode":
		c.Value = FeedbackRatingCustomCode
	default:
		return errors.New("invalid FeedbackRatingStar code received")
	}

	return nil
}

type SiteCodeType string

const (
	SiteCodeAustralia     SiteCodeType = "Australia"
	SiteCodeAustria       SiteCodeType = "Austria"
	SiteCodeBelgiumDutch  SiteCodeType = "Belgium_Dutch"
	SiteCodeBelgiumFrench SiteCodeType = "Belgium_French"
	SiteCodeCanada        SiteCodeType = "Canada"
	SiteCodeCanadaFrench  SiteCodeType = "CanadaFrench"
	SiteCodeEbayMotors    SiteCodeType = "eBayMotors"
	SiteCodeFrance        SiteCodeType = "France"
	SiteCodeGermany       SiteCodeType = "Germany"
	SiteCodeIreland       SiteCodeType = "Ireland"
	SiteCodeItaly         SiteCodeType = "Italy"
	SiteCodeNetherlands   SiteCodeType = "Netherlands"
	SiteCodePhilippines   SiteCodeType = "Philippines"
	SiteCodePoland        SiteCodeType = "Poland"
	SiteCodeRussia        SiteCodeType = "Russia"
	SiteCodeSpain         SiteCodeType = "Spain"
	SiteCodeSwitzerland   SiteCodeType = "Switzerland"
	SiteCodeTaiwan        SiteCodeType = "Taiwan"
	SiteCodeUK            SiteCodeType = "UK"
	SiteCodeUS            SiteCodeType = "US"

	SiteCodeCustomCode SiteCodeType = "CustomCode"
)

type SiteCode struct {
	Value SiteCodeType `xml:",chardata"`
}

func (c *SiteCode) String() string {
	return string(c.Value)
}

func (c *SiteCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "Australia":
		c.Value = SiteCodeAustralia
	case "Austria":
		c.Value = SiteCodeAustria
	case "Belgium_Dutch":
		c.Value = SiteCodeBelgiumDutch
	case "Belgium_French":
		c.Value = SiteCodeBelgiumFrench
	case "Canada":
		c.Value = SiteCodeCanada
	case "CanadaFrench":
		c.Value = SiteCodeCanadaFrench
	case "eBayMotors":
		c.Value = SiteCodeEbayMotors
	case "France":
		c.Value = SiteCodeFrance
	case "Germany":
		c.Value = SiteCodeGermany
	case "Ireland":
		c.Value = SiteCodeIreland
	case "Italy":
		c.Value = SiteCodeItaly
	case "Netherlands":
		c.Value = SiteCodeNetherlands
	case "Philippines":
		c.Value = SiteCodePhilippines
	case "Poland":
		c.Value = SiteCodePoland
	case "Russia":
		c.Value = SiteCodeRussia
	case "Spain":
		c.Value = SiteCodeSpain
	case "Switzerland":
		c.Value = SiteCodeSwitzerland
	case "Taiwan":
		c.Value = SiteCodeTaiwan
	case "UK":
		c.Value = SiteCodeUK
	case "US":
		c.Value = SiteCodeUS
	case "CustomCode":
		c.Value = SiteCodeCustomCode
	default:
		return errors.New("invalid Site code received")
	}

	return nil
}

type SellerBusinessCodeType string

const (
	SellerBusinessCodeCommercial SellerBusinessCodeType = "Commercial"
	SellerBusinessCodePrivate    SellerBusinessCodeType = "Private"
	SellerBusinessCodeUndefined  SellerBusinessCodeType = "Undefined"

	SellerBusinessCodeCustomCode SellerBusinessCodeType = "CustomCode"
)

type SellerBusinessCode struct {
	Value SellerBusinessCodeType `xml:",chardata"`
}

func (c *SellerBusinessCode) String() string {
	return string(c.Value)
}

func (c *SellerBusinessCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "Commercial":
		c.Value = SellerBusinessCodeCommercial

	case "Private":
		c.Value = SellerBusinessCodePrivate

	case "Undefined":
		c.Value = SellerBusinessCodeUndefined

	case "CustomCode":
		c.Value = SellerBusinessCodeCustomCode

	default:
		return errors.New("invalid Site code received")
	}

	return nil
}

type SellerLevelCodeType string

const (
	SellerLevelBronze   SellerLevelCodeType = "Bronze"
	SellerLevelDiamond  SellerLevelCodeType = "Diamond"
	SellerLevelGold     SellerLevelCodeType = "Gold"
	SellerLevelNone     SellerLevelCodeType = "None"
	SellerLevelPlatinum SellerLevelCodeType = "Platinum"
	SellerLevelSilver   SellerLevelCodeType = "Silver"
	SellerLevelTitanium SellerLevelCodeType = "Titanium"

	SellerLevelCustomCode SellerLevelCodeType = "CustomCode"
)

type SellerLevelCode struct {
	Value SellerLevelCodeType `xml:",chardata"`
}

func (c *SellerLevelCode) String() string {
	return string(c.Value)
}

func (c *SellerLevelCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "Bronze":
		c.Value = SellerLevelBronze
	case "Diamond":
		c.Value = SellerLevelDiamond
	case "Gold":
		c.Value = SellerLevelGold
	case "None":
		c.Value = SellerLevelNone
	case "Platinum":
		c.Value = SellerLevelPlatinum
	case "Silver":
		c.Value = SellerLevelSilver
	case "Titanium":
		c.Value = SellerLevelTitanium
	case "CustomCode":
		c.Value = SellerLevelCustomCode

	default:
		return errors.New("invalid SellerLevel code received")
	}

	return nil
}

type UserStatusCodeType string

const (
	UserStatusAccountOnHold           UserStatusCodeType = "AccountOnHold"
	UserStatusConfirmed               UserStatusCodeType = "Confirmed"
	UserStatusCreditCardVerify        UserStatusCodeType = "CreditCardVerify"
	UserStatusDeleted                 UserStatusCodeType = "Deleted"
	UserStatusGhost                   UserStatusCodeType = "Ghost"
	UserStatusInMaintenance           UserStatusCodeType = "InMaintenance"
	UserStatusMerged                  UserStatusCodeType = "Merged"
	UserStatusRegistrationCodeMailOut UserStatusCodeType = "RegistrationCodeMailOut"
	UserStatusSuspended               UserStatusCodeType = "Suspended"
	UserStatusTermPending             UserStatusCodeType = "TermPending"
	UserStatusUnconfirmed             UserStatusCodeType = "Unconfirmed"
	UserStatusUnknown                 UserStatusCodeType = "Unknown"

	UserStatusCustomCode UserStatusCodeType = "CustomCode"
)

type UserStatusCode struct {
	Value UserStatusCodeType `xml:",chardata"`
}

func (c *UserStatusCode) String() string {
	return string(c.Value)
}

func (c *UserStatusCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "AccountOnHold":
		c.Value = UserStatusAccountOnHold
	case "Confirmed":
		c.Value = UserStatusConfirmed
	case "CreditCardVerify":
		c.Value = UserStatusCreditCardVerify
	case "Deleted":
		c.Value = UserStatusDeleted
	case "Ghost":
		c.Value = UserStatusGhost
	case "InMaintenance":
		c.Value = UserStatusInMaintenance
	case "Merged":
		c.Value = UserStatusMerged
	case "RegistrationCodeMailOut":
		c.Value = UserStatusRegistrationCodeMailOut
	case "Suspended":
		c.Value = UserStatusSuspended
	case "TermPending":
		c.Value = UserStatusTermPending
	case "Unconfirmed":
		c.Value = UserStatusUnconfirmed
	case "Unknown":
		c.Value = UserStatusUnknown

	case "CustomCode":
		c.Value = UserStatusCustomCode

	default:
		return errors.New("invalid ListingStatus code received")
	}

	return nil
}

type CharityStatusCodeType string

const (
	CharityStatusValid         CharityStatusCodeType = "Valid"
	CharityStatusNoLongerValid CharityStatusCodeType = "NoLongerValid"

	CharityStatusCustomCode CharityStatusCodeType = "CustomCode"
)

type CharityStatusCode struct {
	Value CharityStatusCodeType `xml:",chardata"`
}

func (c *CharityStatusCode) String() string {
	return string(c.Value)
}

func (c *CharityStatusCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "Valid":
		c.Value = CharityStatusValid
	case "NoLongerValid":
		c.Value = CharityStatusNoLongerValid

	case "CustomCode":
		c.Value = CharityStatusCustomCode

	default:
		return errors.New("invalid CharityStatusCode received")
	}

	return nil
}

type MAPExposureCodeType string

const (
	MAPExposureNone           MAPExposureCodeType = "None"
	MAPExposurePreCheckout    MAPExposureCodeType = "PreCheckout"
	MAPExposureDuringCheckout MAPExposureCodeType = "DuringCheckout"

	MAPExposureCustomCode MAPExposureCodeType = "CustomCode"
)

type MAPExposureCode struct {
	Value MAPExposureCodeType `xml:",chardata"`
}

func (c *MAPExposureCode) String() string {
	return string(c.Value)
}

func (c *MAPExposureCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "None":
		c.Value = MAPExposureNone
	case "PreCheckout":
		c.Value = MAPExposurePreCheckout
	case "DuringCheckout":
		c.Value = MAPExposureDuringCheckout

	case "CustomCode":
		c.Value = MAPExposureCustomCode

	default:
		return errors.New("invalid MAPExposureCode received")
	}

	return nil
}

type PricingTreatmentCodeType string

const (
	PricingTreatmentMAP  PricingTreatmentCodeType = "MAP"
	PricingTreatmentNone PricingTreatmentCodeType = "None"
	PricingTreatmentSTP  PricingTreatmentCodeType = "STP"

	PricingTreatmentCustomCode PricingTreatmentCodeType = "CustomCode"
)

type PricingTreatmentCode struct {
	Value PricingTreatmentCodeType `xml:",chardata"`
}

func (c *PricingTreatmentCode) String() string {
	return string(c.Value)
}

func (c *PricingTreatmentCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "MAP":
		c.Value = PricingTreatmentMAP
	case "None":
		c.Value = PricingTreatmentNone
	case "STP":
		c.Value = PricingTreatmentSTP
	case "CustomCode":
		c.Value = PricingTreatmentCustomCode

	default:
		return errors.New("invalid PricingTreatmentCode received")
	}

	return nil
}

type ListingTypeCodeType string

const (
	ListingTypeAdType         ListingTypeCodeType = "AdType"
	ListingTypeChinese        ListingTypeCodeType = "Chinese"
	ListingTypeFixedPriceItem ListingTypeCodeType = "FixedPriceItem"
	ListingTypeLeadGeneration ListingTypeCodeType = "LeadGeneration"
	ListingTypePersonalOffer  ListingTypeCodeType = "PersonalOffer"

	ListingTypeCustomCode ListingTypeCodeType = "CustomCode"
)

type ListingTypeCode struct {
	Value ListingTypeCodeType `xml:",chardata"`
}

func (c *ListingTypeCode) String() string {
	return string(c.Value)
}

func (c *ListingTypeCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "AdType":
		c.Value = ListingTypeAdType
	case "Chinese":
		c.Value = ListingTypeChinese
	case "FixedPriceItem":
		c.Value = ListingTypeFixedPriceItem
	case "LeadGeneration":
		c.Value = ListingTypeLeadGeneration
	case "PersonalOffer":
		c.Value = ListingTypePersonalOffer
	case "CustomCode":
		c.Value = ListingTypeCustomCode

	default:
		return errors.New("invalid ListingTypeCode received")
	}

	return nil
}

type BuyerPaymentMethodCodeType string

const (
	BuyerPaymentMethodAmEx                        BuyerPaymentMethodCodeType = "AmEx"
	BuyerPaymentMethodCashOnPickup                BuyerPaymentMethodCodeType = "CashOnPickup"
	BuyerPaymentMethodCCAccepted                  BuyerPaymentMethodCodeType = "CCAccepted"
	BuyerPaymentMethodCOD                         BuyerPaymentMethodCodeType = "COD"
	BuyerPaymentMethodCreditCard                  BuyerPaymentMethodCodeType = "CreditCard"
	BuyerPaymentMethodDirectDebit                 BuyerPaymentMethodCodeType = "DirectDebit"
	BuyerPaymentMethodDiscover                    BuyerPaymentMethodCodeType = "Discover"
	BuyerPaymentMethodELV                         BuyerPaymentMethodCodeType = "ELV"
	BuyerPaymentMethodLoanCheck                   BuyerPaymentMethodCodeType = "LoanCheck"
	BuyerPaymentMethodMOCC                        BuyerPaymentMethodCodeType = "MOCC"
	BuyerPaymentMethodMoneyXferAccepted           BuyerPaymentMethodCodeType = "MoneyXferAccepted"
	BuyerPaymentMethodMoneyXferAcceptedInCheckout BuyerPaymentMethodCodeType = "MoneyXferAcceptedInCheckout"
	BuyerPaymentMethodNone                        BuyerPaymentMethodCodeType = "None"
	BuyerPaymentMethodOther                       BuyerPaymentMethodCodeType = "Other"
	BuyerPaymentMethodOtherOnlinePayments         BuyerPaymentMethodCodeType = "OtherOnlinePayments"
	BuyerPaymentMethodPaymentSeeDescription       BuyerPaymentMethodCodeType = "PaymentSeeDescription"
	BuyerPaymentMethodPayPal                      BuyerPaymentMethodCodeType = "PayPal"
	BuyerPaymentMethodPersonalCheck               BuyerPaymentMethodCodeType = "PersonalCheck"
	BuyerPaymentMethodCodeVisa                    BuyerPaymentMethodCodeType = "Visa"

	BuyerPaymentMethodCustomCode BuyerPaymentMethodCodeType = "CustomCode"
)

type BuyerPaymentMethodCode struct {
	Value BuyerPaymentMethodCodeType `xml:",chardata"`
}

func (c *BuyerPaymentMethodCode) String() string {
	return string(c.Value)
}

func (c *BuyerPaymentMethodCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "AmEx":
		c.Value = BuyerPaymentMethodAmEx
	case "CashOnPickup":
		c.Value = BuyerPaymentMethodCashOnPickup
	case "CCAccepted":
		c.Value = BuyerPaymentMethodCCAccepted
	case "COD":
		c.Value = BuyerPaymentMethodCOD
	case "CreditCard":
		c.Value = BuyerPaymentMethodCreditCard
	case "DirectDebit":
		c.Value = BuyerPaymentMethodDirectDebit
	case "Discover":
		c.Value = BuyerPaymentMethodDiscover
	case "ELV":
		c.Value = BuyerPaymentMethodELV
	case "LoanCheck":
		c.Value = BuyerPaymentMethodLoanCheck
	case "MOCC":
		c.Value = BuyerPaymentMethodMOCC
	case "MoneyXferAccepted":
		c.Value = BuyerPaymentMethodMoneyXferAccepted
	case "MoneyXferAcceptedInCheckout":
		c.Value = BuyerPaymentMethodMoneyXferAcceptedInCheckout
	case "None":
		c.Value = BuyerPaymentMethodNone
	case "Other":
		c.Value = BuyerPaymentMethodOther
	case "OtherOnlinePayments":
		c.Value = BuyerPaymentMethodOtherOnlinePayments
	case "PaymentSeeDescription":
		c.Value = BuyerPaymentMethodPaymentSeeDescription
	case "PayPal":
		c.Value = BuyerPaymentMethodPayPal
	case "PersonalCheck":
		c.Value = BuyerPaymentMethodPersonalCheck
	case "Visa":
		c.Value = BuyerPaymentMethodCodeVisa
	default:
		return errors.New("invalid BuyerPaymentMethodCode received")
	}

	return nil
}

type QuantityAvailableHintCodeType string

const (
	QuantityAvailableHintLimited    QuantityAvailableHintCodeType = "Limited"
	QuantityAvailableHintMoreThan   QuantityAvailableHintCodeType = "MoreThan"
	QuantityAvailableHintCustomCode QuantityAvailableHintCodeType = "CustomCode"
)

type QuantityAvailableHintCode struct {
	Value QuantityAvailableHintCodeType `xml:",chardata"`
}

func (c *QuantityAvailableHintCode) String() string {
	return string(c.Value)
}

func (c *QuantityAvailableHintCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "Limited":
		c.Value = QuantityAvailableHintLimited
	case "MoreThan":
		c.Value = QuantityAvailableHintMoreThan
	case "CustomCode":
		c.Value = QuantityAvailableHintCustomCode

	default:
		return errors.New("invalid QuantityAvailableHintCode received")
	}

	return nil
}

type ShippingTypeCodeType string

const (
	ShippingTypeCalculated                          ShippingTypeCodeType = "Calculated"
	ShippingTypeCalculatedDomesticFlatInternational ShippingTypeCodeType = "CalculatedDomesticFlatInternational"
	ShippingTypeFlat                                ShippingTypeCodeType = "Flat"
	ShippingTypeFlatDomesticCalculatedInternational ShippingTypeCodeType = "FlatDomesticCalculatedInternational"
	ShippingTypeFreight                             ShippingTypeCodeType = "Freight"
	ShippingTypeNotSpecified                        ShippingTypeCodeType = "NotSpecified"

	ShippingTypeCustomCode ShippingTypeCodeType = "CustomCode"
)

type ShippingTypeCode struct {
	Value ShippingTypeCodeType `xml:",chardata"`
}

func (c *ShippingTypeCode) String() string {
	return string(c.Value)
}

func (c *ShippingTypeCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "Calculated":
		c.Value = ShippingTypeCalculated
	case "CalculatedDomesticFlatInternational":
		c.Value = ShippingTypeCalculatedDomesticFlatInternational
	case "Flat":
		c.Value = ShippingTypeFlat
	case "FlatDomesticCalculatedInternational":
		c.Value = ShippingTypeFlatDomesticCalculatedInternational
	case "Freight":
		c.Value = ShippingTypeFreight
	case "NotSpecified":
		c.Value = ShippingTypeNotSpecified
	case "CustomCode":
		c.Value = ShippingTypeCustomCode

	default:
		return errors.New("invalid ShippingTypeCode received")
	}

	return nil
}

type CommentTypeCodeType string

const (
	CommentTypeIndependentlyWithdrawn CommentTypeCodeType = "IndependentlyWithdrawn"
	CommentTypeNegative               CommentTypeCodeType = "Negative"
	CommentTypeNeutral                CommentTypeCodeType = "Neutral"
	CommentTypePositive               CommentTypeCodeType = "Positive"
	CommentTypeWithdrawn              CommentTypeCodeType = "Withdrawn"

	CommentTypeCustomCode CommentTypeCodeType = "CustomCode"
)

type CommentTypeCode struct {
	Value CommentTypeCodeType `xml:",chardata"`
}

func (c *CommentTypeCode) String() string {
	return string(c.Value)
}

func (c *CommentTypeCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "IndependentlyWithdrawn":
		c.Value = CommentTypeIndependentlyWithdrawn
	case "Negative":
		c.Value = CommentTypeNegative
	case "Neutral":
		c.Value = CommentTypeNeutral
	case "Positive":
		c.Value = CommentTypePositive
	case "Withdrawn":
		c.Value = CommentTypeWithdrawn

	case "CustomCode":
		c.Value = CommentTypeCustomCode

	default:
		return errors.New("invalid CommentTypeCode received")
	}

	return nil
}

type TradingRoleCodeType string

const (
	TradingRoleBuyer  TradingRoleCodeType = "Buyer"
	TradingRoleSeller TradingRoleCodeType = "Seller"

	TradingRoleCustomCode TradingRoleCodeType = "CustomCode"
)

type TradingRoleCode struct {
	Value TradingRoleCodeType `xml:",chardata"`
}

func (c *TradingRoleCode) String() string {
	return string(c.Value)
}

func (c *TradingRoleCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "Buyer":
		c.Value = TradingRoleBuyer
	case "Seller":
		c.Value = TradingRoleSeller
	case "CustomCode":
		c.Value = TradingRoleCustomCode
	default:
		return errors.New("invalid TradingRoleCode received")
	}

	return nil
}

type FeedbackRatingDetailCodeType string

const (
	FeedbackRatingDetailCodeCommunication              FeedbackRatingDetailCodeType = "Communication"
	FeedbackRatingDetailCodeItemAsDescribed            FeedbackRatingDetailCodeType = "ItemAsDescribed"
	FeedbackRatingDetailCodeShippingAndHandlingCharges FeedbackRatingDetailCodeType = "ShippingAndHandlingCharges"
	FeedbackRatingDetailCodeShippingTime               FeedbackRatingDetailCodeType = "ShippingTime"

	FeedbackRatingDetailCodeCustomCode FeedbackRatingDetailCodeType = "CustomCode"
)

type FeedbackRatingDetailCode struct {
	Value FeedbackRatingDetailCodeType `xml:",chardata"`
}

func (c *FeedbackRatingDetailCode) String() string {
	return string(c.Value)
}

func (c *FeedbackRatingDetailCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)

	switch v {
	case "Communication":
		c.Value = FeedbackRatingDetailCodeCommunication
	case "ItemAsDescribed":
		c.Value = FeedbackRatingDetailCodeItemAsDescribed
	case "ShippingAndHandlingCharges":
		c.Value = FeedbackRatingDetailCodeShippingAndHandlingCharges
	case "ShippingTime":
		c.Value = FeedbackRatingDetailCodeShippingTime
	case "CustomCode":
		c.Value = FeedbackRatingDetailCodeCustomCode
	default:
		return errors.New("invalid FeedbackRatingDetailCode received")
	}

	return nil
}
