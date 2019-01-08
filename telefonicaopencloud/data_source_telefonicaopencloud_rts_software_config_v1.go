package telefonicaopencloud

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/huaweicloud/golangsdk/openstack/rts/v1/softwareconfig"
)

func dataSourceRtsSoftwareConfigV1() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRtsSoftwareConfigV1Read,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"input_values": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
			"output_values": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
			"config": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"options": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceRtsSoftwareConfigV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	orchestrationClient, err := config.orchestrationV1Client(GetRegion(d, config))

	Config, err := softwareconfig.Get(orchestrationClient, d.Get("id").(string)).Extract()
	if err != nil {
		return fmt.Errorf("Unable to retrieve RTS Software Configs: %s", err)
	}

	log.Printf("[INFO] Retrieved RTS Software Config using given filter %s: %+v", Config.Id, Config)
	d.SetId(Config.Id)

	d.Set("name", Config.Name)
	d.Set("group", Config.Group)
	d.Set("region", GetRegion(d, config))
	d.Set("config", Config.Config)
	d.Set("options", Config.Options)
	if err := d.Set("input_values", Config.Inputs); err != nil {
		return fmt.Errorf("[DEBUG] Error saving inputs to state for RTS Software Config (%s): %s", d.Id(), err)
	}
	if err := d.Set("output_values", Config.Outputs); err != nil {
		return fmt.Errorf("[DEBUG] Error saving outputs to state for RTS Software Config (%s): %s", d.Id(), err)
	}

	return nil
}
