package main

import (
  "fmt"
  "os"
  "reflect"
  "strings"

  "github.com/nathanielwheeler/goalg/algorithms"
  
  "github.com/thatisuday/commando"
)

func main() {
  cfg := LoadConfig()
  command := strings.ToLower(os.Args[1])

  // Initialize commando
  commando.
    SetExecutableName("goalg").
    SetVersion("v1.0.0").
    SetDescription("This CLI tool calls an algorithm and inputs values into it.")

  // Configure sub-commands by looping through a yaml file
  for k, v := range cfg.Algs {
    fmt.Println(k, v)
  }
  
  // Args need to be of type []interface{} instead of []string
  strArgs := os.Args[2:]
  args := make([]interface{}, len(strArgs))
  for i, v := range strArgs {
    args[i] = v
  }

  algorithms := map[string]interface{} {
    "fizzbuzz": algorithms.Fizzbuzz,
    "numinlist": algorithms.NumInList,
    "spongetext": algorithms.Spongetext,
  }

  Call(algorithms, command, args...)
}

// Call will take in funcmap, funcname, and arguments, and reflect them to the appropriate algorithms
func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
  f := reflect.ValueOf(m[name])
  if len(params) != f.Type().NumIn() {
      panic("The number of params is not adapted.")
  }
  in := make([]reflect.Value, len(params))
  for k, param := range params {
      in[k] = reflect.ValueOf(param)
  }
  result = f.Call(in)

  fmt.Printf("%v\n", result)
  return
}