package example

import (
	"fmt"
	"log"
	"testing"
)

func init(){
	queryTestInitData()
}

func queryTestInitData() {
	demos := []demo{
		{
			Name: "TestQueryFromWhereModel",
		},
		{
			Name: "TestQueryFromRawSql",
		},
		{
			Name: "TestQueryListFromRawSql.1",
		},
		{
			Name: "TestQueryListFromRawSql.2",
		},
		{
			Name: "TestQueryListFromRawSql.3",
		},
	}
	db1 := db.Create(demos)
	if db1.Error != nil {
		log.Fatal("queryTestInitData fail: ", db1.Error)
	}
}



func  TestQueryFromWhereModel(t *testing.T) {
	d := &demo{
	}
	db1 := db.Model(&demo{}).Where("name like ?", "TestQueryFromWhereModel").Find(d)
	if db1.Error != nil {
		t.Error("TestQueryFromWhereModel fail: ", db1.Error)
	}
	fmt.Printf("%#v", d)
}

func TestQueryFromRawSql(t *testing.T) {
	d := []demo{{}}
	sql := "select * from demo where `name` like ?"
	db1 := db.Raw(sql, "TestQueryFromRawSql").Scan(&d)
	if db1.Error != nil {
		t.Error("TestQueryFromRawSql fail: ", db1.Error)
	}
	fmt.Printf("%#v", d)
}

func TestQueryListFromRawSql(t *testing.T) {
	sql := "select * from demo where `name` like ?"
	rows, err := db.Raw(sql, "TestQueryListFromRawSql.%").Rows()
	if err != nil {
		t.Error("TestQueryListFromRawSql fail: ", err)
	}
	d := []demo{{}}
	for rows.Next() {
		dt := &demo{}
		err := db.ScanRows(rows, dt)
		if err != nil {
			t.Error("TestQueryListFromRawSql ScanRows fail: ", err)
		}
		d = append(d, *dt)
		fmt.Printf("%v", d)
	}
}
