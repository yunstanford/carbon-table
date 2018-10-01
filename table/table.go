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
    new_index        *trie.TrieIndex
    mirroring        bool
    mirroringPeriod  int
}

// NewTable
func NewTable(config *cfg.TableConfig) *Table {
    root := trie.NewTrieIndex(INDEX_NAME, '.')
    tbl := &Table {
        index: root,
        ttl:   config.Ttl,
        // for ttl
        new_index: nil,
        mirroring: false,
        mirroringPeriod: 120,
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
    tbl.mirroring = true
    tbl.new_index = trie.NewTrieIndex(INDEX_NAME, '.')
    waiter := time.NewTimer(time.Second * time.Duration(tbl.mirroringPeriod))

    // Cummulate new data
    <- waiter.C

    // Swap and Reset
    tbl.index = tbl.new_index
    tbl.new_index = nil
    tbl.mirroring = false
}
