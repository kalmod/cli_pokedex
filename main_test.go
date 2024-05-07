package main

import (
  "testing"
  "time"
  "fmt"
  "github.com/kalmod/cli_pokedex/internal"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := internal.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := internal.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestReapLoopRemovedEntries(t *testing.T){
  const baseTime = 5 * time.Millisecond
  cache := internal.NewCache(baseTime)
  cache.Add("https://example.com",[]byte("testdata1"))

  _, ok := cache.Get("https://example.com")
  if !ok {
    t.Errorf("expected to find key")
    return
  }
  time.Sleep(2*time.Millisecond)
  cache.Add("https://bootdev.com",[]byte("testdata2"))
  time.Sleep(4*time.Millisecond)

  _, foundfirst := cache.Get("https://example.com")
  _, foundsecond := cache.Get("https://bootdev.com")
  if foundfirst {
    t.Errorf("should have been removed by reap")
    return
  }
  if !foundsecond {
    t.Errorf("should have found second element")
    return
  }
}
