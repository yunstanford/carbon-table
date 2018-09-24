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
