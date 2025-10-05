package storage

import (
	"encoding/json"
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

// Database struct
type Database struct {
	db *bolt.DB
}

const bucketName = "images"

func NewDatabase(path string) (*Database, error) {
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Ensure the bucket exists
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

// Method to save data to one bucket
func (d *Database) SaveData(bucket string, key string, data interface{}) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return b.Put([]byte(key), value)
	})
}

// Method to retrieves data from one bucket
func (d *Database) GetData(bucket string, key string, out interface{}) error {
	return d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		val := b.Get([]byte(key))
		if val == nil {
			return fmt.Errorf("No data found in bucket: %s for key: %s", bucket, key)
		}
		return json.Unmarshal(val, out)
	})
}

// Close closes the database file.
func (d *Database) Close() error {
	return d.db.Close()
}
