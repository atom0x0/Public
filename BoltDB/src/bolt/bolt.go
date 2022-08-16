package bolt

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func GetDB(db_name []byte) *bolt.DB {
	db, err := bolt.Open(string(db_name), 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ReadKV(db *bolt.DB, bucket []byte, key []byte) string {
	value := []byte("")
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket)
		val := bucket.Get(key)
		value = val
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return string(value)
}

func UpdateKV(db *bolt.DB, bucket []byte, key []byte, value []byte) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			log.Fatal(err)
		}

		err = bucket.Put(key, value)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteKV(db *bolt.DB, bucket []byte, key []byte) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket)
		err := bucket.Delete(key)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteBucket(db *bolt.DB, bucket []byte) {
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(bucket)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
