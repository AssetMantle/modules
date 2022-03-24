package cuckoo

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	"math/rand"
	"testing"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var numData = 32
var fpRate = 0.0001
var retries = 500
var entriesPerBucket = uint(2)

var testFPRandomData = numData * 100
var simulations = 1

func TestCuckoo(t *testing.T) {
	totalFPs := 0
	numErrors := 0
	totalNegatives := 0
	c := NewCuckoo(uint(numData), entriesPerBucket, retries, fpRate)
	fmt.Println("entriesPerBucket:", c.entriesPerBucket)
	fmt.Println("numBuckets:", c.numBuckets)
	fmt.Println("fingerprintLength:", c.fingerprintLength)
	fmt.Println("numItems:", c.n)
	for i := 0; i < simulations; i++ {
		fp, allNegatives, errs := simulate()
		totalFPs = totalFPs + fp
		totalNegatives = totalNegatives + allNegatives
		numErrors = numErrors + errs
	}
	fmt.Println("\nSimulation Result (Total):", simulations, "total random data tested per simulation:", testFPRandomData)
	fmt.Println("total false positive:", totalFPs)
	fmt.Println("total negatives:", totalNegatives)
	fmt.Println("total errors:", numErrors)
}

func simulate() (int, int, int) {
	data := generateData(numData)
	c := NewCuckoo(uint(numData), entriesPerBucket, retries, fpRate)

	insertMap := make(map[string]bool)
	full := 0

	for _, v := range data {
		err := c.Insert(v)
		if err == nil {
			insertMap[v] = true
		} else {
			full++
		}
	}

	allNegatives := 0
	for _, v := range data {
		ok := c.Lookup(v)
		if !ok {
			// all should be negative
			allNegatives++
		}
	}

	falsePositives := 0
	for i := 0; i < testFPRandomData; i++ {

		ok := c.Lookup(randStringRunes(5)) //strconv.Itoa(i))
		if ok {
			// all should be negative
			falsePositives++
		}
	}
	return falsePositives, allNegatives, full
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letterRunes[random.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateData(length int) []string {
	result := make([]string, length)
	for i := 0; i < length; i++ {
		result[i] = sdk.AccAddress(crypto.AddressHash([]byte(randStringRunes(45)))).String()
	}
	return result
}
