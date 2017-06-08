package main

import (
  "fmt"
  "encoding/json"
  "reflect"
  "time"
)

var input = `
{
  "created_at": "Thu Jun 01 00:00:01 +0000 2017"
}
`
type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
  v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
  if err != nil {
    return err
  }
  *t = Timestamp(v)
  return nil
}

func main() {
  var val map[string]Timestamp

  if err := json.Unmarshal([]byte(input), &val); err != nil {
    panic(err)
  }

  fmt.Println(val)
  for k, v := range val {
    fmt.Println(k, reflect.TypeOf(v))
  }
  fmt.Println(time.Time(val["created_at"]))
}
