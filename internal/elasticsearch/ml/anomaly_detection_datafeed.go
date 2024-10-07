package ml

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"regexp"

	"github.com/elastic/terraform-provider-elasticstack/internal/clients"
	"github.com/elastic/terraform-provider-elasticstack/internal/clients/elasticsearch"
	"github.com/elastic/terraform-provider-elasticstack/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceAnomalyDetectionDatafeed() *schema.Resource {
	datafeedSchema := map[string]*schema.Schema{
		"id": {
			Description: "Internal ID of the resource",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"feed_id": {
			Description: "A numerical character string that uniquely identifies the datafeed. This identifier can contain lowercase alphanumeric characters (a-z and 0-9), hyphens, and underscores. It must start and end with alphanumeric characters.",
			Type:        schema.TypeString,
			Required:    true,
			ValidateFunc: validation.All(
				validation.StringMatch(regexp.MustCompile(`^[a-z0-9][a-z0-9_-]*[a-z0-9]$`),
					"must start and end with alphanumeric characters and contain only lowercase alphanumeric characters, hyphens, and underscores"),
			),
		},
	}

	utils.AddConnectionSchema(datafeedSchema)

	return &schema.Resource{
		Description: "Adds and updates datafeeds for anomaly detection jobs.",

		CreateContext: resourceDatafeedPut,
		UpdateContext: resourceDatafeedPut,
		ReadContext:   resourceDatafeedRead,
		DeleteContext: resourceDatafeedDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: datafeedSchema,
	}
}

func resourceDatafeedPut(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := clients.NewApiClientFromSDKResource(d, meta)
	if diags.HasError() {
		return diags
	}

	return resourceDatafeedRead(ctx, d, meta)
}

func resourceDatafeedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := clients.NewApiClientFromSDKResource(d, meta)
	if diags.HasError() {
		return diags
	}

	compId, diags := clients.CompositeIdFromStr(d.Id())
	if diags.HasError() {
		return diags
	}

	datafeedID := compId.ResourceId

	datafeed, diags := elasticsearch.GetDatafeed(ctx, client, datafeedID)
	if datafeed == nil && diags == nil {
		tflog.Warn(ctx, fmt.Sprintf("Datafeed %s not found", datafeedID))
		d.SetId("")
		return diags
	}
	if diags.HasError() {
		return diags
	}

	// Set the feed ID
	if err := d.Set("feed_id", datafeedID); err != nil {
		return diag.FromErr(err)
	}
}

func resourceDatafeedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

}
