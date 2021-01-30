package mdscloud

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/MadDonkeySoftware/mdsCloudSdkGo/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
)

func resourceStateMachine() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStateMachineCreate,
		ReadContext:   resourceStateMachineRead,
		UpdateContext: resourceStateMachineUpdate,
		DeleteContext: resourceStateMachineDelete,
		Schema: map[string]*schema.Schema{
			"orid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"definition": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// https://www.terraform.io/docs/extend/schemas/schema-behaviors.html
					oldJSON, _ := structure.NormalizeJsonString(old)
					newJSON, _ := structure.NormalizeJsonString(new)
					return newJSON == oldJSON
				},
			},
		},
	}
}

func resourceStateMachineCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	definition := d.Get("definition").(string)
	mdsSdk := m.(*sdk.Sdk)
	smClient := mdsSdk.GetStateMachineServiceClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	createResponse, err := smClient.CreateStateMachine(&sdk.CreateStateMachineArgs{
		Definition: definition,
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

	resourceStateMachineRead(ctx, d, m)

	return diags
}

func resourceStateMachineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// NOTE: There is not really a "read the container metadata" call right now. List contents to verify the container exists.
	orid := d.Id()
	mdsSdk := m.(*sdk.Sdk)
	smClient := mdsSdk.GetStateMachineServiceClient()

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

	details, err := smClient.GetStateMachineDetails(&sdk.GetStateMachineDetailsArgs{
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

	definitionBytes, err := json.Marshal(details.Definition)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get resource definition",
			Detail:   "Could not get resource definition from sdk",
		})
		return diags
	}

	if err := d.Set("definition", string(definitionBytes)); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get resource",
			Detail:   "Could not get resource from sdk",
		})
	}

	return diags
}

func resourceStateMachineUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Has Change - update",
		Detail:   strconv.FormatBool(d.HasChange("definition")),
	})

	if d.HasChange("definition") {
		mdsSdk := m.(*sdk.Sdk)
		orid := d.Get("orid").(string)
		definition := d.Get("definition").(string)
		smClient := mdsSdk.GetStateMachineServiceClient()

		_, err := smClient.UpdateStateMachine(&sdk.UpdateStateMachineArgs{
			Orid:       orid,
			Definition: definition,
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

	return resourceStateMachineRead(ctx, d, m)

	// return diags
}

func resourceStateMachineDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	orid := d.Get("orid").(string)
	mdsSdk := m.(*sdk.Sdk)
	smClient := mdsSdk.GetStateMachineServiceClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := smClient.DeleteStateMachine(&sdk.DeleteStateMachineArgs{
		Orid: orid,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete resource",
			Detail:   "Could not remove resource from sdk",
		})
		return diags
	}

	d.SetId("")

	return diags
}
