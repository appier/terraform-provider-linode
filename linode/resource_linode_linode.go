package linode

import (
	"errors"
	"fmt"

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
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

	if value, ok := d.GetOk("group"); ok {
		group := value.(string)
		res.Label = &group
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

func (l *Linode) fillResourceData(d *schema.ResourceData) {
	d.SetId(fmt.Sprintf("%d", *l.ID))

	if l.Hypervisor != nil {
		d.Set("hypervisor", *l.Hypervisor)
	}

	if l.Group != nil {
		d.Set("group", *l.Group)
	}

	if l.Label != nil {
		d.Set("label", *l.Label)
	}

	if l.Region != nil {
		d.Set("region", *l.Region)
	}

	if l.Type != nil {
		d.Set("type", *l.Type)
	}

	if l.Status != nil {
		d.Set("status", *l.Status)
	}

	if l.IPv4 != nil {
		d.Set("ipv4", *l.IPv4)
	}

	if l.IPv6 != nil {
		d.Set("ipv6", *l.IPv6)
	}

	if l.StackscriptID != nil {
		d.Set("stackscript_id", *l.StackscriptID)
	}

	if l.StackscriptData != nil {
		d.Set("stackscript_id", *l.StackscriptData)
	}

	if l.Booted != nil {
		d.Set("booted", *l.Booted)
	}

	if l.RootPass != nil {
		d.Set("root_pass", *l.RootPass)
	}

	if l.Image != nil {
		d.Set("image", *l.Image)
	}

	if l.AuthorizedKeys != nil {
		d.Set("authorized_keys", *l.AuthorizedKeys)
	}

	if l.BackupID != nil {
		d.Set("backup_id", *l.BackupID)
	}

	if l.BackupsEnabled != nil {
		d.Set("backups_enabled", *l.BackupsEnabled)
	}

	if l.SwapSize != nil {
		d.Set("swap_size", *l.SwapSize)
	}
}

func createLinodeLinode(d *schema.ResourceData, meta interface{}) error {
	client := meta.(LinodeClient)

	linode := toLinode(d)

	res := &Linode{}

	// https://developers.linode.com/api/v4#operation/createLinodeInstance
	if err := client.Request("POST", fmt.Sprintf("linode/instances"), linode, res); err != nil {
		return err
	}

	res.fillResourceData(d)

	return nil
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
