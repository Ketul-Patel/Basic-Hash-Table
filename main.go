package main

import (
	"fmt"
	"hashTable/hashMap"
	"strconv"
)

// Basic Hash Table implementation that uses the following assumptions/strategies
// Keys and Values are both strings 
// We only need to access values in the hash table using their corresponding keys
// We can use an array to handle collisions (if two keys hash to the same number then we put them both in an array in the hashed key's bucket)

func main() {
	// array of letters to test the hashmap each letter will be the key and its index in the array will be the value for a basic test
	var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "aa", "ab", "ac", "ad", "ae", "af", "ag", "ah", "ai", "aj", "ak", "al", "am", "an", "ao", "ap", "aq", "ar", "as", "at", "au", "av", "aw", "ax", "ay", "az", "ba", "bb", "bc", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bk", "bl", "bm", "bn", "bo", "bp", "bq", "br", "bs", "bt", "bu", "bv", "bw", "bx", "by", "bz"}
	// put all the entries into the hashMap this should print information about the hashed key and if there was a collision
	for index, letter := range letters {
		hashMap.Put(letter, strconv.Itoa(index))
	}
	// get the values corresponding to the letter keys and check if the value is equal to the index
	for index, letter := range letters {
		value, err := hashMap.Get(letter)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(value)
		if value == strconv.Itoa(index) {
			fmt.Println("success")
		} else {
			fmt.Println("failure" + letter)
		}
	}
	// check that a nonexistent key returns an error
	fmt.Println(hashMap.Get("blah"))
}