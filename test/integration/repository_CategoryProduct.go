package integration

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"

	"challenge-test-synapsis/helper"
	"challenge-test-synapsis/repository"
)

var categoryProduct1 = &repository.CategoryProduct{
	ID:    "cp1",
	Name:  "cp1",
	Audit: auditDefault,
}
var categoryProduct2 = &repository.CategoryProduct{
	ID:    "cp2",
	Name:  "cp2",
	Audit: auditDefault,
}
var categoryProduct3 = &repository.CategoryProduct{
	ID:    "cp3",
	Name:  "cp3",
	Audit: auditDefault,
}
var categoryProduct4 = &repository.CategoryProduct{
	ID:    "cp46",
	Name:  "cp46",
	Audit: auditDefault,
}
var categoryProduct5 = &repository.CategoryProduct{
	ID:    "cp56",
	Name:  "cp56",
	Audit: auditDefault,
}
var categoryProduct6 = &repository.CategoryProduct{
	ID:    "cp6",
	Name:  "cp6",
	Audit: auditDefault,
}

func CategoryProductRepositoryImplCreate(t *testing.T) {
	err := CategoryProductRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
		err := CategoryProductRepository.Create(context.Background(), categoryProduct1)
		err = CategoryProductRepository.Create(context.Background(), categoryProduct2)
		err = CategoryProductRepository.Create(context.Background(), categoryProduct3)
		err = CategoryProductRepository.Create(context.Background(), categoryProduct4)
		err = CategoryProductRepository.Create(context.Background(), categoryProduct5)
		err = CategoryProductRepository.Create(context.Background(), categoryProduct6)
		return err
	})
	assert.NoError(t, err)

}

func CategoryProductRepositoryImplCheckOne(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               "cp1",
			Operator:            repository.Equality,
			NextConditionColumn: "",
		},
	}

	err := CategoryProductRepository.StartTx(context.Background(), pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: "",
		BeginQuery:     "",
	}, func() error {
		res, err := CategoryProductRepository.CheckOne(context.Background(), &filters)
		assert.Equal(t, true, res)
		return err
	})

	assert.NoError(t, err)
}

func CategoryProductRepositoryImplUpdate(t *testing.T) {
	categoryProduct1 = &repository.CategoryProduct{
		ID:   "cp1",
		Name: "baju",
		Audit: repository.Audit{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper.NewNullString("user id"),
		},
	}

	err := CategoryProductRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
		err := CategoryProductRepository.Update(context.Background(), categoryProduct1)
		return err
	})
	assert.NoError(t, err)
}

func CategoryProductRepositoryImplDelete(t *testing.T) {
	err := CategoryProductRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
		err := CategoryProductRepository.Delete(context.Background(), "cp3", "admin1")
		return err
	})
	assert.NoError(t, err)
}

func CategoryProductRepositoryImplFindOne(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               categoryProduct2.ID,
			Operator:            repository.Equality,
			NextConditionColumn: "",
		},
	}

	t.Run("categoryProduct2", func(t *testing.T) {
		err := CategoryProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			categoryProduct, err := CategoryProductRepository.FindOne(context.Background(), &filters)
			assert.Equal(t, categoryProduct2, categoryProduct)
			return err
		})

		assert.NoError(t, err)
	})

	t.Run("after_update_categoryProduct1", func(t *testing.T) {
		filters[0].Value = categoryProduct1.ID
		err := CategoryProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			categoryProduct, err := CategoryProductRepository.FindOne(context.Background(), &filters)
			assert.Equal(t, "baju", categoryProduct.Name)
			return err
		})

		assert.NoError(t, err)
	})
}

func CategoryProductRepositoryImplFindAll(t *testing.T) {
	filters := []repository.Filter{
		{
			Column:   "id",
			Value:    "",
			Operator: repository.Equality,
		},
		{
			Column:   "deleted_at",
			Value:    categoryProduct2.ID,
			Operator: repository.IsNULL,
		},
	}

	page := 1
	pageSize := 2
	paginate := repository.Pagination{
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
		Orders: map[string]string{
			"id": "DESC",
		},
		PrefixOrder: "",
	}

	findAllAndSearch := repository.FPSParam{
		Filters:    &filters,
		Pagination: paginate,
		Search:     "",
	}
	err := CategoryProductRepository.StartTx(context.Background(), pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: "",
		BeginQuery:     "",
	}, func() error {
		categoryProducts, total, err := CategoryProductRepository.FindAllAndSearch(context.Background(), findAllAndSearch)
		t.Log(categoryProducts)
		t.Log(total)
		return err
	})
	assert.NoError(t, err)
}

func CategoryProductRepositoryImplCreateError(t *testing.T) {
	t.Run("tx_is_closed", func(t *testing.T) {
		err := CategoryProductRepository.Create(context.Background(), categoryProduct1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, pgx.ErrTxClosed)
	})
}

func CategoryProductRepositoryImplUpdateError(t *testing.T) {
	t.Run("tx_is_closed", func(t *testing.T) {
		err := CategoryProductRepository.Update(context.Background(), categoryProduct1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, pgx.ErrTxClosed)
	})

}

func CategoryProductRepositoryImplCheckOneError(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               "asal",
			Operator:            repository.Equality,
			NextConditionColumn: "",
		},
		{
			Column:              "deleted_at",
			Value:               "asal",
			Operator:            repository.IsNULL,
			NextConditionColumn: "AND",
		},
	}

	t.Run("not_found", func(t *testing.T) {
		err := CategoryProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			res, err := CategoryProductRepository.CheckOne(context.Background(), &filters)
			assert.Equal(t, false, res)
			return err
		})

		assert.NoError(t, err)
	})
	t.Run("after_delete_categoryProduct3", func(t *testing.T) {
		filters[0].Value = categoryProduct3.ID
		err := CategoryProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			res, err := CategoryProductRepository.CheckOne(context.Background(), &filters)
			assert.Equal(t, false, res)
			return err
		})

		assert.NoError(t, err)
	})
}

func CategoryProductRepositoryImplFindOneError(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               "asal",
			Operator:            repository.Equality,
			NextConditionColumn: "AND",
		},
		{
			Column:              "deleted_at",
			Value:               "asal",
			Operator:            repository.IsNULL,
			NextConditionColumn: "AND",
		},
	}

	t.Run("not_found", func(t *testing.T) {
		filters[0].Value = "asal"
		err := CategoryProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			user, err := CategoryProductRepository.FindOne(context.Background(), &filters)
			assert.Nil(t, user)
			return err
		})

		assert.ErrorIs(t, err, pgx.ErrNoRows)
	})

	t.Run("after_delete_categoryProduct3", func(t *testing.T) {
		filters[0].Value = categoryProduct3.ID
		err := CategoryProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			categoryProduct, err := CategoryProductRepository.FindOne(context.Background(), &filters)
			assert.Nil(t, categoryProduct)
			return err
		})

		assert.ErrorIs(t, err, pgx.ErrNoRows)
	})
}
