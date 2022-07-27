package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-competitor-product-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToCompetitorProductCollection(raw []byte, l *logger.Logger) ([]CompetitorProductCollection, error) {
	pm := &responses.CompetitorProductCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CompetitorProductCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	competitorProductCollection := make([]CompetitorProductCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		competitorProductCollection = append(competitorProductCollection, CompetitorProductCollection{
			ObjectID:              data.ObjectID,
			CompetitorProductID:   data.CompetitorProductID,
			CompetitorProductUUID: data.CompetitorProductUUID,
			CompetitorProductName: data.CompetitorProductName,
			CompetitorID:          data.CompetitorID,
			ListPrice:             data.ListPrice,
			Currency:              data.Currency,
			BestSellerIndicator:   data.BestSellerIndicator,
			ProductComparison:     data.ProductComparison,
			OwnProductID:          data.OwnProductID,
			OwnProductName:        data.OwnProductName,
			OwnProductCategoryID:  data.OwnProductCategoryID,
			BaseUOM:               data.BaseUOM,
			Status:                data.Status,
			CreatedBy:             data.CreatedBy,
			LastChangedBy:         data.LastChangedBy,
			CreatedOn:             data.CreatedOn,
			LastChangedOn:         data.LastChangedOn,
			EntityLastChangedOn:   data.EntityLastChangedOn,
			ETag:                  data.ETag,
		})
	}

	return competitorProductCollection, nil
}
