# mdsCloudTerraformProvider

## Manual Installation of compiled release

* Download a [release](https://github.com/MadDonkeySoftware/mdsCloudTerraformProvider/releases).
* Unpack the binary, Ex: `terraform-provider-mdscloud_0.2_linux_amd64`, to the terraform plugins directory on the system running terraform. Ex: `~/.terraform.d/plugins/maddonkeysoftware.com/tf/mdscloud/0.2/linux_amd64`. You may need to also rename the file you use to `terraform-provider-cloud`.
  * The full path will be a folder `linux_amd64` with a file named `terraform-provider-mdscloud`


## First Time Setup

After pulling the code for the first time run `make first-time` to set up all dependencies. After that you can `make` or `make build` to perform your dev tasks
