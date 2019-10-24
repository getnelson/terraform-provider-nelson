package provider

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"

	"github.com/getnelson/terraform-provider-nelson/nelson"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceBlueprint() *schema.Resource {
	return &schema.Resource{
		Create: resourceBlueprintCreate,
		Read:   resourceBlueprintRead,
		Update: resourceBlueprintUpdate,
		Delete: resourceBlueprintDelete,
		Importer: &schema.ResourceImporter{
			State: resourceBlueprintImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Nelson blueprint",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of the blueprint",
			},
			"file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Blueprint as a string",
			},
		},
	}
}

// Create and Update will be the same thing in this case. Blueprints
// are an update/create only resource.
func resourceBlueprintCreate(d *schema.ResourceData, meta interface{}) error {
	value, exists := d.GetOk("file")
	if !exists {
		return fmt.Errorf("file was not specified")
	}

	name := d.Get("name").(string)
	templateBytes := []byte(value.(string))
	cbr := nelson.CreateBlueprintRequest{
		Name:        name,
		Description: d.Get("description").(string),
		Sha256:      fmt.Sprintf("%x", sha256.Sum256(templateBytes)),
		Template:    b64.StdEncoding.EncodeToString(templateBytes),
	}

	nelsonClient, ok := meta.(*nelson.Nelson)
	if !ok {
		return fmt.Errorf("provider configuration could not create Nelson client")
	}
	if err := nelsonClient.CreateBlueprint(cbr); err != nil {
		return err
	}

	d.SetId(name)
	return nil
}

// https://www.terraform.io/docs/plugins/provider.html#read
func resourceBlueprintRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Id()

	nelsonClient, ok := meta.(*nelson.Nelson)
	if !ok {
		return fmt.Errorf("provider configuration could not create Nelson client")
	}
	if blueprint, err := nelsonClient.GetBlueprint(name); err != nil {
		d.Set("name", blueprint.Name)
		d.Set("description", blueprint.Description)

		decode, err := b64.StdEncoding.DecodeString(blueprint.Template)
		template, err := b64.StdEncoding.DecodeString(string(decode))
		if err != nil {
			return err
		}
		d.Set("file", template)
		return err
	}
	return nil
}

// Create and Update will be the same thing in this case. Blueprints
// are an update/create only resource.
func resourceBlueprintUpdate(d *schema.ResourceData, meta interface{}) error {
	value, exists := d.GetOk("file")
	if !exists {
		return fmt.Errorf("file was not specified")
	}

	templateBytes := []byte(value.(string))
	cbr := nelson.CreateBlueprintRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Sha256:      fmt.Sprintf("%x", sha256.Sum256(templateBytes)),
		Template:    b64.StdEncoding.EncodeToString(templateBytes),
	}

	nelsonClient, ok := meta.(*nelson.Nelson)
	if !ok {
		return fmt.Errorf("provider configuration could not create Nelson client")
	}
	if err := nelsonClient.CreateBlueprint(cbr); err != nil {
		return err
	}
	return nil
}

func resourceBlueprintDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceBlueprintImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	name := d.Id()

	nelsonClient, ok := meta.(*nelson.Nelson)
	if !ok {
		return nil, fmt.Errorf("provider configuration could not create Nelson client")
	}
	blueprint, err := nelsonClient.GetBlueprint(name)
	if err != nil {
		return nil, err
	}

	d.SetId(name)
	d.Set("name", blueprint.Name)
	d.Set("description", blueprint.Description)

	decode, err := b64.StdEncoding.DecodeString(blueprint.Template)
	template, err := b64.StdEncoding.DecodeString(string(decode))
	if err != nil {
		return nil, err
	}
	d.Set("file", template)
	return nil, nil
}
