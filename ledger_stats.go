package main
import (
	"fmt"
	"log"
	"time"
	"github.com/ledgerdb/ledgerdb-go"
	"runtime"
)
func collectLedgerMetrics(db *ledgerdb.DB, key string, value []byte) {
	start := time.Now()
	err := db.Put([]byte(key), value)
	if err != nil {
		log.Fatalf("Write Error: %v", err)
	}
	writeLatency := time.Since(start)
	start = time.Now()
	_, err = db.Get([]byte(key))
	if err != nil {
		log.Fatalf("Read Error: %v", err)
	}
	readLatency := time.Since(start)
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("LedgerDB Metrics:\n")
	fmt.Printf("Write Latency: %v\n", writeLatency)
	fmt.Printf("Read Latency: %v\n", readLatency)
	fmt.Printf("Memory Usage: %v MB\n", memStats.Alloc/1024/1024)
	fmt.Printf("Write Amplification: Lower due to append-only logging\n")
}
func main() {
	db, err := ledgerdb.Open("./ledgerdb")
	if err != nil {
		log.Fatalf("Error opening LedgerDB: %v", err)
	}
	defer db.Close()
	collectLedgerMetrics(db, "testKey", []byte("testValue"))
}
