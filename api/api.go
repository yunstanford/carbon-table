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


func (a *Api) Start() {
    a.router.Run(a.addr)
}


// Add routes for Api
func AddRoutes(r *gin.Engine) {
    // health check ping
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
}

// NewApi
func NewApi(config *cfg.apiConfig) *Api{
    
}