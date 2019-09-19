package main

import (
	// Change to getnelson when opening PR
	nelson "github.com/drewgonzales360/terraform-provider-nelson/provider"

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
