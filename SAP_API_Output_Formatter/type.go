package sap_api_output_formatter

type CompetitorProduct struct {
	ConnectionKey         string `json:"connection_key"`
	Result                bool   `json:"result"`
	RedisKey              string `json:"redis_key"`
	Filepath              string `json:"filepath"`
	APISchema             string `json:"api_schema"`
	CompetitorProductCode string `json:"competitor_product_code"`
	Deleted               bool   `json:"deleted"`
}

type CompetitorProductCollection struct {
	ObjectID              string `json:"ObjectID"`
	CompetitorProductID   string `json:"CompetitorProductID"`
	CompetitorProductUUID string `json:"CompetitorProductUUID"`
	CompetitorProductName string `json:"CompetitorProductName"`
	CompetitorID          string `json:"CompetitorID"`
	ListPrice             string `json:"ListPrice"`
	Currency              string `json:"Currency"`
	BestSellerIndicator   bool   `json:"BestSellerIndicator"`
	ProductComparison     string `json:"ProductComparison"`
	OwnProductID          string `json:"OwnProductID"`
	OwnProductName        string `json:"OwnProductName"`
	OwnProductCategoryID  string `json:"OwnProductCategoryID"`
	BaseUOM               string `json:"BaseUOM"`
	Status                string `json:"Status"`
	CreatedBy             string `json:"CreatedBy"`
	LastChangedBy         string `json:"LastChangedBy"`
	CreatedOn             string `json:"CreatedOn"`
	LastChangedOn         string `json:"LastChangedOn"`
	EntityLastChangedOn   string `json:"EntityLastChangedOn"`
	ETag                  string `json:"ETag"`
}
