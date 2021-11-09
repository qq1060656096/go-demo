package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	b, err := InitializeBusiness()
	fmt.Println(err)
	b.Exist("test")
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


func NewBusiness(dao Dao) (*Business, error) {
	var err error
	if time.Now().Unix()%2 == 0 {
		err = errors.New("随机错误")
	}
	return &Business{
		dao: dao,
	}, err
}

func NewDao(db DB) Dao {
	return Dao{
		db: db,
	}
}

func NewDB() DB {
	return DB{}
}