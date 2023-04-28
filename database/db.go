package database

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	DB *leveldb.DB // a global DB instance, must call InitDB before accessing
)

// Initialize a leveldb instance
func InitDB(data_folder string) *leveldb.DB {
	var err error
	DB, err = leveldb.OpenFile(data_folder, nil)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return DB
}
