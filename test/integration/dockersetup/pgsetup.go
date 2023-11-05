package dockersetup

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"challenge-test-synapsis/conf"
)

func PostgresContainer(pool *dockertest.Pool) (*pgxpool.Pool, *dockertest.Resource, string) {
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=root",
			"POSTGRES_USER=root",
			"POSTGRES_DB=postgres",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	// databaseUrl := fmt.Sprintf("postgres://postgres:postgres@%s/voting?sslmode=disable", hostAndPort)
	var url string

	resource.Expire(120)

	pool.MaxWait = 120 * time.Second
	var db *pgxpool.Pool
	if err = pool.Retry(func() error {
		port, _ := strconv.Atoi(strings.Split(hostAndPort, ":")[1])
		pgConf := conf.PostgresConf{
			User:     "root",
			Password: "root",
			Host:     strings.Split(hostAndPort, ":")[0],
			Port:     port,
			DB:       "postgres",
			SSL:      "disable",
		}
		url = pgConf.DBUrl()

		db, err = pgxpool.New(context.Background(), pgConf.DBUrl())
		if err != nil {
			return err
		}

		return db.Ping(context.Background())
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return db, resource, url
}
