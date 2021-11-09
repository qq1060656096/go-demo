package main

import "fmt"

func main() {
	InitializeBusiness().Exist("test")
}


type Business struct {
	dao Dao
}

func (b *Business) Exist(name string) {
	fmt.Printf("Business.Exist call\n")
	b.dao.Exist(name)
}

type Dao struct {
	db DB
}

func (d *Dao) Exist(name string) bool {
	fmt.Printf("Dao.Exist call\n")
	return d.db.Exist(name)
}

type DB struct {

}

func (d *DB) Exist(name string) bool {
	fmt.Printf("DB.Exist call\n")
	if name != "" {
		return true
	}
	return false
}


func NewBusiness(dao Dao) *Business {
	return &Business{
		dao: dao,
	}
}

func NewDao(db DB) Dao {
	return Dao{
		db: db,
	}
}

func NewDB() DB {
	return DB{}
}