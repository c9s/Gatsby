package gatsby

import (
	"database/sql"
	"errors"
	"gatsby/sqlutils"
)

// Delete from DB connection object or a transaction object (pointer)
func Delete(e interface{}, val interface{}) *Result {
	var executor, ok = e.(Executor)
	if !ok {
		panic("Not an Executor type")
	}

	pkName := sqlutils.GetPrimaryKeyColumnName(val)

	if pkName == nil {
		return NewErrorResult(errors.New("PrimaryKey column is not defined."), "")
	}

	sqlStr := "DELETE FROM " + sqlutils.GetTableName(val) + " WHERE " + *pkName + " = $1"

	if val.(sqlutils.PrimaryKey) == nil {
		return NewErrorResult(errors.New("PrimaryKey interface is required."), sqlStr)
	}

	id := val.(sqlutils.PrimaryKey).GetPkId()

	var err error
	var res sql.Result

	res, err = executor.Exec(sqlStr, id)
	if err != nil {
		return NewErrorResult(err, sqlStr)
	}

	var r = NewResult(sqlStr)
	r.Result = res
	r.Id = id
	return r
}
