package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/tweag/terraform-provider-statevar/statevar"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: statevar.Provider})
}
