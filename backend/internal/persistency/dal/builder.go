package dal

import (
	"fmt"
	"strings"
)

type queryParameterEnumerate struct {
	enumerator *parameterEnumerate
	columns    []string
	parameters []string
}

type parameterEnumerate struct {
	index int
	args  []interface{}
}

func (enumerator *parameterEnumerate) WithParameterEnumerate() (*parameterEnumerate, *queryParameterEnumerate) {
	return enumerator, &queryParameterEnumerate{enumerator: enumerator}
}

func (enumerator *parameterEnumerate) Enumerate(value interface{}) string {
	enumerator.index++
	enumerator.args = append(enumerator.args, value)
	return fmt.Sprintf("$%d", enumerator.index)
}

func (enumerator *parameterEnumerate) CreateLikeCondition(columnName, value string) string { //nolint
	enumerator.index++
	enumerator.args = append(enumerator.args, value)
	return fmt.Sprintf(" AND %s ILIKE $%d", columnName, enumerator.index)
}

func (enumerator *parameterEnumerate) CreateExactCondition(columnName string, value interface{}) string { //nolint
	enumerator.index++
	enumerator.args = append(enumerator.args, value)
	return fmt.Sprintf(" AND %s = $%d", columnName, enumerator.index)
}

func (enumerator *queryParameterEnumerate) AppendParameter(column string, value interface{}) {
	enumerator.columns = append(enumerator.columns, column)
	parameter := enumerator.enumerator.Enumerate(value)
	enumerator.parameters = append(enumerator.parameters, parameter)
}

func (enumerator *queryParameterEnumerate) GetColumns() string {
	return strings.Join(enumerator.columns, ", ")
}

func (enumerator *queryParameterEnumerate) GetParameters() string {
	return strings.Join(enumerator.parameters, ", ")
}

func (enumerator *queryParameterEnumerate) GetAssignedParameters() string {
	assignParameters := make([]string, 0)
	for i := 0; i < len(enumerator.parameters); i++ {
		assignParameters = append(assignParameters, fmt.Sprintf("%s = %s", enumerator.columns[i], enumerator.parameters[i]))
	}

	return strings.Join(assignParameters, ", ")
}
