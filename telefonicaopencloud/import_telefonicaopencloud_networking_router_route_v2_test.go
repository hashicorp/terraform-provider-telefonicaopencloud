package telefonicaopencloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccNetworkingV2RouterRoute_importBasic(t *testing.T) {
	resourceName := "telefonicaopencloud_networking_router_route_v2.router_route_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkingV2RouterRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingV2RouterRoute_create,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
