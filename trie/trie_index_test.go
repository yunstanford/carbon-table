package trie

import (
    "testing"
    "reflect"
    "sort"
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

// Test TrieIndex Insert
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

// Test TrieIndex ExpandQuery
func TestTrieIndexExpandQuery(t *testing.T) {
    // Setup
    var index *TrieIndex

    index = NewTrieIndex("carbon-table", '.')

    index.Insert("zillow.seattle.velocity")
    index.Insert("zillow.seattle.velo1city")
    index.Insert("zillow.seattle.velo2city")
    index.Insert("zillow.seattle.rentalsConsumer")
    index.Insert("zillow.seattle.rentalsRevenue")
    index.Insert("zillow.seattle.pa")
    index.Insert("zillow.seattle.data")

    index.Insert("zillow.sf1.pa")
    index.Insert("zillow.sf1.data")

    index.Insert("zillow.sf2.pa")
    index.Insert("zillow.sf2.data")

    index.Insert("zillow.nyc.rentalsConsumer")
    index.Insert("zillow.nyc.rentalsRevenue")
    index.Insert("zillow.nyc.data")

    // ExpandQuery
    testCases := []struct {
        pattern string
        expectedQueries []string
    }{
        {"zillow.seattle.velocity", []string{"zillow.seattle.velocity"}},
        {"zillow.sf1.*", []string{"zillow.sf1.data", "zillow.sf1.pa"}},
        {"zillow.*.data", []string{"zillow.nyc.data", "zillow.seattle.data", "zillow.sf1.data", "zillow.sf2.data"}},
        {"zillow.sf[0-9].data", []string{"zillow.sf1.data", "zillow.sf2.data"}},
        {"zillow.seattle.velo[1-9]city", []string{"zillow.seattle.velo1city", "zillow.seattle.velo2city"}},
        {"zillow.*.rentals{Revenue,Consumer}", []string{"zillow.nyc.rentalsConsumer", "zillow.nyc.rentalsRevenue", "zillow.seattle.rentalsConsumer", "zillow.seattle.rentalsRevenue"}},
    }

    // Verify
    for _, testCase := range testCases {
        queries := index.ExpandQuery(testCase.pattern)
        sort.Strings(queries)
        if !reflect.DeepEqual(queries, testCase.expectedQueries) {
            t.Errorf("expected %s, got: %s", testCase.expectedQueries, queries)
        }
    }
}

// Test TrieIndex ExpandPattern
func TestTrieIndexExpandPattern(t *testing.T) {
    // Setup
    var index *TrieIndex

    index = NewTrieIndex("carbon-table", '.')

    index.Insert("zillow.seattle.velocity")
    index.Insert("zillow.seattle.velo1city")
    index.Insert("zillow.seattle.velo2city")
    index.Insert("zillow.seattle.rentalsConsumer")
    index.Insert("zillow.seattle.rentalsRevenue")
    index.Insert("zillow.seattle.pa")
    index.Insert("zillow.seattle.data")

    index.Insert("zillow.sf1.pa")
    index.Insert("zillow.sf1.data")

    index.Insert("zillow.sf2.pa")
    index.Insert("zillow.sf2.data")

    index.Insert("zillow.nyc.rentalsConsumer")
    index.Insert("zillow.nyc.rentalsRevenue")
    index.Insert("zillow.nyc.data")

    // ExpandPattern
    testCases := []struct {
        pattern string
        expectedQueries []*QueryResult
    }{
        {
            "zillow.seattle.velocity",
            []*QueryResult{
                &QueryResult{query: "zillow.seattle.velocity", isLeaf: true},
            },
        },
        {
            "zillow.sf1.*",
            []*QueryResult{
                &QueryResult{query: "zillow.sf1.data", isLeaf: true},
                &QueryResult{query: "zillow.sf1.pa", isLeaf: true},
            },
        },
        {
            "zillow.*.data",
            []*QueryResult{
                &QueryResult{query: "zillow.nyc.data", isLeaf: true},
                &QueryResult{query: "zillow.seattle.data", isLeaf: true},
                &QueryResult{query: "zillow.sf1.data", isLeaf: true},
                &QueryResult{query: "zillow.sf2.data", isLeaf: true},
            },
        },
        {
            "zillow.sf[0-9].data",
            []*QueryResult{
                &QueryResult{query: "zillow.sf1.data", isLeaf: true},
                &QueryResult{query: "zillow.sf2.data", isLeaf: true},
            },
        },
        {
            "zillow.seattle.velo[1-9]city",
            []*QueryResult{
                &QueryResult{query: "zillow.seattle.velo1city", isLeaf: true},
                &QueryResult{query: "zillow.seattle.velo2city", isLeaf: true},
            },
        },
        {
            "zillow.*.rentals{Revenue,Consumer}",
            []*QueryResult{
                &QueryResult{query: "zillow.nyc.rentalsConsumer", isLeaf: true},
                &QueryResult{query: "zillow.nyc.rentalsRevenue", isLeaf: true},
                &QueryResult{query: "zillow.seattle.rentalsConsumer", isLeaf: true},
                &QueryResult{query: "zillow.seattle.rentalsRevenue", isLeaf: true},
            },
        },
    }

    // Verify
    for _, testCase := range testCases {
        queryResults := index.ExpandPattern(testCase.pattern)
        sort.Slice(queryResults, func(i, j int) bool {
          return queryResults[i].query < queryResults[j].query
        })
        if !reflect.DeepEqual(queryResults, testCase.expectedQueries) {
            t.Errorf("failed with pattern %s", testCase.pattern)
        }
    }
}