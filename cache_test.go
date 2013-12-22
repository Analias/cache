// Tests for github.com/analias/cache.
package cache

import (
	"fmt"
	"testing"
)

type DbConnection struct {
	server   string
	port     uint
	password string
	scheme   string
}

var dbConnectionCache Cache

func TestNewCache(t *testing.T) {
	dbConnectionCache = newCache(DbConnection{port: 1521}, 10, 10*1000)

	fmt.Printf("TestNewCache: dbConnectionCache type: %T\n", dbConnectionCache)
	fmt.Printf("TestNewCache: dbConnectionCache value: %#v\n", dbConnectionCache)
}

func TestStore(t *testing.T) {
	err := dbConnectionCache.Store("server1", DbConnection{server: "server1", port: 1521, password: "secret", scheme: "test"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("TestStore: dbConnectionCache type: %T\n", dbConnectionCache)
	fmt.Printf("TestStore: dbConnectionCache value: %#v\n", dbConnectionCache)
}

func TestFetch(t *testing.T) {
	server1, err := dbConnectionCache.Fetch("store1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("TestFetch: server1 type: %T\n", server1)
	fmt.Printf("TestFetch: server1 value: %#v\n", server1)
}
