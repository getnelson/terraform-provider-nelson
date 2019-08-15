package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceRepositoryCreate,
		Read:   resourceRepositoryRead,
		Update: resourceRepositoryUpdate,
		Delete: resourceRepositoryDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Nelson repository",
			},
			"repository_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URL of the repository",
			},
		},
	}
}

// Call nelson enable
func resourceRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// Call Nelson repos list
func resourceRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

//
func resourceRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// Call Nelson repos sync
func resourceRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
