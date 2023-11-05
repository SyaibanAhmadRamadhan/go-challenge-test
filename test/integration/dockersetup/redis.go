package dockersetup

import (
	"context"
	"fmt"
	"log"

	"github.com/ory/dockertest/v3"
	"github.com/redis/go-redis/v9"
)

func RedisContainer(pool *dockertest.Pool) (rc *redis.Client, dr *dockertest.Resource) {
	dr, err := pool.Run("redis", "6.2", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		rc = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("localhost:%s", dr.GetPort("6379/tcp")),
		})

		return rc.Ping(context.TODO()).Err()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return
}
