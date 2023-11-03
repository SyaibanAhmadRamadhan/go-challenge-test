package unit

import (
	"testing"

	"challenge-test-synapsis/repository"
)

func TestSpesificColumnToString(t *testing.T) {
	SpesificColumns := &[]repository.Filter{
		{
			Prefix:              "",
			Column:              "column1",
			Value:               "value1",
			Operator:            repository.Equality,
			NextConditionColumn: repository.AND,
		},
		{
			Prefix:              "user.",
			Column:              "column2",
			Value:               "value2",
			Operator:            repository.Inequality,
			NextConditionColumn: "",
		},
	}

	t.Log(repository.GenerateFilters(SpesificColumns))
}

func TestPagination_OrderBy(t *testing.T) {
	paginate := repository.Pagination{
		Page:   1,
		Offset: 2,
		Orders: map[string]string{
			"colum n1": "value2",
			"column2":  "ASC",
		},
		PrefixOrder: "user.",
	}

	t.Log(paginate.GenerateOrderBy())
}

// func TestValidateColumnFromStruct(t *testing.T) {
// 	user := masterRepository.User{}
//
// 	err := repository.ValidateColumnFromStruct(&user, "id")
// 	t.Log(err)
// }
