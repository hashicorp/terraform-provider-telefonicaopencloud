package telefonicaopencloud

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/vbs/v2/policies"
)

func resourceVBSBackupPolicyV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceVBSBackupPolicyV2Create,
		Read:   resourceVBSBackupPolicyV2Read,
		Update: resourceVBSBackupPolicyV2Update,
		Delete: resourceVBSBackupPolicyV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateVBSPolicyName,
			},

			"start_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"frequency": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateVBSPolicyFrequency,
			},
			"rentention_num": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateVBSPolicyRetentionNum,
			},
			"retain_first_backup": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateVBSPolicyRetainBackup,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateVBSPolicyStatus,
			},
			"policy_resource_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceVBSBackupPolicyV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vbsClient, err := config.vbsV2Client(GetRegion(d, config))

	if err != nil {
		return fmt.Errorf("Error creating Telefonicaopencloud VBS client: %s", err)
	}

	createOpts := policies.CreateOpts{
		Name: d.Get("name").(string),
		ScheduledPolicy: policies.ScheduledPolicy{
			StartTime:         d.Get("start_time").(string),
			Frequency:         d.Get("frequency").(int),
			RententionNum:     d.Get("rentention_num").(int),
			RemainFirstBackup: d.Get("retain_first_backup").(string),
			Status:            d.Get("status").(string),
		},
	}

	create, err := policies.Create(vbsClient, createOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error creating Telefonicaopencloud Backup Policy: %s", err)
	}
	d.SetId(create.ID)

	log.Printf("[DEBUG] Waiting for Telefonicaopencloud Backup Policy (%s) to become available", d.Id())

	return resourceVBSBackupPolicyV2Read(d, meta)

}

func resourceVBSBackupPolicyV2Read(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)
	vbsClient, err := config.vbsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating Telefonicaopencloud VBS client: %s", err)
	}

	PolicyOpts := policies.ListOpts{ID: d.Id()}
	policies, err := policies.List(vbsClient, PolicyOpts)
	if err != nil {
		if _, ok := err.(golangsdk.ErrDefault404); ok {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error retrieving Telefonicaopencloud Backup Policy: %s", err)
	}

	n := policies[0]

	d.Set("name", n.Name)
	d.Set("start_time", n.ScheduledPolicy.StartTime)
	d.Set("frequency", n.ScheduledPolicy.Frequency)
	d.Set("rentention_num", n.ScheduledPolicy.RententionNum)
	d.Set("retain_first_backup", n.ScheduledPolicy.RemainFirstBackup)
	d.Set("status", n.ScheduledPolicy.Status)
	d.Set("policy_resource_count", n.ResourceCount)

	return nil
}

func resourceVBSBackupPolicyV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vbsClient, err := config.vbsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error updating Telefonicaopencloud VBS client: %s", err)
	}
	var updateOpts policies.UpdateOpts

	if d.HasChange("name") {
		updateOpts.Name = d.Get("name").(string)
	}
	if d.HasChange("start_time") {
		updateOpts.ScheduledPolicy.StartTime = d.Get("start_time").(string)
	}
	if d.HasChange("frequency") {
		updateOpts.ScheduledPolicy.Frequency = d.Get("frequency").(int)
	}
	if d.HasChange("rentention_num") {
		updateOpts.ScheduledPolicy.RententionNum = d.Get("rentention_num").(int)
	}
	if d.HasChange("retain_first_backup") {
		updateOpts.ScheduledPolicy.RemainFirstBackup = d.Get("retain_first_backup").(string)
	}
	if d.HasChange("status") {
		updateOpts.ScheduledPolicy.Status = d.Get("status").(string)
	}
	_, err = policies.Update(vbsClient, d.Id(), updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating Telefonicaopencloud backup policy: %s", err)
	}

	return resourceVBSBackupPolicyV2Read(d, meta)
}

func resourceVBSBackupPolicyV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vbsClient, err := config.vbsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating Telefonicaopencloud VBS client: %s", err)
	}

	delete := policies.Delete(vbsClient, d.Id())
	if delete.Err != nil {
		if _, ok := err.(golangsdk.ErrDefault404); ok {
			log.Printf("[INFO] Successfully deleted Telefonicaopencloud VBS Backup Policy %s", d.Id())

		}
		if errCode, ok := err.(golangsdk.ErrUnexpectedResponseCode); ok {
			if errCode.Actual == 409 {
				log.Printf("[INFO] Error deleting Telefonicaopencloud VBS Backup Policy %s", d.Id())
			}
		}
		log.Printf("[INFO] Successfully deleted Telefonicaopencloud VBS Backup Policy %s", d.Id())
	}

	d.SetId("")
	return nil
}
