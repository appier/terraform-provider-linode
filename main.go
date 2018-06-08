package main

import (
	"github.com/appier/terraform-provider-linode/linode"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: linode.Provider,
	})
}
