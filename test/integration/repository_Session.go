package integration

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"

	"challenge-test-synapsis/helper"
	"challenge-test-synapsis/repository"
)

var session1 = &repository.Session{
	ID:      "1",
	Token:   "token",
	Device:  "device",
	LoginAt: timeUnix,
	IP:      "127.0.0.1",
	Audit:   auditDefault,
	UserID:  user2.ID,
}

var session2 = &repository.Session{
	ID:      "2",
	Token:   "token",
	Device:  "device",
	LoginAt: timeUnix,
	IP:      "127.0.0.1",
	Audit:   auditDefault,
	UserID:  user3.ID,
}

var session3 = &repository.Session{
	ID:      "3",
	Token:   "token",
	Device:  "device",
	LoginAt: timeUnix,
	IP:      "127.0.0.1",
	Audit:   auditDefault,
	UserID:  user3.ID,
}

var sessionExp = &repository.Session{
	ID:      "4",
	Token:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTkxNzExOTksInN1YiI6IjAxSEVGNlM2RTdTOUJaSDBQUUYyNDQzNFE4In0.6kDtcMj6nBZW-5rCrRzg1f3l2NdPJogKvOV2UGzI8js",
	Device:  "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36",
	LoginAt: timeUnix,
	IP:      "127.0.0.1",
	Audit:   auditDefault,
	UserID:  user3.ID,
}

func TestSessionRepositoryImplCreate(t *testing.T) {
	_ = SessionRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
		err := SessionRepository.Create(context.Background(), session1)
		assert.NoError(t, err)

		err = SessionRepository.Create(context.Background(), session2)
		assert.NoError(t, err)

		err = SessionRepository.Create(context.Background(), session3)
		assert.NoError(t, err)

		err = SessionRepository.Create(context.Background(), sessionExp)
		assert.NoError(t, err)
		return nil
	})
}

func TestSessionRepositoryImplUpdate(t *testing.T) {
	session1Update := &repository.Session{
		ID:      "1",
		Token:   "tokenUpdate",
		Device:  "device",
		LoginAt: timeUnix,
		IP:      "127.0.0.1",
		Audit: repository.Audit{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper.NewNullString(user2.ID),
		},
		UserID: user2.ID,
	}

	t.Run("fulfill_the_conditions", func(t *testing.T) {
		_ = SessionRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
			err := SessionRepository.Update(context.Background(), session1Update)
			assert.NoError(t, err)
			return nil
		})
	})

	t.Run("does_not_meet_the_conditions", func(t *testing.T) {
		session1Update.UserID = "asal"
		_ = SessionRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
			err := SessionRepository.Update(context.Background(), session1Update)
			assert.NoError(t, err)
			return nil
		})
	})
}

func TestSessionRepositoryImplDelete(t *testing.T) {
	t.Run("fulfill_the_conditions", func(t *testing.T) {
		_ = SessionRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
			err := SessionRepository.Delete(context.Background(), session3.ID, session3.UserID)
			assert.NoError(t, err)
			return nil
		})
	})

	t.Run("does_not_meet_the_conditions_1", func(t *testing.T) {
		_ = SessionRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
			err := SessionRepository.Delete(context.Background(), "1", "asal")
			assert.NoError(t, err)
			return nil
		})
	})

	t.Run("after_delete_and_delete_again", func(t *testing.T) {
		_ = SessionRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
			err := SessionRepository.Delete(context.Background(), session3.ID, session3.UserID)
			assert.NoError(t, err)
			return nil
		})
	})
}

func TestSessionRepositoryImplCheckOne(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               session1.ID,
			Operator:            repository.Equality,
			NextConditionColumn: "AND",
		},
		{
			Prefix:              "",
			Column:              "user_id",
			Value:               session1.UserID,
			Operator:            repository.Equality,
			NextConditionColumn: "AND",
		},
		{
			Prefix:              "",
			Column:              "token",
			Value:               "tokenUpdate",
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

	err := SessionRepository.StartTx(context.Background(), pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: "",
		BeginQuery:     "",
	}, func() error {
		res, err := SessionRepository.CheckOne(context.Background(), &filters)
		assert.Equal(t, true, res)
		return err
	})

	assert.NoError(t, err)
}

func TestSessionRepositoryImplFindOne(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               session1.ID,
			Operator:            repository.Equality,
			NextConditionColumn: "AND",
		},
		{
			Prefix:              "",
			Column:              "user_id",
			Value:               session1.UserID,
			Operator:            repository.Equality,
			NextConditionColumn: "AND",
		},
		{
			Prefix:              "",
			Column:              "token",
			Value:               "tokenUpdate",
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

	err := SessionRepository.StartTx(context.Background(), pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: "",
		BeginQuery:     "",
	}, func() error {
		session, err := SessionRepository.FindOne(context.Background(), &filters)
		assert.NotNil(t, session)
		assert.Equal(t, session1.ID, session.ID)
		return err
	})

	assert.NoError(t, err)
}

func TestSessionRepositoryImplUpdateError(t *testing.T) {
	t.Run("tx_is_closed", func(t *testing.T) {
		err := SessionRepository.Update(context.Background(), session1)
		assert.ErrorIs(t, err, pgx.ErrTxClosed)
	})
}

func TestSessionRepositoryImplCreateError(t *testing.T) {
	t.Run("tx_is_closed", func(t *testing.T) {
		err := SessionRepository.Create(context.Background(), session1)
		assert.ErrorIs(t, err, pgx.ErrTxClosed)
	})
}

func TestSessionRepositoryImplCheckOneError(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               strconv.Itoa(12345),
			Operator:            repository.Equality,
			NextConditionColumn: "",
		},
	}

	err := SessionRepository.StartTx(context.Background(), pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: "",
		BeginQuery:     "",
	}, func() error {
		res, err := SessionRepository.CheckOne(context.Background(), &filters)
		assert.Equal(t, false, res)
		return err
	})

	assert.NoError(t, err)
}

func TestSessionRepositoryImplFindOneError(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               strconv.Itoa(12345),
			Operator:            repository.Equality,
			NextConditionColumn: "",
		},
	}

	err := SessionRepository.StartTx(context.Background(), pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: "",
		BeginQuery:     "",
	}, func() error {
		session, err := SessionRepository.FindOne(context.Background(), &filters)
		assert.Nil(t, session)
		assert.Equal(t, pgx.ErrNoRows, err)
		return nil
	})

	assert.NoError(t, err)
}
