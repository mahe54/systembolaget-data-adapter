package systembolaget

type Category1 string
type Category2 string
type WebUrl string

const (
	Ol                   Category1 = "Öl"
	Sprit                Category1 = "Sprit"
	Cider_o_blanddrycker Category1 = "Cider & blanddrycker"
	Alkoholfritt         Category1 = "Alkoholfritt"
	Presentartiklar      Category1 = "Presentartiklar"
	Vin                  Category1 = "Vin"

	Rott_vin        Category2 = "Rött vin"
	Vitt_vin        Category2 = "Vitt vin"
	Mousserande_vin Category2 = "Mousserande vin"
	Rosevin         Category2 = "Rosévin"
	Vinlada         Category2 = "Vinlåda"
	Starkvin        Category2 = "Starkvin"

	URL WebUrl = "https://www.systembolaget.se/api/gateway/productsearch/search/"
)

type Client struct {
	Client ClientInterface
}

// NewClient acts as a factory and creates a Systembolaget Client to use for retrieving data from www.systembolaget.se
// func NewClient(clientImpl ClientInterface) (*Client, error) {
// 	client := &Client{Client: clientImpl}
// 	return client, nil
// }

func NewClient(clientImpl ClientInterface) (*Client, error) {
	client := &Client{Client: clientImpl}
	return client, nil
}

// The interface/func to implement
type ClientInterface interface {
	GetProductPages(category1 Category1, category2 Category2) []ProductPage
}

