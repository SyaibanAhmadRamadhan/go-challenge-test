package integration

import (
	"testing"
)

func TestInit(t *testing.T) {
	defer func() {
		db.Close()
	}()
	t.Run("UserRepositoryImpl", func(t *testing.T) {
		t.Run("Create", UserRepositoryImplCreate)
		t.Run("CheckOne", UserRepositoryImplCheckOne)
		t.Run("Update", UserRepositoryImplUpdate)
		t.Run("Delete", UserRepositoryImplDelete)
		t.Run("FindOne", UserRepositoryImplFindOne)
		t.Run("Create_error", UserRepositoryImplCreateError)
		t.Run("Update_error", UserRepositoryImplUpdateError)
		t.Run("FindOne_error", UserRepositoryImplFindOneError)
		t.Run("CheckOne_error", UserRepositoryImplCheckOneError)
	})

	t.Run("SessionRepositoryImpl", func(t *testing.T) {
		t.Run("Create", TestSessionRepositoryImplCreate)
		t.Run("Update", TestSessionRepositoryImplUpdate)
		t.Run("Delete", TestSessionRepositoryImplDelete)
		t.Run("CheckOne", TestSessionRepositoryImplCheckOne)
		t.Run("FindOne", TestSessionRepositoryImplFindOne)
		t.Run("Create_error", TestSessionRepositoryImplCreateError)
		t.Run("Update_error", TestSessionRepositoryImplUpdateError)
		t.Run("CheckOne_error", TestSessionRepositoryImplCheckOneError)
		t.Run("FindOne_error", TestSessionRepositoryImplFindOneError)
	})

	t.Run("CategoryProductRepositoryImpl", func(t *testing.T) {
		t.Run("Create", CategoryProductRepositoryImplCreate)
		t.Run("CheckOne", CategoryProductRepositoryImplCheckOne)
		t.Run("Update", CategoryProductRepositoryImplUpdate)
		t.Run("Delete", CategoryProductRepositoryImplDelete)
		t.Run("FindOne", CategoryProductRepositoryImplFindOne)
		t.Run("FindAll", CategoryProductRepositoryImplFindAll)
		t.Run("Create_error", CategoryProductRepositoryImplCreateError)
		t.Run("Update_error", CategoryProductRepositoryImplUpdateError)
		t.Run("FindOne_error", CategoryProductRepositoryImplFindOneError)
		t.Run("CheckOne_error", CategoryProductRepositoryImplCheckOneError)
	})

}
