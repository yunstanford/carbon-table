package table

import (
    "testing"
    "reflect"
    "sort"
    "github.com/yunstanford/carbon-table/cfg"
    "github.com/yunstanford/carbon-table/trie"
)

// Test NewTrieIndex
func TestNewTable(t *testing.T) {
    tbl := NewTable(&cfg.TableConfig{Ttl: 3600})
    index := tbl.GetIndex()
    root := index.GetRoot()

    if root.IsLeaf() != false {
        t.Errorf("expected isLeaf false, got: %t", root.IsLeaf())
    }
    if root.Count() != 0 {
        t.Errorf("expected size 0, got: %d", root.Count())
    }
    if tbl.GetTtl() != 3600 {
        t.Errorf("expected ttl 3600, got: %d", tbl.GetTtl())
    }
}

// Test Insert
func TestInsert(t *testing.T) {
    tbl := NewTable(&cfg.TableConfig{Ttl: 3600})
    tbl.Insert("carbon.cache.a")
    index := tbl.GetIndex()
    root := index.GetRoot()

    carbon := root.Get("carbon")
    if carbon == nil {
        t.Errorf("not found carbon")
    }

    cache := carbon.Get("cache")
    if cache == nil {
        t.Errorf("not found cache")
    }

    a := cache.Get("a")
    if a == nil {
        t.Errorf("not found a")
    }
}


// Test ExpandQuery
func TestExpandQuery(t *testing.T) {
    tbl := NewTable(&cfg.TableConfig{Ttl: 3600})

    tbl.Insert("zillow.seattle.velocity")
    tbl.Insert("zillow.seattle.velo1city")
    tbl.Insert("zillow.seattle.velo2city")
    tbl.Insert("zillow.seattle.rentalsConsumer")
    tbl.Insert("zillow.seattle.rentalsRevenue")
    tbl.Insert("zillow.seattle.pa")
    tbl.Insert("zillow.seattle.data")

    tbl.Insert("zillow.sf1.pa")
    tbl.Insert("zillow.sf1.data")

    tbl.Insert("zillow.sf2.pa")
    tbl.Insert("zillow.sf2.data")

    tbl.Insert("zillow.nyc.rentalsConsumer")
    tbl.Insert("zillow.nyc.rentalsRevenue")
    tbl.Insert("zillow.nyc.data")

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
        queries := tbl.ExpandQuery(testCase.pattern)
        sort.Strings(queries)
        if !reflect.DeepEqual(queries, testCase.expectedQueries) {
            t.Errorf("expected %s, got: %s", testCase.expectedQueries, queries)
        }
    }
}

// Test ExpandPattern
func TestExpandPattern(t *testing.T) {
    tbl := NewTable(&cfg.TableConfig{Ttl: 3600})

    tbl.Insert("zillow.seattle.velocity")
    tbl.Insert("zillow.seattle.velo1city")
    tbl.Insert("zillow.seattle.velo2city")
    tbl.Insert("zillow.seattle.rentalsConsumer")
    tbl.Insert("zillow.seattle.rentalsRevenue")
    tbl.Insert("zillow.seattle.pa")
    tbl.Insert("zillow.seattle.data")

    tbl.Insert("zillow.sf1.pa")
    tbl.Insert("zillow.sf1.data")

    tbl.Insert("zillow.sf2.pa")
    tbl.Insert("zillow.sf2.data")

    tbl.Insert("zillow.nyc.rentalsConsumer")
    tbl.Insert("zillow.nyc.rentalsRevenue")
    tbl.Insert("zillow.nyc.data")

    // ExpandPattern
    testCases := []struct {
        pattern string
        expectedQueries []*trie.QueryResult
    }{
        {
            "zillow.seattle.velocity",
            []*trie.QueryResult{
                &trie.QueryResult{Query: "zillow.seattle.velocity", IsLeaf: true},
            },
        },
        {
            "zillow.sf1.*",
            []*trie.QueryResult{
                &trie.QueryResult{Query: "zillow.sf1.data", IsLeaf: true},
                &trie.QueryResult{Query: "zillow.sf1.pa", IsLeaf: true},
            },
        },
        {
            "zillow.*.data",
            []*trie.QueryResult{
                &trie.QueryResult{Query: "zillow.nyc.data", IsLeaf: true},
                &trie.QueryResult{Query: "zillow.seattle.data", IsLeaf: true},
                &trie.QueryResult{Query: "zillow.sf1.data", IsLeaf: true},
                &trie.QueryResult{Query: "zillow.sf2.data", IsLeaf: true},
            },
        },
        {
            "zillow.sf[0-9].data",
            []*trie.QueryResult{
                &trie.QueryResult{Query: "zillow.sf1.data", IsLeaf: true},
                &trie.QueryResult{Query: "zillow.sf2.data", IsLeaf: true},
            },
        },
        {
            "zillow.seattle.velo[1-9]city",
            []*trie.QueryResult{
                &trie.QueryResult{Query: "zillow.seattle.velo1city", IsLeaf: true},
                &trie.QueryResult{Query: "zillow.seattle.velo2city", IsLeaf: true},
            },
        },
        {
            "zillow.*.rentals{Revenue,Consumer}",
            []*trie.QueryResult{
                &trie.QueryResult{Query: "zillow.nyc.rentalsConsumer", IsLeaf: true},
                &trie.QueryResult{Query: "zillow.nyc.rentalsRevenue", IsLeaf: true},
                &trie.QueryResult{Query: "zillow.seattle.rentalsConsumer", IsLeaf: true},
                &trie.QueryResult{Query: "zillow.seattle.rentalsRevenue", IsLeaf: true},
            },
        },
    }

    // Verify
    for _, testCase := range testCases {
        queryResults := tbl.ExpandPattern(testCase.pattern)
        sort.Slice(queryResults, func(i, j int) bool {
          return queryResults[i].Query < queryResults[j].Query
        })
        if !reflect.DeepEqual(queryResults, testCase.expectedQueries) {
            t.Errorf("failed with pattern %s", testCase.pattern)
        }
    }
}

// Test IndexRefresh
func TestIndexRefresh(t *testing.T) {
    tbl := NewTable(&cfg.TableConfig{Ttl: 3600})
    tbl.mirroringPeriod = 1
    tbl.Insert("carbon.cache.a")

    IndexRefresh(tbl)

    if tbl.new_index != nil {
        t.Errorf("new_index is not nil")
    }

    if tbl.mirroring != false {
        t.Errorf("tbl.mirroring == true")
    }

    queries := tbl.ExpandQuery("carbon.*.a")

    if len(queries) != 0 {
        t.Errorf("queries is not empty!")
    }
}
