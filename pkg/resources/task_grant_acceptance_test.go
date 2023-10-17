package resources_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAcc_TaskGrant(t *testing.T) {
	name := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))

	resource.ParallelTest(t, resource.TestCase{
		Providers:    acc.TestAccProviders(),
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: taskGrantConfig(name, 8, normal, "OPERATE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "database_name", acc.TestDatabaseName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "schema_name", acc.TestSchemaName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "task_name", name),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "with_grant_option", "false"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "privilege", "OPERATE"),
					resource.TestCheckResourceAttr("snowflake_warehouse.test", "max_concurrency_level", "8"),
					resource.TestCheckResourceAttr("snowflake_warehouse.test", "statement_timeout_in_seconds", "86400"),
				),
			},
			// UPDATE MAX_CONCURRENCY_LEVEL
			{
				Config: taskGrantConfig(name, 10, normal, "OPERATE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "database_name", acc.TestDatabaseName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "schema_name", acc.TestSchemaName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "task_name", name),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "with_grant_option", "false"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "privilege", "OPERATE"),
					resource.TestCheckResourceAttr("snowflake_warehouse.test", "max_concurrency_level", "10"),
					resource.TestCheckResourceAttr("snowflake_warehouse.test", "statement_timeout_in_seconds", "86400"),
				),
			},
			// UPDATE PRIVILEGE
			{
				Config: taskGrantConfig(name, 10, normal, "ALL PRIVILEGES"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "database_name", acc.TestDatabaseName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "schema_name", acc.TestSchemaName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "task_name", name),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "privilege", "ALL PRIVILEGES"),
				),
			},
			// IMPORT
			{
				ResourceName:      "snowflake_task_grant.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"enable_multiple_grants", // feature flag attribute not defined in Snowflake, can't be imported
				},
			},
		},
	})
}

func TestAcc_TaskGrant_onAll(t *testing.T) {
	name := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))

	resource.ParallelTest(t, resource.TestCase{
		Providers:    acc.TestAccProviders(),
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: taskGrantConfig(name, 8, onAll, "OPERATE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "database_name", acc.TestDatabaseName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "schema_name", acc.TestSchemaName),
					resource.TestCheckNoResourceAttr("snowflake_task_grant.test", "task_name"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "on_all", "true"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "with_grant_option", "false"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "privilege", "OPERATE"),
					resource.TestCheckResourceAttr("snowflake_warehouse.test", "max_concurrency_level", "8"),
					resource.TestCheckResourceAttr("snowflake_warehouse.test", "statement_timeout_in_seconds", "86400"),
				),
			},
			// IMPORT
			{
				ResourceName:      "snowflake_task_grant.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"enable_multiple_grants", // feature flag attribute not defined in Snowflake, can't be imported
				},
			},
		},
	})
}

func TestAcc_TaskGrant_onFuture(t *testing.T) {
	name := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))

	resource.ParallelTest(t, resource.TestCase{
		Providers:    acc.TestAccProviders(),
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: taskGrantConfig(name, 8, onFuture, "OPERATE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "database_name", acc.TestDatabaseName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "schema_name", acc.TestSchemaName),
					resource.TestCheckNoResourceAttr("snowflake_task_grant.test", "task_name"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "on_future", "true"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "with_grant_option", "false"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "privilege", "OPERATE"),
					resource.TestCheckResourceAttr("snowflake_warehouse.test", "max_concurrency_level", "8"),
					resource.TestCheckResourceAttr("snowflake_warehouse.test", "statement_timeout_in_seconds", "86400"),
				),
			},
			// IMPORT
			{
				ResourceName:      "snowflake_task_grant.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"enable_multiple_grants", // feature flag attribute not defined in Snowflake, can't be imported
				},
			},
		},
	})
}

func taskGrantConfig(name string, concurrency int32, grantType grantType, privilege string) string {
	var taskNameConfig string
	switch grantType {
	case normal:
		taskNameConfig = "task_name \t= snowflake_task.test.name"
	case onFuture:
		taskNameConfig = "on_future = true"
	case onAll:
		taskNameConfig = "on_all = true"
	}

	s := `
resource "snowflake_role" "test" {
  name = "%v"
}

resource "snowflake_task" "test" {
  name     	    = "%s"
  database  	= "terraform_test_database"
  schema   	  	= "terraform_test_schema"
  warehouse 	= "terraform_test_warehouse"
  sql_statement = "SHOW FUNCTIONS"
  enabled  	  	= true
  schedule 	  	= "15 MINUTES"
  lifecycle {
    ignore_changes = [session_parameters]
  }
}

resource "snowflake_task_grant" "test" {
  %s
  database_name = "terraform_test_database"
  roles         = [snowflake_role.test.name]
  schema_name   = "terraform_test_schema"
  privilege 	= "%s"
}
`
	return fmt.Sprintf(s, name, name, taskNameConfig, privilege)
}

func TestAcc_TaskOwnershipGrant_onFuture(t *testing.T) {
	name := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	new_name := name + "_NEW"

	resource.ParallelTest(t, resource.TestCase{
		Providers:    acc.TestAccProviders(),
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			// CREATE SCHEMA level FUTURE ownership grant to role <name>
			{
				Config: taskOwnershipGrantConfig(name, onFuture, "OWNERSHIP", name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "database_name", acc.TestDatabaseName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "schema_name", acc.TestSchemaName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "on_future", "true"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "with_grant_option", "false"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "privilege", "OWNERSHIP"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "roles.0", name),
				),
			},
			// UPDATE SCHEMA level FUTURE OWNERSHIP grant to role <new_name>
			{
				Config: taskOwnershipGrantConfig(name, onFuture, "OWNERSHIP", new_name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "database_name", acc.TestDatabaseName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "schema_name", acc.TestSchemaName),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "on_future", "true"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "with_grant_option", "false"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "privilege", "OWNERSHIP"),
					resource.TestCheckResourceAttr("snowflake_task_grant.test", "roles.0", new_name),
				),
			},
			// IMPORT
			{
				ResourceName:      "snowflake_task_grant.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"enable_multiple_grants", // feature flag attribute not defined in Snowflake, can't be imported
				},
			},
		},
	})
}

func taskOwnershipGrantConfig(name string, grantType grantType, privilege string, rolename string) string {
	var taskNameConfig string
	switch grantType {
	case normal:
		taskNameConfig = "task_name \t= snowflake_task.test.name"
	case onFuture:
		taskNameConfig = "on_future = true"
	case onAll:
		taskNameConfig = "on_all = true"
	}

	s := `
resource "snowflake_role" "test" {
  name = "%v"
}

resource "snowflake_role" "test_new" {
	name = "%v_NEW"
  }

resource "snowflake_task_grant" "test" {
  %s
  database_name 	= "terraform_test_database"
  roles             = [ "%s" ]
  schema_name       = "terraform_test_schema"
  privilege 	    = "%s"
  with_grant_option = false
}
`
	return fmt.Sprintf(s, name, name, taskNameConfig, rolename, privilege)
}
