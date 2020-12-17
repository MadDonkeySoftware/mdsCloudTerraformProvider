package mdscloud

import (
	"context"
	"fmt"

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
			"user_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MDS_USER_ID", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MDS_PASSWORD", nil),
			},
			"allow_self_cert": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MDS_ALLOW_SELF_CERT", false),
			},
			"identity_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MDS_IDENTITY_URL", nil),
			},
			"qs_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MDS_QS_URL", nil),
			},
			"fs_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MDS_FS_URL", nil),
			},
			"sf_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MDS_SF_URL", nil),
			},
			"sm_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MDS_SM_URL", nil),
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
	userID := d.Get("user_id").(string)
	password := d.Get("password").(string)
	allowSelfCert := d.Get("allow_self_cert").(bool)
	identityURL := d.Get("identity_url").(string)
	qsURL := d.Get("qs_url").(string)
	fsURL := d.Get("fs_url").(string)
	sfURL := d.Get("sf_url").(string)
	smURL := d.Get("sm_url").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	checkProviderInput(&diags, account, "account number", "account", "MDS_ACCOUNT")
	checkProviderInput(&diags, userID, "user id", "user_id", "MDS_USER_ID")
	checkProviderInput(&diags, password, "user password", "password", "MDS_PASSWORD")

	checkProviderInput(&diags, sfURL, "the serverless functions URL", "sf_url", "MDS_SF_URL")
	checkProviderInput(&diags, identityURL, "the identity URL", "identity_url", "MDS_IDENTITY_URL")

	sdk := sdk.NewSdk(account, userID, password, allowSelfCert, map[string]string{
		"identityUrl": identityURL,
		"qsUrl":       qsURL,
		"fsUrl":       fsURL,
		"sfUrl":       sfURL,
		"smUrl":       smURL,
	})
	if sdk == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create mdsCloud sdk",
			Detail:   "Unable to create mdsCloud sdk",
		})
	}

	if len(diags) > 0 {
		return nil, diags
	}

	return sdk, diags
}

func checkProviderInput(diags *diag.Diagnostics, value string, label string, providerKey string, envKey string) {
	if len(value) == 0 {
		*diags = append(*diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Unable to determine %s", label),
			Detail:   fmt.Sprintf("Unable to determine %s. Please use key \"%s\" on provider or environment variable %s.", label, providerKey, envKey),
		})
	}
}
