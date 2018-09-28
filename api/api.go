package api

import (
    "github.com/gin-gonic/gin"
    "github.com/yunstanford/carbon-table/cfg"
    "github.com/yunstanford/carbon-table/table"
)

type Api struct {
    addr   string
    router *gin.Engine
    table  *table.Table
}


// Handlers
///////////

// health check handler
func AddHealthPing(a *Api) {
    a.router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
}

// metric expand query handler
func AddExpandQuery(a *Api) {
    a.router.GET("/metric/query/:query/", func(c *gin.Context) {
        query := c.Param("query")
        // trie expand
        queryResults := a.table.ExpandQuery(query)
        c.JSON(200, queryResults)
    })
}

// metric expand query handler
func AddExpandPattern(a *Api) {
    a.router.GET("/metric/pattern/:pattern/", func(c *gin.Context) {
        pattern := c.Param("pattern")
        // trie expand
        queryResults := a.table.ExpandPattern(pattern)
        c.JSON(200, queryResults)
    })
}

// Add routes for Api
func AddRoutes(a *Api) {
    AddHealthPing(a)
    AddExpandQuery(a)
    AddExpandPattern(a)

    // Add more routes here...
}

// Start
func (a *Api) Start() {
    a.router.Run(a.addr)
}

// NewApi
func NewApi(c *cfg.ApiConfig, t *table.Table) *Api{
    a := &Api {
        addr:   c.ApiAddr,
        router: gin.Default(),
        table:  t,
    }
    // add routes
    AddRoutes(a)
    return a
}