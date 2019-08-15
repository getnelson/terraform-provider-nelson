package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceDataCenter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDataCenterRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Nelson datacenter",
			},
			"datacenter_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URL of the datacenter",
			},
		},
	}
}

func dataSourceDataCenterRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
