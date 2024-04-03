package testint

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/collections"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestInt_ApplicationRoles setup is a little bit different from usual integration test, because of how native apps work.
// I will try to explain it in a short form, but check out this article for more detailed description (https://docs.snowflake.com/en/developer-guide/native-apps/tutorials/getting-started-tutorial#introduction)
//   - create stage - it is where we will be keeping our application files
//   - put native app specific stuff onto our stage (manifest.yml and setup.sql)
//   - create an application package and a new version of our application
//   - create an application with the application package and the version we just created
//   - while creating the application, the setup.sql script will be run in our application context (and that is where application roles for our tests are created)
//   - we're ready to query application roles we have just created
func TestInt_ApplicationRoles(t *testing.T) {
	client := testClient(t)
	ctx := testContext(t)

	stageName := "stage_" + random.AlphaN(4)
	stage, cleanupStage := createStage(t, client, sdk.NewSchemaObjectIdentifier(TestDatabaseName, TestSchemaName, stageName))
	t.Cleanup(cleanupStage)

	putOnStage(t, client, stage, "manifest.yml")
	putOnStage(t, client, stage, "setup.sql")

	appPackageName := "application_package_" + random.AlphaN(4)
	versionName := "v1"
	cleanupAppPackage := createApplicationPackage(t, client, appPackageName)
	t.Cleanup(cleanupAppPackage)
	addApplicationPackageVersion(t, client, stage, appPackageName, versionName)

	appName := "application_" + random.AlphaN(4)
	cleanupApp := createApplication(t, client, appName, appPackageName, versionName)
	t.Cleanup(cleanupApp)

	assertApplicationRole := func(t *testing.T, appRole *sdk.ApplicationRole, name string, comment string) {
		t.Helper()
		assert.Equal(t, name, appRole.Name)
		assert.Equal(t, appName, appRole.Owner)
		assert.Equal(t, comment, appRole.Comment)
		assert.Equal(t, "APPLICATION", appRole.OwnerRoleType)
	}

	assertApplicationRoles := func(t *testing.T, appRoles []sdk.ApplicationRole, name string, comment string) {
		t.Helper()
		appRole, err := collections.FindOne(appRoles, func(role sdk.ApplicationRole) bool {
			return role.Name == name
		})
		require.NoError(t, err)
		assertApplicationRole(t, appRole, name, comment)
	}

	assertGrantToRoles := func(t *testing.T, grants []sdk.Grant, id sdk.DatabaseObjectIdentifier, grantee sdk.AccountObjectIdentifier, ot sdk.ObjectType) {
		t.Helper()

		grant, err := collections.FindOne(grants, func(grant sdk.Grant) bool {
			return grant.Name.FullyQualifiedName() == id.FullyQualifiedName()
		})
		require.NoError(t, err)
		require.Equal(t, ot, grant.GrantedOn)
		require.Equal(t, grantee.FullyQualifiedName(), grant.GranteeName.FullyQualifiedName())
	}

	t.Run("Show by id", func(t *testing.T) {
		name := "app_role_1"
		id := sdk.NewDatabaseObjectIdentifier(appName, name)

		appRole, err := client.ApplicationRoles.ShowByID(ctx, sdk.NewAccountObjectIdentifier(appName), id)
		require.NoError(t, err)

		assertApplicationRole(t, appRole, name, "some comment")
	})

	t.Run("Show", func(t *testing.T) {
		req := sdk.NewShowApplicationRoleRequest().
			WithApplicationName(sdk.NewAccountObjectIdentifier(appName)).
			WithLimit(&sdk.LimitFrom{
				Rows: sdk.Int(2),
			})
		appRoles, err := client.ApplicationRoles.Show(ctx, req)
		require.NoError(t, err)

		assertApplicationRoles(t, appRoles, "app_role_1", "some comment")
		assertApplicationRoles(t, appRoles, "app_role_2", "some comment2")
	})

	t.Run("Grant and Revoke role: application", func(t *testing.T) {
		role, cleanupRole := createRole(t, client)
		t.Cleanup(cleanupRole)

		name := "app_role_1"
		id := sdk.NewDatabaseObjectIdentifier(appName, name)

		kindOfRole := sdk.NewKindOfRoleRequest().WithRoleName(sdk.Pointer(role.ID()))
		request := sdk.NewGrantApplicationRoleRequest(id).WithTo(*kindOfRole)
		err := client.ApplicationRoles.Grant(ctx, request)
		require.NoError(t, err)

		grants, err := client.Grants.Show(ctx, &sdk.ShowGrantOptions{
			To: &sdk.ShowGrantsTo{
				Role: role.ID(),
			},
		})
		require.NoError(t, err)
		assertGrantToRoles(t, grants, id, role.ID(), sdk.ObjectTypeApplicationRole)
	})
}
