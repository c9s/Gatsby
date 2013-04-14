package sqlutils
import "testing"

func TestCreate(t *testing.T) {

	db, err := openDB()
	if err != nil {
		t.Fatal(err)
	}

	staff := Staff{Name: "John", Gender: "m", Phone: "1234567"}
	sql, args := BuildInsertClause(&staff)

	if len(sql) == 0 {
		t.Fatal("Empty SQL")
	}
	if len(args) == 0 {
		t.Fatal("Empty argument")
	}

	t.Log(sql,args)
	Create(db, &staff, DriverPg)
}

