package mdscloud

import (
	"context"

	"github.com/MadDonkeySoftware/mdsCloudSdkGo/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceContainer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceContainerCreate,
		ReadContext:   resourceContainerRead,
		// UpdateContext: resourceContainerUpdate,
		DeleteContext: resourceContainerDelete,
		Schema: map[string]*schema.Schema{
			"orid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceContainerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	mdsSdk := m.(*sdk.Sdk)
	fsClient := mdsSdk.GetFileServiceClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	createResponse, err := fsClient.CreateContainer(&sdk.CreateContainerArgs{
		Name: name,
	})

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create resource",
			Detail:   err.Error(),
			// Detail:   "Could not create resource from sdk",
		})
		return diags
	}

	fieldMappings := map[string]string{
		"Orid": "orid",
	}
	mapDataFields(&diags, fieldMappings, d, *createResponse)

	d.SetId(createResponse.Orid)

	// resourceContainerRead(ctx, d, m)

	return diags
}

func resourceContainerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// NOTE: There is not really a "read the container metadata" call right now. List contents to verify the container exists.
	orid := d.Id()
	mdsSdk := m.(*sdk.Sdk)
	fsClient := mdsSdk.GetFileServiceClient()

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

	_, err := fsClient.ListContainerContents(&sdk.ListContainerContentsArgs{
		Orid: orid,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get resource",
			Detail:   "Could not get resource from sdk",
		})
		return diags
	}

	// The ORID field is not in the list container contents response. Map it manually
	if err := d.Set("orid", orid); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get resource",
			Detail:   "Could not get resource from sdk",
		})
	}

	return diags
}

// func resourceContainerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	// NOTE: In the event the name has changed this container will be destroyed and recreated. There is no modify as of this implementation.
// 	return resourceContainerRead(ctx, d, m)
// }

func resourceContainerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	orid := d.Get("orid").(string)
	mdsSdk := m.(*sdk.Sdk)
	fsClient := mdsSdk.GetFileServiceClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := fsClient.DeleteContainerOrPath(&sdk.DeleteContainerArgs{
		Orid: orid,
	})
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
