package main

import (
	"fmt"
	"log"
	"time"
	"github.com/dgraph-io/badger/v3"
)
type BadgerETCD struct {
	db *badger.DB
}
func NewBadgerETCD(dbPath string) *BadgerETCD {
	opts := badger.DefaultOptions(dbPath).WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("Failed to open BadgerDB: %v", err)
	}
	return &BadgerETCD{db: db}
}
func (b *BadgerETCD) Put(key, value []byte) error {
	start := time.Now()
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
	log.Printf("Insertion Time: %v µs", time.Since(start).Microseconds())
	return err
}
func (b *BadgerETCD) Get(key []byte) ([]byte, error) {
	start := time.Now()
	var value []byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			value = append([]byte{}, val...)
			return nil
		})
	})
	log.Printf("Search Time: %v µs", time.Since(start).Microseconds())
	return value, err
}

func (b *BadgerETCD) Delete(key []byte) error {
	start := time.Now()
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
	log.Printf("Deletion Time: %v µs", time.Since(start).Microseconds())
	return err
}

func (b *BadgerETCD) Close() {
	b.db.Close()
}

func main() {
	dbPath := "./badger_etcd"
	etcdDB := NewBadgerETCD(dbPath)
	defer etcdDB.Close()

	key := []byte("test_key")
	value := []byte("test_value")

	if err := etcdDB.Put(key, value); err != nil {
		log.Fatalf("Put Error: %v", err)
	}
	
	val, err := etcdDB.Get(key)
	if err != nil {
		log.Fatalf("Get Error: %v", err)
	}
	fmt.Printf("Retrieved Value: %s\n", val)
	
	if err := etcdDB.Delete(key); err != nil {
		log.Fatalf("Delete Error: %v", err)
	}
}
