package api

import (
    "fmt"
    "testing"
    "time"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "go.uber.org/zap"
    "github.com/stretchr/testify/assert"
    "github.com/yunstanford/carbon-table/cfg"
    "github.com/yunstanford/carbon-table/table"
    "github.com/yunstanford/carbon-table/trie"
)

// Test Api Full
func TestApiFull(t *testing.T) {
    // Setup
    config := cfg.NewConfig()
    tbl := table.NewTable(config.Table)
    logger, _ := zap.NewProduction()
    a := NewApi(config.Api, tbl, logger)
    now := tbl.IndexVersion

    tbl.Insert("carbon.cache.a")

    // Test Ping
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/ping", nil)
    a.router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())

    // Test Version
    w = httptest.NewRecorder()
    req, _ = http.NewRequest("GET", "/version", nil)
    a.router.ServeHTTP(w, req)

    version := fmt.Sprintf("{\"version\":\"%s\"}", now.Format(time.RFC3339))

    assert.Equal(t, 200, w.Code)
    assert.Equal(t, version, w.Body.String())

    // Test ExpandQuery
    w = httptest.NewRecorder()
    req, _ = http.NewRequest("GET", "/metric/query/carbon.cache.*/", nil)
    a.router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Equal(t, "[\"carbon.cache.a\"]", w.Body.String())

    // Test ExpandPattern
    w = httptest.NewRecorder()
    req, _ = http.NewRequest("GET", "/metric/pattern/carbon.cache.*/", nil)
    a.router.ServeHTTP(w, req)

    var queries []trie.QueryResult
    if err := json.Unmarshal(w.Body.Bytes(), &queries); err != nil {
        panic(err)
    }

    assert.Equal(t, 200, w.Code)
    assert.Equal(t, []trie.QueryResult{trie.QueryResult{Query:"carbon.cache.a", IsLeaf:true}}, queries)
}
