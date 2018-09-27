package trie

import (
    "testing"
    // "reflect"
    // "sort"
)

// Test NewTrieIndex
func TestNewTrieIndex(t *testing.T) {
    trieIndex := NewTrieIndex("carbon-table", '.')

    if trieIndex.name != "carbon-table" {
        t.Errorf("expected name carbon-table, got: %s", trieIndex.name)
    }
    if trieIndex.sep != '.' {
        t.Errorf("expected sep '.', got: %s", string(trieIndex.sep))
    }
    if trieIndex.root.isLeaf != false {
        t.Errorf("expected isLeaf false, got: %t", trieIndex.root.isLeaf)
    }
    if trieIndex.root.Count() != 0 {
        t.Errorf("expected size 0, got: %d", trieIndex.root.Count())
    }
}

func TestTrieIndexInsert(t *testing.T) {
    var index *TrieIndex
    index = NewTrieIndex("carbon-table", '.')

    // Insert
    index.Insert("carbon.cache.a")
    index.Insert("carbon.cache.b")
    index.Insert("carbon.relay.a")

    // Verify
    if index.root.Count() != 1 {
        t.Errorf("expected size 1, got: %d", index.root.Count())
    }
    if index.root.Get("carbon").Count() != 2 {
        t.Errorf("expected size 2, got: %d", index.root.Get("carbon").Count())
    }

    cache := index.root.Get("carbon").Get("cache")
    if cache.Count() != 2 {
        t.Errorf("expected size 2, got: %d", cache.Count())
    }
    if cache.isLeaf != false {
         t.Errorf("expected isLeaf false, got: %t", cache.isLeaf)
    }
    relay := index.root.Get("carbon").Get("relay")
    if relay.Count() != 1 {
        t.Errorf("expected size 1, got: %d", relay.Count())
    }
    if relay.isLeaf != false {
         t.Errorf("expected isLeaf false, got: %t", relay.isLeaf)
    }
    cacheA := cache.Get("a")
    if cacheA.Count() != 0 {
        t.Errorf("expected size 0, got: %d", cacheA.Count())
    }
    if cacheA.isLeaf != true {
         t.Errorf("expected isLeaf true, got: %t", cacheA.isLeaf)
    }

}
