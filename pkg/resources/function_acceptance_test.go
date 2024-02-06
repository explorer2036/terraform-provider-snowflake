package resources_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	acc "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAcc_Function(t *testing.T) {
	if _, ok := os.LookupEnv("SKIP_FUNCTION_TESTS"); ok {
		t.Skip("Skipping TestAcc_Function")
	}

	functName := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))

	expBody1 := "3.141592654::FLOAT"
	expBody2 := "var X=3\nreturn X"
	expBody3 := "select 1, 2\nunion all\nselect 3, 4\n"
	expBody4 := `class CoolFunc {public static String test(int n) {return "hello!";}}`

	resource.Test(t, resource.TestCase{
		Providers:    acc.TestAccProviders(),
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: functionConfig(functName, acc.TestDatabaseName, acc.TestSchemaName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("snowflake_function.test_function", "name", functName),
					resource.TestCheckResourceAttr("snowflake_function.test_function", "comment", "Terraform acceptance test"),
					resource.TestCheckResourceAttr("snowflake_function.test_function", "statement", expBody2),
					resource.TestCheckResourceAttr("snowflake_function.test_function", "arguments.#", "1"),
					resource.TestCheckResourceAttr("snowflake_function.test_function", "arguments.0.name", "ARG1"),
					resource.TestCheckResourceAttr("snowflake_function.test_function", "arguments.0.type", "VARCHAR"),

					resource.TestCheckResourceAttr("snowflake_function.test_function_simple", "name", functName),
					resource.TestCheckResourceAttr("snowflake_function.test_function_simple", "comment", "user-defined function"),
					resource.TestCheckResourceAttr("snowflake_function.test_function_simple", "statement", expBody1),

					resource.TestCheckResourceAttr("snowflake_function.test_function_complex", "name", functName),
					resource.TestCheckResourceAttr("snowflake_function.test_function_complex", "comment", "Table func with 2 args"),
					resource.TestCheckResourceAttr("snowflake_function.test_function_complex", "statement", expBody3),
					resource.TestCheckResourceAttr("snowflake_function.test_function_complex", "arguments.#", "2"),
					resource.TestCheckResourceAttr("snowflake_function.test_function_complex", "arguments.1.name", "ARG2"),
					resource.TestCheckResourceAttr("snowflake_function.test_function_complex", "arguments.1.type", "DATE"),

					resource.TestCheckResourceAttr("snowflake_function.test_function_java", "name", functName),
					resource.TestCheckResourceAttr("snowflake_function.test_function_java", "comment", "Terraform acceptance test for java"),
					resource.TestCheckResourceAttr("snowflake_function.test_function_java", "statement", expBody4),
					resource.TestCheckResourceAttr("snowflake_function.test_function_java", "arguments.#", "1"),
					resource.TestCheckResourceAttr("snowflake_function.test_function_java", "arguments.0.name", "ARG1"),
					resource.TestCheckResourceAttr("snowflake_function.test_function_java", "arguments.0.type", "NUMBER"),
					checkBool("snowflake_function.test_function_java", "is_secure", false), // this is from user_acceptance_test.go

					// TODO: temporarily remove unit tests to allow for urgent release
					// resource.TestCheckResourceAttr("snowflake_function.test_function_python", "name", functName),
					// resource.TestCheckResourceAttr("snowflake_function.test_function_python", "comment", "Terraform acceptance test for python"),
					// resource.TestCheckResourceAttr("snowflake_function.test_function_python", "statement", expBody5),
					// resource.TestCheckResourceAttr("snowflake_function.test_function_python", "arguments.#", "2"),
					// resource.TestCheckResourceAttr("snowflake_function.test_function_python", "arguments.0.name", "ARG1"),
					// resource.TestCheckResourceAttr("snowflake_function.test_function_python", "arguments.0.type", "NUMBER"),
				),
			},
		},
	})
}

func functionConfig(name string, databaseName string, schemaName string) string {
	return fmt.Sprintf(`
	resource "snowflake_function" "test_function_simple" {
		name = "%s"
		database = "%s"
		schema   = "%s"
		return_type = "float"
		statement = "3.141592654::FLOAT"
	}

	resource "snowflake_function" "test_function" {
		name = "%s"
		database = "%s"
		schema   = "%s"
		arguments {
			name = "arg1"
			type = "varchar"
		}
		comment = "Terraform acceptance test"
		return_type = "varchar"
		language = "javascript"
		statement = "var X=3\nreturn X"
	}

	resource "snowflake_function" "test_function_java" {
		name = "%s"
		database = "%s"
		schema   = "%s"
		arguments {
			name = "arg1"
			type = "number"
		}
		comment = "Terraform acceptance test for java"
		return_type = "varchar"
		language = "java"
		handler = "CoolFunc.test"
		statement = "class CoolFunc {public static String test(int n) {return \"hello!\";}}"
	}

	resource "snowflake_function" "test_function_complex" {
		name = "%s"
		database = "%s"
		schema   = "%s"
		arguments {
			name = "arg1"
			type = "varchar"
		}
		arguments {
			name = "arg2"
			type = "DATE"
		}
		comment = "Table func with 2 args"
		return_type = "table (x number, y number)"
		statement = <<EOT
select 1, 2
union all
select 3, 4
EOT
	}
	`, name, databaseName, schemaName, name, databaseName, schemaName, name, databaseName, schemaName, name, databaseName, schemaName)
}

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
					resource.TestCheckResourceAttrSet(resourceName, "return_type"),
					resource.TestCheckResourceAttrSet(resourceName, "statement"),
					resource.TestCheckResourceAttrSet(resourceName, "is_secure"),
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

func TestAcc_Function_SQL(t *testing.T) {
	testAccFunction(t)
}
