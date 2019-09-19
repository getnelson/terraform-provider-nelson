package provider

import (
	"fmt"

	// Change to getnelson when opening PR
	"github.com/drewgonzales360/terraform-provider-nelson/nelson"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// ProviderConfig is used to configure the creation of the client
type ProviderConfig struct {

	// Address is the address of the Nelson server
	Address string

	// Token is the Github Auth token to use when communicating with Nelson
	Path string

	// API version is the Nelson version
	APIVersion string

	// Github Token will be used by the provider to authenticate with Nelson
	GithubToken string
}

// Provider is the entrypoint to this terraform provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NELSON_ADDR", nil),
				Description: "URL of the Nelson server",
			},
			"path": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Default:     "~/.nelson/config.yml",
				Description: "Path to Nelson Config. Defaults to ~/.nelson/config.yml",
			},
			"api_version": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Default:     "v1",
				Description: "Version for Nelson API. Defaults to v1",
			},
			"github_token": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_TOKEN", ""),
				Description: "Github token used to authenticate with Nelson",
			},
		},

		ConfigureFunc: providerConfigure,

		DataSourcesMap: map[string]*schema.Resource{
			"nelson_datacenter": dataSourceDataCenter(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"nelson_repository": resourceRepository(),
			"nelson_blueprint":  resourceBlueprint(),
		},
	}
}

// providerConfigure reads the provider block in the terraform code and creates a config.
// The config will have the necessary information to find and authenicate with the Nelson
// server.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	clientConfig := defaultProviderConfig()

	if address, ok := d.GetOk("address"); ok {
		clientConfig.Address = address.(string)
	} else {
		return nil, fmt.Errorf("error parsing address")
	}
	if path, ok := d.GetOk("path"); ok {
		clientConfig.Path = path.(string)
	} else {
		return nil, fmt.Errorf("error parsing config path")
	}
	if apiVersion, ok := d.GetOk("api_version"); ok {
		clientConfig.APIVersion = apiVersion.(string)
	} else {
		return nil, fmt.Errorf("error parsing api_version")
	}
	if githubToken, ok := d.GetOk("github_token"); ok {
		clientConfig.GithubToken = githubToken.(string)
	} else {
		return nil, fmt.Errorf("error parsing github token")
	}

	nelsonClient, err := nelson.CreateNelson(
		clientConfig.Address,
		clientConfig.APIVersion,
		clientConfig.Path,
	)
	if err != nil {
		return nil, multierror.Prefix(err, "couldn't create config")
	}

	if err := nelsonClient.Login(clientConfig.GithubToken); err != nil {
		return nil, multierror.Prefix(err, "couldn't create provider config:")
	}
	return nelsonClient, nil
}

// DefaultProviderConfig returns a default configuration for the client
func defaultProviderConfig() *ProviderConfig {
	providerConfig := &ProviderConfig{
		Address:     "https://nelson.local:9000",
		Path:        "/home/nelson/.nelson/config.yml",
		APIVersion:  "v1",
		GithubToken: "",
	}

	return providerConfig
}
