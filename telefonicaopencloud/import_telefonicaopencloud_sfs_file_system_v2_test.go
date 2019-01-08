package telefonicaopencloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccSFSFileSystemV2_importBasic(t *testing.T) {
	resourceName := "telefonicaopencloud_sfs_file_system_v2.sfs_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSFSFileSystemV2Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSFSFileSystemV2_basic,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
