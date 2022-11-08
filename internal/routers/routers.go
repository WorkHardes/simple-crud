package routers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Init() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", IndexHandler)

	return router
}

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello, World!")
}
