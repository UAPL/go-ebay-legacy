package shopping

type ListingStatusCode string

const (
	ListingStatusActive    ListingStatusCode = "Active"
	ListingStatusCompleted ListingStatusCode = "Completed"
	ListingStatusEnded     ListingStatusCode = "Ended"

	ListingStatusCustomCode ListingStatusCode = "CustomCode"
)

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

type SiteCode string

const (
	SiteCodeAustralia     SiteCode = "Australia"
	SiteCodeAustria       SiteCode = "Austria"
	SiteCodeBelgiumDutch  SiteCode = "Belgium_Dutch"
	SiteCodeBelgiumFrench SiteCode = "Belgium_French"
	SiteCodeCanada        SiteCode = "Canada"
	SiteCodeCanadaFrench  SiteCode = "CanadaFrench"
	SiteCodeEbayMotors    SiteCode = "eBayMotors"
	SiteCodeFrance        SiteCode = "France"
	SiteCodeGermany       SiteCode = "Germany"
	SiteCodeIreland       SiteCode = "Ireland"
	SiteCodeItaly         SiteCode = "Italy"
	SiteCodeNetherlands   SiteCode = "Netherlands"
	SiteCodePhilippines   SiteCode = "Philippines"
	SiteCodePoland        SiteCode = "Poland"
	SiteCodeRussia        SiteCode = "Russia"
	SiteCodeSpain         SiteCode = "Spain"
	SiteCodeSwitzerland   SiteCode = "Switzerland"
	SiteCodeTaiwan        SiteCode = "Taiwan"
	SiteCodeUK            SiteCode = "UK"
	SiteCodeUS            SiteCode = "US"

	SiteCodeCustomCode SiteCode = "CustomCode"
)

type SellerBusinessCode string

const (
	SellerBusinessCodeCommercial SellerBusinessCode = "Commercial"
	SellerBusinessCodePrivate    SellerBusinessCode = "Private"
	SellerBusinessCodeUndefined  SellerBusinessCode = "Undefined"

	SellerBusinessCodeCustomCode SellerBusinessCode = "CustomCode"
)

type SellerLevelCode string

const (
	SellerLevelBronze   SellerLevelCode = "Bronze"
	SellerLevelDiamond  SellerLevelCode = "Diamond"
	SellerLevelGold     SellerLevelCode = "Gold"
	SellerLevelNone     SellerLevelCode = "None"
	SellerLevelPlatinum SellerLevelCode = "Platinum"
	SellerLevelSilver   SellerLevelCode = "Silver"
	SellerLevelTitanium SellerLevelCode = "Titanium"

	SellerLevelCustomCode SellerLevelCode = "CustomCode"
)

type UserStatusCode string

const (
	UserStatusAccountOnHold           UserStatusCode = "AccountOnHold"
	UserStatusConfirmed               UserStatusCode = "Confirmed"
	UserStatusCreditCardVerify        UserStatusCode = "CreditCardVerify"
	UserStatusDeleted                 UserStatusCode = "Deleted"
	UserStatusGhost                   UserStatusCode = "Ghost"
	UserStatusInMaintenance           UserStatusCode = "InMaintenance"
	UserStatusMerged                  UserStatusCode = "Merged"
	UserStatusRegistrationCodeMailOut UserStatusCode = "RegistrationCodeMailOut"
	UserStatusSuspended               UserStatusCode = "Suspended"
	UserStatusTermPending             UserStatusCode = "TermPending"
	UserStatusUnconfirmed             UserStatusCode = "Unconfirmed"
	UserStatusUnknown                 UserStatusCode = "Unknown"

	UserStatusCustomCode UserStatusCode = "CustomCode"
)

type CharityStatusCode string

const (
	CharityStatusCustomCode    CharityStatusCode = "CustomCode"
	CharityStatusNoLongerValid CharityStatusCode = "NoLongerValid"
	CharityStatusValid         CharityStatusCode = "Valid"
)

type MinimumAdvertisedPriceExposureCode string

const (
	MinimumAdvertisedPriceExposureDuringCheckout MinimumAdvertisedPriceExposureCode = "DuringCheckout"
	MinimumAdvertisedPriceExposureNone           MinimumAdvertisedPriceExposureCode = "None"
	MinimumAdvertisedPriceExposurePreCheckout    MinimumAdvertisedPriceExposureCode = "PreCheckout"

	MinimumAdvertisedPriceExposureCustomCode MinimumAdvertisedPriceExposureCode = "CustomCode"
)

type PricingTreatmentCode string

const (
	PricingTreatmentMAP  PricingTreatmentCode = "MAP"
	PricingTreatmentNone PricingTreatmentCode = "None"
	PricingTreatmentSTP  PricingTreatmentCode = "STP"

	PricingTreatmentCustomCode PricingTreatmentCode = "CustomCode"
)

type ListingTypeCode string

