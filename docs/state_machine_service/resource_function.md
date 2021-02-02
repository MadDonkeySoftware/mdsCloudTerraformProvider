# Resource: mdscloud_state_machine

## Example

[See the example app](https://github.com/MadDonkeySoftware/mdsCloudSampleApp/tree/main/terraform)
for more details about the state machine json file.

```hcl
data "template_file" "test_state_machine_definition" {
  template = file("${path.module}/sample-state-machine.json")
  vars = {
    sf_one_orid   = mdscloud_function.sf_one.orid
    sf_two_orid   = mdscloud_function.sf_two.orid
    sf_three_orid = mdscloud_function.sf_three.orid
  }
}

resource "mdscloud_state_machine" "test_sm" {
  definition = data.template_file.test_state_machine_definition.rendered
}
```

## Argument Reference

* `definition` - (Required) The definition of the state machine adhering to the
  [State Machine Language Specification](https://states-language.net).

## Attribute Reference

* `orid` - The orid for this specific state machine. Ex: `orid:1:mdsCloud:::1001:sm:test`

###### Last updated February 2020
