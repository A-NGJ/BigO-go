package hash_table

import "fmt"

func (ht HashTable) hash(key string) int {
	hsh := 0
	for i := 0; i < len(key); i++ {
		hsh = (hsh + int(key[i])*i) % len(ht.data)
	}
	return hsh
}

type HashTable struct {
	data [][]interface{}
}

func New(size int) HashTable {
	return HashTable{
		data: make([][]interface{}, size),
	}
}

func (ht HashTable) Len() int {
	return len(ht.data)
}

func (ht *HashTable) Set(key string, val interface{}) {
	address := ht.hash(key)
	if ht.data[address] == nil {
		ht.data[address] = []interface{}{}
		ht.data[address] = append(ht.data[address], []interface{}{key, val})
		return
	}

	for idx, i := range ht.data[address] {
		if i.([]interface{})[0].(string) == key {
			ht.data[address][idx] = []interface{}{key, val}
			return
		}
	}
	ht.data[ht.hash(key)] = append(ht.data[ht.hash(key)], []interface{}{key, val})
	return
}

func (ht HashTable) Get(key string) (interface{}, error) {
	data := ht.data[ht.hash(key)]
	if len(data) != 0 {
		for _, k := range data {
			bucket := k.([]interface{})
			if bucket[0] == key {
				return bucket[1], nil
			}
		}
	}
	return nil, fmt.Errorf("key %s not found in map", key)
}

func (ht HashTable) Keys() []string {
	keysArray := []string{}
	for _, d := range ht.data {
		if len(d) > 0 {
			for _, k := range d {
				keysArray = append(keysArray, k.([]interface{})[0].(string))
			}
		}
	}
	return keysArray
}

func FirstRecurringCharacter(characters []int) (int, error) {
	set := make(map[int]struct{})
	for _, c := range characters {
		if _, ok := set[c]; ok {
			return c, nil
		}
		set[c] = struct{}{}
	}
	return -1, fmt.Errorf("no recurring characters found")
}
