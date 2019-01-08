package telefonicaopencloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDcsAZV1DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDcsAZV1DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcsAZV1DataSourceID("data.telefonicaopencloud_dcs_az_v1.az1"),
					resource.TestCheckResourceAttr(
						"data.telefonicaopencloud_dcs_az_v1.az1", "name", "AZ1"),
					resource.TestCheckResourceAttr(
						"data.telefonicaopencloud_dcs_az_v1.az1", "port", "8004"),
					resource.TestCheckResourceAttr(
						"data.telefonicaopencloud_dcs_az_v1.az1", "code", "sa-chile-1a"),
				),
			},
		},
	})
}

func testAccCheckDcsAZV1DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find Dcs az data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Dcs az data source ID not set")
		}

		return nil
	}
}

var testAccDcsAZV1DataSource_basic = fmt.Sprintf(`
data "telefonicaopencloud_dcs_az_v1" "az1" {
name = "AZ1"
port = "8004"
code = "sa-chile-1a"
}
`)
