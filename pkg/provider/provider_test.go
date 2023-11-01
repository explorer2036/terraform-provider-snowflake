package provider

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/snowflakedb/gosnowflake"
	"github.com/stretchr/testify/require"
)

var TestAccProvider *schema.Provider

func TestProvider_impl(t *testing.T) {
	_ = Provider()
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_counts(t *testing.T) {
	// @tombuildsstuff: this is less a unit test and more a useful placeholder tbh
	provider := Provider()
	log.Printf("Data Sources: %d", len(provider.DataSourcesMap))
	log.Printf("Resources:    %d", len(provider.ResourcesMap))
	log.Printf("-----------------")
	log.Printf("Total:        %d", len(provider.ResourcesMap)+len(provider.DataSourcesMap))
}

func TestProvider_clientUsernamePasswordAuth(t *testing.T) {
	if os.Getenv("SNOWFLAKE_USER") == "" {
		t.Skip("SNOWFLAKE_USER must be set")
	}
	if os.Getenv("SNOWFLAKE_PASSWORD") == "" {
		t.Skip("SNOWFLAKE_PASSWORD must be set")
	}
	if os.Getenv("SNOWFLAKE_ACCOUNT") == "" {
		t.Skip("SNOWFLAKE_ACCOUNT must be set")
	}

	logging.SetOutput(t)

	provider := Provider()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	provider.ConfigureFunc = func(rd *schema.ResourceData) (interface{}, error) {
		account := rd.Get("account").(string)
		user := rd.Get("user").(string)
		password := rd.Get("password").(string)

		client, err := sdk.NewClient(&gosnowflake.Config{
			User:     user,
			Account:  account,
			Password: password,
		})
		require.NoError(t, err)
		require.NoError(t, client.Ping())
		return client, nil
	}
	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("configure: %+v", d)
	}
}

func TestProvider_clientProfileAuth(t *testing.T) {
	if os.Getenv("SNOWFLAKE_CONFIG_PATH") == "" {
		t.Skip("SNOWFLAKE_CONFIG_PATH must be set")
	}
	if os.Getenv("SNOWFLAKE_PROFILE") == "" {
		t.Skip("SNOWFLAKE_PROFILE must be set")
	}

	logging.SetOutput(t)

	provider := Provider()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	provider.ConfigureFunc = func(rd *schema.ResourceData) (interface{}, error) {
		profile := rd.Get("profile").(string)

		var config *gosnowflake.Config
		if profile == "default" {
			config = sdk.DefaultConfig()
		} else {
			pc, err := sdk.ProfileConfig(profile)
			require.NoError(t, err)
			require.NotNil(t, pc)
			config = pc
		}
		client, err := sdk.NewClient(config)
		require.NoError(t, err)
		require.NoError(t, client.Ping())
		return client, nil
	}
	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("configure: %+v", d)
	}
}

func userSetRsaPublicKey(t *testing.T, account, user, password, key string) {
	t.Helper()

	client, err := sdk.NewClient(&gosnowflake.Config{
		User:     user,
		Account:  account,
		Password: password,
	})
	require.NoError(t, err)
	t.Cleanup(func() { client.Close() })
	// set the public key on the user
	setOptions := sdk.AlterUserOptions{
		Set: &sdk.UserSet{
			ObjectProperties: &sdk.UserObjectProperties{
				RSAPublicKey: sdk.String(key),
			},
		},
	}
	id := sdk.NewAccountObjectIdentifier(user)
	err = client.Users.Alter(context.Background(), id, &setOptions)
	require.NoError(t, err)
	t.Cleanup(func() {
		// unset the public key on the user
		unsetOptions := sdk.AlterUserOptions{
			Unset: &sdk.UserUnset{
				ObjectProperties: &sdk.UserObjectPropertiesUnset{
					RSAPublicKey: sdk.Bool(true),
				},
			},
		}
		err = client.Users.Alter(context.Background(), id, &unsetOptions)
		require.NoError(t, err)
	})
}

