package resources_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func testAccProcedure(t *testing.T) {
	t.Helper()

	name := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	resourceName := "snowflake_procedure.p"
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
					resource.TestCheckResourceAttrSet(resourceName, "null_input_behavior"),
					resource.TestCheckResourceAttrSet(resourceName, "return_type"),
					resource.TestCheckResourceAttrSet(resourceName, "statement"),
					resource.TestCheckResourceAttrSet(resourceName, "execute_as"),
					resource.TestCheckResourceAttrSet(resourceName, "secure"),
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

			// test - import - TODO: fix the error: ImportStateVerify attributes not equivalent
			// {
			// 	ConfigDirectory:   config.TestStepDirectory(),
			// 	ConfigVariables:   variableSet2,
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func TestAcc_Procedure_SQL(t *testing.T) {
	testAccProcedure(t)
}

func TestAcc_Procedure_Python(t *testing.T) {
	testAccProcedure(t)
}

func TestAcc_Procedure_Javascript(t *testing.T) {
	testAccProcedure(t)
}

func TestAcc_Procedure_Java(t *testing.T) {
	testAccProcedure(t)
}

func procedureConfig(name string, databaseName string, schemaName string) string {
	return fmt.Sprintf(`
	resource "snowflake_procedure" "test_proc_simple" {
		name = "%s"
		database = "%s"
		schema   = "%s"
		return_type = "varchar"
		language = "javascript"
		statement = <<-EOF
			return "Hi"
		EOF
	}

	resource "snowflake_procedure" "test_proc" {
		name = "%s"
		database = "%s"
		schema   = "%s"
		arguments {
			name = "arg1"
			type = "varchar"
		}
		comment = "Terraform acceptance test"
		language = "javascript"
		return_type = "varchar"
		statement = <<-EOF
			var X=3
			return X
		EOF
	}

	resource "snowflake_procedure" "test_proc_complex" {
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
		comment = "Proc with 2 args"
		return_type = "VARCHAR"
		execute_as = "CALLER"
		null_input_behavior = "RETURNS NULL ON NULL INPUT"
		language = "javascript"
		statement = <<-EOF
			var X=1
			return X
		EOF
	}

	resource "snowflake_procedure" "test_proc_sql" {
		name = "%s_sql"
		database = "%s"
		schema   = "%s"
		language = "SQL"
		return_type         = "INTEGER"
		execute_as          = "CALLER"
		null_input_behavior = "RETURNS NULL ON NULL INPUT"
		statement           = <<EOT
			declare
				x integer;
				y integer;
			begin
				x := 3;
				y := x * x;
				return y;
			end;
		EOT
	  }
	`, name, databaseName, schemaName, name, databaseName, schemaName, name, databaseName, schemaName, name, databaseName, schemaName,
	)
}
