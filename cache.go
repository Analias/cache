// cache project cache.go
package cache

import (
)

type cache interface {
	Fetch (key string) *interface{}
	
	Store (key string, value interface{}) error
	StoreKeyValues (keys []string, values []interface{}) error
	StoreMap (values map[interface{}]interface{}) error

	Remove (key string) (*interface{}, error)
	RemoveAll () error
	
	Contains (key string) (bool, error)
}

