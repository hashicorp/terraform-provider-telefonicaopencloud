package telefonicaopencloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccVpcSubnetIdsV2DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSubnetIdV2DataSource_vpcsubnet,
			},
			{
				Config: testAccSubnetIdV2DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccSubnetIdV2DataSourceID("data.telefonicaopencloud_vpc_subnet_ids_v1.subnet_ids"),
					resource.TestCheckResourceAttr("data.telefonicaopencloud_vpc_subnet_ids_v1.subnet_ids", "ids.#", "1"),
				),
			},
		},
	})
}
func testAccSubnetIdV2DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find vpc subnet data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Vpc Subnet data source ID not set")
		}

		return nil
	}
}

const testAccSubnetIdV2DataSource_vpcsubnet = `
resource "telefonicaopencloud_vpc_v1" "vpc_1" {
	name = "test_vpc"
	cidr= "192.168.0.0/16"
}

resource "telefonicaopencloud_vpc_subnet_v1" "subnet_1" {
  name = "subnetTest02"
  cidr = "192.168.0.0/24"
  gateway_ip = "192.168.0.1"
  vpc_id = "${telefonicaopencloud_vpc_v1.vpc_1.id}"
}
`

var testAccSubnetIdV2DataSource_basic = fmt.Sprintf(`
%s
data "telefonicaopencloud_vpc_subnet_ids_v1" "subnet_ids" {
  vpc_id = "${telefonicaopencloud_vpc_v1.vpc_1.id}"
}
`, testAccSubnetIdV2DataSource_vpcsubnet)
