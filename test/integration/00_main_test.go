package integration

import (
	"log"
	"os"
	"testing"

	"challenge-test-synapsis/infra"
	"challenge-test-synapsis/repository/categoryProduct_repository"
	"challenge-test-synapsis/repository/session_repository"
	"challenge-test-synapsis/repository/uow_repository"
	"challenge-test-synapsis/repository/user_repository"
	"challenge-test-synapsis/test/integration/dockersetup"
)

func TestMain(m *testing.M) {
	dockerPool := dockersetup.DockerPool()
	pgConn, resource, url := dockersetup.PostgresSetup(dockerPool)

	db = pgConn

	infra.MigrateMaster("up", "", url)

	UOW = uow_repository.NewUnitOfWorkRepositoryImpl(db)
	UserRepository = user_repository.NewUserRepositoryImpl(UOW)
	SessionRepository = session_repository.NewSessionRepositoryImpl(UOW)
	CategoryProductRepository = categoryProduct_repository.NewCategoryProductRepositoryImpl(UOW)

	code := m.Run()

	if err := dockerPool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
