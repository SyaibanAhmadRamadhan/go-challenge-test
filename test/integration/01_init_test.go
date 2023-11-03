package integration

import (
	"testing"
)

func TestInit(t *testing.T) {
	defer func() {
		db.Close()
	}()
	t.Run("UserRepositoryImplCreate", func(t *testing.T) {
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

}
