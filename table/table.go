package table

import (
	"time"
	"github.com/yunstanford/carbon-table/api"
	"github.com/yunstanford/carbon-table/trie"
)


// Table
type Table struct {
    index  		 *trie.Node
    tableConfig  *TableConfig
    apiHandler	 *api.Api
}

// NewTable
func NewTable() *Table {

}