package gin_graphql

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	ginMiddlewares "github.com/Laisky/gin-middlewares"
)

var (
	auth *ginMiddlewares.Auth
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

func setupAuth() (err error) {
	auth, err = ginMiddlewares.NewAuth([]byte("settings.secret"))
	return
}

func RunServer() {
	if err := setupAuth(); err != nil {
		log.Fatal("try to setup auth got error", err)
	}
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	engine.Use(CORSMiddleware())
	// engine.Use(cors.Default())
	h := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))
	// h := handler.New(NewExecutableSchema(Config{Resolvers: &Resolver{}}))
	engine.POST("/query", FromStd(h.ServeHTTP))
	engine.GET("/", FromStd(playground.Handler("GraphQL", "/query")))

	engine.Run(":8090")
}
