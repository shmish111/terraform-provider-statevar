package statevar

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func secretResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Delete: resourceDelete,
		Importer: &schema.ResourceImporter{
			State: resourceImporter,
		},

		Schema: map[string]*schema.Schema{
			"value": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func stringResource() *schema.Resource {
	return &schema.Resource{
		Create: defaultResourceCreate,
		Read:   resourceRead,
		Delete: resourceDelete,
		Update: defaultResourceUpdate,
		Importer: &schema.ResourceImporter{
			State: resourceImporter,
		},

		Schema: map[string]*schema.Schema{
			"value": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"default": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An optional value that will be used if no value is stored in state",
				Computed:    true,
				ForceNew:    false,
			},
		},
	}
}

func defaultResourceCreate(d *schema.ResourceData, meta interface{}) error {
	value := d.Get("default")
	d.SetId("-")
	d.Set("value", value)
	return nil
}

func defaultResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCreate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("this resource can only be created on import")
}

func resourceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

func resourceImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	value := d.Id()
	d.SetId("-")
	d.Set("value", value)
	return []*schema.ResourceData{d}, nil
}
