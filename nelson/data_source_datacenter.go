package nelson

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataCenterDataSource() *schema.Resource {
	return &schema.Resource {
		Read: dataCenterDataSourceRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				Description: "Name of the Nelson datacenter"
			},
			"datacenter_url": {
				Type: schema.TypeString,
				Required: true,
				Description: "URL of the datacenter",
			}
		}
	}
}

func dataCenterDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)

	name := d.Get("name").(string)

	namespace, err := nelsonNamespace(name, client)
	if err != nil {
		return fmt.Errorf("error reading Nelson: %s", err)
	}
	if namespace == nil {
		return fmt.Errorf("no namespace %s found", name)
	}

	d.Set("datacenter_url", namespace.DatacenterUrl)

	return nil
}