package repository

import (
	"fmt"
	"strings"
)

type Logical string

const AND Logical = "AND"
const OR Logical = "OR"

type Operator string

const Equality Operator = "="
const Inequality Operator = "<>"
const IsNULL Operator = "IS NULL"
const IsNotNULL Operator = "IS NOT NULL"

type Filter struct {
	Prefix              string // example user. check unit test for detail. TestSpesificColumnToString
	Column              string
	Value               string
	Operator            Operator // by defautl Equality
	NextConditionColumn Logical  // by default AND
}

type Pagination struct {
	Limit       int
	Offset      int
	Orders      map[string]string // key: column, value: ASC OR DESC. check unit test for detail. TestPagination_OrderBy
	PrefixOrder string            // check unit test for detail. TestPagination_OrderBy
}

type FindAllAndSearchParam struct {
	Filters    *[]Filter
	Pagination Pagination
	Search     string
}

func (p *Pagination) GenerateOrderBy() (str string) {
	str += "ORDER BY "
	if p.Orders != nil {
		lenMaps := len(p.Orders)
		i := 0
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
	} else {
		str += fmt.Sprintf("%sid DESC", p.PrefixOrder)
	}

	return
}

func GenerateFilters(filters *[]Filter) (str string, values []any, lastPlaceHolder int) {
	lastPlaceHolder = 1
	if filters != nil {
		str += "WHERE "
		latest := (*filters)[len(*filters)-1]
		for i, filter := range *filters {
			if filter.Operator == IsNULL || filter.Operator == IsNotNULL {
				str += fmt.Sprintf("%s%s %s",
					filter.Prefix, filter.Column, filter.Operator)
			} else {
				if filter.Value != "" {
					if filter.Operator == "" {
						filter.Operator = Equality
					}
					values = append(values, filter.Value)
					str += fmt.Sprintf("%s%s %s $%d",
						filter.Prefix, filter.Column, filter.Operator, i+1)

					if filter != latest {
						if filter.NextConditionColumn == "" {
							filter.NextConditionColumn = AND
						}
						str += fmt.Sprintf(" %s ", filter.NextConditionColumn)
					}
					lastPlaceHolder += i + 1
				}
			}

		}
	}

	return
}
