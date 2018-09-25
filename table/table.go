package table

import (
    "github.com/yunstanford/carbon-table/cfg"
    "github.com/yunstanford/carbon-table/trie"
)

const (
    INDEX_NAME = "carbon-table"
)

// Table
type Table struct {
    index   *trie.Node
    ttl     int
}

// NewTable
func NewTable(config *cfg.tableConfig) *Table {
    root = trie.NewTrieIndex(INDEX_NAME, '.')
    return &Table {
        index: root,
        ttl:   config.Ttl,
    }
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
