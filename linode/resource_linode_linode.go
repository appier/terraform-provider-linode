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

func toLinode(d *schema.ResourceData) *Linode {
	res := &Linode{}

	if value, ok := d.GetOk("hypervisor"); ok {
		hypervisor := value.(string)
		res.Hypervisor = &hypervisor
	}

	if value, ok := d.GetOk("label"); ok {
		label := value.(string)
		res.Label = &label
	}

	if value, ok := d.GetOk("region"); ok {
		region := value.(string)
		res.Region = &region
	}

	if value, ok := d.GetOk("type"); ok {
		type_ := value.(string)
		res.Type = &type_
	}

	if value, ok := d.GetOk("status"); ok {
		status := value.(string)
		res.Status = &status
	}

	if value, ok := d.GetOk("ipv4"); ok {
		ipv4 := value.([]string)
		res.IPv4 = &ipv4
	}

	if value, ok := d.GetOk("ipv6"); ok {
		ipv6 := value.([]string)
		res.IPv6 = &ipv6
	}

	if value, ok := d.GetOk("stackscript_id"); ok {
		stackscriptID := value.(string)
		res.StackscriptID = &stackscriptID
	}

	if value, ok := d.GetOk("stackscript_data"); ok {
		stackscriptData := value.(string)
		res.StackscriptData = &stackscriptData
	}

	if value, ok := d.GetOk("booted"); ok {
		booted := value.(bool)
		res.Booted = &booted
	}

	if value, ok := d.GetOk("root_pass"); ok {
		rootPass := value.(string)
		res.RootPass = &rootPass
	}

	if value, ok := d.GetOk("image"); ok {
		image := value.(string)
		res.Image = &image
	}

	if value, ok := d.GetOk("authorized_keys"); ok {
		authorizedKeys := value.([]string)
		res.AuthorizedKeys = &authorizedKeys
	}

	if value, ok := d.GetOk("backup_id"); ok {
		backupID := value.(string)
		res.BackupID = &backupID
	}

	if value, ok := d.GetOk("backups_enableds"); ok {
		backupsEnabled := value.(bool)
		res.BackupsEnabled = &backupsEnabled
	}

	if value, ok := d.GetOk("swap_size"); ok {
		swapSize := value.(int)
		res.SwapSize = &swapSize
	}

	return res
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
