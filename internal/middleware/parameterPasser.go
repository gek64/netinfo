package middleware

import (
    "context"
    "github.com/gin-gonic/gin"
    "netinfo/ent"
)

const (
    Client  = "database"
    Context = "context"
)

func ParameterPasser(client *ent.Client, ctx context.Context) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set(Client, client)
        c.Set(Context, ctx)
        c.Next()
    }
}
