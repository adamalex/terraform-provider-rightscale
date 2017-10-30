package rightscale

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccRightScaleDeployment_basic(t *testing.T) {
	resourceName := "rightscale_deployment.deployment"
	rString := acctest.RandString(4)
	basicConfig := testAccRightScaleDeployment_basic(rString)
	updatedConfig := testAccRightScaleDeployment_update(rString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckRightScaleResourcesDestroyed("rightscale_deployment"),
		Steps: []resource.TestStep{
			{
				Config: basicConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckRightScaleResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "created-deployment-"+rString),
					resource.TestCheckResourceAttr(resourceName, "description", "created"),
				),
			},

			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "updated-deployment-"+rString),
					resource.TestCheckResourceAttr(resourceName, "description", "updated"),
				),
			},
		},
	})
}

func testAccRightScaleDeployment_basic(rString string) string {
	return fmt.Sprintf(`
resource "rightscale_deployment" "deployment"  {
	name = "created-deployment-%s"
	description = "created"
}
`, rString)
}

func testAccRightScaleDeployment_update(rString string) string {
	return fmt.Sprintf(`
resource "rightscale_deployment" "deployment"  {
	name = "updated-deployment-%s"
	description = "updated"
}
`, rString)
}
