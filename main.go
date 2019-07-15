package main

import (
	util "github.com/eagle7410/go_util/libs"
	"github.com/gorilla/handlers"
	"go_frame/lib"
	"log"
	"net/http"
)

const port = ":8080"

func init() {

	util.OpenLogFile()

	if err := lib.ENV.Init(); err != nil {
		util.LogFatalf("Error on initializing environment : %s", err)
	}

	util.Env.SetEnv(&lib.ENV.IsDev, &lib.ENV.AllowedMethods)
}

func main() {

	router := lib.GetRouter()

	util.LogAppRun(port)

	middleware := util.SetCorsMiddleware(
		util.LogRequest(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"*"}),
				handlers.AllowedMethods(lib.ENV.AllowedMethods),
				handlers.AllowedOrigins([]string{"*"}))(router)))

	log.Fatal(http.ListenAndServe(port, middleware))
}
