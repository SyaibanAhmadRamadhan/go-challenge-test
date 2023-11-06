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

var product1 = &repository.Product{
	ID:                "cp1",
	CategoryProductID: "cp1",
	Name:              "cp1",
	Stock:             10,
	Price:             123,
	Description:       "apa aja",
	Audit:             auditDefault,
}
var product2 = &repository.Product{
	ID:                "cp2",
	CategoryProductID: "cp2",
	Name:              "cp2",
	Stock:             10,
	Price:             123,
	Description:       "apa aja2",
	Audit:             auditDefault,
}
var product3 = &repository.Product{
	ID:                "cp3",
	CategoryProductID: "cp1",
	Name:              "cp3",
	Stock:             10,
	Price:             123,
	Description:       "apa aja3",
	Audit:             auditDefault,
}
var product4 = &repository.Product{
	ID:                "cp46",
	CategoryProductID: "cp2",
	Name:              "cp46",
	Stock:             10,
	Price:             123,
	Description:       "apa aja4",
	Audit:             auditDefault,
}
var product5 = &repository.Product{
	ID:                "cp56",
	CategoryProductID: "cp1",
	Name:              "cp56",
	Stock:             10,
	Price:             123,
	Description:       "apa aja5",
	Audit:             auditDefault,
}
var product6 = &repository.Product{
	ID:                "cp6",
	CategoryProductID: "cp1",
	Name:              "cp6",
	Stock:             10,
	Price:             123,
	Description:       "apa aja6",
	Audit:             auditDefault,
}

func ProductRepositoryImplCreate(t *testing.T) {
	err := ProductRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
		err := ProductRepository.Create(context.Background(), product1)
		err = ProductRepository.Create(context.Background(), product2)
		err = ProductRepository.Create(context.Background(), product3)
		err = ProductRepository.Create(context.Background(), product4)
		err = ProductRepository.Create(context.Background(), product5)
		err = ProductRepository.Create(context.Background(), product6)
		return err
	})
	assert.NoError(t, err)

}

func ProductRepositoryImplCheckOne(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               "cp1",
			Operator:            repository.Equality,
			NextConditionColumn: "",
		},
	}

	err := ProductRepository.StartTx(context.Background(), pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: "",
		BeginQuery:     "",
	}, func() error {
		res, err := ProductRepository.CheckOne(context.Background(), &filters)
		assert.Equal(t, true, res)
		return err
	})

	assert.NoError(t, err)
}

func ProductRepositoryImplUpdate(t *testing.T) {
	product1 = &repository.Product{
		ID:                "cp1",
		CategoryProductID: "cp1",
		Name:              "baju keren",
		Stock:             10,
		Price:             1000000,
		Description:       "baju mahal",
		Audit: repository.Audit{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper.NewNullString("user id"),
		},
	}

	err := ProductRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
		err := ProductRepository.Update(context.Background(), product1)
		return err
	})
	assert.NoError(t, err)
}

func ProductRepositoryImplDelete(t *testing.T) {
	err := ProductRepository.StartTx(context.Background(), repository.LevelReadCommitted(), func() error {
		err := ProductRepository.Delete(context.Background(), "cp3", "admin1")
		return err
	})
	assert.NoError(t, err)
}

func ProductRepositoryImplFindOne(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "id",
			Value:               product2.ID,
			Operator:            repository.Equality,
			NextConditionColumn: "",
		},
	}

	t.Run("product2", func(t *testing.T) {
		err := ProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			product, err := ProductRepository.FindOne(context.Background(), &filters)
			assert.Equal(t, product2, product)
			return err
		})

		assert.NoError(t, err)
	})

	t.Run("after_update_product1", func(t *testing.T) {
		filters[0].Value = product1.ID
		err := ProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			product, err := ProductRepository.FindOne(context.Background(), &filters)
			assert.Equal(t, "baju keren", product.Name)
			return err
		})

		assert.NoError(t, err)
	})
}

func ProductRepositoryImplFindAll(t *testing.T) {
	filters := []repository.Filter{
		{
			Prefix:              "",
			Column:              "deleted_at",
			Value:               product2.ID,
			Operator:            repository.IsNULL,
			NextConditionColumn: "",
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
		Search:     "apa",
	}
	err := ProductRepository.StartTx(context.Background(), pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: "",
		BeginQuery:     "",
	}, func() error {
		products, total, err := ProductRepository.FindAllAndSearch(context.Background(), findAllAndSearch)
		t.Log(products)
		t.Log(total)
		return err
	})
	assert.NoError(t, err)
}

func ProductRepositoryImplCreateError(t *testing.T) {
	t.Run("tx_is_closed", func(t *testing.T) {
		err := ProductRepository.Create(context.Background(), product1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, pgx.ErrTxClosed)
	})
}

func ProductRepositoryImplUpdateError(t *testing.T) {
	t.Run("tx_is_closed", func(t *testing.T) {
		err := ProductRepository.Update(context.Background(), product1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, pgx.ErrTxClosed)
	})

}

func ProductRepositoryImplCheckOneError(t *testing.T) {
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
		err := ProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			res, err := ProductRepository.CheckOne(context.Background(), &filters)
			assert.Equal(t, false, res)
			return err
		})

		assert.NoError(t, err)
	})
	t.Run("after_delete_product3", func(t *testing.T) {
		filters[0].Value = product3.ID
		err := ProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			res, err := ProductRepository.CheckOne(context.Background(), &filters)
			assert.Equal(t, false, res)
			return err
		})

		assert.NoError(t, err)
	})
}

func ProductRepositoryImplFindOneError(t *testing.T) {
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
		err := ProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			user, err := ProductRepository.FindOne(context.Background(), &filters)
			assert.Nil(t, user)
			return err
		})

		assert.ErrorIs(t, err, pgx.ErrNoRows)
	})

	t.Run("after_delete_product3", func(t *testing.T) {
		filters[0].Value = product3.ID
		err := ProductRepository.StartTx(context.Background(), pgx.TxOptions{
			IsoLevel:       pgx.ReadCommitted,
			AccessMode:     pgx.ReadOnly,
			DeferrableMode: "",
			BeginQuery:     "",
		}, func() error {
			product, err := ProductRepository.FindOne(context.Background(), &filters)
			assert.Nil(t, product)
			return err
		})

		assert.ErrorIs(t, err, pgx.ErrNoRows)
	})
}
