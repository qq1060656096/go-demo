package example

import "testing"

func TestInsertFromModel(t *testing.T) {
	d := &demo{
		Name: "TestInsertModel",
	}
	db1 := db.Create(&d)
	if db1.Error != nil {
		t.Error("TestInsertModel fail: ", db1.Error)
	}
}

func TestBatchInsertFromModels(t *testing.T) {
	demos := []demo{
		{
			Name: "TestBatchInsertFromModels.1",
		},
		{
			Name: "TestBatchInsertFromModels.2",
		},
	}
	db1 := db.Create(demos)
	if db1.Error != nil {
		t.Error("TestBatchInsertFromModels fail: ", db1.Error)
	}
}

func TestBatchInsertFromModels2(t *testing.T) {
	demos := []demo{
		{
			Name: "TestBatchInsertFromModels2222.111",
		},
		{
			Name: "TestBatchInsertFromModels2222.222",
		},
		{
			Name: "TestBatchInsertFromModels2222.333",
		},
	}


	db1 := db.CreateInBatches(demos, 2)
	if db1.Error != nil {
		t.Error("TestBatchInsertFromModels fail: ", db1.Error)
	}
}

func TestInsertFromRawSql(t *testing.T) {
	sql := "insert into demo(`name`) values(?)"
	db1 := db.Exec(sql, "TestInsertRawSql")
	if db1.Error != nil {
		t.Error("TestInsertRawSql fail: ", db1.Error)
	}
}