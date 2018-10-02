package table

import (
    "time"
    "github.com/yunstanford/carbon-table/cfg"
    "github.com/yunstanford/carbon-table/trie"
)

const (
    INDEX_NAME = "carbon-table"
)

// Table
type Table struct {
    index   *trie.TrieIndex
    ttl     int

    // for ttl
    newIndex         *trie.TrieIndex
    mirroring        bool
    mirroringPeriod  int

    // Index Version
    IndexVersion     time.Time
}

// NewTable
func NewTable(config *cfg.TableConfig) *Table {
    root := trie.NewTrieIndex(INDEX_NAME, '.')
    tbl := &Table {
        index: root,
        ttl:   config.Ttl,
        // for ttl
        newIndex: nil,
        mirroring: false,
        mirroringPeriod: config.Period,
        // IndexVersion
        IndexVersion: time.Now(),
    }

    // Setup tick
    quit := make(chan bool, 1)

    go func() {
        ttl := time.NewTicker(time.Second * time.Duration(tbl.ttl))

        for { //ever
            select {
            case <-quit:
                return
            case <-ttl.C:
                go IndexRefresh(tbl)
            }
        }

    }()

    return tbl
}

// Insert
func (t *Table) Insert(metric string) {
    t.index.Insert(metric)
}

// ExpandQuery
func (t *Table) ExpandQuery(query string) []string {
    return t.index.ExpandQuery(query)
}

// ExpandPattern
func (t *Table) ExpandPattern(pattern string) []*trie.QueryResult {
    return t.index.ExpandPattern(pattern)
}

// GetIndex
func (t *Table) GetIndex() *trie.TrieIndex {
    return t.index
}

// GetTtl
func (t *Table) GetTtl() int {
    return t.ttl
}

// IndexRefresh
func IndexRefresh(tbl *Table) {
    // Setup new index
    tbl.mirroring = true
    tbl.newIndex = trie.NewTrieIndex(INDEX_NAME, '.')
    waiter := time.NewTimer(time.Second * time.Duration(tbl.mirroringPeriod))

    // Cummulate new data
    <- waiter.C

    // Swap and Reset
    tbl.index = tbl.newIndex
    tbl.IndexVersion = time.Now()
    tbl.newIndex = nil
    tbl.mirroring = false
}
