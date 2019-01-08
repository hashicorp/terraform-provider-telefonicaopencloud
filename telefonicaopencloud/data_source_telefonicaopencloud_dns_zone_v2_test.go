package telefonicaopencloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var zoneName = fmt.Sprintf("ACPTTEST%s.com.", acctest.RandString(5))

func TestAccTelefonicaOpenCloudDNSZoneV2DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckDNS(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTelefonicaOpenCloudDNSZoneV2DataSource_zone,
			},
			{
				Config: testAccTelefonicaOpenCloudDNSZoneV2DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDNSZoneV2DataSourceID("data.telefonicaopencloud_dns_zone_v2.z1"),
					resource.TestCheckResourceAttr(
						"data.telefonicaopencloud_dns_zone_v2.z1", "name", zoneName),
					resource.TestCheckResourceAttr(
						"data.telefonicaopencloud_dns_zone_v2.z1", "type", "PRIMARY"),
					resource.TestCheckResourceAttr(
						"data.telefonicaopencloud_dns_zone_v2.z1", "ttl", "7200"),
				),
			},
		},
	})
}

func testAccCheckDNSZoneV2DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find DNS Zone data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DNS Zone data source ID not set")
		}

		return nil
	}
}

var testAccTelefonicaOpenCloudDNSZoneV2DataSource_zone = fmt.Sprintf(`
resource "telefonicaopencloud_dns_zone_v2" "z1" {
  name = "%s"
  email = "terraform-dns-zone-v2-test-name@example.com"
  type = "PRIMARY"
  ttl = 7200
}`, zoneName)

var testAccTelefonicaOpenCloudDNSZoneV2DataSource_basic = fmt.Sprintf(`
%s
data "telefonicaopencloud_dns_zone_v2" "z1" {
	name = "${telefonicaopencloud_dns_zone_v2.z1.name}"
}
`, testAccTelefonicaOpenCloudDNSZoneV2DataSource_zone)
