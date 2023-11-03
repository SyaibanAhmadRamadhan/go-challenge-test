package unit

import (
	"fmt"
	"testing"

	"challenge-test-synapsis/repository"
)

func testVal(str ...any) {
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
}

func TestAudit(t *testing.T) {
	user := repository.User{}
	str := fmt.Sprintf(`SELECT id, role_id, username, email, password, phone_number, %s
	FROM m_user WHERE LIMIT 1`, user.Audit.ToQuery(""))
	t.Log(str)
}

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

	filterStr, values := repository.GenerateFilters(SpesificColumns)
	res := fmt.Sprintf("SELECT id FROM m_user WHERE %s LIMIT 1", filterStr)
	t.Log(res)
	testVal(values...)
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
