package pokecache_test

import (
	"bytes"
	"testing"
	"time"

	pokecache "github.com/zig-gy/pokedexcli/internal/pokeCache"
)

func TestAddGet(t *testing.T) {
	cache := pokecache.NewCache(1 * time.Millisecond)
	cases := []struct {
		inputKey		string
		inputBytes		[]byte
	}{
		{
			inputKey: "thing",
			inputBytes: []byte{1},
		},
		{
			inputKey: "thangz",
			inputBytes: []byte{2},
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputBytes)
		val, ok := cache.Get(c.inputKey)
		if !ok {
			t.Errorf("%s does not exist in cache", c.inputKey)
		}
		if !bytes.Equal(val, c.inputBytes){
			t.Errorf("stored entry does not match input entry")
		}
	}
}

func TestReap(t *testing.T) {
	cache := pokecache.NewCache(10 * time.Millisecond)
	cases := []struct {
		inputKey		string
		inputBytes		[]byte
	}{
		{
			inputKey: "thing",
		},
		{
			inputKey: "thangz",
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputBytes)
		time.Sleep(30 * time.Millisecond)
		_, ok := cache.Get(c.inputKey)
		if ok {
			t.Errorf("%s still exists after interval", c.inputKey)
		}
	}
}
