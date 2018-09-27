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
    // trie expand
    queryResults := a.table.ExpandQuery("blablabla")

    a.router.GET("/metric/:query/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
}


// Add routes for Api
func AddRoutes(a *Api) {
    AddHealthPing(a)
    AddExpandQuery(a)

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