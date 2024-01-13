package testint

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
	"github.com/stretchr/testify/require"
)

func TestInt_Applications(t *testing.T) {
	client := testClient(t)
	ctx := testContext(t)

	databaseTest, schemaTest := testDb(t), testSchema(t)
	tagTest, tagCleanup := createTag(t, client, databaseTest, schemaTest)
	t.Cleanup(tagCleanup)

	cleanupApplicationHandle := func(id sdk.AccountObjectIdentifier) func() {
		return func() {
			err := client.Applications.Drop(ctx, sdk.NewDropApplicationRequest(id))
			if errors.Is(err, sdk.ErrObjectNotExistOrAuthorized) {
				return
			}
			require.NoError(t, err)
		}
	}

	putOnStageHandle := func(t *testing.T, id sdk.SchemaObjectIdentifier, name string, content string) {
		t.Helper()
		tempFile := fmt.Sprintf("/tmp/%s", name)
		f, err := os.Create(tempFile)
		require.NoError(t, err)
		if content != "" {
			_, err = f.Write([]byte(content))
			require.NoError(t, err)
		}
		f.Close()
		defer os.Remove(f.Name())

		_, err = client.ExecForTests(ctx, fmt.Sprintf(`PUT file://%s @%s AUTO_COMPRESS = FALSE OVERWRITE = TRUE`, f.Name(), id.FullyQualifiedName()))
		require.NoError(t, err)
		t.Cleanup(func() {
			_, err = client.ExecForTests(ctx, fmt.Sprintf(`REMOVE @%s/%s`, id.FullyQualifiedName(), name))
			require.NoError(t, err)
		})
	}

	createApplicationPackageHandle := func(t *testing.T, applicationPackageName, version string, patch int, defaultReleaseDirective bool) *sdk.Stage {
		t.Helper()

		stage, cleanupStage := createStage(t, client, databaseTest, schemaTest, "dev_stage_test")
		t.Cleanup(cleanupStage)
		putOnStageHandle(t, stage.ID(), "manifest.yml", "")
		putOnStageHandle(t, stage.ID(), "setup.sql", "CREATE APPLICATION ROLE IF NOT EXISTS APP_HELLO_SNOWFLAKE;")
		cleanupApplicationPackage := createApplicationPackage(t, client, applicationPackageName)
		t.Cleanup(cleanupApplicationPackage)
		addApplicationPackageVersion(t, client, stage, applicationPackageName, version)

		// set default release directive for application package
		if defaultReleaseDirective {
			_, err := client.ExecForTests(ctx, fmt.Sprintf(`ALTER APPLICATION PACKAGE "%s" SET DEFAULT RELEASE DIRECTIVE VERSION = %s PATCH = %d`, applicationPackageName, version, patch))
			require.NoError(t, err)
		}
		return stage
	}

	t.Run("create application: without version", func(t *testing.T) {
		applicationPackageName, version, patch := "hello_snowflake_package_test", "V001", 0
		createApplicationPackageHandle(t, applicationPackageName, version, patch, true)

		id := sdk.NewAccountObjectIdentifier(random.StringN(4))
		pid := sdk.NewAccountObjectIdentifier(applicationPackageName)
		comment := random.StringN(4)
		request := sdk.NewCreateApplicationRequest(id, pid).
			WithComment(&comment).
			WithTag([]sdk.TagAssociation{
				{
					Name:  tagTest.ID(),
					Value: "v1",
				},
			})
		err := client.Applications.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupApplicationHandle(id))

		e, err := client.Applications.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, id.Name(), e.Name)
		require.Equal(t, "ACCOUNTADMIN", e.Owner)
		require.Equal(t, comment, e.Comment)
		require.Equal(t, "APPLICATION PACKAGE", e.SourceType)
		require.Equal(t, applicationPackageName, e.Source)
		require.Equal(t, version, e.Version)
		require.Equal(t, patch, e.Patch)
	})

	t.Run("create application: version and patch", func(t *testing.T) {
		applicationPackageName, version, patch := "hello_snowflake_package_test", "V001", 0
		createApplicationPackageHandle(t, applicationPackageName, version, patch, false)

		id := sdk.NewAccountObjectIdentifier(random.StringN(4))
		pid := sdk.NewAccountObjectIdentifier(applicationPackageName)
		vr := sdk.NewApplicationVersionRequest().WithVersionAndPatch(sdk.NewVersionAndPatchRequest(version, &patch))
		comment := random.StringN(4)
		request := sdk.NewCreateApplicationRequest(id, pid).
			WithDebugMode(sdk.Bool(true)).
			WithComment(&comment).
			WithVersion(vr)
		err := client.Applications.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupApplicationHandle(id))

		e, err := client.Applications.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, id.Name(), e.Name)
		require.Equal(t, "ACCOUNTADMIN", e.Owner)
		require.Equal(t, comment, e.Comment)
		require.Equal(t, "APPLICATION PACKAGE", e.SourceType)
		require.Equal(t, applicationPackageName, e.Source)
		require.Equal(t, version, e.Version)
		require.Equal(t, patch, e.Patch)
	})

	t.Run("create application: version directory", func(t *testing.T) {
		applicationPackageName, version, patch := "hello_snowflake_package_test", "V001", 0
		stage := createApplicationPackageHandle(t, applicationPackageName, version, patch, false)

		id := sdk.NewAccountObjectIdentifier(random.StringN(4))
		pid := sdk.NewAccountObjectIdentifier(applicationPackageName)
		vr := sdk.NewApplicationVersionRequest().WithVersionDirectory(sdk.String("@" + stage.ID().FullyQualifiedName()))
		comment := random.StringN(4)
		request := sdk.NewCreateApplicationRequest(id, pid).
			WithDebugMode(sdk.Bool(true)).
			WithComment(&comment).
			WithVersion(vr).
			WithTag([]sdk.TagAssociation{
				{
					Name:  tagTest.ID(),
					Value: "v1",
				},
			})
		err := client.Applications.Create(ctx, request)
		require.NoError(t, err)
		t.Cleanup(cleanupApplicationHandle(id))

		e, err := client.Applications.ShowByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, id.Name(), e.Name)
		require.Equal(t, "ACCOUNTADMIN", e.Owner)
		require.Equal(t, comment, e.Comment)
		require.Equal(t, "APPLICATION PACKAGE", e.SourceType)
		require.Equal(t, applicationPackageName, e.Source)
		require.Equal(t, "UNVERSIONED", e.Version)
		require.Equal(t, patch, e.Patch)
	})
}
