package gin_graphql

import (
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RunServer() {
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	engine.Use(CORSMiddleware())
	// engine.Use(cors.Default())
	h := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

	engine.POST("/query", FromStd(h.ServeHTTP))
	engine.GET("/", FromStd(playground.Handler("GraphQL", "/query")))

	engine.Run(":8090")
}
