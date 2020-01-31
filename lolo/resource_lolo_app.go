package lolo

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/lolocompany/lolo-sdk-go"
)

func resourceLoloApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoloAppCreate,
		Read:   resourceLoloAppRead,
		Update: resourceLoloAppUpdate,
		Delete: resourceLoloAppDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
		},
	}
}

/*
 * Create
 */
func resourceLoloAppCreate(d *schema.ResourceData, m interface{}) error {
	app := &lolo.App{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	if err := m.(*lolo.Client).CreateApp(app); err != nil {
		return fmt.Errorf("Failed to create app: %s", err.Error())
	}

	d.SetId(app.Id)
	return resourceLoloAppRead(d, m)
}

/*
 * Read
 */
func resourceLoloAppRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	app, err := m.(*lolo.Client).GetApp(id)

	if err != nil {
		return err
	}

	if err = d.Set("name", app.Name); err != nil {
		return err
	}

	if err = d.Set("description", app.Description); err != nil {
		return err
	}

	return nil
}

/*
 * Update
 */
func resourceLoloAppUpdate(d *schema.ResourceData, m interface{}) error {
	app := &lolo.App{
		Id:          d.Id(),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	if err := m.(*lolo.Client).UpdateApp(app); err != nil {
		return fmt.Errorf("Failed to update app: %s", err.Error())
	}

	return resourceLoloAppRead(d, m)
}

/*
 * Delete
 */
func resourceLoloAppDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	if err := m.(*lolo.Client).DeleteApp(id); err != nil {
		return err
	}
	return nil
}
