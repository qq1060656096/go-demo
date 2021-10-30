package example

import (
	"log"
	"testing"
)

func init()  {
	deleteTestInitData()
}

func deleteTestInitData() {
	demos := []demo{
		{
			Name: "TestDeleteFromWhereModel",
		},
		{
			Name: "TestDeleteFromWhereModel.2",
		},
		{
			Name: "TestDeleteFromRawSql",
		},
		{
			Name: "TestDeleteFromRawSql.2",
		},
	}
	db1 := db.Create(demos)
	if db1.Error != nil {
		log.Fatal("deleteTestInitData fail: ", db1.Error)
	}
}


func  TestDeleteFromWhereModel(t *testing.T) {
	db1 := db.Where("name = ?", "TestDeleteFromWhereModel").Delete(demo{})
	if db1.Error != nil {
		t.Error("TestDeleteFromWhereModel fail: ", db1.Error)
	}
}

func TestDeleteFromRawSql(t *testing.T) {
	sql := "delete from demo where `name` = ? limit 1"
	db1 := db.Exec(sql, "TestDeleteFromRawSql")
	if db1.Error != nil {
		t.Error("TestDeleteFromRawSql fail: ", db1.Error)
	}
}