const (
	ListingTypeAdType         ListingTypeCode = "AdType"
	ListingTypeChinese        ListingTypeCode = "Chinese"
	ListingTypeFixedPriceItem ListingTypeCode = "FixedPriceItem"
	ListingTypeLeadGeneration ListingTypeCode = "LeadGeneration"
	ListingTypePersonalOffer  ListingTypeCode = "PersonalOffer"

	ListingTypeCustomCode ListingTypeCode = "CustomCode"
)

type BuyerPaymentMethodCode string

const (
	BuyerPaymentMethodAmEx                        BuyerPaymentMethodCode = "AmEx"
	BuyerPaymentMethodCashOnPickup                BuyerPaymentMethodCode = "CashOnPickup"
	BuyerPaymentMethodCCAccepted                  BuyerPaymentMethodCode = "CCAccepted"
	BuyerPaymentMethodCOD                         BuyerPaymentMethodCode = "COD"
	BuyerPaymentMethodCreditCard                  BuyerPaymentMethodCode = "CreditCard"
	BuyerPaymentMethodDirectDebit                 BuyerPaymentMethodCode = "DirectDebit"
	BuyerPaymentMethodDiscover                    BuyerPaymentMethodCode = "Discover"
	BuyerPaymentMethodELV                         BuyerPaymentMethodCode = "ELV"
	BuyerPaymentMethodLoanCheck                   BuyerPaymentMethodCode = "LoanCheck"
	BuyerPaymentMethodMOCC                        BuyerPaymentMethodCode = "MOCC"
	BuyerPaymentMethodMoneyXferAccepted           BuyerPaymentMethodCode = "MoneyXferAccepted"
	BuyerPaymentMethodMoneyXferAcceptedInCheckout BuyerPaymentMethodCode = "MoneyXferAcceptedInCheckout"
	BuyerPaymentMethodNone                        BuyerPaymentMethodCode = "None"
	BuyerPaymentMethodOther                       BuyerPaymentMethodCode = "Other"
	BuyerPaymentMethodOtherOnlinePayments         BuyerPaymentMethodCode = "OtherOnlinePayments"
	BuyerPaymentMethodPaymentSeeDescription       BuyerPaymentMethodCode = "PaymentSeeDescription"
	BuyerPaymentMethodPayPal                      BuyerPaymentMethodCode = "PayPal"
	BuyerPaymentMethodPersonalCheck               BuyerPaymentMethodCode = "PersonalCheck"
	BuyerPaymentMethodCodeVisa                    BuyerPaymentMethodCode = "Visa"

	BuyerPaymentMethodCustomCode BuyerPaymentMethodCode = "CustomCode"
)

type QuantityAvailableHintCode string

const (
	QuantityAvailableHintLimited    QuantityAvailableHintCode = "Limited"
	QuantityAvailableHintMoreThan   QuantityAvailableHintCode = "MoreThan"
	QuantityAvailableHintCustomCode QuantityAvailableHintCode = "CustomCode"
)

type ShippingTypeCode string

const (
	ShippingTypeCalculated                          ShippingTypeCode = "Calculated"
	ShippingTypeCalculatedDomesticFlatInternational ShippingTypeCode = "CalculatedDomesticFlatInternational"
	ShippingTypeFlat                                ShippingTypeCode = "Flat"
	ShippingTypeFlatDomesticCalculatedInternational ShippingTypeCode = "FlatDomesticCalculatedInternational"
	ShippingTypeFreight                             ShippingTypeCode = "Freight"
	ShippingTypeNotSpecified                        ShippingTypeCode = "NotSpecified"

	ShippingTypeCustomCode ShippingTypeCode = "CustomCode"
)

type CommentTypeCode string

const (
	CommentTypeIndependentlyWithdrawn CommentTypeCode = "IndependentlyWithdrawn"
	CommentTypeNegative               CommentTypeCode = "Negative"
	CommentTypeNeutral                CommentTypeCode = "Neutral"
	CommentTypePositive               CommentTypeCode = "Positive"
	CommentTypeWithdrawn              CommentTypeCode = "Withdrawn"

	CommentTypeCustomCode CommentTypeCode = "CustomCode"
)

type TradingRoleCode string

const (
	TradingRoleBuyer  = "Buyer"
	TradingRoleSeller = "Seller"

	TradingRoleCustomCode = "CustomCode"
)

type FeedbackRatingDetailCode string

const (
	FeedbackRatingDetailCodeCommunication              FeedbackRatingDetailCode = "Communication"
	FeedbackRatingDetailCodeItemAsDescribed            FeedbackRatingDetailCode = "ItemAsDescribed"
	FeedbackRatingDetailCodeShippingAndHandlingCharges FeedbackRatingDetailCode = "ShippingAndHandlingCharges"
	FeedbackRatingDetailCodeShippingTime               FeedbackRatingDetailCode = "ShippingTime"

	FeedbackRatingDetailCodeCustomCode FeedbackRatingDetailCode = "CustomCode"
)
