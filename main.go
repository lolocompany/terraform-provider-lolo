package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/lolocompany/terraform-provider-lolo/lolo"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: lolo.Provider})
}
