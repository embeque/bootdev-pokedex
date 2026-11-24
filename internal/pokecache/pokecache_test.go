package pokecache

import (
	"testing"
	"time"
)

var cache = Cache{
	data: map[string]CachedEntry{
		"api": {
			createdAt: time.Now(),
			val:       []byte("testingapi"),
		},
		"localhost": {
			createdAt: time.Now().Add(time.Hour),
			val:       []byte("8080"),
		},
	},
}

func TestAdd(t *testing.T) {
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "addtest",
			value: []byte("added"),
		}, {
			key:   "addtest",
			value: []byte("changed"),
		}, {
			key:   "localhost",
			value: []byte("127.0.0.1:8080"),
		},
	}

	for i, c := range cases {
		// ooperation
		cache.Add(c.key, (c.value))

		expected := c.value
		actual, ok := cache.data[c.key]
		if !ok {
			t.Fatalf("Test %d: Insertion Failed: Key %s not added in Cache Data", i, c.key)
		}

		if string(actual.val) != string(expected) {
			t.Fatalf("Test %d: Assertion Failed: (got: %s, expected: %s)", i, string(actual.val), string(expected))
		}
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		key      string
		expected []byte
	}{
		{
			key:      "api",
			expected: []byte("testingapi"),
		}, {
			key:      "localhost",
			expected: []byte("127.0.0.1:8080"), // updated value
		},
	}

	for i, c := range cases {
		// ooperation
		actual, ok := cache.Get(c.key)
		if !ok {
			t.Fatalf("Test %d: Extraction Failed: Key %s not present in Cache Data", i, c.key)
		}

		expected := c.expected

		if string(actual) != string(expected) {
			t.Fatalf("Test %d: Assertion Failed: (got: %s, expected: %s)", i, string(actual), string(expected))
		}
	}
}

func TestExpiry(t *testing.T) {
	cache = NewCache(time.Second * 5)
	cache.Add("apikey", []byte("personal"))
	cache.Add("localhost", []byte("8080"))
	cache.Add("port", []byte("443"))

	if len(cache.data) != 3 {
		t.Fatal("Insertion Error: data is lost not all inserted")
	}
	time.Sleep(time.Second * 5)
	actual := len(cache.data)
	if actual >= 3 {
		t.Fatalf("Expiration Error: values should expire got %d", actual)
	}
	time.Sleep(time.Second * 12)
	actual = len(cache.data)
	if actual != 0 {
		t.Fatalf("Expiration Error: values should expire got %d", actual)
	}
}
