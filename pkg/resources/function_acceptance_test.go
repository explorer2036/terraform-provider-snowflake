package resources_test

import (
	"strings"
	"testing"

	acc "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func testAccFunction(t *testing.T) {
	t.Helper()

	name := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	resourceName := "snowflake_function.f"
	m := func() map[string]config.Variable {
		return map[string]config.Variable{
			"name":     config.StringVariable(name),
			"database": config.StringVariable(acc.TestDatabaseName),
			"schema":   config.StringVariable(acc.TestSchemaName),
			"comment":  config.StringVariable("Terraform acceptance test"),
		}
	}
	variableSet2 := m()
	variableSet2["comment"] = config.StringVariable("Terraform acceptance test - updated")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(tfversion.Version1_5_0),
		},
		CheckDestroy: testAccCheckDynamicTableDestroy,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.TestStepDirectory(),
				ConfigVariables: m(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "database", acc.TestDatabaseName),
					resource.TestCheckResourceAttr(resourceName, "schema", acc.TestSchemaName),
					resource.TestCheckResourceAttr(resourceName, "comment", "Terraform acceptance test"),

					// computed attributes
					// resource.TestCheckResourceAttrSet(resourceName, "return_type"),
					// resource.TestCheckResourceAttrSet(resourceName, "statement"),
					// resource.TestCheckResourceAttrSet(resourceName, "is_secure"),
				),
			},

			// test - change comment
			{
				ConfigDirectory: config.TestStepDirectory(),
				ConfigVariables: variableSet2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "database", acc.TestDatabaseName),
					resource.TestCheckResourceAttr(resourceName, "schema", acc.TestSchemaName),
					resource.TestCheckResourceAttr(resourceName, "comment", "Terraform acceptance test - updated"),
				),
			},

			// test - import
			{
				ConfigDirectory:   config.TestStepDirectory(),
				ConfigVariables:   variableSet2,
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAcc_Function_Javascript(t *testing.T) {
	testAccFunction(t)
}
