package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/mateors/mql"
	"github.com/mateors/mql/database/couchbase"
	"github.com/mateors/mtool"
)

var err error
var PORT string

var DRIVER, HOST, DBNAME, DBUSER, DBPASS, DBPORT, DBSCOPE string
var db *sql.DB

type Account struct {
	ID          string `json:"id"`           //system auto generated
	Type        string `json:"type"`         //account
	AccountType string `json:"account_type"` //
	AccountName string `json:"account_name"` //
	FirstName   string `json:"first_name"`   //
	LastName    string `json:"last_name"`    //
	Phone       string `json:"phone"`        //
	Email       string `json:"email"`        //username
	Remarks     string `json:"remarks"`      //
	IpAddress   string `json:"ip_address"`
	CreateDate  string `json:"create_date"`
	Status      int    `json:"status"`
}

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT = os.Getenv("PORT")
	DBUSER = os.Getenv("DBUSER")
	DBPASS = os.Getenv("DBPASS")
	HOST = os.Getenv("HOST")
	DBNAME = os.Getenv("DBNAME")
	DBPORT = os.Getenv("DBPORT")
	DRIVER = os.Getenv("DRIVER")
	DBSCOPE = os.Getenv("DBSCOPE")

	fmt.Println("Bismillah, graphqlgo running on", PORT)
	//Couchbase database connection string
	dataSourceName := fmt.Sprintf("http://%s:%s@%s:%s", DBUSER, DBPASS, HOST, DBPORT)
	fmt.Println(dataSourceName)
	pdb, err := couchbase.New(dataSourceName)
	if err != nil {
		fmt.Println("Error connecting to Couchbase:", err)
		return
	}
	db = pdb.DB
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db connection successful.")

	mql.BUCKET = DBNAME
	mql.DRIVER = DRIVER
	mql.SCOPE = DBSCOPE
	mql.RegisterModel(Account{})

	// sql := fmt.Sprintf("SELECT * FROM %s.%s.%s WHERE email=%q", DBNAME, DBSCOPE, "account", "admin@bosemann.com")
	// fmt.Println(sql)
	// rows, err := mql.GetRows(sql, db)
	// fmt.Println(err, rows)
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
		Name: "Account", //account
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
		Description: "Create new account or signup",
		Args: graphql.FieldConfigArgument{
			"first_name": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"last_name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"email":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)}, //username
			"passwd":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			//smap := p.Source.(map[string]interface{})
			//ipAddress := cleanIp(smap["RemoteAddr"].(string))
			//log.Println(">>", smap["RemoteAddr"], smap["UserAgent"], smap["Referer"], smap["ip"], mtool.TimeNow())
			return nil, nil
		},
	}
}

func GetAcountField() *graphql.Field {

	return &graphql.Field{
		Type:        AccountType,
		Description: "Get Account details",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)}, //username
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			//smap := p.Source.(map[string]interface{})
			//ipAddress := cleanIp(smap["RemoteAddr"].(string))
			//log.Println(">>", smap["RemoteAddr"], smap["UserAgent"], smap["Referer"], smap["ip"], mtool.TimeNow())

			//colmap := tableFields(p.Info.FieldASTs)
			//funclabel := fmt.Sprint(p.Info.Path.Key)
			//cols := colmap[funclabel].([]string)
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

func tableFields(FieldASTs []*ast.Field) map[string]interface{} {

	var fields = make(map[string]interface{})
	for _, val := range FieldASTs {
		var cols []string
		for _, sel := range val.SelectionSet.Selections {
			field, ok := sel.(*ast.Field)
			if ok {
				if field.Name.Kind == "Name" {
					cols = append(cols, field.Name.Value)
				}
			}
		}
		fields[val.Name.Value] = cols
	}
	return fields
}
