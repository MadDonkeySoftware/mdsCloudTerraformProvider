package mdscloud

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func mapDataFields(diags *diag.Diagnostics, fieldMappings map[string]string, d *schema.ResourceData, sourceObj interface{}) {
	v := reflect.ValueOf(sourceObj)
	for key, funcKey := range fieldMappings {
		if err := d.Set(funcKey, v.FieldByName(key).String()); err != nil {
			*diags = append(*diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to map resources",
				Detail:   fmt.Sprintf("Could not map response field (%s) from API of resource property \"%s\"", key, funcKey),
			})
		}
	}
}
