package telefonicaopencloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccVBSBackupPolicyV2DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVBSBackupPolicyV2DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVBSBackupPolicyV2DataSource("data.telefonicaopencloud_vbs_backup_policy_v2.policies"),
					resource.TestCheckResourceAttr("data.telefonicaopencloud_vbs_backup_policy_v2.policies", "name", "policy_001"),
					resource.TestCheckResourceAttr("data.telefonicaopencloud_vbs_backup_policy_v2.policies", "status", "ON"),
				),
			},
		},
	})
}

func testAccCheckVBSBackupPolicyV2DataSource(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find backup policy data source: %s ", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("backup policy ID not set ")
		}

		return nil
	}
}

var testAccVBSBackupPolicyV2DataSource_basic = `
resource "telefonicaopencloud_vbs_backup_policy_v2" "vbs_1" {
  name = "policy_001"
  start_time  = "12:00"
  status  = "ON"
  retain_first_backup = "N"
  rentention_num = 2
  frequency = 1
}

data "telefonicaopencloud_vbs_backup_policy_v2" "policies" {
  id = "${telefonicaopencloud_vbs_backup_policy_v2.vbs_1.id}"
}
`
