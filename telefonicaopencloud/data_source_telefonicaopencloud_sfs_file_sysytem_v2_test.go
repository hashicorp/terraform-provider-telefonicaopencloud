package telefonicaopencloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccSFSFileSystemV2DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSFSFileSystemV2DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSFSFileSystemV2DataSourceID("data.telefonicaopencloud_sfs_file_system_v2.shares"),
					resource.TestCheckResourceAttr("data.telefonicaopencloud_sfs_file_system_v2.shares", "name", "sfs-c2c-1"),
					resource.TestCheckResourceAttr("data.telefonicaopencloud_sfs_file_system_v2.shares", "status", "available"),
					resource.TestCheckResourceAttr("data.telefonicaopencloud_sfs_file_system_v2.shares", "size", "1"),
				),
			},
		},
	})
}

func testAccCheckSFSFileSystemV2DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find share file data source: %s ", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("share file data source ID not set ")
		}

		return nil
	}
}

var testAccSFSFileSystemV2DataSource_basic = fmt.Sprintf(`
resource "telefonicaopencloud_sfs_file_system_v2" "sfs_1" {
        share_proto = "NFS"
        size=1
        name="sfs-c2c-1"
        availability_zone="%s"
        access_to="%s"
        access_type="cert"
        access_level="rw"
        description="sfs_c2c_test-file"
}
data "telefonicaopencloud_sfs_file_system_v2" "shares" {
  id = "${telefonicaopencloud_sfs_file_system_v2.sfs_1.id}"
}
`, OS_AVAILABILITY_ZONE, OS_VPC_ID)
