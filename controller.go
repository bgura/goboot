package goboot

import (
  "io"
  "strconv"
  "net/http"

  "encoding/json"
  
  "github.com/gorilla/mux"
)

const MAX_JSON_SIZE int64 = 1048576

// Struct which handles http endpoints.
// Has helpers which aid in send & recieve of JSON data
type Controller struct {

}

// Send back an object 'v' serialized in json to the ResponseWriter
func (c *Controller) SendJson(w http.ResponseWriter, v interface{}) error {
  w.Header().Add("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  return json.NewEncoder(w).Encode(v)
}

// Send an empty json response with http status code 's'
func(c *Controller) SendError(w http.ResponseWriter, s int)  {
  w.Header().Add("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(s)
}

// Read the json body of a request and parse into provided interface 'v'
// Number of bytes to be read can be limited with 's'
func(c *Controller) ReadJson(r *http.Request, s int64, v interface{}) error {
  d := json.NewDecoder( io.LimitReader(r.Body, s))

  return d.Decode(&v);
}

// Send a json response back to the client with a given request id and payload
func(c *Controller) SendPayload(w http.ResponseWriter, s string, v interface{}) error {
  msg := JsonResponse{
    RequestId: s,
    Payload: v,
  }
  return c.SendJson(w, msg)
}
/*
This function can determine if a given string (the body of a request) is a json
array or a single json object. We can't safely parse one over the other and must
know which type it is
func(c *Controller) IsJsonArray(b string) err {
    b, err := ioutil.ReadAll(r)
    if err != nil {
      return nil, err
    }

    isArray := false
    for _, c := range b {
        if ' ' == c || '\t' == c || '\r' == c || '\n' == c {
          continue
        }
        isArray = c = '['
        break
    }
}
*/

// Get the api version which was requested
func (c *Controller) GetVersion(r *http.Request) string {
  return mux.Vars(r)["version"]
}

// Read a var uint64. Mux should provide a facility for validating
// that the var is present as its a required piece of the routing
func (c *Controller) ReadVarUint64(r *http.Request, s string) uint64 {
	if v, ok := mux.Vars(r)[s]; ok {
		// Try to convert the parameter
		if i, err := strconv.ParseUint(v, 10, 64); err != nil {
			panic("Failed to convert parameter to Uint64")
		} else {
			return i
		}
	} else {
		panic("Parameter not found and no default defined")
	}
}