# Using MdsCloud File Service to Track Terraform State

Much like using AWS S3, Google Cloud Storage, or Azure Blob Storage, mdsCloud
can provide distributed Terraform state tracking through File Service. At the
time of this writing mdsCloud utilizes 
[the http backend](https://www.terraform.io/docs/language/settings/backends/http.html)
type provided by Terraform. While the experience may not seem as "first class"
the functionality is no different from using one of the big cloud specific
providers when it comes to running the various terraform commands.

Assuming the following:

* File Service Base URL: `http://127.0.0.1:8084`
* A container named `sample-tf-state` with the ORID
  `orid:1:mdsCloud:::1001:fs:sample-tf-state`

A basic terraform backend configuration would look like:

```hcl
terraform {
  required_providers {
    mdscloud = {
      version = "0.2"
      source = "maddonkeysoftware.com/tf/mdscloud"
    }
  }
  backend "http" {
    address = "http://127.0.0.1:8084/tf/orid:1:mdsCloud:::1001:fs:sample-tf-state"
    lock_address = "http://127.0.0.1:8084/tf/orid:1:mdsCloud:::1001:fs:sample-tf-state"
    unlock_address = "http://127.0.0.1:8084/tf/orid:1:mdsCloud:::1001:fs:sample-tf-state"
    # NOTE: use TF_HTTP_USERNAME environment variable so your username and
    # password are not exposed if you store your Terraform documents in source
    # control. At the time of this writing it appears that Terraform variables
    # cannot be used to populate these values.
    # See https://www.terraform.io/docs/backends/types/http.html for more env vars
    # username="user"
    # password="password"
  }
}
```

###### Last updated February 2021
