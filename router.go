package goboot

import (
    "log"

    "net/http"
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    
    return router.PathPrefix("/{version}/").Subrouter()
}

func RegisterRoutes(m *mux.Router, r Routes) {
  for _, route := range r {
    var handler http.Handler
    handler = route.Handler
    handler = Logger(handler, route.Name)
    //handler = SecureHeaders(handler, route.Name)

    log.Println("Register Route", route.Name, route.Pattern, route.Method)

    m.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)
  }
}
