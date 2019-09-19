package main

import (
	nelson "github.com/getnelson/terraform-provider-nelson/provider"

	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return nelson.Provider()
		},
	})
}
