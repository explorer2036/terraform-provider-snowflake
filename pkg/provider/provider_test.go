package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/snowflakedb/gosnowflake"
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
		if err != nil {
			return nil, fmt.Errorf("new client: %w", err)
		}
		if err := client.Ping(); err != nil {
			return nil, fmt.Errorf("ping: %w", err)
		}
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

		config := sdk.DefaultConfig()
		if profile != "default" {
			pc, err := sdk.ProfileConfig(profile)
			if err != nil {
				return nil, fmt.Errorf("retrieve profile config %s: %w", profile, err)
			}
			if pc == nil {
				return nil, fmt.Errorf("profile with name %s not found", profile)
			}
			config = pc
		}
		client, err := sdk.NewClient(config)
		if err != nil {
			return nil, fmt.Errorf("new client: %w", err)
		}
		if err := client.Ping(); err != nil {
			return nil, fmt.Errorf("ping: %w", err)
		}
		return client, nil
	}
	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("configure: %+v", d)
	}
}