type ProductPage struct {
	Metadata        Metadata      `json:"metadata,omitempty"`
	Products        []Products    `json:"products,omitempty"`
	Filters         []Filters     `json:"filters,omitempty"`
	FilterMenuItems []interface{} `json:"filterMenuItems,omitempty"`
}
type PriceRange struct {
	Min float64 `json:"min,omitempty"`
	Max float64 `json:"max,omitempty"`
}
type VolumeRange struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}
type AlcoholPercentageRange struct {
	Min float64 `json:"min,omitempty"`
	Max float64 `json:"max,omitempty"`
}
type SugarContentRange struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}
type SugarContentGramPer100MlRange struct {
	Min int     `json:"min,omitempty"`
	Max float64 `json:"max,omitempty"`
}
type Metadata struct {
	DocCount                      int                           `json:"docCount,omitempty"`
	FullAssortmentDocCount        int                           `json:"fullAssortmentDocCount,omitempty"`
	NextPage                      int                           `json:"nextPage,omitempty"`
	PriceRange                    PriceRange                    `json:"priceRange,omitempty"`
	VolumeRange                   VolumeRange                   `json:"volumeRange,omitempty"`
	AlcoholPercentageRange        AlcoholPercentageRange        `json:"alcoholPercentageRange,omitempty"`
	SugarContentRange             SugarContentRange             `json:"sugarContentRange,omitempty"`
	SugarContentGramPer100MlRange SugarContentGramPer100MlRange `json:"sugarContentGramPer100mlRange,omitempty"`
	DidYouMeanQuery               interface{}                   `json:"didYouMeanQuery,omitempty"`
}
type Images struct {
	ImageURL string      `json:"imageUrl,omitempty"`
	FileType interface{} `json:"fileType,omitempty"`
	Size     interface{} `json:"size,omitempty"`
}
type TasteClocks struct {
	Key   string `json:"key,omitempty"`
	Value int    `json:"value,omitempty"`
}
type Products struct {
	ProductID                       string        `json:"productId,omitempty"`
	ProductNumber                   string        `json:"productNumber,omitempty"`
	ProductNameBold                 string        `json:"productNameBold,omitempty"`
	ProductNameThin                 string        `json:"productNameThin,omitempty"`
	Category                        interface{}   `json:"category,omitempty"`
	ProductNumberShort              string        `json:"productNumberShort,omitempty"`
	ProducerName                    string        `json:"producerName,omitempty"`
	SupplierName                    string        `json:"supplierName,omitempty"`
	IsKosher                        bool          `json:"isKosher,omitempty"`
	BottleText                      string        `json:"bottleText,omitempty"`
	RestrictedParcelQuantity        int           `json:"restrictedParcelQuantity,omitempty"`
	IsOrganic                       bool          `json:"isOrganic,omitempty"`
	IsSustainableChoice             bool          `json:"isSustainableChoice,omitempty"`
	IsClimateSmartPackaging         bool          `json:"isClimateSmartPackaging,omitempty"`
	IsEthical                       bool          `json:"isEthical,omitempty"`
	EthicalLabel                    interface{}   `json:"ethicalLabel,omitempty"`
	IsWebLaunch                     bool          `json:"isWebLaunch,omitempty"`
	ProductLaunchDate               string        `json:"productLaunchDate,omitempty"`
	IsCompletelyOutOfStock          bool          `json:"isCompletelyOutOfStock,omitempty"`
	IsTemporaryOutOfStock           bool          `json:"isTemporaryOutOfStock,omitempty"`
	AlcoholPercentage               float64       `json:"alcoholPercentage,omitempty"`
	VolumeText                      string        `json:"volumeText,omitempty"`
	Volume                          int           `json:"volume,omitempty"`
	Price                           float64       `json:"price,omitempty"`
	Country                         string        `json:"country,omitempty"`
	OriginLevel1                    string        `json:"originLevel1,omitempty"`
	OriginLevel2                    string        `json:"originLevel2,omitempty"`
	CategoryLevel1                  string        `json:"categoryLevel1,omitempty"`
	CategoryLevel2                  string        `json:"categoryLevel2,omitempty"`
	CategoryLevel3                  string        `json:"categoryLevel3,omitempty"`
	CategoryLevel4                  interface{}   `json:"categoryLevel4,omitempty"`
	CustomCategoryTitle             string        `json:"customCategoryTitle,omitempty"`
	AssortmentText                  string        `json:"assortmentText,omitempty"`
	Usage                           string        `json:"usage,omitempty"`
	Taste                           string        `json:"taste,omitempty"`
	TasteSymbols                    []string      `json:"tasteSymbols,omitempty"`
	TasteClockGroupBitter           interface{}   `json:"tasteClockGroupBitter,omitempty"`
	TasteClockGroupSmokiness        interface{}   `json:"tasteClockGroupSmokiness,omitempty"`
	TasteClockBitter                int           `json:"tasteClockBitter,omitempty"`
	TasteClockFruitacid             int           `json:"tasteClockFruitacid,omitempty"`
	TasteClockBody                  int           `json:"tasteClockBody,omitempty"`
	TasteClockRoughness             int           `json:"tasteClockRoughness,omitempty"`
	TasteClockSweetness             int           `json:"tasteClockSweetness,omitempty"`
	TasteClockSmokiness             int           `json:"tasteClockSmokiness,omitempty"`
	TasteClockCasque                int           `json:"tasteClockCasque,omitempty"`
	Assortment                      string        `json:"assortment,omitempty"`
	RecycleFee                      int           `json:"recycleFee,omitempty"`
	IsManufacturingCountry          bool          `json:"isManufacturingCountry,omitempty"`
	IsRegionalRestricted            bool          `json:"isRegionalRestricted,omitempty"`
	PackagingLevel1                 string        `json:"packagingLevel1,omitempty"`
	IsNews                          bool          `json:"isNews,omitempty"`
	Images                          []Images      `json:"images,omitempty"`
	IsDiscontinued                  bool          `json:"isDiscontinued,omitempty"`
	IsSupplierTemporaryNotAvailable bool          `json:"isSupplierTemporaryNotAvailable,omitempty"`
	SugarContent                    int           `json:"sugarContent,omitempty"`
	SugarContentGramPer100Ml        int           `json:"sugarContentGramPer100ml,omitempty"`
	Seal                            []interface{} `json:"seal,omitempty"`
	Vintage                         interface{}   `json:"vintage,omitempty"`
	Grapes                          []interface{} `json:"grapes,omitempty"`
	OtherSelections                 interface{}   `json:"otherSelections,omitempty"`
	TasteClocks                     []TasteClocks `json:"tasteClocks,omitempty"`
	Color                           string        `json:"color,omitempty"`
	DishPoints                      interface{}   `json:"dishPoints,omitempty"`
}
type SearchModifiers struct {
	Value        string      `json:"value,omitempty"`
	Count        int         `json:"count,omitempty"`
	IsActive     bool        `json:"isActive,omitempty"`
	SubtitleText interface{} `json:"subtitleText,omitempty"`
	FriendlyURL  interface{} `json:"friendlyUrl,omitempty"`
}
type Filters struct {
	Name                  string            `json:"name,omitempty"`
	Type                  string            `json:"type,omitempty"`
	DisplayName           string            `json:"displayName,omitempty"`
	Description           string            `json:"description,omitempty"`
	Summary               interface{}       `json:"summary,omitempty"`
	LegalText             interface{}       `json:"legalText,omitempty"`
	IsMultipleChoice      bool              `json:"isMultipleChoice,omitempty"`
	IsActive              bool              `json:"isActive,omitempty"`
	IsSubtitleTextVisible bool              `json:"isSubtitleTextVisible,omitempty"`
	SearchModifiers       []SearchModifiers `json:"searchModifiers,omitempty"`
	Child                 interface{}       `json:"child,omitempty"`
}
