package keycloak

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"
)

func TestKeycloakIntegration(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err)
	err = pool.Client.Ping()
	require.NoError(t, err)
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "quay.io/keycloak/keycloak",
		Tag:        "latest",
		Env: []string{
			"KEYCLOAK_ADMIN=admin",
			"KEYCLOAK_ADMIN_PASSWORD=admin",
		},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"8080/tcp": {
				{
					HostIP:   "0.0.0.0",
					HostPort: "8080",
				},
			},
		},
		Mounts: []string{
			os.Getenv("PWD") + "/realm-export.json:/opt/keycloak/data/import/realm-export.json",
		},
		Cmd: []string{
			"start-dev",
			"--http-port=8080",
			"--import-realm",
		},
	}, func(hc *docker.HostConfig) {
		hc.AutoRemove = true
		hc.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	require.NoError(t, err)
	defer func() {
		require.NoError(t, pool.Purge(resource))
	}()

	c, err := sdk.NewDefaultClient()
	require.NoError(t, err)
	sql := `CREATE OR REPLACE USER "test" LOGIN_NAME = 'test@snowflake.com' EMAIL = 'test@snowflake.com'`
	_, err = c.ExecForTests(context.Background(), sql)
	require.NoError(t, err)
	sql = `GRANT ROLE SECURITYADMIN TO USER "test"`
	_, err = c.ExecForTests(context.Background(), sql)
	require.NoError(t, err)
	sql = `
CREATE OR REPLACE SECURITY INTEGRATION "keycloak_saml"
    TYPE = SAML2
    ENABLED = TRUE
    SAML2_ENABLE_SP_INITIATED = TRUE
    SAML2_ISSUER = 'http://localhost:8080/realms/snowflake'
    SAML2_SSO_URL = 'http://localhost:8080/realms/snowflake/protocol/saml'
    SAML2_PROVIDER = 'Custom'
    SAML2_X509_CERT = 'MIICoTCCAYkCBgGMRnSjkTANBgkqhkiG9w0BAQsFADAUMRIwEAYDVQQDDAlzbm93Zmxha2UwHhcNMjMxMjA3MjI0MzE4WhcNMzMxMjA3MjI0NDU4WjAUMRIwEAYDVQQDDAlzbm93Zmxha2UwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCCtHH9Ws9xhYddJmcwhsgj7hxgj5iL8q+1NpsdyPYYTM8lfnfjUEGzfYsZzFpxqMr5l/m9fhvkmmoOlwdo2d8BDIxQNlAZ/ChjthFiaxCr8SiRosk5lK7riylSA6po6iToq9fi4ehV0j66ulFfLcZqeTDIXzO9eLq9YpAmlTaBMr6tmOSlkCCHr8cpDqJLPnN3Vb4mVsHOu5RXVKauqDt7nN1TuO0ZvultIFPHnk7o4Yv83kyegTyNXhO/kXN44mmufpG6kg+h8FbOscp+fAJQto91r42HtsEG2X+qkzzDqpzOxZf7reFtOn6KyTFIFo0N987N5srIva4G0F7kbqHHAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAAiwFjINNHDOkqusHl82+a/XaMP4o6mKBiMdYrodeKpE193NhaHCh2AVlVCxbOcTtaYvBauF1a8q9I59CDn/hCgXV+dtjA3fRgdUXJMnk8Wf81RRPjvLb1VJxgekFwczOChXE5bDmJ7hPyPA7mjrbmJd4q88yL2UucL5meO/Hhyw4ZvJ5+8DkWv9YL+cLZlBAmPNw8CAzK0AQ8pNodPMwrbka88eBE3e4tnWBmrpD8/hWGYgjYssXELYP3zYYEQPoMvvf2NtlYkZr4Wy0A8C36WqeM8GbrRavS3K3T791htuauAxlTnWjT9cTEBViCLHAQdre17aonXkOM0y24RYUMA='
    SAML2_SNOWFLAKE_ACS_URL = 'https://yxa30390.snowflakecomputing.com/fed/login'
    SAML2_SNOWFLAKE_ISSUER_URL = 'https://yxa30390.snowflakecomputing.com'
    SAML2_SP_INITIATED_LOGIN_PAGE_LABEL = 'Keycloak SSO'`
	_, err = c.ExecForTests(context.Background(), sql)
	require.NoError(t, err)

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
	// Wait for Keycloak to be ready
	<-s
}
