// cache project cache.go
package cache

import (
	"fmt"
	"log"
	//"reflect"
)

type CacheOps interface {
	Fetch(key string) (*interface{}, error)

	Store(key string, value interface{}) error
	//StoreKeyValues(keys []string, values []interface{}) error
	StoreMap(values map[string]interface{}) error

	Remove(key string) error
	RemoveAll() error

	Contains(key string) (bool, error)
}

type CacheEntry struct {
	Value interface{}

	lastReferenced uint64
	referenceCount uint
}

type Cache struct {
	CacheOps

	store    map[string]CacheEntry
	size     uint
	TTL      uint64
	defaults interface{}
}

type CacheError struct {
	message string
	key     string
	value   interface{}
}

func (ce CacheError) Error() string {
	return fmt.Sprintf("%s", ce.message)
}

func (c Cache) Fetch(key string) (interface{}, error) {
	match, present := c.store[key]
	if !present {
		return nil, CacheError{"key not found!", key, match.Value}
	}

	return match.Value, nil
}

func (c Cache) Store(key string, v interface{}) error {
	ce := c.store[key]
	ce.Value = v

	return nil
}

//func (c Cache) StoreKeyValues(values map[string]interface{}) error {
//	for key, value := range values {
//		c.store[key].value = value
//	}

//	return nil
//}

func newCache(s interface{}, size uint, ttl uint) Cache {

	log.Printf("s type: %T\n", s)
	log.Printf("s value: %#v\n", s)

	//st := reflect.TypeOf(s)

	//store := make(map[string]interface{})

	//log.Printf("stMap type: %t\n", store)
	//log.Printf("stMap value: %#v\n", store)

	return Cache{size: size, store: make(map[string]CacheEntry, size)} //, err
}
