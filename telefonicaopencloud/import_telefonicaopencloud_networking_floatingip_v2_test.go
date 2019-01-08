package telefonicaopencloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccNetworkingV2FloatingIP_importBasic(t *testing.T) {
	resourceName := "telefonicaopencloud_networking_floatingip_v2.fip_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkingV2FloatingIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingV2FloatingIP_basic,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
