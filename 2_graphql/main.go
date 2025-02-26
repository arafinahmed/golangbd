package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/mateors/mtool"
)

var err error
var PORT string

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT = os.Getenv("PORT")
	fmt.Println("Bismillah, graphqlgo running on", PORT)
}

var rootQueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{

			// http://localhost:8080/graphql?query={getAccount(email:"mostain@lxroot.com"){first_name,last_name,status}}
			"getAccount": GetAcountField(),
		},
	})

var rootMutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{

		// http://localhost:8080/graphql?query=mutation+_{createAccount(first_name:"mostain",last_name:"billah",email:"mostain@lxroot.com",passwd:"test321"){id,email,status}}
		"createAccount": CreateAcountField(), //
	},
})

var AccountType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Account",
		Fields: graphql.Fields{
			"id":           &graphql.Field{Type: graphql.String},
			"account_type": &graphql.Field{Type: graphql.String},
			"account_name": &graphql.Field{Type: graphql.String},
			"first_name":   &graphql.Field{Type: graphql.String},
			"last_name":    &graphql.Field{Type: graphql.String},
			"phone":        &graphql.Field{Type: graphql.String},
			"email":        &graphql.Field{Type: graphql.String},
			"remarks":      &graphql.Field{Type: graphql.String},
			"status":       &graphql.Field{Type: graphql.Int},
		},
	},
)

func CreateAcountField() *graphql.Field {

	return &graphql.Field{
		Type:        AccountType,
		Description: "Create new account",
		Args: graphql.FieldConfigArgument{
			"first_name": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"last_name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"email":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)}, //username
			"passwd":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return nil, nil
		},
	}
}

func GetAcountField() *graphql.Field {

	return &graphql.Field{
		Type:        AccountType,
		Description: "Create new account",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)}, //username
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			return nil, nil
		},
	}
}

func rootObjectFn(ctx context.Context, r *http.Request) map[string]interface{} {
	var rmap = make(map[string]interface{})
	rmap["ip"] = mtool.ReadUserIP(r)
	rmap["UserAgent"] = r.UserAgent()
	rmap["RemoteAddr"] = r.RemoteAddr
	rmap["Referer"] = r.Referer()
	return rmap
}

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second)) //processing should be stopped.

	var schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootQueryType,
			Mutation: rootMutationType,
		},
	)
	if err != nil {
		log.Fatal("SCHEMA_ERROR:", err)
	}
	// simplest relay-compliant graphql server HTTP handler
	h := handler.New(&handler.Config{
		Schema:       &schema,
		Pretty:       true,
		RootObjectFn: rootObjectFn,
	})

	// static file server to serve Graphiql in-browser editor
	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	r.Handle("/graphql", h)

	http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
}
