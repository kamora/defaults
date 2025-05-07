# Defaults

[![License](https://img.shields.io/github/license/kamora/defaults)](./LICENSE)

The simplest way to initialize struct fields with default values.

# References

Originally forked from: https://github.com/creasty

Simplified. Redesigned. Reorganized. Refactored.

# Compatibility
No backward compatibility preserved!

# Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kamora/defaults"
	"github.com/kamora/fluid"
	"math/rand/v2"
	"strconv"
)

type Gender string

type EmbeddedKey struct {
	StringUIDCombined    string  `default:"%fluid32%.%fluid64%"`
	StringUIDCombinedPtr *string `default:"%fluid64%.%fluid64%"`
}

type OtherStruct struct {
	Hello  string `default:"world"` 
	Foo    int    `default:"99"`
	Random int    `default:"%rand%"`
}

type Sample struct {
	Name    string `default:"John Smith"`
	Age     int    `default:"27"`
	Gender  Gender `default:"m"`
	Working bool   `default:"true"`

	EmbeddedKey

	Struct    OtherStruct  `default:"."`
	StructPtr *OtherStruct `default:"."`

	Uid    string `default:"%fluid64%"`
	Random int    `default:"%rand%"`
	NoTag  OtherStruct
}

func main() {
	configuration := map[string]func(string) string{
		"fluid32": func(s string) string {
			return fluid.Encode(uint32(0))
		},
		"fluid64": func(s string) string {
			return fluid.Encode(uint64(0))
		},
		"rand": func(s string) string {
			fmt.Println(strconv.Itoa(rand.N[int](99) + 1))
			return strconv.Itoa(rand.N[int](99) + 1)
		},
	}

	if err := defaults.Configure(configuration); err != nil {
		panic(err)
	}

	obj := &Sample{
		Gender: "f",
		Random: 124,
	}
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
	//	"Gender": "f",
	//	"Working": true,
	//	"StringUIDCombined": "6dmpPQe.11pIrkgbKrQ8v",
	//	"StringUIDCombinedPtr": "11pIrkgbKrQ8v.11pIrkgbKrQ8v",
	//	"Struct": {
	//		"Hello": "world",
	//		"Foo": 99,
	//		"Random": 51
	//	},
	//	"StructPtr": {
	//		"Hello": "world",
	//		"Foo": 99,
	//		"Random": 6
	//	},
	//	"Uid": "11pIrkgbKrQ8v",
	//	"Random": 124,
	//	"NoTag": {
	//		"Hello": "",
	//		"Foo": 0,
	//		"Random": 0
	//	}
	//}
}
```
