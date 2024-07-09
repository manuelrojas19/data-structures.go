package hashmap

import (
	"fmt"
	"hash/fnv"
)

var defaultCapacity uint64 = 1 << 10

type node[K any, V any] struct {
	key   K
	value V
	next  *node[K, V]
}

type HashMap[K any, V any] struct {
	capacity uint64
	size     uint64
	table    []*node[K, V]
}

func createHashMap[K any, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		capacity: defaultCapacity,
		table:    make([]*node[K, V], defaultCapacity),
	}
}

func (hashmap *HashMap[K, V]) hash(key K) uint64 {
	hash := fnv.New64a()
	_, _ = hash.Write([]byte(fmt.Sprintf("%v", key)))
	hashValue := hash.Sum64()
	return (hashmap.capacity - 1) & (hashValue ^ (hashValue >> 16))
}

func (hashmap *HashMap[K, V]) putValue(key any, value any) {
}
