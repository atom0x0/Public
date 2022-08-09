package main

import (
	"Go-pro/src/bolt"
	"fmt"
)

var db_name = []byte("xray.db")
var bucket = []byte("bucket")
var key = []byte("foo")
var value = []byte("bar")

func main() {
	db := bbolt.GetDB(db_name)
	bbolt.UpdateKV(db, bucket, key, value)
	val_1 := bbolt.ReadKV(db, bucket, key)
	fmt.Printf("Value is {%s} \n", val_1)

	bbolt.DeleteKV(db, bucket, key)
	val_2 := bbolt.ReadKV(db, bucket, key)
	fmt.Printf("Value is {%s} \n", val_2)
	defer db.Close()
}
