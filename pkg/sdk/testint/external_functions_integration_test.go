package testint

import (
	"context"
	"fmt"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/internal/random"
	"github.com/stretchr/testify/require"
)

func TestInt_ExternalFunctions(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	databaseTest, schemaTest := testDb(t), testSchema(t)

	// cleanupExternalFuncionHandle := func(id sdk.SchemaObjectIdentifier, dts []sdk.DataType) func() {
	// 	return func() {
	// 		err := client.Functions.Drop(ctx, sdk.NewDropFunctionRequest(id, dts).WithIfExists(sdk.Bool(true)))
	// 		require.NoError(t, err)
	// 	}
	// }

	createApiIntegrationHandle := func(t *testing.T, id sdk.AccountObjectIdentifier) {
		t.Helper()

		_, err := client.ExecForTests(ctx, fmt.Sprintf(`CREATE API INTEGRATION %s API_PROVIDER = aws_api_gateway API_AWS_ROLE_ARN = 'arn:aws:iam::123456789012:role/hello_cloud_account_role' API_ALLOWED_PREFIXES = ('https://xyz.execute-api.us-west-2.amazonaws.com/production') ENABLED = true`, id.FullyQualifiedName()))
		require.NoError(t, err)
		t.Cleanup(func() {
			_, err = client.ExecForTests(ctx, fmt.Sprintf(`DROP API INTEGRATION %s`, id.FullyQualifiedName()))
			require.NoError(t, err)
		})
	}

	t.Run("create external function", func(t *testing.T) {
		integration := sdk.NewAccountObjectIdentifier(random.AlphaN(4))
		createApiIntegrationHandle(t, integration)

		name := random.StringN(4)
		id := sdk.NewSchemaObjectIdentifier(databaseTest.Name, schemaTest.Name, name)
		argument := sdk.NewExternalFunctionArgumentRequest("x", sdk.DataTypeVARCHAR)

		as := "https://xyz.execute-api.us-west-2.amazonaws.com/production/remote_echo"
		request := sdk.NewCreateExternalFunctionRequest(id, sdk.DataTypeVariant, &integration, as).
			WithOrReplace(sdk.Bool(true)).
			WithArguments([]sdk.ExternalFunctionArgumentRequest{*argument}).
			WithNullInputBehavior(sdk.NullInputBehaviorPointer(sdk.NullInputBehaviorCalledOnNullInput))
	})
}
