package suites

import (
	"fmt"
	"time"
)

var mariadbSuiteName = "Mariadb"

func init() {
	dockerEnvironment := NewDockerEnvironment([]string{
		"internal/suites/docker-compose.yml",
		"internal/suites/Mariadb/docker-compose.yml",
		"internal/suites/example/compose/authelia/docker-compose.backend.{}.yml",
		"internal/suites/example/compose/authelia/docker-compose.frontend.{}.yml",
		"internal/suites/example/compose/nginx/backend/docker-compose.yml",
		"internal/suites/example/compose/nginx/portal/docker-compose.yml",
		"internal/suites/example/compose/smtp/docker-compose.yml",
		"internal/suites/example/compose/mariadb/docker-compose.yml",
		"internal/suites/example/compose/ldap/docker-compose.yml",
	})

	setup := func(suitePath string) error {
		if err := dockerEnvironment.Up(); err != nil {
			return err
		}

		return waitUntilAutheliaBackendIsReady(dockerEnvironment)
	}

	onSetupTimeout := func() error {
		backendLogs, err := dockerEnvironment.Logs("authelia-backend", nil)
		if err != nil {
			return err
		}
		fmt.Println(backendLogs)

		frontendLogs, err := dockerEnvironment.Logs("authelia-frontend", nil)
		if err != nil {
			return err
		}
		fmt.Println(frontendLogs)
		return nil
	}

	teardown := func(suitePath string) error {
		err := dockerEnvironment.Down()
		return err
	}

	GlobalRegistry.Register(mariadbSuiteName, Suite{
		SetUp:           setup,
		SetUpTimeout:    5 * time.Minute,
		OnSetupTimeout:  onSetupTimeout,
		TearDown:        teardown,
		TearDownTimeout: 2 * time.Minute,
	})
}
