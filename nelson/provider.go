package nelson

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("NELSON_ADDR", nil),
				Description: "URL of the Nelson server",
			},
			"token": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("NELSON_TOKEN", nil),
				Description: "",
			}
		}.

		ConfigureFunc: providerConfigure,

		DataSourcesMap: map[string]*schema.Resource{
			"nelson_datacenter":    dataCenterDataSource(),
			"nelson_load_balancer": loadBalancerDataSource()
		},

		ResourcesMap: map[string]*schema.Resource{
			"nelson_repository":    repositoryResource(),
			"nelson_datacenter":    dataCenterResource(),
			"nelson_load_balancer": loadBalancerResource()
		}
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	clientConfig := api.DefaultConfig()

	clientConfig.Address = d.Get("address").(string)
	clientConfig.Token = d.Get("token").(string)

	client := api.NewClient(clientConfig)

	return client, nil
}