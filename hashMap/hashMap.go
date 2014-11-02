package hashMap

import (
	"fmt"
	"errors"
	"hash/fnv"
)

// Define an entry struct that will represent each entry into the hash table (key value pairs). Also hang on to the Hashedkey, it may help when implementing the binary tree for collisions.
type entry struct {
	Key string
	Hashedkey int
	Value string
}

// The array with 100 buckets each of which has an array of entries. The append function that will be used to add entries to the bucket array doubles the size when it hits the capacity limit. Next step is to implement a binary tree for each bucket to make it faster to find the entry. 
var table	[100][]entry

// Hashing function that takes in a string and returns an int.
// Uses one of golang's built in hashing functions. Mod 100 so that we can use an array with 100 buckets (0-99)
func hash(str string) int {
	h := fnv.New32a()
	h.Write([]byte(str))
	return int(h.Sum32()) % 100
}

// Make an entry from a key and value and place it into a bucket in the hash table 
func Put(key string, value string) entry {
	hashedkey := hash(key)
	fmt.Println("key is: "+key+" value is: "+value)
	fmt.Println("Hashed key is :")
	fmt.Println(hashedkey)
	e := entry{key, hashedkey, value}
	if table[e.Hashedkey] != nil {
		fmt.Println("We have a collision")
		fmt.Println(table[e.Hashedkey])
	}
	table[e.Hashedkey] = append(table[e.Hashedkey], e)
	return e
}

// Get an entry given the key, returns nil if the entry doesn't exist.
func Get(k string) (string, error) {
	entries := table[hash(k)]
	if entries != nil {
		for _ , e := range entries {
			if k == e.Key {
				return e.Value, nil
			}
		}
	}
	return "" , errors.New("Key does not exist") 
}

