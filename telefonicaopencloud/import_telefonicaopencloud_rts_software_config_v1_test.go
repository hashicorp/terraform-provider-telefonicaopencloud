package telefonicaopencloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccRtsSoftwareConfigV1_importBasic(t *testing.T) {
	resourceName := "telefonicaopencloud_rts_software_config_v1.config_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRtsConfigV1Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRtsSoftwareConfigV1_basic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
