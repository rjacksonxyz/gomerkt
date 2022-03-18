package gomerkt

import (
	"math/rand"
	"testing"
	"time"
)

var contents []Content
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var iterations = 10000

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10000; i++ {
		contents = append(contents,
			TestSHA256Content{
				RandStringRunes(5),
			})
	}
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func BenchmarkNewTree(b *testing.B) {
	for i := 0; i < iterations; i++ {
		NewTree(contents)
	}
}

func BenchmarkNewTreeCC2(b *testing.B) {
	for i := 0; i < iterations; i++ {
		NewTreeCC(2, contents)
	}
}
