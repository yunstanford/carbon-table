package api


import (
    "github.com/gin-gonic/gin"
)


type Api struct {
    addr   string
    router *gin.Engine
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
