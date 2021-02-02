# Provider Configuration

The mdsCloud Terraform provider requires configuration similar to the mdCloud
command line interface. In order to communicate with the various mdsCloud
services not only do the various service urls need to be set but the credentials
for the account need to be provided as well. It is highly suggested to use 
[Terraform variables](https://www.terraform.io/docs/language/values/variables.html)
to configure the provider section as this allows for a greater degree of
flexibility when it comes to changing accounts for different deployment
environments such as DEV, QA, or PROD.

Another added boon is that an encryption tool can be used to encrypt your
Terraform variables file (`tfvars`) for storage in your source control of
choice. Some 

`main.tf`:

```hcl
provider "mdscloud" {
  account=var.account
  user_id=var.user_id
  password=var.password
  allow_self_cert=var.allow_self_signed_cert

  sf_url=var.sf_url
  qs_url=var.qs_url
  fs_url=var.fs_url
  sm_url=var.sm_url
  identity_url=var.identity_url
}
```

`local.tfvars`:

```ini
identity_url="https://127.0.0.1:8081"
qs_url="http://127.0.0.1:8083"
fs_url="http://127.0.0.1:8084"
sf_url="http://127.0.0.1:8085"
sm_url="http://127.0.0.1:8086"
account="1001"
user_id="myUser"
password="password"
allow_self_signed_cert=true
```

###### Last updated February 2020