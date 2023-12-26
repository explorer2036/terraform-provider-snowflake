package keycloak

import (
	"testing"

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
		ExposedPorts: []string{"8080"},
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

	// hostAndPort := resource.GetHostPort("8080/tcp")

	t.Skip("Skipping keycloak integration test")
}
