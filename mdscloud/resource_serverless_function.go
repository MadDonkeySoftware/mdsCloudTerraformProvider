package mdscloud

import (
	"context"

	"github.com/MadDonkeySoftware/mdsCloudSdkGo/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerlessFunction() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFunctionCreate,
		ReadContext:   resourceFunctionRead,
		UpdateContext: resourceFunctionUpdate,
		DeleteContext: resourceFunctionDelete,
		Schema: map[string]*schema.Schema{
			"orid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"runtime": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"entry_point": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
			"file_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_code_hash": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFunctionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	mdsSdk := m.(*sdk.Sdk)
	sfClient := mdsSdk.GetServerlessFunctionsClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	summary, err := sfClient.CreateFunction(name)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create resource",
			Detail:   err.Error(),
			// Detail:   "Could not create resource from sdk",
		})
		return diags
	}

	runtime := d.Get("runtime").(string)
	entryPoint := d.Get("entry_point").(string)
	fileName := d.Get("file_name").(string)

	err = updateFunctionCode(mdsSdk, summary.Orid, runtime, entryPoint, fileName)
	if err != nil {
		sfClient.DeleteFunction(summary.Orid)

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to upload resource code",
			Detail:   "Could not upload resource code with sdk",
		})
		return diags
	}

	d.SetId(summary.Orid)

	resourceFunctionRead(ctx, d, m)

	return diags
}

func resourceFunctionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	orid := d.Id()
	mdsSdk := m.(*sdk.Sdk)
	sfClient := mdsSdk.GetServerlessFunctionsClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if orid == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to determine Id",
			Detail:   "Unable to determine Id",
		})
		return diags
	}

	summary, err := sfClient.GetFunctionDetails(orid)
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
	mapDataFields(&diags, fieldMappings, d, *summary)

	return diags
}

func updateFunctionCode(mdsSdk *sdk.Sdk, orid string, runtime string, entryPoint string, fileName string) error {
	sfClient := mdsSdk.GetServerlessFunctionsClient()
	err := sfClient.UpdateFunctionCode(orid, runtime, entryPoint, fileName)
	return err
}

func resourceFunctionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChange("source_code_hash") || d.HasChange("runtime") || d.HasChange("entry_point") {
		mdsSdk := m.(*sdk.Sdk)
		orid := d.Get("orid").(string)
		runtime := d.Get("runtime").(string)
		entryPoint := d.Get("entry_point").(string)
		fileName := d.Get("file_name").(string)

		err := updateFunctionCode(mdsSdk, orid, runtime, entryPoint, fileName)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to upload resource code",
				Detail:   "Could not upload resource code with sdk",
			})
			return diags
		}
	}

	// return diags
	return resourceFunctionRead(ctx, d, m)
}

func resourceFunctionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("orid").(string)
	mdsSdk := m.(*sdk.Sdk)
	sfClient := mdsSdk.GetServerlessFunctionsClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := sfClient.DeleteFunction(name)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get resource",
			Detail:   "Could not get resource from sdk",
		})
		return diags
	}

	d.SetId("")

	return diags
}
