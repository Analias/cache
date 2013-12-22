// Tests for github.com/analias/cache.
package cache

import (
	"fmt"
	"testing"
	
	_ "github.com/analias/cache"
)

type stringIntCache struct {
	store map[string]int
}

type stringStringCache struct {
	store map[string]string
}

type location struct {
	x int
	y int
	z int
}

type locationCache struct {
	store map[location]string
}

func TestTests (t *testing.T) {
	
}

func TestIntCache (t *testing.T) {
	
}
