package internal

import (
	"encoding/json"
	"fmt"
	"strconv"

	bolt "go.etcd.io/bbolt"
)

type DB struct {
	conn       *bolt.DB
	bucketName string
}

func OpenDB(dbPath, bucketName string) (*DB, error) {
	conn, err := bolt.Open(dbPath, 0o600, nil)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := conn.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket %q: %w", bucketName, err)
		}
		return nil
	}); err != nil {
		if closeErr := conn.Close(); closeErr != nil {
			return nil, fmt.Errorf(
				"initialize database: %w (close database: %v)",
				err,
				closeErr,
			)
		}
		return nil, fmt.Errorf("initialize database: %w", err)
	}

	return &DB{
		conn:       conn,
		bucketName: bucketName,
	}, nil
}

func (db *DB) CreateTask(task Task) error {
	return db.conn.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(db.bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", db.bucketName)
		}

		id, _ := bucket.NextSequence()
		task.ID = strconv.FormatUint(id, 10)

		data, err := json.Marshal(task)
		if err != nil {
			return fmt.Errorf("marshal task: %w", err)
		}

		return bucket.Put([]byte(task.ID), data)
	})
}

func (db *DB) ListAllTasks() ([]Task, error) {
	var tasks []Task

	err := db.conn.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(db.bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", db.bucketName)
		}

		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			var task Task
			if err := json.Unmarshal(value, &task); err != nil {
				return fmt.Errorf("decode task %q: %w", key, err)
			}

			tasks = append(tasks, task)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}

	return tasks, nil
}

func (db *DB) DoTask(id string) error {
	return db.conn.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(db.bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", db.bucketName)
		}

		key := []byte(id)
		data := bucket.Get(key)
		if data == nil {
			return fmt.Errorf("task %q not found", id)
		}

		var task Task
		if err := json.Unmarshal(data, &task); err != nil {
			return fmt.Errorf("decode task %q: %w", id, err)
		}

		task.Done = true
		fmt.Printf("You have completed the \"%s\" task.\n", task.Description)
		updated, _ := json.Marshal(task)
		return bucket.Put(key, updated)
	})
}

func (db *DB) Close() error {
	if err := db.conn.Close(); err != nil {
		return fmt.Errorf("close database: %w", err)
	}
	return nil
}
