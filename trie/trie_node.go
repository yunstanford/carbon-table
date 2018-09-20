package trie

import (
    "path/filepath"
    "strings"
    "bytes"
    "regexp"
)


var EXPAND_BRACES_RE = regexp.MustCompile(`.*(\{.*?[^\\]?\})`)


// Node trie tree node container carries name and children.
type Node struct {
    isLeaf bool
    name  string
    sep    rune
    children  map[string]*Node
}


// NewNode
func NewNode(isLeaf bool, name string, sep rune) *Node {
    return &Node{
        isLeaf: isLeaf,
        name: name,
        sep: sep,
        children: make(map[string]*Node),
    }
}


// Get - get a child node.
func (n *Node) Get(childName string) *Node {
    if child, ok := n.children[childName]; ok {
        return child
    } else {
        return nil
    }
}

// Insert - insert a child node.
func (n *Node) Insert(childNode *Node) {
    if child, ok := n.children[childNode.name]; ok {
        child.isLeaf = child.isLeaf || childNode.isLeaf
    } else {
        n.children[childNode.name] = childNode
    }
}

// Delete - delete a child_node.
func (n *Node) Delete(childName string) bool {
    if _, ok := n.children[childName]; ok {
        delete(n.children, childName)
        return true
    } else {
        return false
    }
}

// IsLeaf - check if node is leaf node.
func (n *Node) IsLeaf() bool {
    return n.isLeaf
}

// RemoveOuterBraces - remove outer braces
func RemoveOuterBraces(s string) string {
    if s[0] == '{' && s[len(s) - 1] == '}' {
        return s[1:len(s) - 1]
    }
    return s
}

// ExpandBraces - expand braces
func ExpandBraces(s string, sep string) []string {
    var res []string
    m := EXPAND_BRACES_RE.FindStringSubmatchIndex(s)
    if len(m) == 0 { // can't find match
        res = append(res, strings.Replace(s, "\\}", "}", -1))
    } else {
        openBrace, closeBrace := m[2], m[3]
        sub := s[openBrace:closeBrace]
        if strings.Contains(sub, sep) {
            for _, pat := range strings.Split(strings.Trim(sub, "{}"), sep) {
                var buffer bytes.Buffer
                buffer.WriteString(s[:openBrace])
                buffer.WriteString(pat)
                buffer.WriteString(s[closeBrace:])
                sub_res := ExpandBraces(buffer.String(), sep)
                res = append(res, sub_res...)
            }
        } else {
            var buffer bytes.Buffer
            buffer.WriteString(s[:openBrace])
            buffer.WriteString(RemoveOuterBraces(sub))
            buffer.WriteString(s[closeBrace:])
            sub_res := ExpandBraces(buffer.String(), sep)
            res = append(res, sub_res...)
        }
    }
    return res
}

// GetAllNode - get all child nodes based on wild card query.
func (n *Node) GetAllNode(pattern string) []*Node {
    // TODO: Add expand braces logic
    var matches []*Node
    patterns := ExpandBraces(pattern, ",")

    for childName, childNode := range n.children {
        for _, p := range patterns {
            matched, err := filepath.Match(p, childName)
            if err != nil {
                // logging
                continue
            }

            if matched {
                matches = append(matches, childNode)
            }
        }
    }
    return matches
}

// ExpandQuery - expand a wildcard query.
func (n *Node) ExpandQuery(query string) []string{
    var queries []string
    sepIndex := strings.IndexRune(query, n.sep)

    if sepIndex < 0 {
        for _, node := range n.GetAllNode(query) {
            if node.isLeaf {
                queries = append(queries, node.name)
            }
        }
    } else {
        curPart := query[:sepIndex]
        curMatches := n.GetAllNode(curPart)
        subQuery := query[sepIndex + 1:]
        for _, m := range curMatches {
            subQueries := m.ExpandQuery(subQuery)
            for _, sq := range subQueries {
                var buffer bytes.Buffer
                buffer.WriteString(m.name)
                buffer.WriteRune(n.sep)
                buffer.WriteString(sq)
                queries = append(queries, buffer.String())
            }   
        }
    }
    return queries
}

// Count - return children count
func (n *Node) Count() int {
    return len(n.children)
}
