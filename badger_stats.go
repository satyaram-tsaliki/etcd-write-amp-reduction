package main

import (
	"fmt"
	"log"
	"time"
	"github.com/dgraph-io/badger/v3"
	"runtime"
)

func collectBadgerMetrics(db *badger.DB, key string, value []byte) {
	start := time.Now()
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
	if err != nil {
		log.Fatalf("Write Error: %v", err)
	}
	writeLatency := time.Since(start)

	start = time.Now()
	err = db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		return err
	})
	if err != nil {
		log.Fatalf("Read Error: %v", err)
	}
	readLatency := time.Since(start)

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("BadgerDB Metrics:\n")
	fmt.Printf("Write Latency: %v\n", writeLatency)
	fmt.Printf("Read Latency: %v\n", readLatency)
	fmt.Printf("Memory Usage: %v MB\n", memStats.Alloc/1024/1024)
	fmt.Printf("Write Amplification: Approximate factor depends on LSM levels\n")
}
func main() {
	// Open BadgerDB
	opts := badger.DefaultOptions("./badgerdb")
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("Error opening BadgerDB: %v", err)
	}
	defer db.Close()
	collectBadgerMetrics(db, "testKey", []byte("testValue"))
}
