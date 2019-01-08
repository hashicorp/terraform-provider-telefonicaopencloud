package telefonicaopencloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCTSTrackerV1DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCTSTrackerV1DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCTSTrackerV1DataSourceID("data.telefonicaopencloud_cts_tracker_v1.tracker_v1"),
					resource.TestCheckResourceAttr("data.telefonicaopencloud_cts_tracker_v1.tracker_v1", "bucket_name", "tf-test-bucket"),
					resource.TestCheckResourceAttr("data.telefonicaopencloud_cts_tracker_v1.tracker_v1", "status", "enabled"),
				),
			},
		},
	})
}

func testAccCheckCTSTrackerV1DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find cts tracker data source: %s ", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("tracker data source not set ")
		}

		return nil
	}
}

const testAccCTSTrackerV1DataSource_basic = `
resource "telefonicaopencloud_s3_bucket" "bucket" {
  bucket		= "tf-test-bucket"
  acl			= "public-read"
  force_destroy = true
}

resource "telefonicaopencloud_cts_tracker_v1" "tracker_v1" {
  bucket_name		= "${telefonicaopencloud_s3_bucket.bucket.bucket}"
  file_prefix_name  = "yO8Q"
}

data "telefonicaopencloud_cts_tracker_v1" "tracker_v1" {  
  tracker_name = "${telefonicaopencloud_cts_tracker_v1.tracker_v1.id}"
}
`
