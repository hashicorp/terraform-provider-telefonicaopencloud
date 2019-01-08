package telefonicaopencloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccComputeV2VolumeAttach_importBasic(t *testing.T) {
	resourceName := "telefonicaopencloud_compute_volume_attach_v2.va_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2VolumeAttachDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeV2VolumeAttach_basic,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
