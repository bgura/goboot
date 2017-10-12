package goboot

import (
    "net/http"
)

type Route struct {
    Name      string
    Method    string
    Pattern   string
    Handler   http.HandlerFunc
}

type Routes []Route
