package telefonicaopencloud

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/huaweicloud/golangsdk/openstack/dms/v1/queues"
)

func resourceDmsQueuesV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceDmsQueuesV1Create,
		Read:   resourceDmsQueuesV1Read,
		Delete: resourceDmsQueuesV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"queue_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"redrive_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"max_consume_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"retention_hours": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"reservation": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_msg_size_byte": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"produced_messages": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"group_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceDmsQueuesV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	dmsV1Client, err := config.dmsV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating Telefonicaopencloud dms queue client: %s", err)
	}

	description := d.Get("description").(string)
	if description == "" {
		description = " "
	}
	createOpts := &queues.CreateOps{
		Name:            d.Get("name").(string),
		QueueMode:       d.Get("queue_mode").(string),
		Description:     description,
		RedrivePolicy:   d.Get("redrive_policy").(string),
		MaxConsumeCount: d.Get("max_consume_count").(int),
		RetentionHours:  d.Get("retention_hours").(int),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	v, err := queues.Create(dmsV1Client, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating Telefonicaopencloud queue: %s", err)
	}
	log.Printf("[INFO] Queue ID: %s", v.ID)

	// Store the queue ID now
	d.SetId(v.ID)

	return resourceDmsQueuesV1Read(d, meta)
}

func resourceDmsQueuesV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	dmsV1Client, err := config.dmsV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating Telefonicaopencloud dms queue client: %s", err)
	}
	v, err := queues.Get(dmsV1Client, d.Id(), true).Extract()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Dms queue %s: %+v", d.Id(), v)

	d.SetId(v.ID)
	d.Set("name", v.Name)
	d.Set("created", v.Created)
	d.Set("description", v.Description)
	d.Set("queue_mode", v.QueueMode)
	d.Set("reservation", v.Reservation)
	d.Set("max_msg_size_byte", v.MaxMsgSizeByte)
	d.Set("produced_messages", v.ProducedMessages)
	d.Set("redrive_policy", v.RedrivePolicy)
	d.Set("max_consume_count", v.MaxConsumeCount)
	d.Set("group_count", v.GroupCount)

	return nil
}

func resourceDmsQueuesV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	dmsV1Client, err := config.dmsV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating Telefonicaopencloud dms queue client: %s", err)
	}

	v, err := queues.Get(dmsV1Client, d.Id(), false).Extract()
	if err != nil {
		return CheckDeleted(d, err, "queue")
	}

	err = queues.Delete(dmsV1Client, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting Telefonicaopencloud queue: %s", err)
	}

	log.Printf("[DEBUG] Dms queue %s: %+v deactivated.", d.Id(), v)
	d.SetId("")
	return nil
}
