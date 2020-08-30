# Terraform Provider Hashicups

Run the following command to build the provider

```shell
go build -o terraform-provider-hashicups
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
terraform init && terraform apply
```

(WIP tutorial from here)[https://learn.hashicorp.com/tutorials/terraform/provider-complex-read?in=terraform/providers]
(Backend Vault tutorial here)[https://learn.hashicorp.com/tutorials/vault/plugin-backends]
(Other backends here)[https://www.terraform.io/docs/backends/types/http.html]