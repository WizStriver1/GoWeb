package database

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/boltdb/bolt"
)

var (
	// BoltDB wrapper
	BoltDB *bolt.DB
	// Databases info
	databases Info
)

// Type is the type of database from a Type* constant
type Type string

// Info contains the database configurations
type Info struct {
	// Database type
	Type Type
	// Bolt info
	Bolt BoltInfo
}

// BoltInfo is the details for the database connectioin
type BoltInfo struct {
	Path string
}

// Connect to the database
func Connect(d Info) {
	var err error

	// Store the config
	databases = d

	// Connect to Bolt
	if BoltDB, err = bolt.Open(d.Bolt.Path, 0600, nil); err != nil {
		log.Println("Bolt Driver Error", err)
	}
}

// Update makes a modification to Bolt
func Update(bucketName string, key string, dataStruct interface{}) error {
	err := BoltDB.Update(func(tx *bolt.Tx) error {
		// Create the bucket
		bucket, e := tx.CreateBucketIfNotExists([]byte(bucketName))
		if e != nil {
			return e
		}

		// Encode the record
		encodedRecord, e := json.Marshal(dataStruct)
		if e != nil {
			return e
		}

		// Store the record
		if e = bucket.Put([]byte(key), encodedRecord); e != nil {
			return e
		}
		return nil
	})
	return err
}

// View retrieves a record in Bolt
func View(bucketName string, key string, dataStruct interface{}) error {
	// dataStruct.Deleted = 1
	rType := reflect.TypeOf(dataStruct)
	rVal := reflect.ValueOf(dataStruct)
	if rType.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else if rType.Kind() == reflect.Struct {
		panic("inStructPtr must be ptr to struct")
	}
	// field, _ := rType.FieldByName("Deleted")
	value := rVal.FieldByName("Deleted")
	value.Set(reflect.ValueOf(uint8(1)))
	err := BoltDB.View(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		// Retrieve the record
		v := b.Get([]byte(key))
		if len(v) < 1 {
			return bolt.ErrInvalid
		}

		// Decode the record
		e := json.Unmarshal(v, &dataStruct)
		if e != nil {
			return e
		}

		return nil
	})

	return err
}

// Delete removes a record from Bolt
func Delete(bucketName string, key string) error {
	err := BoltDB.Update(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		return b.Delete([]byte(key))
	})
	return err
}

// ReadConfig returns the database information
func ReadConfig() Info {
	return databases
}
