package retail

type Pagination struct {
	Limit          int `json:"limit,omitempty"`
	TotalCount     int `json:"totalCount,omitempty"`
	CurrentPage    int `json:"currentPage,omitempty"`
	TotalPageCount int `json:"totalPageCount,omitempty"`
}

type Address struct {
	Index      string `json:"index,omitempty"`
	CountryIso string `json:"countryIso,omitempty"`
	Region     string `json:"region,omitempty"`
	RegionID   int    `json:"regionId,omitempty"`
	City       string `json:"city,omitempty"`
	CityID     int    `json:"cityId,omitempty"`
	CityType   string `json:"cityType,omitempty"`
	Street     string `json:"street,omitempty"`
	StreetID   int    `json:"streetId,omitempty"`
	StreetType string `json:"streetType,omitempty"`
	Building   string `json:"building,omitempty"`
	Flat       string `json:"flat,omitempty"`
	Floor      int    `json:"floor,omitempty"`
	Block      int    `json:"block,omitempty"`
	House      string `json:"house,omitempty"`
	Metro      string `json:"metro,omitempty"`
	Notes      string `json:"notes,omitempty"`
	Text       string `json:"text,omitempty"`
}

type GeoHierarchyRow struct {
	Country  string `json:"country,omitempty"`
	Region   string `json:"region,omitempty"`
	RegionID int    `json:"regionId,omitempty"`
	City     string `json:"city,omitempty"`
	CityID   int    `json:"cityId,omitempty"`
}

type Source struct {
	Source   string `json:"source,omitempty"`
	Medium   string `json:"medium,omitempty"`
	Campaign string `json:"campaign,omitempty"`
	Keyword  string `json:"keyword,omitempty"`
	Content  string `json:"content,omitempty"`
}

type Contragent struct {
	ContragentType    string `json:"contragentType,omitempty"`
	LegalName         string `json:"legalName,omitempty"`
	LegalAddress      string `json:"legalAddress,omitempty"`
	INN               string `json:"INN,omitempty"`
	OKPO              string `json:"OKPO,omitempty"`
	KPP               string `json:"KPP,omitempty"`
	OGRN              string `json:"OGRN,omitempty"`
	OGRNIP            string `json:"OGRNIP,omitempty"`
	CertificateNumber string `json:"certificateNumber,omitempty"`
	CertificateDate   string `json:"certificateDate,omitempty"`
	BIK               string `json:"BIK,omitempty"`
	Bank              string `json:"bank,omitempty"`
	BankAddress       string `json:"bankAddress,omitempty"`
	CorrAccount       string `json:"corrAccount,omitempty"`
	BankAccount       string `json:"bankAccount,omitempty"`
}

type Property struct {
	Code  string   `json:"code,omitempty"`
	Name  string   `json:"name,omitempty"`
	Value string   `json:"value,omitempty"`
	Sites []string `json:"Sites,omitempty"`
}

type Customer struct {
	ID                           int         `json:"id,omitempty"`
	ExternalID                   string      `json:"externalId,omitempty"`
	FirstName                    string      `json:"firstName,omitempty"`
	LastName                     string      `json:"lastName,omitempty"`
	Patronymic                   string      `json:"patronymic,omitempty"`
	Sex                          string      `json:"sex,omitempty"`
	Email                        string      `json:"email,omitempty"`
	Phones                       []Phone     `json:"phones,omitempty"`
	Address                      *Address    `json:"address,omitempty"`
	CreatedAt                    string      `json:"createdAt,omitempty"`
	Birthday                     string      `json:"birthday,omitempty"`
	ManagerID                    int         `json:"managerId,omitempty"`
	Vip                          bool        `json:"vip,omitempty"`
	Bad                          bool        `json:"bad,omitempty"`
	Site                         string      `json:"site,omitempty"`
	Source                       *Source     `json:"source,omitempty"`
	Contragent                   *Contragent `json:"contragent,omitempty"`
	PersonalDiscount             float32     `json:"personalDiscount,omitempty"`
	CumulativeDiscount           float32     `json:"cumulativeDiscount,omitempty"`
	DiscountCardNumber           string      `json:"discountCardNumber,omitempty"`
	EmailMarketingUnsubscribedAt string      `json:"emailMarketingUnsubscribedAt,omitempty"`
	AvgMarginSumm                float32     `json:"avgMarginSumm,omitempty"`
	MarginSumm                   float32     `json:"marginSumm,omitempty"`
	TotalSumm                    float32     `json:"totalSumm,omitempty"`
	AverageSumm                  float32     `json:"averageSumm,omitempty"`
	OrdersCount                  int         `json:"ordersCount,omitempty"`
	CostSumm                     float32     `json:"costSumm,omitempty"`
	MaturationTime               int         `json:"maturationTime,omitempty"`
	FirstClientID                string      `json:"firstClientId,omitempty"`
	LastClientID                 string      `json:"lastClientId,omitempty"`
	BrowserID                    string      `json:"browserId,omitempty"`
	MgCustomerID                 string      `json:"mgCustomerId,omitempty"`
	PhotoURL                     string      `json:"photoUrl,omitempty"`
	Tags                         []Tag       `json:"tags,omitempty"`
}

