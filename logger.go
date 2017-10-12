package goboot

import (
    "log"
    "time"
    "fmt"

    "net/http"
    "net/http/httputil"
)

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        debugHttp(r)

        inner.ServeHTTP(w, r)

        log.Printf( "%s\t%s\t%s\t%s\n", r.Method, r.RequestURI,
            name, time.Since(start),
        )
    })
}

func debugHttp(r *http.Request) {
  if d, err := httputil.DumpRequest(r, true); err != nil {
    fmt.Println(err)
    fmt.Println(string(d))
  } else {
    fmt.Println(string(d))
  }
}
