package mdscloud

import (
	"context"
	"strconv"
	"time"

	"github.com/MadDonkeySoftware/mdsCloudSdkGo/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceListServerlessFunctions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceListServerlessFunctionsRead,
		Schema: map[string]*schema.Schema{
			"functions": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"orid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceListServerlessFunctionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	mdsSdk := m.(*sdk.Sdk)
	sfClient := mdsSdk.GetServerlessFunctionsClient("", "")

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	functions, err := sfClient.ListFunctions()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get resources",
			Detail:   "Could not get resources from sdk",
		})
		return diags
	}

	data := make([]map[string]interface{}, 0)
	for _, f := range *functions {
		e := map[string]interface{}{
			"orid": f.Orid,
			"name": f.Name,
		}
		data = append(data, e)
	}
	if err := d.Set("functions", data); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to map resources",
			Detail:   "Could not map response from API of resources",
		})
		return diags
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
