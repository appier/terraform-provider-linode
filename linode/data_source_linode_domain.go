package linode

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceLinodeDomain() *schema.Resource {
	return &schema.Resource{
		Read: readLinodeDomain,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func readLinodeDomain(d *schema.ResourceData, meta interface{}) error {
	var err error

	client := meta.(LinodeClient)

	domain := d.Get("domain")

	for page := 1; ; page++ {
		res := &Domains{}
		// https://developers.linode.com/api/v4#operation/getDomains
		err = client.Request("GET", fmt.Sprintf("domains?page=%d", page), nil, res)
		if err != nil {
			return err
		}

		for _, v := range res.Data {
			if v.Domain == domain {
				d.SetId(fmt.Sprintf("%d", v.ID))
				return nil
			}
		}

		if res.Pages == page {
			break
		}
	}

	return fmt.Errorf("Cannot find domain `%s`", domain)
}
