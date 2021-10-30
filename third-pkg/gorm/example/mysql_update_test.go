package example

import (
	"log"
	"testing"
)
func init(){
	updateTestInitData()
}

func updateTestInitData() {
	demos := []demo{
		{
			Name: "TestUpdateFromWhereModel",
		},
		{
			Name: "TestUpdateFromRawSql",
		},
	}
	db1 := db.Create(demos)
	if db1.Error != nil {
		log.Fatal("updateTestInitData fail: ", db1.Error)
	}
}

func  TestUpdateFromWhereModel(t *testing.T) {
	d := demo{
		Name: "TestUpdateFromWhereModel.update",
	}
	db1 := db.Model(&demo{}).Where("name = ?", "TestUpdateFromWhereModel").Updates(d)
	if db1.Error != nil {
		t.Error("TestUpdateFromWhereModel fail: ", db1.Error)
	}
}

func TestUpdateFromRawSql(t *testing.T) {
	sql := "update demo set `name` = ? where `name` like ?"
	db1 := db.Exec(sql, "TestUpdateFromRawSql.update", "TestUpdateFromRawSql")
	if db1.Error != nil {
		t.Error("TestUpdateFromRawSql fail: ", db1.Error)
	}
}