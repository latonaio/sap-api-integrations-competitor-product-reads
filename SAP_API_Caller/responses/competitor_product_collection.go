package responses

type CompetitorProductCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                          string `json:"ObjectID"`
			CompetitorProductID               string `json:"CompetitorProductID"`
			CompetitorProductUUID             string `json:"CompetitorProductUUID"`
			CompetitorProductName             string `json:"CompetitorProductName"`
			CompetitorID                      string `json:"CompetitorID"`
			ListPrice                         string `json:"ListPrice"`
			Currency                          string `json:"Currency"`
			BestSellerIndicator               bool   `json:"BestSellerIndicator"`
			ProductComparison                 string `json:"ProductComparison"`
			OwnProductID                      string `json:"OwnProductID"`
			OwnProductName                    string `json:"OwnProductName"`
			OwnProductCategoryID              string `json:"OwnProductCategoryID"`
			BaseUOM                           string `json:"BaseUOM"`
			Status                            string `json:"Status"`
			CreatedBy                         string `json:"CreatedBy"`
			LastChangedBy                     string `json:"LastChangedBy"`
			CreatedOn                         string `json:"CreatedOn"`
			LastChangedOn                     string `json:"LastChangedOn"`
			EntityLastChangedOn               string `json:"EntityLastChangedOn"`
			ETag                              string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
