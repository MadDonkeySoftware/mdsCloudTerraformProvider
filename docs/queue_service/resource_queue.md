# Resource: mdscloud_queue

## Examples

### Stand alone queue

```hcl
resource "mdscloud_queue" "example_queue" {
  name = "exampleQueue"
}
```

### Queue that invokes function

```hcl
resource "mdscloud_function" "example_function" {
  name = "test"
  file_name = "code.zip"
  runtime = "node"
  entry_point = "src/test:main"
  source_code_hash = filebase64sha256("code.zip")
}

resource "mdscloud_queue" "example_queue" {
  name = "exampleQueue"
  resource = mdscloud_function.example_function.orid
}
```

## Argument Reference

* `name` - (Required) The name of the queue to be created or maintained. If the
  name is changed the queue, along with its contents, will be destroyed.
* `resource` - (Optional) The resource to invoke when an item is enqueued.
  Currently this only works with a serverless function.

## Attribute Reference

* `orid` - The orid for this specific queue. Ex: `orid:1:mdsCloud:::1001:qs:exampleQueue`

###### Last updated February 2021
