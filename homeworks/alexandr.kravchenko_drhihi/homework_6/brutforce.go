package main

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
)

// BrutSha256 returns source (we would call it `key`) of hashsum sha256
// params:
// hash - sha256 hash, represented as HEX string
// availableChars - chars that could be used in key (if it PIN than available chars would be - 0,1,2,3,4,5,6,7,8,9)
// length - length of key
// e.g.:
// BrutSha256("888df25ae35772424a560c7152a1de794440e0ea5cfee62828333a456a506e05", []byte("1234567890"), 4) should return "9999", nil
func BrutSha256(hash string, availableChars []byte, length uint8) (string, error) {

	var wg sync.WaitGroup
	ch := new(string)
	q := new(bool)

	testKey := make([]byte, length)
	for i := uint8(0); i < length; i++ {
		testKey[i] = availableChars[0]
	}
	wg.Add(1)
	go createKey(&wg, testKey, 0, ch, q, &hash, availableChars, length)

	wg.Wait()

	return *ch, nil
}

func main() {
	key := "9999"
	hasher := sha256.New()
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hash)

	result, err := BrutSha256(hashHex, []byte("1234567890"), 4)
	if err != nil {
		panic(err)
	}

	if result != key {
		panic("hash not decrypted")
	} else {
		println(result)
	}
}

func createKey(wg *sync.WaitGroup, testKey []byte, index uint8, ch *string, q *bool, hash *string, availableChars []byte, length uint8) {

	//defer wg.Done()

	if *q {
		wg.Done()
		return
	}

	for _, v := range availableChars {
		testKey[index] = v
		wg.Add(1)
		copyTestKey := make([]byte, length)
		copy(copyTestKey, testKey)
		go checkKey(wg, copyTestKey, ch, q, hash)
		if index+1 < length {
			wg.Add(1)
			go createKey(wg, copyTestKey, index+1, ch, q, hash, availableChars, length)
		}
	}

	wg.Done()

}

func checkKey(wg *sync.WaitGroup, testKey []byte, ch *string, q *bool, hash *string) {

	//defer wg.Done()

	if *q {
		wg.Done()
		return
	}

	hasherTest := sha256.New()
	hasherTest.Write(testKey)
	hashTest := hasherTest.Sum(nil)
	hashHexTest := hex.EncodeToString(hashTest)

	if hashHexTest == *hash {
		*ch = string(testKey)
		*q = true
	}

	wg.Done()

}
