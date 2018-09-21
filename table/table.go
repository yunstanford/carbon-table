package table

import (
	"github.com/yunstanford/carbon-table/api"
	"github.com/yunstanford/carbon-table/trie"
	"time"
)


// Table
type Table struct {
    index  		 *trie.Node
    tableConfig  *TableConfig
    apiHandler	 *api.Api
}

