package api

import (
    "testing"
    "net/http"
	"net/http/httptest"
	"go.uber.org/zap"
	"github.com/stretchr/testify/assert"
	"github.com/yunstanford/carbon-table/cfg"
	"github.com/yunstanford/carbon-table/table"
)


// Test Api Full
func TestApiFull(t *testing.T) {
	// Setup
	config := cfg.NewConfig()
	tbl := table.NewTable(config.Table)
	logger, _ := zap.NewProduction()
	a := NewApi(config.Api, tbl, logger)

	// Test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	a.router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}
