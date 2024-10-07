package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elastic/terraform-provider-elasticstack/internal/clients"
	"github.com/elastic/terraform-provider-elasticstack/internal/models"
	"github.com/elastic/terraform-provider-elasticstack/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func GetDatafeed(ctx context.Context, apiClient *clients.ApiClient, feedID string) (*models.Datafeed, diag.Diagnostics) {
	var diags diag.Diagnostics

	esClient, err := apiClient.GetESClient()
	if err != nil {
		return nil, diag.FromErr(err)
	}

	req := esClient.ML.GetDatafeeds.WithDatafeedID(feedID)
	res, err := esClient.ML.GetDatafeeds(req, esClient.ML.GetDatafeeds.WithContext(ctx))
	if err != nil {
		return nil, diag.FromErr(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if diags := utils.CheckError(res, "Unable to find datafeed on cluster."); diags.HasError() {
		return nil, diags
	}

	datafeeds := make(map[string]models.Datafeed)
	if err := json.NewDecoder(res.Body).Decode(&datafeeds); err != nil {
		return nil, diag.FromErr(err)
	}

	if datafeed, ok := datafeeds[feedID]; ok {
		return &datafeed, diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to find datafeed on cluster.",
		Detail:   fmt.Sprint("Unable to find ", feedID, " datafeed on cluster"),
	})

	return nil, diags
}