type Phone struct {
	Number string `json:"number,omitempty"`
}

type Order struct {
	ID                            int                     `json:"id,omitempty"`
	ExternalID                    string                  `json:"externalId,omitempty"`
	Number                        string                  `json:"number,omitempty"`
	FirstName                     string                  `json:"firstName,omitempty"`
	LastName                      string                  `json:"lastName,omitempty"`
	Patronymic                    string                  `json:"patronymic,omitempty"`
	Email                         string                  `json:"email,omitempty"`
	Phone                         string                  `json:"phone,omitempty"`
	AdditionalPhone               string                  `json:"additionalPhone,omitempty"`
	CreatedAt                     string                  `json:"createdAt,omitempty"`
	StatusUpdatedAt               string                  `json:"statusUpdatedAt,omitempty"`
	ManagerID                     int                     `json:"managerId,omitempty"`
	Mark                          int                     `json:"mark,omitempty"`
	Call                          bool                    `json:"call,omitempty"`
	Expired                       bool                    `json:"expired,omitempty"`
	FromAPI                       bool                    `json:"fromApi,omitempty"`
	MarkDatetime                  string                  `json:"markDatetime,omitempty"`
	CustomerComment               string                  `json:"customerComment,omitempty"`
	ManagerComment                string                  `json:"managerComment,omitempty"`
	Status                        string                  `json:"status,omitempty"`
	StatusComment                 string                  `json:"statusComment,omitempty"`
	FullPaidAt                    string                  `json:"fullPaidAt,omitempty"`
	Site                          string                  `json:"site,omitempty"`
	OrderType                     string                  `json:"orderType,omitempty"`
	OrderMethod                   string                  `json:"orderMethod,omitempty"`
	CountryIso                    string                  `json:"countryIso,omitempty"`
	Summ                          float32                 `json:"summ,omitempty"`
	TotalSumm                     float32                 `json:"totalSumm,omitempty"`
	PrepaySum                     float32                 `json:"prepaySum,omitempty"`
	PurchaseSumm                  float32                 `json:"purchaseSumm,omitempty"`
	DiscountManualAmount          float32                 `json:"discountManualAmount,omitempty"`
	DiscountManualPercent         float32                 `json:"discountManualPercent,omitempty"`
	Weight                        float32                 `json:"weight,omitempty"`
	Length                        int                     `json:"length,omitempty"`
	Width                         int                     `json:"width,omitempty"`
	Height                        int                     `json:"height,omitempty"`
	ShipmentStore                 string                  `json:"shipmentStore,omitempty"`
	ShipmentDate                  string                  `json:"shipmentDate,omitempty"`
	ClientID                      string                  `json:"clientId,omitempty"`
	Shipped                       bool                    `json:"shipped,omitempty"`
	UploadedToExternalStoreSystem bool                    `json:"uploadedToExternalStoreSystem,omitempty"`
	Source                        *Source                 `json:"source,omitempty"`
	Contragent                    *Contragent             `json:"contragent,omitempty"`
	Customer                      *Customer               `json:"customer,omitempty"`
	Delivery                      *OrderDelivery          `json:"delivery,omitempty"`
	Marketplace                   *OrderMarketplace       `json:"marketplace,omitempty"`
	Items                         []OrderItem             `json:"items,omitempty"`
	Payments                      map[string]OrderPayment `json:"payments,omitempty"`
}

type OrderDelivery struct {
	Code            string                `json:"code,omitempty"`
	IntegrationCode string                `json:"integrationCode,omitempty"`
	Cost            float32               `json:"cost,omitempty"`
	NetCost         float32               `json:"netCost,omitempty"`
	VatRate         string                `json:"vatRate,omitempty"`
	Date            string                `json:"date,omitempty"`
	Time            *OrderDeliveryTime    `json:"time,omitempty"`
	Address         *Address              `json:"address,omitempty"`
	Service         *OrderDeliveryService `json:"service,omitempty"`
	Data            *OrderDeliveryData    `json:"data,omitempty"`
}

