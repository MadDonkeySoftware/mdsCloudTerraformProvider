# Resource: mdscloud_function

## Example

```hcl
resource "mdscloud_function" "example_function" {
  name = "test"
  file_name = "code.zip"
  runtime = "node"
  entry_point = "src/test:main"
  source_code_hash = filebase64sha256("code.zip")
  context = "some string for your context"
}
```

## Argument Reference

* `name` - (Required) The name of the function to be created or maintained.
* `file_name` - (Required) The zip file containing the code to upload.
* `runtime` - (Required) The runtime to use for this serverless function.
* `entry_point` - (Required) The path to the module/method to execute.
* `source_code_hash` - (Required) The hash of the file referenced by `file_name`
* `context` - (Optional) A user defined string that will be supplied to each execution of the serverless function
  

## Attribute Reference

* `orid` - The orid for this specific function. Ex: `orid:1:mdsCloud:::1001:sf:test`
* `version` - A unique version number for the code that was uploaded.
* `created` - A timestamp for when the function was created.
* `last_update` - A timestamp for the last code push to the function.
* `last_invoke` - A timestamp for when the function was last executed.

###### Last updated February 2021
