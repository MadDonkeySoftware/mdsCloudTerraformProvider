# Resource: mdscloud_function

## Example

```hcl
data "mdscloud_list_functions" "all_functions" {
}
```

## Argument Reference

* `orid` - (Required) The orid for this specific function. Ex: `orid:1:mdsCloud:::1001:sf:test`

## Attribute Reference

* `name` - The name of the function to be created or maintained.
* `version` - A unique version number for the code that was uploaded.
* `runtime` - The runtime to use for this serverless function.
* `entry_point` - The path to the module/method to execute.
* `created` - A timestamp for when the function was created.
* `last_update` - A timestamp for the last code push to the function.
* `last_invoke` - A timestamp for when the function was last executed.

###### Last updated February 2021
