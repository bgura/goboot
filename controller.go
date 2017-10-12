package goboot

import (
  "io"
  "io/ioutil"

  "net/http"

  "encoding/json"
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
