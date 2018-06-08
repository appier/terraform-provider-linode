package linode

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LINODE_TOKEN", nil),
				Description: "Linode Token",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"linode_domain_record": resourceLinodeDomainRecord(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"linode_domain": dataSourceLinodeDomain(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := NewLinodeClient(d.Get("token").(string))
	return client, nil
}
