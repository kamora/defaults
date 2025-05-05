# Defaults

[![License](https://img.shields.io/github/license/kamora/defaults)](./LICENSE)

The simplest way to initialize struct fields with default values.

# References

Originally forked from: https://github.com/creasty

Simplified. Reorganized. Refactored.

# Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kamora/defaults"
	"math/rand"
)

type Gender string

type Sample struct {
	Name    string `default:"John Smith"`
	Age     int    `default:"27"`
	Gender  Gender `default:"m"`
	Working bool   `default:"true"`

	Struct    OtherStruct  `default:"."`
	StructPtr *OtherStruct `default:"."`

	NoTag OtherStruct
}

type OtherStruct struct {
	Hello  string `default:"world"` // Tags in a nested struct also work
	Foo    int
	Random int
}

// SetDefaults implements defaults.Setter interface
func (s *OtherStruct) SetDefaults() {
	s.Random = rand.Int() // Set a dynamic value
}

func main() {
	obj := &Sample{}
	if err := defaults.Set(obj); err != nil {
		panic(err)
	}

	out, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))

	// Output:
	//{
	//	"Name": "John Smith",
	//	"Age": 27,
	//	"Gender": "m",
	//	"Working": true,
	//	"Struct": {
	//		"Hello": "world",
	//			"Foo": 0,
	//			"Random": 0
	//	},
	//	"StructPtr": {
	//		"Hello": "world",
	//			"Foo": 0,
	//			"Random": 0
	//	},
	//	"NoTag": {
	//		"Hello": "",
	//			"Foo": 0,
	//			"Random": 0
	//	}
	//}
}
```
