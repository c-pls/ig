package utils

import (
	nanoid "github.com/matoous/go-nanoid"
	"log"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	res, err := nanoid.ID(n)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func RandomNumber(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func UniqueId() string {
	res, err := nanoid.ID(10)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func UniquePostID() string {
	res, err := nanoid.ID(8)
	if err != nil {
		log.Fatal(err)
	}
	return "CP" + res
}

const (
	COMMENT = "comment"
	POST    = "post"
)
