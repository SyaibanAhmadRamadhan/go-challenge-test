package integration

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"

	"challenge-test-synapsis/repository"
)

var user1 = &repository.User{
	ID:          "rama",
	Username:    "rama",
	Email:       "rama",
	Password:    "rama",
	PhoneNumber: "rama",
	RoleID:      1,
	Audit:       auditDefault,
}
var user2 = repository.User{
	ID:          "iban2",
	Username:    "iban2",
	Email:       "iban2",
	Password:    "iban2",
	PhoneNumber: "iban2",
	RoleID:      1,
	Audit:       auditDefault,
}
var user3 = repository.User{
	ID:          "user3",
	Username:    "user3",
	Email:       "user3",
	Password:    "user3",
	PhoneNumber: "user3",
	RoleID:      1,
	Audit:       auditDefault,
}
var user4 = repository.User{
	ID:          "user4",
	Username:    "user3",
	Email:       "user3",
	Password:    "user3",
	PhoneNumber: "user3",
	RoleID:      1,
	Audit:       auditDefault,
}

func UserRepositoryImplCreate(t *testing.T) {
	err := UserRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
		err := UserRepository.Create(context.Background(), user1)
		err = UserRepository.Create(context.Background(), &user2)
		err = UserRepository.Create(context.Background(), &user3)
		return err
	})
	assert.NoError(t, err)

}

func UserRepositoryImplCheckOne(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               "rama",
			Operator:            repository.Equality,
			NextConditionColumn: "AND",
		},
		{
			Prefix:              "",
			Column:              "username",
			Value:               "rama",
			Operator:            repository.Equality,
			NextConditionColumn: "",
		},
	}

	err := UserRepository.StartTx(context.Background(), pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: "",
		BeginQuery:     "",
	}, func() error {
		res, err := UserRepository.CheckOne(context.Background(), &filters)
		assert.Equal(t, true, res)
		return err
	})

	assert.NoError(t, err)
}

func UserRepositoryImplUpdate(t *testing.T) {
	user1 = &repository.User{
		ID:          "rama",
		Username:    "ibanrama",
		Email:       "tes@gmail.com",
		Password:    "rama123",
		PhoneNumber: "088295007524",
		RoleID:      1,
		Audit: repository.Audit{
			CreatedAt: timeUnix,
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: sql.NullString{},
		},
	}

	err := UserRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
		err := UserRepository.Update(context.Background(), user1)
		return err
	})
	assert.NoError(t, err)
}

func UserRepositoryImplDelete(t *testing.T) {
	err := UserRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
		err := UserRepository.Delete(context.Background(), "user4")
		return err
	})
	assert.NoError(t, err)
}

func UserRepositoryImplFindOne(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               user2.ID,
			Operator:            repository.Equality,
			NextConditionColumn: "AND",
		},
		{
			Prefix:              "",
			Column:              "username",
			Value:               user2.Username,
			Operator:            repository.Equality,
			NextConditionColumn: "",
		},
	}

	t.Run("user2", func(t *testing.T) {
		err := UserRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			user, err := UserRepository.FindOne(context.Background(), &filters)
			assert.Equal(t, &user2, user)
			return err
		})

		assert.NoError(t, err)
	})

	t.Run("after_update_user1", func(t *testing.T) {
		filters[0].Value = user1.ID
		filters[1].Value = "ibanrama"
		err := UserRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			user, err := UserRepository.FindOne(context.Background(), &filters)
			assert.Equal(t, user1, user)
			return err
		})

		assert.NoError(t, err)
	})
}

func UserRepositoryImplCreateError(t *testing.T) {
	user := repository.User{
		ID:          "rama2",
		Username:    "rama2",
		Email:       "rama2",
		Password:    "rama2",
		PhoneNumber: "rama2",
		RoleID:      1,
		Audit: repository.Audit{
			CreatedAt: time.Now().Unix(),
			CreatedBy: "",
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: sql.NullString{},
			DeletedAt: sql.NullInt64{},
			DeletedBy: sql.NullString{},
		},
	}

	t.Run("tx_is_closed", func(t *testing.T) {
		err := UserRepository.Create(context.Background(), &user)
		assert.Error(t, err)
		assert.ErrorIs(t, err, pgx.ErrTxClosed)
	})
}

func UserRepositoryImplUpdateError(t *testing.T) {
	user := repository.User{
		ID:          "rama2",
		Username:    "ibanrama",
		Email:       "ibanrama@gmail.com",
		Password:    "rama123",
		PhoneNumber: "088295007524",
		RoleID:      1,
		Audit:       auditDefault,
	}

	t.Run("tx_is_closed", func(t *testing.T) {
		err := UserRepository.Update(context.Background(), &user)
		assert.Error(t, err)
		assert.ErrorIs(t, err, pgx.ErrTxClosed)
	})

}

func UserRepositoryImplCheckOneError(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               "rama123",
			Operator:            repository.Equality,
			NextConditionColumn: "",
		},
		{
			Prefix:              "",
			Column:              "deleted_at",
			Operator:            repository.IsNULL,
			NextConditionColumn: "",
		},
	}

	t.Run("not_found", func(t *testing.T) {
		err := UserRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			res, err := UserRepository.CheckOne(context.Background(), &filters)
			assert.Equal(t, false, res)
			return err
		})

		assert.NoError(t, err)
	})
	t.Run("after_delete_user4", func(t *testing.T) {
		filters[0].Value = user4.ID
		err := UserRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			res, err := UserRepository.CheckOne(context.Background(), &filters)
			assert.Equal(t, false, res)
			return err
		})

		assert.NoError(t, err)
	})
}

func UserRepositoryImplFindOneError(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               user2.ID,
			Operator:            repository.Equality,
			NextConditionColumn: "AND",
		},
		{
			Prefix:              "",
			Column:              "username",
			Value:               user2.Username,
			Operator:            repository.Equality,
			NextConditionColumn: "",
		},
		{
			Prefix:              "",
			Column:              "deleted_at",
			Operator:            repository.IsNULL,
			NextConditionColumn: "",
		},
	}

	t.Run("user_not_found", func(t *testing.T) {
		filters[1].Value = "asal"
		err := UserRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			user, err := UserRepository.FindOne(context.Background(), &filters)
			assert.Nil(t, user)
			return err
		})

		assert.ErrorIs(t, err, pgx.ErrNoRows)
	})

	t.Run("after_delete_user4", func(t *testing.T) {
		filters[0].Value = user4.ID
		filters[1].Value = user4.Username
		err := UserRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			user, err := UserRepository.FindOne(context.Background(), &filters)
			assert.Nil(t, user)
			return err
		})

		assert.ErrorIs(t, err, pgx.ErrNoRows)
	})
}
