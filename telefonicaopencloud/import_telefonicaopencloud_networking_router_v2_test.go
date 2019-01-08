package telefonicaopencloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccNetworkingV2Router_importBasic(t *testing.T) {
	resourceName := "telefonicaopencloud_networking_router_v2.router_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkingV2RouterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingV2Router_basic,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
