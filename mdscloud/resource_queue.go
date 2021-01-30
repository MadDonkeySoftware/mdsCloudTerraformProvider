package mdscloud

import (
	"context"

	"github.com/MadDonkeySoftware/mdsCloudSdkGo/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceQueue() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceQueueCreate,
		ReadContext:   resourceQueueRead,
		UpdateContext: resourceQueueUpdate,
		DeleteContext: resourceQueueDelete,
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
			"resource": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
		},
	}
}

func resourceQueueCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	resource := d.Get("resource").(string)
	mdsSdk := m.(*sdk.Sdk)
	qsClient := mdsSdk.GetQueueServiceClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	createResponse, err := qsClient.CreateQueue(&sdk.CreateQueueArgs{
		Name:     name,
		Resource: resource,
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
		"Name": "name",
	}
	mapDataFields(&diags, fieldMappings, d, *createResponse)

	d.SetId(createResponse.Orid)

	// resourceQueueRead(ctx, d, m)

	return diags
}

func resourceQueueRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	orid := d.Id()
	mdsSdk := m.(*sdk.Sdk)
	qsClient := mdsSdk.GetQueueServiceClient()

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

	summary, err := qsClient.GetQueueDetails(&sdk.GetQueueDetailsArgs{
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

	fieldMappings := map[string]string{
		"Orid":     "orid",
		"Resource": "resource",
	}
	mapDataFields(&diags, fieldMappings, d, *summary)

	return diags
}

func resourceQueueUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChange("resource") {
		mdsSdk := m.(*sdk.Sdk)
		orid := d.Get("orid").(string)
		resource := d.Get("resource").(string)
		qsClient := mdsSdk.GetQueueServiceClient()

		// // The default value of empty string is causing issue with a bug in the SDK. Map to "NULL" as work around
		// if resource == "" {
		// 	resource = "NULL"
		// }

		err := qsClient.UpdateQueue(&sdk.UpdateQueueArgs{
			Orid:     orid,
			Resource: resource,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update resource",
				Detail:   "Could not upload resource code with sdk",
			})
			return diags
		}
	}

	// return diags
	return resourceQueueRead(ctx, d, m)
}

func resourceQueueDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	orid := d.Get("orid").(string)
	mdsSdk := m.(*sdk.Sdk)
	qsClient := mdsSdk.GetQueueServiceClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := qsClient.DeleteQueue(&sdk.DeleteQueueArgs{
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
