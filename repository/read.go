package repository

import (
	"fmt"
	"strings"
)

type Operator string
type Logical string

const AND Logical = "AND"
const OR Logical = "OR"
const Equality Operator = "="
const Inequality Operator = "<>"

type Filter struct {
	Prefix              string // example user. check unit test for detail. TestSpesificColumnToString
	Column              string
	Value               string
	Operator            Operator
	NextConditionColumn Logical
}

type Pagination struct {
	Page        int
	Offset      int
	Orders      map[string]string // key: column, value: ASC OR DESC. check unit test for detail. TestPagination_OrderBy
	PrefixOrder string            // check unit test for detail. TestPagination_OrderBy
}

type SearchParam struct {
	Filters    []Filter
	Pagination Pagination
	Search     string
}

func (p *Pagination) GenerateOrderBy() (str string) {
	if p.Orders != nil {
		lenMaps := len(p.Orders)
		i := 0
		str += "ORDER BY "
		for k, order := range p.Orders {
			i++
			if order != "DESC" && order != "ASC" {
				order = "DESC"
			}
			if i == lenMaps {
				str += fmt.Sprintf("%s%s %s", p.PrefixOrder, strings.ReplaceAll(k, " ", ""), order)
				continue
			}
			str += fmt.Sprintf("%s%s %s, ", p.PrefixOrder, strings.ReplaceAll(k, " ", ""), order)
		}
	}

	return
}

func GenerateFilters(filters *[]Filter) (str string, values []string) {
	if filters != nil {
		latest := (*filters)[len(*filters)-1]
		for i, filter := range *filters {
			values = append(values, filter.Value)
			if filter == latest {
				str += fmt.Sprintf("%s%s %s $%d", filter.Prefix, filter.Column, filter.Operator, i+1)
				continue
			}
			// column=value AND
			str += fmt.Sprintf("%s%s %s $%d %s ", filter.Prefix, filter.Column, filter.Operator, i+1, filter.NextConditionColumn)
		}
	}
	return
}