func TestProvider_clientUnencryptedPrivateKeyPathAuth(t *testing.T) {
	if os.Getenv("SNOWFLAKE_USER") == "" {
		t.Skip("SNOWFLAKE_USER must be set")
	}
	if os.Getenv("SNOWFLAKE_PASSWORD") == "" {
		t.Skip("SNOWFLAKE_PASSWORD must be set")
	}
	if os.Getenv("SNOWFLAKE_ACCOUNT") == "" {
		t.Skip("SNOWFLAKE_ACCOUNT must be set")
	}
	if os.Getenv("SNOWFLAKE_PRIVATE_KEY_PATH") == "" {
		t.Skip("SNOWFLAKE_PRIVATE_KEY_PATH must be set")
	}

	logging.SetOutput(t)

	provider := Provider()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	provider.ConfigureFunc = func(rd *schema.ResourceData) (interface{}, error) {
		account := rd.Get("account").(string)
		user := rd.Get("user").(string)
		password := rd.Get("password").(string)

		privateKeyPath := rd.Get("private_key_path").(string)
		privateKey, err := getPrivateKey(privateKeyPath, "", "")
		require.NoError(t, err)
		require.NotNil(t, privateKey)
		publicKey, err := x509.MarshalPKIXPublicKey(privateKey.Public())
		require.NoError(t, err)

		// set the public key on the user
		userSetRsaPublicKey(t, account, user, password, base64.StdEncoding.EncodeToString(publicKey))

		// create a new client with user, account and private key
		client, err := sdk.NewClient(&gosnowflake.Config{
			User:          user,
			Account:       account,
			Authenticator: gosnowflake.AuthTypeJwt,
			PrivateKey:    privateKey,
		})
		require.NoError(t, err)
		require.NoError(t, client.Ping())
		return client, nil
	}
	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("configure: %+v", d)
	}
}

func TestProvider_clientEncryptedPrivateKeyPathAuth(t *testing.T) {
	if os.Getenv("SNOWFLAKE_USER") == "" {
		t.Skip("SNOWFLAKE_USER must be set")
	}
	if os.Getenv("SNOWFLAKE_PASSWORD") == "" {
		t.Skip("SNOWFLAKE_PASSWORD must be set")
	}
	if os.Getenv("SNOWFLAKE_ACCOUNT") == "" {
		t.Skip("SNOWFLAKE_ACCOUNT must be set")
	}
	if os.Getenv("SNOWFLAKE_PRIVATE_KEY_PATH") == "" {
		t.Skip("SNOWFLAKE_PRIVATE_KEY_PATH must be set")
	}
	if os.Getenv("SNOWFLAKE_PRIVATE_KEY_PASSPHRASE") == "" {
		t.Skip("SNOWFLAKE_PRIVATE_KEY_PASSPHRASE must be set")
	}

	logging.SetOutput(t)

	provider := Provider()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	provider.ConfigureFunc = func(rd *schema.ResourceData) (interface{}, error) {
		account := rd.Get("account").(string)
		user := rd.Get("user").(string)
		password := rd.Get("password").(string)

		privateKeyPath := rd.Get("private_key_path").(string)
		privateKeyPassphrase := rd.Get("private_key_passphrase").(string)
		privateKey, err := getPrivateKey(privateKeyPath, "", privateKeyPassphrase)
		require.NoError(t, err)
		require.NotNil(t, privateKey)
		publicKey, err := x509.MarshalPKIXPublicKey(privateKey.Public())
		require.NoError(t, err)

		// set the public key on the user
		userSetRsaPublicKey(t, account, user, password, base64.StdEncoding.EncodeToString(publicKey))

		// create a new client with user, account and private key
		client, err := sdk.NewClient(&gosnowflake.Config{
			User:          user,
			Account:       account,
			Authenticator: gosnowflake.AuthTypeJwt,
			PrivateKey:    privateKey,
		})
		require.NoError(t, err)
		require.NoError(t, client.Ping())
		return client, nil
	}
	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("configure: %+v", d)
	}
}
