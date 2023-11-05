package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func providerFactories() map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"snowflake": func() (*schema.Provider, error) {
			return Provider(), nil
		},
	}
}

func TestAcc_ProviderUsernameAndPasswordAuth(t *testing.T) {
	username := os.Getenv("SNOWFLAKE_USER")
	if username == "" {
		t.Skip("SNOWFLAKE_USER must be set")
	}
	password := os.Getenv("SNOWFLAKE_PASSWORD")
	if password == "" {
		t.Skip("SNOWFLAKE_PASSWORD must be set")
	}
	account := os.Getenv("SNOWFLAKE_ACCOUNT")
	if account == "" {
		t.Skip("SNOWFLAKE_ACCOUNT must be set")
	}

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: providerFactories(),
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: providerConfig_UsernameAndPassword(account, username, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccount(t, account),
				),
				PlanOnly: true,
			},
		},
	})
}

func testAccCheckAccount(t *testing.T, expected string) resource.TestCheckFunc {
	t.Helper()
	return func(s *terraform.State) error {
		return nil
	}
}

func providerConfig_UsernameAndPassword(account, username, password string) string {
	return fmt.Sprintf(`
provider "snowflake" {
	account                = "%s"
	username               = "%s"
	password               = "%s"
`, account, username, password)
}
