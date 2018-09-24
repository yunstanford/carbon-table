package trie

const (
    INDEX_ROOT_NAME = "ROOT"
)

// Trie Index
type TrieIndex struct {
    name  string
    sep   rune
    root  *Node
}

// NewTrieIndex
func NewTrieIndex(name string, sep rune) *TrieIndex {
    root = NewNode(false, INDEX_ROOT_NAME, '.')
    return &TrieIndex{
        name: name,
        sep: sep,
        root: root,
    }
}

// Insert
func (trieIndex *TrieIndex) Insert(metric string) {
    // split

    // insert recursively

}

// ExpandQuery
func (trieIndex *TrieIndex) ExpandQuery(query string) []string {
    return trieIndex.root.ExpandQuery(query)
}

// ExpandPattern
func (trieIndex *TrieIndex) ExpandPattern(pattern string) []string {
    return trieIndex.root.ExpandPattern(pattern)
}

// insert
func insert(parent *TrieIndex, metricParts []string) {
    
}