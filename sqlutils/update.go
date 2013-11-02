package sqlutils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// This function generates "UPDATE {table} SET name = $1, name2 = $2"
func BuildUpdateClause(val interface{}, placeholder int) (string, []interface{}) {
	tableName := GetTableName(val)
	sql, values := BuildUpdateColumns(val, placeholder)
	return "UPDATE " + tableName + " SET " + sql, values
}

// This function builds update columns from a map
// which generates SQL like "name = $1, phone = $2".
func BuildUpdateColumnsFromMap(cols map[string]interface{}, placeholder int) (string, []interface{}) {
	var setFields []string
	var values []interface{}
	var i int = 1
	for col, arg := range cols {
		if placeholder == QMARK_HOLDER {
			setFields = append(setFields, fmt.Sprintf("%s = ?"))
		} else {
			setFields = append(setFields, fmt.Sprintf("%s = $%d", col, i))
		}
		values = append(values, arg)
		i++
	}
	return strings.Join(setFields, ", "), values
}

// This function generate update columns from a struct object.
func BuildUpdateColumns(val interface{}, placeholder int) (string, []interface{}) {
	var t = reflect.ValueOf(val).Elem()
	var typeOfT = t.Type()
	var values []interface{}
	var tag reflect.StructTag
	var field reflect.Value
	var columnName *string
	var clauseSQL string = ""
	var value interface{}

	for i := 0; i < t.NumField(); i++ {
		field = t.Field(i)
		if value = field.Interface(); value == nil {
			continue
		}

		tag = typeOfT.Field(i).Tag
		if columnName = GetColumnNameFromTag(&tag); columnName == nil {
			continue
		}

		if placeholder == QMARK_HOLDER {
			clauseSQL += *columnName + " = ?" + ", "
		} else {
			clauseSQL += *columnName + " = $" + strconv.Itoa(len(values)+1) + ", "
		}
		// setFields = append(setFields, *columnName+" = $"+strconv.Itoa(len(values)+1))
		values = append(values, value)
	}
	return clauseSQL[:len(clauseSQL)-2], values
}
