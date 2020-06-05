package main

import (
	"github.com/icrowley/fake"
	"os"
	"testing"
)

var uri = os.Getenv("MONGODB_URI")

func BenchmarkInsertOne(b *testing.B) {
	emails := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		emails[i] = fake.EmailAddress()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertOne(uri, emails[i])
	}
}

func BenchmarkFind(b *testing.B) {
	email := fake.EmailAddress()
	for i := 0; i < b.N; i++ {
		InsertOne(uri, email)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Find(uri, email)
	}
}
