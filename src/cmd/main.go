package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ansel1/merry/v2"
	"github.com/bcc-code/brunstadtv-lite-api/src/generated"
	"github.com/bcc-code/brunstadtv-lite-api/src/graph"
	"github.com/bcc-code/brunstadtv-lite-api/src/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// Defining the Graphql handler
func graphqlHandler(queries *sqlc.Queries) gin.HandlerFunc {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					Queries: queries,
				},
			},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	dbPath := os.Getenv("DB_PATH")
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if dbPath == "" {
		log.Panic().Msg("Usage: DB_PATH=/path/to/db.sqlite3 server")
	}

	if _, err := os.Stat(dbPath); err != nil {
		log.Panic().Err(merry.Wrap(err)).Msg("unable to open db")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		err = merry.Wrap(err)
		log.Panic().Err(err).Msg("unable to open DB")
	}

	// TODO: Validate DB version

	queries := sqlc.New(db)

	r := gin.Default()
	r.POST("/query", graphqlHandler(queries))
	r.GET("/", playgroundHandler())
	r.Run()
}
