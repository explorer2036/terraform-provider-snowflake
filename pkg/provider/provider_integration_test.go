package provider

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
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

	t.Run("auth with key pair", func(t *testing.T) {
		key, err := rsa.GenerateKey(rand.Reader, 2048)
		require.NoError(t, err)
		publicKey, err := x509.MarshalPKIXPublicKey(key.Public())
		require.NoError(t, err)
		c1, err := sdk.NewClient(&gosnowflake.Config{
			User:     user,
			Account:  account,
			Password: password,
		})
		require.NoError(t, err)
		encoded := base64.StdEncoding.EncodeToString(publicKey)
		setOptions := sdk.AlterUserOptions{
			Set: &sdk.UserSet{
				ObjectProperties: &sdk.UserObjectProperties{
					RSAPublicKey: sdk.String(encoded),
				},
			},
		}
		id := sdk.NewAccountObjectIdentifier(user)
		err = c1.Users.Alter(context.Background(), id, &setOptions)
		require.NoError(t, err)

		u, err := c1.Users.Describe(context.Background(), id)
		require.NoError(t, err)
		require.Equal(t, encoded, u.RsaPublicKey.Value)

		c2, err := sdk.NewClient(&gosnowflake.Config{
			User:          user,
			Account:       account,
			Authenticator: gosnowflake.AuthTypeJwt,
			PrivateKey:    key,
		})
		require.NoError(t, err)
		require.NoError(t, c2.Ping())

		unsetOptions := sdk.AlterUserOptions{
			Unset: &sdk.UserUnset{
				ObjectProperties: &sdk.UserObjectPropertiesUnset{
					RSAPublicKey: sdk.Bool(true),
				},
			},
		}
		err = c1.Users.Alter(context.Background(), id, &unsetOptions)
		require.NoError(t, err)

		u, err = c1.Users.Describe(context.Background(), id)
		require.NoError(t, err)
		require.Equal(t, "", u.RsaPublicKey.Value)
	})
}
