package rightscale

import (
	"testing"

	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"rightscale": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = Provider()
}

func testAccPreCheck(t *testing.T) {
	variables := []string{
		"RS_ACCOUNT_NUMBER",
		"RS_HOSTNAME",
		"RS_REFRESH_TOKEN",
	}

	for _, variable := range variables {
		value := os.Getenv(variable)
		if value == "" {
			t.Fatalf("Environment variable '%s' must be set for acceptance tests!", variable)
		}
	}
}

func testCheckRightScaleResourceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(ProviderConfiguration)
		client := config.client
		res, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		href := res.Primary.ID

		exists, err := client.resourceExists(href, config.accountNumber, config.apiHostname)
		if err != nil {
			return fmt.Errorf("could not check existence of %s: %+v", resourceName, err)
		}
		if !exists {
			return fmt.Errorf("does not exist: %s", resourceName)
		}

		return nil
	}
}

func testCheckRightScaleResourcesDestroyed(resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(ProviderConfiguration)
		client := config.client

		for _, res := range s.RootModule().Resources {
			if res.Type != resourceType {
				continue
			}

			resourceName := res.Primary.Attributes["name"]
			href := res.Primary.ID

			exists, err := client.resourceExists(href, config.accountNumber, config.apiHostname)
			if err != nil {
				return fmt.Errorf("could not check existence of %s: %+v", resourceName, err)
			}
			if exists {
				return fmt.Errorf("still exists: %s", resourceName)
			}
		}

		return nil
	}
}
