package trie

import (
    "testing"
    "reflect"
    "sort"
)

// Simply make sure creating a new tree works.
func TestNewTrieNode(t *testing.T) {
    var node *Node
    node = NewNode(true, "carbon", '.')

    if node.Count() != 0 {
        t.Errorf("expected size 0, got: %d", node.Count())
    }

    if node.isLeaf != true {
        t.Errorf("expected isLeaf true, got: %t", node.isLeaf)
    }

    if node.name != "carbon" {
        t.Errorf("expected name carbon, got: %s", node.name)
    }
}

// Testing node Insert works.
func TestNodeInsertAndGet(t *testing.T) {
    var root *Node
    var childNode *Node
    root = NewNode(false, "root", '.')
    childNode = NewNode(true, "carbon", '.')

    // Insert
    root.Insert(childNode)

    // Verify
    fromRoot := root.Get("carbon")
    if fromRoot != childNode {
        t.Errorf("childNode carbon doens't exist")
    }
}

// Testing node Delete works.
func TestNodeDelete(t *testing.T) {
    var root *Node
    var childNode *Node
    root = NewNode(false, "root", '.')
    childNode = NewNode(true, "carbon", '.')

    // Insert
    root.Insert(childNode)

    // Delete
    root.Delete("carbon")

    // Verify
    fromRoot := root.Get("carbon")
    if fromRoot != nil {
        t.Errorf("childNode carbon still exists")
    }
}

// Testing util function RemoveOuterBraces
func TestRemoveOuterBraces(t *testing.T) {
    testCases := []struct {
        input string
        expectedOutput string
    }{
        {"{hello,world}", "hello,world"},
        {"hello,world}", "hello,world}"},
        {"hello,{world}", "hello,{world}"},
    }

    for _, testCase := range testCases {
        output := RemoveOuterBraces(testCase.input)
        if output != testCase.expectedOutput {
            t.Errorf("expected %s, got: %s", testCase.expectedOutput, output)
        }
    }
}

// Testing util function ExpandBraces
func TestExpandBraces(t *testing.T) {
    testCases := []struct {
        input string
        expectedOutput []string
    }{
        {"foo{hello,world}", []string{"foohello", "fooworld"}},
        {"hello,world}", []string{"hello,world}"}},
        {"{foo,bar}", []string{"foo", "bar"}},
    }

    var output []string
    for _, testCase := range testCases {
        output = ExpandBraces(testCase.input, ",")
        if !reflect.DeepEqual(output, testCase.expectedOutput) {
            t.Errorf("expected %s, got: %s", testCase.expectedOutput, output)
        }
    }
}

// Testing node GetAllNode works
func TestNodeGetAllNode(t *testing.T) {
    // Setup
    var root, velocity, velocityOne, velocityTwo, rentalsConsumer, rentalsRevenue *Node
    root = NewNode(false, "zillow", '.')
    velocity = NewNode(true, "velocity", '.')
    velocityOne = NewNode(true, "velo1city", '.')
    velocityTwo = NewNode(true, "velo2city", '.')
    rentalsConsumer = NewNode(true, "rentalsConsumer", '.')
    rentalsRevenue = NewNode(true, "rentalsRevenue", '.')

    // Insert
    root.Insert(velocity)
    root.Insert(velocityOne)
    root.Insert(velocityTwo)
    root.Insert(rentalsRevenue)
    root.Insert(rentalsConsumer)

    // GetAllNode
    testCases := []struct {
        pattern string
        expectedCount int
    }{
        {"*", 5},
        {"velo[0-9]city", 2},
        {"rentals*", 2},
        {"rentals{Consumer,Revenue}", 2},
    }

    for _, testCase := range testCases {
        allNodes := root.GetAllNode(testCase.pattern)
        if len(allNodes) != testCase.expectedCount {
            t.Errorf("expected count %d, got: %d", testCase.expectedCount, len(allNodes))
        }
    }
}

// Testing node ExpandQuery works
func TestNodeExpandQuery(t *testing.T) {
    // Setup
    var root *Node
    var zillow *Node
    var seattle, sf1, sf2, nyc *Node
    var rentalsConsumer, rentalsRevenue, velocity, velocityOne, velocityTwo, pa, data *Node

    root = NewNode(false, "root", '.')
    zillow = NewNode(false, "zillow", '.')
    seattle = NewNode(false, "seattle", '.')
    sf1 = NewNode(false, "sf1", '.')
    sf2 = NewNode(false, "sf2", '.')
    nyc = NewNode(false, "nyc", '.')
    rentalsConsumer = NewNode(true, "rentalsConsumer", '.')
    rentalsRevenue = NewNode(true, "rentalsRevenue", '.')
    velocity = NewNode(true, "velocity", '.')
    velocityOne = NewNode(true, "velo1city", '.')
    velocityTwo = NewNode(true, "velo2city", '.')
    pa = NewNode(true, "pa", '.')
    data = NewNode(true, "data", '.')

    seattle.Insert(velocity)
    seattle.Insert(velocityOne)
    seattle.Insert(velocityTwo)
    seattle.Insert(rentalsRevenue)
    seattle.Insert(rentalsConsumer)
    seattle.Insert(pa)
    seattle.Insert(data)
    sf1.Insert(pa)
    sf1.Insert(data)
    sf2.Insert(pa)
    sf2.Insert(data)
    nyc.Insert(rentalsConsumer)
    nyc.Insert(rentalsRevenue)
    nyc.Insert(data)
    zillow.Insert(seattle)
    zillow.Insert(sf1)
    zillow.Insert(sf2)
    zillow.Insert(nyc)
    root.Insert(zillow)

    // ExpandQuery
    testCases := []struct {
        pattern string
        expectedQueries []string
    }{
        {"zillow.seattle.velocity", []string{"zillow.seattle.velocity"}},
        {"zillow.sf1.*", []string{"zillow.sf1.data", "zillow.sf1.pa"}},
        {"zillow.*.data", []string{"zillow.nyc.data", "zillow.seattle.data", "zillow.sf1.data", "zillow.sf2.data"}},
        {"zillow.sf[0-9].data", []string{"zillow.sf1.data", "zillow.sf2.data"}},
        {"zillow.seattle.velo[1-9]city", []string{"zillow.seattle.velo1city", "zillow.seattle.velo2city"}},
        {"zillow.*.rentals{Revenue,Consumer}", []string{"zillow.nyc.rentalsConsumer", "zillow.nyc.rentalsRevenue", "zillow.seattle.rentalsConsumer", "zillow.seattle.rentalsRevenue"}},
    }

    // Verify
    for _, testCase := range testCases {
        queries := root.ExpandQuery(testCase.pattern)
        sort.Strings(queries)
        if !reflect.DeepEqual(queries, testCase.expectedQueries) {
            t.Errorf("expected %s, got: %s", testCase.expectedQueries, queries)
        }
    }
}