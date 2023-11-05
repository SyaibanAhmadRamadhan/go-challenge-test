package integration

import (
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"

	"challenge-test-synapsis/conf"
	"challenge-test-synapsis/infra"
	"challenge-test-synapsis/repository/categoryProduct_repository"
	"challenge-test-synapsis/repository/product_repository"
	"challenge-test-synapsis/repository/redis_repository"
	"challenge-test-synapsis/repository/session_repository"
	"challenge-test-synapsis/repository/uow_repository"
	"challenge-test-synapsis/repository/user_repository"
	"challenge-test-synapsis/test/integration/dockersetup"
	"challenge-test-synapsis/usecase/auth_usecase"
)

func TestMain(m *testing.M) {
	conf.LoadEnv()
	dockerPool := dockersetup.DockerPool()
	pgConn, resourcePg, url := dockersetup.PostgresContainer(dockerPool)
	rcConn, resourceRedis := dockersetup.RedisContainer(dockerPool)
	db = pgConn
	rc = rcConn
	resources := []*dockertest.Resource{
		resourcePg, resourceRedis,
	}

	infra.MigrateMaster("up", "", url)

	UOW = uow_repository.NewUnitOfWorkRepositoryImpl(db)
	UserRepository = user_repository.NewUserRepositoryImpl(UOW)
	SessionRepository = session_repository.NewSessionRepositoryImpl(UOW)
	CategoryProductRepository = categoryProduct_repository.NewCategoryProductRepositoryImpl(UOW)
	ProductRepository = product_repository.NewProductRepositoryImpl(UOW)
	RedisRepository = redis_repository.NewRedisRepositoryImpl(rc)

	AuthUsecase = auth_usecase.NewAuthUsecaseImpl(UserRepository, SessionRepository, RedisRepository)

	code := m.Run()

	for _, resource := range resources {
		if err := dockerPool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resourcePg: %s", err)
		}
	}

	os.Exit(code)
}
