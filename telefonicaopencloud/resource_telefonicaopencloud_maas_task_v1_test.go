package telefonicaopencloud

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/huaweicloud/golangsdk/openstack/maas/v1/task"
)

func TestAccMaasTask_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckMaas(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMaasTaskV1Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMaasTaskV1_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMaasTaskV1Exists("telefonicaopencloud_maas_task_v1.task_1"),
					resource.TestCheckResourceAttr("telefonicaopencloud_maas_task_v1.task_1", "description", "migration task"),
				),
			},
		},
	})
}

func testAccCheckMaasTaskV1Destroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	maasClient, err := config.maasV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud maas client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "telefonicaopencloud_maas_task_v1" {
			continue
		}

		_, err := task.Get(maasClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Maas task still exists")
		}
	}

	return nil
}

func testAccCheckMaasTaskV1Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		maasClient, err := config.maasV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating TelefonicaOpenCloud maas client: %s", err)
		}

		found, err := task.Get(maasClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if strconv.FormatInt(found.ID, 10) != rs.Primary.ID {
			return fmt.Errorf("Task not found")
		}

		return nil
	}
}

var testAccMaasTaskV1_basic = fmt.Sprintf(`
resource "telefonicaopencloud_maas_task_v1" "task_1" {
  description = "migration task"
  enable_kms = false
  thread_num = 1
  src_node {
    region = "ap-northeast-1"
	ak = "%s"
	sk = "%s"
    object_key = "123.txt"
    bucket = "tommy-bucket"
  }
  dst_node {
    region = "eu-de"
	ak = "%s"
	sk = "%s"
    object_key = "maas/"
    bucket = "test-maas"
  }
}
`, OS_SRC_ACCESS_KEY, OS_SRC_SECRET_KEY, OS_ACCESS_KEY, OS_SECRET_KEY)
