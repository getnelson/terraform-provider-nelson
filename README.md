# Terraform Nelson provider

This project is a [Terraform Provider](https://www.terraform.io/docs/plugins/provider.html) for an open source project called [Nelson](https://getnelson.io/)

## Installation

Installing plugins is well documented on the [Terraform site](https://www.terraform.io/docs/plugins/basics.html#installing-plugins), but it's explained here as well just in case. To install this plugin, build it from source (see CONTRIBUTING.md). Once the binary has been built for a given OS, move the binary to the directory corresponding to the OS. For example, unix systems would move the binary to `~/.terraform.d/plugins/terraform-provider-nelson`.

## Usage

See the /examples directory for full examples on usage with more detailed comments and explanations.

Setting up the provider can be done with this block

```hcl
provider nelson {
  address = "https://nelson.local"         # Address of the Nelson server
  path = "/home/nelson/.nelson/config.yml" # Local path used for storing session tokens
}
```

## Resources

This section attempts to explain the resources that can be managed by Nelson. As this provider grows, more Nelson primitives will get managed by this provider. Pull requests are welcome!

#### Blueprints

```hcl
resource "nelson_blueprint" "deployment" {
  name = "nelson"
  description = "deployment blueprint"
  file = "${file("${path.module}/deployment.bp")}"
}
```
