package main

import (
	"BoltDB/src/bolt"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var db_name = []byte("xray.db")
var bucket = []byte("bucket")
var key = []byte("foo")
var wg sync.WaitGroup

func main() {

	// 起始时间
	t_start := time.Now()

	// 获取数据库对象
	db := bolt.GetDB(db_name)
	defer db.Close()
	// 插入记录
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			value := []byte(strconv.Itoa(i))
			bolt.UpdateKV(db, bucket, key, value)
			wg.Done()
		}(i)

		// value := []byte(strconv.Itoa(i))
		// bolt.UpdateKV(db, bucket, key, value)
	}
	wg.Wait()

	// 删除前值
	val_befor := bolt.ReadKV(db, bucket, key)
	// 删除记录
	bolt.DeleteKV(db, bucket, key)
	// 删除后值
	val_after := bolt.ReadKV(db, bucket, key)
	// 结束时间
	t_end := time.Now()
	// 时间差
	diff := t_end.Sub(t_start)

	fmt.Printf("Total is {%v} \n", diff)
	fmt.Printf("Before delete value is {%v} \n", val_befor)
	fmt.Printf("After delete value is {%v} \n", val_after)
}
