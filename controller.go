package goboot

import (
  "strconv"
  "io"
  "fmt"
  "io/ioutil"

  "net/http"

  "encoding/json"
  
  "github.com/gorilla/mux"
)

const MAX_JSON_SIZE int64 = 1048576

type Controller struct {

}

func (c *Controller) SendJson(w http.ResponseWriter, r *http.Request, v interface{}) error {
  w.Header().Add("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  return json.NewEncoder(w).Encode(v)
}

func(c *Controller) SendError(w http.ResponseWriter, r *http.Request, s int)  {
  w.Header().Add("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(s)
}

func(c *Controller) ReadJson(w http.ResponseWriter, r *http.Request, s int64, v interface{}) error {
  d := json.NewDecoder( io.LimitReader(r.Body, s))

  return d.Decode(&v);
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

func (c *Controller) DecodeJson(w http.ResponseWriter, r *http.Request, s int64, v interface{}) error {
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, s))
  if err != nil {
    return err
  }
  return json.Unmarshal(body, &v)
}

// Get the api version which was requested
func (c *Controller) GetVersion(r *http.Request) string {
  if val, ok := mux.Vars(r)["version"]; ok {
    return val
  } else {
    return ""
  }
}

// Get a uint param with given name
func (c *Controller) GetParamUint(s string, r *http.Request) (uint64, error) {
  v := r.URL.Query().Get(s)
  if len(v) == 0 {
    // Default to 0
    return 0, nil
  }
  return strconv.ParseUint(v, 0, 64)
}

// Get a query parameter with given name
func (c *Controller) GetQueryParamUint(s string, r *http.Request) (uint64, error) {
  if val, ok := mux.Vars(r)[s]; ok {
    return strconv.ParseUint(val, 0, 64)
  } else {
    return 0, fmt.Errorf("Missing parameter")
  }
}