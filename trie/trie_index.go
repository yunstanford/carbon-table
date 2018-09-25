package trie

import (
    "strings"
)


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
    root := NewNode(false, INDEX_ROOT_NAME, '.')
    return &TrieIndex{
        name: name,
        sep: sep,
        root: root,
    }
}

// Insert
func (trieIndex *TrieIndex) Insert(metric string) {
    // split
    metricParts := strings.Split(metric, string(trieIndex.sep))

    // insert recursively
    insert(trieIndex.root, metricParts)
}

// ExpandQuery
func (trieIndex *TrieIndex) ExpandQuery(query string) []string {
    return trieIndex.root.ExpandQuery(query)
}

// ExpandPattern
func (trieIndex *TrieIndex) ExpandPattern(pattern string) []*QueryResult {
    return trieIndex.root.ExpandPattern(pattern)
}

// insert
func insert(parent *Node, metricParts []string) {
    if len(metricParts) == 0 {
        return
    }
    if len(metricParts) == 1 {
        parent.Insert(NewNode(true, metricParts[0], parent.sep))
        return
    }
    if parent.Get(metricParts[0]) != nil {
        parent.Insert(NewNode(false, metricParts[0], parent.sep))
    }
    insert(parent.Get(metricParts[0]), metricParts[1:])
}