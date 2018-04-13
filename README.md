goboot is a simple framework for getting a RESTful JSON web api up and running using Go

### Usage

1. Implement a Controller
```golang

package myapi

// Controller for handling the api functions
type myapiController struct {
  goboot.Controller
}

// Default/Index 'landing' for the user controller
func (c *myapiController) Index(w http.ResponseWriter, r *http.Request) {
  myobject := myObject{ Text: "Hello World!")}
  if err := c.SendPayload(w, "", p); err != nil {
    c.SendError(w, http.StatusInternalServerError)
  }
}

```

2. Export the Routes by Implementing goboot.Api
```golang
package myapi

import (
  "github.com/bgura/goboot"
)

// Define the the api
type myApi struct {
	routes []goboot.Route
}

// Create an instance of this api via a public ctor
func NewMyApi() goboot.Api {
  api := myApi{}

  r := []goboot.Route{
    goboot.Route{
      Name:    "Index",
      Method:  "GET",
      Pattern: "/myapi",
      Handler: api.Index,
    },
    // More routes...
  }

  return &myApi{ routes: r }
}

// Implement the goboot.Api interface for myApi
func (api *myApi) GetRoutes() []goboot.Route {
  return api.routes
}

```
3. Create a Router and Register the Routes

```golang
package main

import (
  "net/http"
  "github.com/bgura/goboot"
)

func main() {
  router     := goboot.NewRouter()
  controller := myapi.NewMyApi()

  gooboot.RegisterRoutes(router, controller.GetRoutes())

  http.ListenAndServe("127.0.0.1:8080", router)
} 
```

4. Compile and Run!
```
go build -o server
./server
```

5. Test the Endpoint
```
wget http://127.0.0.1:8080/v1/myapi

{ "Text" : "Hello World!"  }
```

### Versioning Schema

Goboot has a required versioning built into the api. All routes defined have a prefix of http://127.0.0.1:8080/"version"/. The version can be queried in the controller by invoking ```GetVersion(r *http.Request) string;``` on a goboot controller. As indicated by the signature, the value is a string and interpretation needs to be managed by the controller

### QueryController vs. Controller

As a bare minimum, all endpoints should be implemented with the goboot.Controller. The goboot.QueryController is an extended implementation which provides coveinences for reading parameters. The query controller has a member, ```ParamReader  goboot.paramReader```.

To use the param reader, you must register the request that that you wish to parse with the ParamReader. An example implementation is below:
```golang
func (c *MyController) Index(w http.ResponseWriter, r *http.Request) {
  c.ParamReader.Context(r)

  c.ParamReader.Optional("id", "0", goboot.Uint64)
  c.ParamReader.Optional("username", "me", goboot.String)

  c.ParamReader.Require("email", goboot.String)

  if c.ParamReader.HasErrors() {
    c.SendError(w, http.StatusBadRequest)
    return;
  }

  id       := c.ParamReader.ReadUint64("id")
  username := c.ParamReader.ReadString("username")

  email    := c.ParamReader.ReadString("email")
  
  // remaining impl...

```

In the above example, the user must provide the query paramter email, while id and username are both optional. If the user does not provide an id in the request, it will default to 0, and username will default to "me" when we read the parameters.

Its important to call the HasError function before attempting to read the parameters. The function will return a true or false based on the input values and the validations provided to it. If you invoke the "Read" functions without validating, its possible a panic will be raise upon an invalid conversion


