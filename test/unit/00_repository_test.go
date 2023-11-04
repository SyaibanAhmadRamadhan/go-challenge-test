package unit

import (
	"fmt"
	"testing"
	"unsafe"

	"challenge-test-synapsis/repository"
)

func testVal(str ...any) {
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
}

func TestAudit(t *testing.T) {
	str := fmt.Sprintf(`SELECT id, role_id, username, email, password, phone_number, %s
	FROM m_user WHERE LIMIT 1`, repository.AuditToQuery(""))
	t.Log(str)
}

func TestSpesificColumnToString(t *testing.T) {
	SpesificColumns := &[]repository.Filter{
		{
			Prefix:              "user.",
			Column:              "deleted_at",
			Value:               "value2",
			Operator:            repository.IsNotNULL,
			NextConditionColumn: "",
		},
	}

	filterStr, values := repository.GenerateFilters(SpesificColumns)
	res := fmt.Sprintf("SELECT id FROM m_user %s LIMIT 1", filterStr)
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

func TestValidateColumnFromStruct(t *testing.T) {
	user := repository.User{}

	err := repository.ValidateColumnFromStruct(&user, "id")
	t.Log(err)
}

func TestStructEmbedMemoryUsage(t *testing.T) {
	var categoryProduct repository.CategoryProduct
	var categoryProductEmbed repository.CategoryProductJoin

	sizeCategoryProduct := int(unsafe.Sizeof(categoryProduct))
	sizeCategoryProductEmbed := int(unsafe.Sizeof(categoryProductEmbed))

	t.Logf("Size of categoryProduct: %d bytes\n", sizeCategoryProduct)
	t.Logf("Size of categoryProductEmbed: %d bytes\n", sizeCategoryProductEmbed)

	if sizeCategoryProduct <= sizeCategoryProductEmbed {
		t.Logf("CategoryProduct should have smaller memory footprint than categoryProductEmbed")
	}

}
