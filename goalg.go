package main

import (
	"fmt"
	"reflect"

  // This is referenced by the FuncMaps stored in config.yml
	_ "github.com/nathanielwheeler/goalg/algorithms"

	"github.com/thatisuday/commando"
)

func main() {
  cfg := LoadConfig()
  
	// Initialize commando
	commando.
		SetExecutableName("alg").
		SetVersion("v1.0.0").
		SetDescription("This CLI tool calls an algorithm and inputs values into it.")

	// Configure sub-commands by looping through a yaml file
	for k, v := range cfg.Algs {
		c := commando.
			Register(k).
			SetShortDescription(v.ShortDesc).
			SetDescription(v.FullDesc)

		for _, v := range v.Arguments {
      c.AddArgument(v.Name, v.Desc, v.Default)
    }
    
    c.SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
      CallAlgorithm(v.FuncMap, k, args)
    })
  }
  
  commando.Parse(nil)
}

// CallAlgorithm will take in funcmap, funcname, and arguments, and reflect them to the appropriate algorithms
func CallAlgorithm(funcMap map[string]interface{}, name string, params ...interface{}) {
	f := reflect.ValueOf(funcMap[name])
	if len(params) != f.Type().NumIn() {
		panic("The number of params is not adapted.")
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result := f.Call(in)

	fmt.Printf("%v\n", result)
	return
}
