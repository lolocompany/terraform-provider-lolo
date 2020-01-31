package lolo

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/lolocompany/lolo-sdk-go"
)

func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Lolo API key.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"lolo_app": resourceLoloApp(),
		},
		ConfigureFunc: providerConfigure,
	}

	return provider
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client, err := lolo.NewClient(d.Get("api_key").(string))

	if err != nil {
		log.Printf("[ERROR] Lolo Client error: %v", err)
		return client, err
	}

	return client, nil
}
