# Resource: mdscloud_container

## Example

```hcl
resource "mdscloud_container" "example_container" {
  name = "example"
}
```

## Argument Reference

* `name` - (Required) The name of the container to be created or maintained. If
  the name is changed the container, along with its contents, will be destroyed.

## Attribute Reference

* `orid` - The orid for this specific container. Ex: `orid:1:mdsCloud:::1001:fs:example`

###### Last updated February 2020
