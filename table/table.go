package table

import (
    "github.com/yunstanford/carbon-table/cfg"
    "github.com/yunstanford/carbon-table/trie"
)


// Table
type Table struct {
    index   *trie.Node
    ttl     int
}

// NewTable
func NewTable(config *cfg.tableConfig) *Table {

}