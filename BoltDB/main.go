package main

import (
	"BoltDB/src/bolt"
	"fmt"
)

var db_name = []byte("xray.db")
var bucket = []byte("bucket")
var key = []byte("foo")
var value = []byte("bar")

func main() {
	db := bolt.GetDB(db_name)
	bolt.UpdateKV(db, bucket, key, value)
	val_1 := bolt.ReadKV(db, bucket, key)
	fmt.Printf("Value is {%s} \n", val_1)

	bolt.DeleteKV(db, bucket, key)
	val_2 := bolt.ReadKV(db, bucket, key)
	fmt.Printf("Value is {%s} \n", val_2)
	defer db.Close()
}
