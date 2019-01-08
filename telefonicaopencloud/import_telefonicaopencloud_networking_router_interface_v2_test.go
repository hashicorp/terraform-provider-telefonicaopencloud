package telefonicaopencloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccNetworkingV2RouterInterface_importBasic_port(t *testing.T) {
	resourceName := "telefonicaopencloud_networking_router_interface_v2.int_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkingV2RouterInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingV2RouterInterface_basic_port,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetworkingV2RouterInterface_importBasic_subnet(t *testing.T) {
	resourceName := "telefonicaopencloud_networking_router_interface_v2.int_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkingV2RouterInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingV2RouterInterface_basic_subnet,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