type OrderDeliveryTime struct {
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Custom string `json:"custom,omitempty"`
}

type OrderDeliveryService struct {
	Name   string `json:"name,omitempty"`
	Code   string `json:"code,omitempty"`
	Active bool   `json:"active,omitempty"`
}

type OrderDeliveryDataBasic struct {
	TrackNumber        string `json:"trackNumber,omitempty"`
	Status             string `json:"status,omitempty"`
	PickuppointAddress string `json:"pickuppointAddress,omitempty"`
	PayerType          string `json:"payerType,omitempty"`
}

type OrderDeliveryData struct {
	OrderDeliveryDataBasic
	AdditionalFields map[string]interface{}
}

type OrderMarketplace struct {
	Code    string `json:"code,omitempty"`
	OrderID string `json:"orderId,omitempty"`
}

type OrderPayment struct {
	ID         int     `json:"id,omitempty"`
	ExternalID string  `json:"externalId,omitempty"`
	Type       string  `json:"type,omitempty"`
	Status     string  `json:"status,omitempty"`
	PaidAt     string  `json:"paidAt,omitempty"`
	Amount     float32 `json:"amount,omitempty"`
	Comment    string  `json:"comment,omitempty"`
}

type OrderItem struct {
	ID                    int                 `json:"id,omitempty"`
	InitialPrice          float32             `json:"initialPrice,omitempty"`
	PurchasePrice         float32             `json:"purchasePrice,omitempty"`
	DiscountTotal         float32             `json:"discountTotal,omitempty"`
	DiscountManualAmount  float32             `json:"discountManualAmount,omitempty"`
	DiscountManualPercent float32             `json:"discountManualPercent,omitempty"`
	ProductName           string              `json:"productName,omitempty"`
	VatRate               string              `json:"vatRate,omitempty"`
	CreatedAt             string              `json:"createdAt,omitempty"`
	Quantity              float32             `json:"quantity,omitempty"`
	Status                string              `json:"status,omitempty"`
	Comment               string              `json:"comment,omitempty"`
	IsCanceled            bool                `json:"isCanceled,omitempty"`
	Offer                 Offer               `json:"offer,omitempty"`
	Properties            map[string]Property `json:"properties,omitempty"`
	PriceType             *PriceType          `json:"priceType,omitempty"`
}

type Offer struct {
	ID            int               `json:"id,omitempty"`
	ExternalID    string            `json:"externalId,omitempty"`
	Name          string            `json:"name,omitempty"`
	XMLID         string            `json:"xmlId,omitempty"`
	Article       string            `json:"article,omitempty"`
	VatRate       string            `json:"vatRate,omitempty"`
	Price         float32           `json:"price,omitempty"`
	PurchasePrice float32           `json:"purchasePrice,omitempty"`
	Quantity      float32           `json:"quantity,omitempty"`
	Height        float32           `json:"height,omitempty"`
	Width         float32           `json:"width,omitempty"`
	Length        float32           `json:"length,omitempty"`
	Weight        float32           `json:"weight,omitempty"`
	Stores        []Inventory       `json:"stores,omitempty"`
	Properties    map[string]string `json:"properties,omitempty"`
	Prices        []OfferPrice      `json:"prices,omitempty"`
	Images        []string          `json:"images,omitempty"`
	Unit          *Unit             `json:"unit,omitempty"`
}

type Inventory struct {
	PurchasePrice float32 `json:"purchasePrice,omitempty"`
	Quantity      float32 `json:"quantity,omitempty"`
	Store         string  `json:"store,omitempty"`
}

type OfferPrice struct {
	Price     float32 `json:"price,omitempty"`
	Ordering  int     `json:"ordering,omitempty"`
	PriceType string  `json:"priceType,omitempty"`
}

type Unit struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Sym     string `json:"sym"`
	Default bool   `json:"default,omitempty"`
	Active  bool   `json:"active,omitempty"`
}

type PriceType struct {
	ID               int               `json:"id,omitempty"`
	Code             string            `json:"code,omitempty"`
	Name             string            `json:"name,omitempty"`
	Active           bool              `json:"active,omitempty"`
	Default          bool              `json:"default,omitempty"`
	Description      string            `json:"description,omitempty"`
	FilterExpression string            `json:"filterExpression,omitempty"`
	Ordering         int               `json:"ordering,omitempty"`
	Groups           []string          `json:"groups,omitempty"`
	Geo              []GeoHierarchyRow `json:"geo,omitempty"`
}

type Tag struct {
	Name     string `json:"name,omitempty"`
	Color    string `json:"color,omitempty"`
	Attached bool   `json:"attached,omitempty"`
}
