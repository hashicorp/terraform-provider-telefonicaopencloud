package telefonicaopencloud

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/huaweicloud/golangsdk/openstack/networking/v2/extensions/elb/healthcheck"
)

const nameELBHC = "ELB-HealthCheck"

func resourceELBHealthCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceELBHealthCheckCreate,
		Read:   resourceELBHealthCheckRead,
		Update: resourceELBHealthCheckUpdate,
		Delete: resourceELBHealthCheckDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"listener_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"healthcheck_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"healthcheck_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"healthcheck_connect_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"healthy_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"unhealthy_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"healthcheck_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"healthcheck_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceELBHealthCheckCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	var createOpts healthcheck.CreateOpts
	_, err = buildCreateParam(&createOpts, d, nil)
	if err != nil {
		return fmt.Errorf("Error creating %s: building parameter failed:%s", nameELBHC, err)
	}
	log.Printf("[DEBUG] Create %s Options: %#v", nameELBHC, createOpts)

	hc, err := healthcheck.Create(networkingClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating %s: %s", nameELBHC, err)
	}
	log.Printf("[DEBUG] Create %s: %#v", nameELBHC, *hc)

	// If all has been successful, set the ID on the resource
	d.SetId(hc.ID)

	return resourceELBHealthCheckRead(d, meta)
}

func resourceELBHealthCheckRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	hc, err := healthcheck.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "healthcheck")
	}
	log.Printf("[DEBUG] Retrieved %s %s: %#v", nameELBHC, d.Id(), hc)

	return refreshResourceData(hc, d, nil)
}

func resourceELBHealthCheckUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	hcId := d.Id()

	var updateOpts healthcheck.UpdateOpts
	_, err = buildUpdateParam(&updateOpts, d, nil)
	if err != nil {
		return fmt.Errorf("Error updating %s(%s): building parameter failed:%s", nameELBHC, hcId, err)
	}
	b, err := updateOpts.IsNeedUpdate()
	if err != nil {
		return err
	}
	if !b {
		log.Printf("[INFO] Updating %s %s with no changes", nameELBHC, hcId)
		return nil
	}
	log.Printf("[DEBUG] Updating healthcheck %s(%s) with options: %#v", nameELBHC, hcId, updateOpts)

	timeout := d.Timeout(schema.TimeoutUpdate)
	err = resource.Retry(timeout, func() *resource.RetryError {
		_, err := healthcheck.Update(networkingClient, hcId, updateOpts).Extract()
		if err != nil {
			return checkForRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Error updating %s(%s): %s", nameELBHC, hcId, err)
	}

	return resourceELBHealthCheckRead(d, meta)
}

func resourceELBHealthCheckDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	hcId := d.Id()
	log.Printf("[DEBUG] Deleting %s %s", nameELBHC, hcId)

	timeout := d.Timeout(schema.TimeoutDelete)
	err = resource.Retry(timeout, func() *resource.RetryError {
		err := healthcheck.Delete(networkingClient, hcId).ExtractErr()
		if err != nil {
			return checkForRetryableError(err)
		}
		return nil
	})
	if err != nil {
		if isResourceNotFound(err) {
			log.Printf("[INFO] deleting an unavailable %s: %s", nameELBHC, hcId)
			return nil
		}
		return fmt.Errorf("Error deleting %s(%s): %s", nameELBHC, hcId, err)
	}

	return nil
}
