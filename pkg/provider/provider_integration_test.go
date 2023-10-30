package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/snowflakedb/gosnowflake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt_Provider(t *testing.T) {
	user, password, account := os.Getenv("TEST_SNOWFLAKE_USER"), os.Getenv("TEST_SNOWFLAKE_PASSWORD"), os.Getenv("TEST_SNOWFLAKE_ACCOUNT")
	if user == "" || password == "" || account == "" {
		t.Skip("TEST_SNOWFLAKE_USER, TEST_SNOWFLAKE_PASSWORD, and TEST_SNOWFLAKE_ACCOUNT must be set")
	}

	t.Run("auth with username and password", func(t *testing.T) {
		c, err := sdk.NewClient(&gosnowflake.Config{
			User:     user,
			Account:  account,
			Password: password,
		})
		require.NoError(t, err)
		require.NoError(t, c.Ping())
		assert.Equal(t, account, c.GetConfig().Account)
	})

	t.Run("auth with profile", func(t *testing.T) {
		content := fmt.Sprintf(`
[default]
account='%s'
user='%s'
password='%s'
		`, account, user, password)

		path := "/tmp/.snowflake"
		err := os.MkdirAll(path, os.ModePerm)
		require.NoError(t, err)
		t.Cleanup(func() {
			os.RemoveAll(path)
		})
		file := path + "/config"
		err = os.WriteFile(file, []byte(content), os.ModePerm)
		require.NoError(t, err)

		os.Setenv("SNOWFLAKE_CONFIG_PATH", file)
		c, err := sdk.NewDefaultClient()
		require.NoError(t, err)
		require.NoError(t, c.Ping())
	})
}

func TestInt_KeyPair(t *testing.T) {
	data, err := os.ReadFile("/home/longwu/.ssh/snowflake_key")
	require.NoError(t, err)
	testPrivKey, err := parsePrivateKey(data, nil)
	require.NoError(t, err)

	// testPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// require.NoError(t, err)

	c, err := sdk.NewClient(&gosnowflake.Config{
		User:          "TERRAFORM_SVC_ACCOUNT",
		Account:       "YXA30390",
		Authenticator: gosnowflake.AuthTypeJwt,
		PrivateKey:    testPrivKey,
	})
	require.NoError(t, err)
	require.NoError(t, c.Ping())
}
