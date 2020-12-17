package mdscloud

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/MadDonkeySoftware/mdsCloudSdkGo/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerlessFunction() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerlessFunctionRead,
		Schema: map[string]*schema.Schema{
			"orid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"runtime": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entry_point": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_update": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_invoke": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceServerlessFunctionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	orid := d.Get("orid").(string)
	mdsSdk := m.(*sdk.Sdk)
	sfClient := mdsSdk.GetServerlessFunctionsClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	funcDetails, err := sfClient.GetFunctionDetails(orid)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get resource",
			Detail:   "Could not get resource from sdk",
		})
		return diags
	}

	fieldMappings := map[string]string{
		"Orid":       "orid",
		"Name":       "name",
		"Version":    "version",
		"EntryPoint": "entry_point",
		"Created":    "created",
		"LastUpdate": "last_update",
		"LastInvoke": "last_invoke",
	}
	v := reflect.ValueOf(*funcDetails)
	for key, funcKey := range fieldMappings {
		fmt.Println(key)
		fmt.Println(funcKey)
		if err := d.Set(funcKey, v.FieldByName(key).String()); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to map resources",
				Detail:   fmt.Sprintf("Could not map response field (%s) from API of resource", key),
			})
			return diags
		}
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
