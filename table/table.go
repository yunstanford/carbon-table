package table

import (
    "github.com/yunstanford/carbon-table/cfg"
    "github.com/yunstanford/carbon-table/trie"
)

const (
    INDEX_ROOT_NAME = "ROOT"
)

// Table
type Table struct {
    index   *trie.Node
    ttl     int
}

// NewTable
func NewTable(config *cfg.tableConfig) *Table {
    root = trie.NewNode(false, INDEX_ROOT_NAME, '.')
    return &Table {
        index: root,
        ttl:   config.Ttl,
    }
}
