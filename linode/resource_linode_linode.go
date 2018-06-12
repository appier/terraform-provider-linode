package linode

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLinodeLinode() *schema.Resource {
	return &schema.Resource{
		Create: createLinodeLinode,
		Read:   readLinodeLinode,
		Update: updateLinodeLinode,
		Delete: deleteLinodeLinode,
		Schema: map[string]*schema.Schema{
			"hypervisor": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv4": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
			},
			"stackscript_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"stackscript_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"booted": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"root_pass": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorized_keys": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
			},
			"backup_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"backups_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"swap_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func createLinodeLinode(d *schema.ResourceData, meta interface{}) error {
	return errors.New("Not implemented")
}

func readLinodeLinode(d *schema.ResourceData, meta interface{}) error {
	return errors.New("Not implemented")
}

func updateLinodeLinode(d *schema.ResourceData, meta interface{}) error {
	return errors.New("Not implemented")
}

func deleteLinodeLinode(d *schema.ResourceData, meta interface{}) error {
	return errors.New("Not implemented")
}
