package sqlutils
import "strings"
import "reflect"
import "github.com/c9s/inflect"
import "database/sql"

// Generate SQL columns string for selecting.
func BuildSelectColumnClause(val interface{}) (string) {
	columns := ParseColumnNames(val)
	return strings.Join(columns,",")
}

func BuildSelectClause(val interface{}) (string) {
	// get table name
	// inflect.Underscore()
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()
	tableName := inflect.Tableize(typeOfT.Name())
	return "SELECT " + BuildSelectColumnClause(val) + " FROM " + tableName;
}


func Select(db *sql.DB, val interface{}) (*sql.Rows, error) {
	sql := BuildSelectClause(val)
	return PrepareAndQuery(db, sql)
}

func SelectWith(db *sql.DB, val interface{}, postSQL string, args ...interface{}) (*sql.Rows, error) {
	sql := BuildSelectClause(val) + " " + postSQL
	return PrepareAndQuery(db, sql, args)
}

