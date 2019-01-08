package telefonicaopencloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccComputeV2SecGroup_importBasic(t *testing.T) {
	resourceName := "telefonicaopencloud_compute_secgroup_v2.sg_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2SecGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeV2SecGroup_basic_orig,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
