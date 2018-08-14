package linode

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sethvargo/go-password/password"
)

func resourceLinodeLinode() *schema.Resource {
	return &schema.Resource{
		Create: createLinodeLinode,
		Read:   readLinodeLinode,
		Update: updateLinodeLinode,
		Delete: deleteLinodeLinode,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"stackscript_id": &schema.Schema{
				Type:     schema.TypeInt,
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
				Optional:  true,
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
			"private_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func toLinode(d *schema.ResourceData) (*Linode, error) {
	res := &Linode{}

	if value, ok := d.GetOk("hypervisor"); ok {
		hypervisor := value.(string)
		res.Hypervisor = &hypervisor
	}

	if value, ok := d.GetOk("group"); ok {
		group := value.(string)
		res.Group = &group
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
		ipv6 := value.(string)
		res.IPv6 = &ipv6
	}

	if value, ok := d.GetOk("stackscript_id"); ok {
		stackscriptID := value.(int)
		res.StackscriptID = &stackscriptID
	}

	if value, ok := d.GetOk("stackscript_data"); ok {
		stackscriptData := value.(string)
		res.StackscriptData = new(map[string]string)
		if err := json.Unmarshal([]byte(stackscriptData), res.StackscriptData); err != nil {
			return nil, err
		}
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

	if value, ok := d.GetOk("private_ip"); ok {
		privateIP := value.(bool)
		res.PrivateIP = &privateIP
	}

	return res, nil
}

func (l *Linode) fillResourceData(d *schema.ResourceData) (err error) {
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
		content, err := json.Marshal(l.StackscriptData)
		if err != nil {
			return err
		}
		d.Set("stackscript_id", content)
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

	if l.PrivateIP != nil {
		d.Set("private_ip", *l.PrivateIP)
	}

	return nil
}

func createLinodeLinode(d *schema.ResourceData, meta interface{}) error {
	client := meta.(LinodeClient)

	linode, err := toLinode(d)
	if err != nil {
		return err
	}

	if linode.RootPass == nil {
		log.Printf("Generate random root_pass")
		rootPass, err := password.Generate(64, 10, 10, false, false)
		if err != nil {
			return err
		}

		linode.RootPass = &rootPass
	}

	res := &Linode{}

	// https://developers.linode.com/api/v4#operation/createLinodeInstance
	if err := client.Request("POST", fmt.Sprintf("linode/instances"), linode, res); err != nil {
		return err
	}

	if err := res.fillResourceData(d); err != nil {
		return err
	}

	id := d.Id()

	for {
		// https://developers.linode.com/api/v4#operation/getLinodeInstance
		if err := client.Request("GET", fmt.Sprintf("linode/instances/%s", id), nil, res); err != nil {
			return err
		}

		if *res.Status == "running" {
			break
		}

		time.Sleep(5 * time.Second)
	}

	return nil
}

func readLinodeLinode(d *schema.ResourceData, meta interface{}) error {
	client := meta.(LinodeClient)

	id := d.Id()

	res := &Linode{}

	// https://developers.linode.com/api/v4#operation/getLinodeInstance
	if err := client.Request("GET", fmt.Sprintf("linode/instances/%s", id), nil, res); err != nil {
		return err
	}

	if err := res.fillResourceData(d); err != nil {
		return err
	}

	return nil
}

func updateLinodeLinode(d *schema.ResourceData, meta interface{}) error {
	client := meta.(LinodeClient)

	id := d.Id()

	linode, err := toLinode(d)
	if err != nil {
		return err
	}

	res := &Linode{}

	// https://developers.linode.com/api/v4#operation/updateLinodeInstance
	if err := client.Request("PUT", fmt.Sprintf("linode/instances/%s", id), linode, res); err != nil {
		return err
	}

	if err := res.fillResourceData(d); err != nil {
		return err
	}

	return nil
}

func deleteLinodeLinode(d *schema.ResourceData, meta interface{}) error {
	client := meta.(LinodeClient)

	id := d.Id()

	res := &Linode{}

	// https://developers.linode.com/api/v4#operation/deleteLinodeInstance
	if err := client.Request("DELETE", fmt.Sprintf("linode/instances/%s", id), nil, res); err != nil {
		return err
	}

	return nil
}
