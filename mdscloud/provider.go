package mdscloud

import (
	"context"

	"github.com/MadDonkeySoftware/mdsCloudSdkGo/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MDS_ACCOUNT", nil),
			},
			"sf_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MDS_SF_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"mdscloud_function": resourceServerlessFunction(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"mdscloud_list_functions": dataSourceListServerlessFunctions(),
			"mdscloud_function":       dataSourceServerlessFunction(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	account := d.Get("account").(string)
	sfURL := d.Get("sf_url").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if len(account) == 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to determine account number",
			Detail:   "Unable to determine account number. Please use key \"account\" on provider or environment variable MDS_ACCOUNT.",
		})
		return nil, diags
	}

	if len(sfURL) == 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to determine account number",
			Detail:   "Unable to determine account number. Please use key \"sfUrl\" on provider or environment variable MDS_SF_URL.",
		})
		return nil, diags
	}

	sdk := sdk.NewSdk(account, map[string]string{"sfUrl": sfURL})
	if sdk == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create mdsCloud sdk",
			Detail:   "Unable to create mdsCloud sdk",
		})
		return nil, diags
	}

	return sdk, diags
}
