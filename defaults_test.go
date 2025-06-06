package defaults

import (
	"math/rand/v2"
	"strconv"
	"testing"
)

type (
	MyInt     int
	MyInt8    int8
	MyInt16   int16
	MyInt32   int32
	MyInt64   int64
	MyUint    uint
	MyUint8   uint8
	MyUint16  uint16
	MyUint32  uint32
	MyUint64  uint64
	MyFloat32 float32
	MyFloat64 float64
	MyBool    bool
	MyString  string
)

type Sample struct {
	Int       int     `default:"1"`
	Int8      int8    `default:"8"`
	Int16     int16   `default:"16"`
	Int32     int32   `default:"32"`
	Int64     int64   `default:"64"`
	Uint      uint    `default:"1"`
	Uint8     uint8   `default:"8"`
	Uint16    uint16  `default:"16"`
	Uint32    uint32  `default:"32"`
	Uint64    uint64  `default:"64"`
	Float32   float32 `default:"1.32"`
	Float64   float64 `default:"1.64"`
	BoolTrue  bool    `default:"true"`
	BoolFalse bool    `default:"false"`
	String    string  `default:"hello"`

	IntOct    int    `default:"0o1"`
	Int8Oct   int8   `default:"0o10"`
	Int16Oct  int16  `default:"0o20"`
	Int32Oct  int32  `default:"0o40"`
	Int64Oct  int64  `default:"0o100"`
	UintOct   uint   `default:"0o1"`
	Uint8Oct  uint8  `default:"0o10"`
	Uint16Oct uint16 `default:"0o20"`
	Uint32Oct uint32 `default:"0o40"`
	Uint64Oct uint64 `default:"0o100"`

	IntHex    int    `default:"0x1"`
	Int8Hex   int8   `default:"0x8"`
	Int16Hex  int16  `default:"0x10"`
	Int32Hex  int32  `default:"0x20"`
	Int64Hex  int64  `default:"0x40"`
	UintHex   uint   `default:"0x1"`
	Uint8Hex  uint8  `default:"0x8"`
	Uint16Hex uint16 `default:"0x10"`
	Uint32Hex uint32 `default:"0x20"`
	Uint64Hex uint64 `default:"0x40"`

	IntBin    int    `default:"0b1"`
	Int8Bin   int8   `default:"0b1000"`
	Int16Bin  int16  `default:"0b10000"`
	Int32Bin  int32  `default:"0b100000"`
	Int64Bin  int64  `default:"0b1000000"`
	UintBin   uint   `default:"0b1"`
	Uint8Bin  uint8  `default:"0b1000"`
	Uint16Bin uint16 `default:"0b10000"`
	Uint32Bin uint32 `default:"0b100000"`
	Uint64Bin uint64 `default:"0b1000000"`

	IntPtr     *int     `default:"1"`
	Float32Ptr *float32 `default:"1"`
	BoolPtr    *bool    `default:"true"`
	StringPtr  *string  `default:"hello"`

	StringUID    string  `default:"%test10%"`
	StringUIDPtr *string `default:"%test20%"`

	RandomInt    int  `default:"%rand%"`
	RandomIntPtr *int `default:"%rand%"`

	EmbeddedKey

	MyInt       MyInt     `default:"1"`
	MyInt8      MyInt8    `default:"8"`
	MyInt16     MyInt16   `default:"16"`
	MyInt32     MyInt32   `default:"32"`
	MyInt64     MyInt64   `default:"64"`
	MyUint      MyUint    `default:"1"`
	MyUint8     MyUint8   `default:"8"`
	MyUint16    MyUint16  `default:"16"`
	MyUint32    MyUint32  `default:"32"`
	MyUint64    MyUint64  `default:"64"`
	MyFloat32   MyFloat32 `default:"1.32"`
	MyFloat64   MyFloat64 `default:"1.64"`
	MyBoolTrue  MyBool    `default:"true"`
	MyBoolFalse MyBool    `default:"false"`
	MyString    MyString  `default:"hello"`

	Empty     string `default:""`
	NoDefault *string

	StructPtrWithNoTag *Struct `default:"."`
	StructWithNoTag    Struct

	Test1 string `default:"Hello"`
	Test2 string `default:"Hello"`
}

type EmbeddedKey struct {
	StringUIDCombined    string  `default:"%test10%.%test20%"`
	StringUIDCombinedPtr *string `default:"%test20%.%test20%"`
}

type Struct struct {
	Foo         int
	Bar         int    `default:"456"`
	WithDefault string `default:"foo"`
}

type Child struct {
	Name string `default:"Tom"`
	Age  int    `default:"20"`
}

type Parent struct {
	Child *Child `default:"."`
}

var (
	charset = "abcdefghijklmnopqrstuvwxyz"
)

func Test(t *testing.T) {
	configuration := map[string]func(string) string{
		"test10": func(s string) string {
			r := ""

			for i := 0; i < 10; i++ {
				r += string(charset[rand.N[int](len(charset))])
			}

			return r
		},
		"test20": func(s string) string {
			r := ""

			for i := 0; i < 20; i++ {
				r += string(charset[rand.N[int](len(charset))])
			}

			return r
		},
		"rand": func(s string) string {
			return strconv.Itoa(rand.N[int](99) + 1)
		},
	}

	if err := Configure(configuration); err != nil {
		panic(err)
	}

	sample := &Sample{
		Test1: "Bye",
	}

	if err := Set(sample); err != nil {
		t.Fatalf("it should not return an error: %v", err)
	}

	nonPtrVal := 1

	if err := Set(nonPtrVal); err == nil {
		t.Fatalf("it should return an error when used for a non-pointer type")
	}

	if err := Set(&nonPtrVal); err == nil {
		t.Fatalf("it should return an error when used for a non-pointer type")
	}

	t.Run("primitive types", func(t *testing.T) {
		if sample.Int != 1 {
			t.Errorf("it should initialize int")
		}
		if sample.Int8 != 8 {
			t.Errorf("it should initialize int8")
		}
		if sample.Int16 != 16 {
			t.Errorf("it should initialize int16")
		}
		if sample.Int32 != 32 {
			t.Errorf("it should initialize int32")
		}
		if sample.Int64 != 64 {
			t.Errorf("it should initialize int64")
		}
		if sample.Uint != 1 {
			t.Errorf("it should initialize uint")
		}
		if sample.Uint8 != 8 {
			t.Errorf("it should initialize uint8")
		}
		if sample.Uint16 != 16 {
			t.Errorf("it should initialize uint16")
		}
		if sample.Uint32 != 32 {
			t.Errorf("it should initialize uint32")
		}
		if sample.Uint64 != 64 {
			t.Errorf("it should initialize uint64")
		}
		if sample.Float32 != 1.32 {
			t.Errorf("it should initialize float32")
		}
		if sample.Float64 != 1.64 {
			t.Errorf("it should initialize float64")
		}
		if sample.BoolTrue != true {
			t.Errorf("it should initialize bool (true)")
		}
		if sample.BoolFalse != false {
			t.Errorf("it should initialize bool (false)")
		}
		if *sample.BoolPtr != true {
			t.Errorf("it should initialize bool (true)")
		}
		if sample.String != "hello" {
			t.Errorf("it should initialize string")
		}

		if sample.IntOct != 0o1 {
			t.Errorf("it should initialize int with octal literal")
		}
		if sample.Int8Oct != 0o10 {
			t.Errorf("it should initialize int8 with octal literal")
		}
		if sample.Int16Oct != 0o20 {
			t.Errorf("it should initialize int16 with octal literal")
		}
		if sample.Int32Oct != 0o40 {
			t.Errorf("it should initialize int32 with octal literal")
		}
		if sample.Int64Oct != 0o100 {
			t.Errorf("it should initialize int64 with octal literal")
		}
		if sample.UintOct != 0o1 {
			t.Errorf("it should initialize uint with octal literal")
		}
		if sample.Uint8Oct != 0o10 {
			t.Errorf("it should initialize uint8 with octal literal")
		}
		if sample.Uint16Oct != 0o20 {
			t.Errorf("it should initialize uint16 with octal literal")
		}
		if sample.Uint32Oct != 0o40 {
			t.Errorf("it should initialize uint32 with octal literal")
		}
		if sample.Uint64Oct != 0o100 {
			t.Errorf("it should initialize uint64 with octal literal")
		}

		if sample.IntHex != 0x1 {
			t.Errorf("it should initialize int with hexadecimal literal")
		}
		if sample.Int8Hex != 0x8 {
			t.Errorf("it should initialize int8 with hexadecimal literal")
		}
		if sample.Int16Hex != 0x10 {
			t.Errorf("it should initialize int16 with hexadecimal literal")
		}
		if sample.Int32Hex != 0x20 {
			t.Errorf("it should initialize int32 with hexadecimal literal")
		}
		if sample.Int64Hex != 0x40 {
			t.Errorf("it should initialize int64 with hexadecimal literal")
		}
		if sample.UintHex != 0x1 {
			t.Errorf("it should initialize uint with hexadecimal literal")
		}
		if sample.Uint8Hex != 0x8 {
			t.Errorf("it should initialize uint8 with hexadecimal literal")
		}
		if sample.Uint16Hex != 0x10 {
			t.Errorf("it should initialize uint16 with hexadecimal literal")
		}
		if sample.Uint32Hex != 0x20 {
			t.Errorf("it should initialize uint32 with hexadecimal literal")
		}
		if sample.Uint64Hex != 0x40 {
			t.Errorf("it should initialize uint64 with hexadecimal literal")
		}

		if sample.IntBin != 0b1 {
			t.Errorf("it should initialize int with binary literal")
		}
		if sample.Int8Bin != 0b1000 {
			t.Errorf("it should initialize int8 with binary literal")
		}
		if sample.Int16Bin != 0b10000 {
			t.Errorf("it should initialize int16 with binary literal")
		}
		if sample.Int32Bin != 0b100000 {
			t.Errorf("it should initialize int32 with binary literal")
		}
		if sample.Int64Bin != 0b1000000 {
			t.Errorf("it should initialize int64 with binary literal")
		}
		if sample.UintBin != 0b1 {
			t.Errorf("it should initialize uint with binary literal")
		}
		if sample.Uint8Bin != 0b1000 {
			t.Errorf("it should initialize uint8 with binary literal")
		}
		if sample.Uint16Bin != 0b10000 {
			t.Errorf("it should initialize uint16 with binary literal")
		}
		if sample.Uint32Bin != 0b100000 {
			t.Errorf("it should initialize uint32 with binary literal")
		}
		if sample.Uint64Bin != 0b1000000 {
			t.Errorf("it should initialize uint64 with binary literal")
		}
	})

	t.Run("parsers", func(t *testing.T) {
		if len(sample.StringUID) != 10 {
			t.Errorf("it should be initialize with generators")
		}
		if len(*sample.StringUIDPtr) != 20 {
			t.Errorf("it should be initialize with generators")
		}
		if len(sample.StringUIDCombined) != 31 {
			t.Errorf("it should be initialize with generators")
		}
		if len(*sample.StringUIDCombinedPtr) != 41 {
			t.Errorf("it should be initialize with generators")
		}
	})

	t.Run("pointer types", func(t *testing.T) {
		if sample.IntPtr == nil || *sample.IntPtr != 1 {
			t.Errorf("it should initialize int pointer")
		}
		if sample.Float32Ptr == nil || *sample.Float32Ptr != 1 {
			t.Errorf("it should initialize float32 pointer")
		}
		if sample.BoolPtr == nil || *sample.BoolPtr != true {
			t.Errorf("it should initialize bool pointer")
		}
		if sample.StringPtr == nil || *sample.StringPtr != "hello" {
			t.Errorf("it should initialize string pointer")
		}
	})

	t.Run("aliased types", func(t *testing.T) {
		if sample.MyInt != 1 {
			t.Errorf("it should initialize int")
		}
		if sample.MyInt8 != 8 {
			t.Errorf("it should initialize int8")
		}
		if sample.MyInt16 != 16 {
			t.Errorf("it should initialize int16")
		}
		if sample.MyInt32 != 32 {
			t.Errorf("it should initialize int32")
		}
		if sample.MyInt64 != 64 {
			t.Errorf("it should initialize int64")
		}
		if sample.MyUint != 1 {
			t.Errorf("it should initialize uint")
		}
		if sample.MyUint8 != 8 {
			t.Errorf("it should initialize uint8")
		}
		if sample.MyUint16 != 16 {
			t.Errorf("it should initialize uint16")
		}
		if sample.MyUint32 != 32 {
			t.Errorf("it should initialize uint32")
		}
		if sample.MyUint64 != 64 {
			t.Errorf("it should initialize uint64")
		}
		if sample.MyFloat64 != 1.64 {
			t.Errorf("it should initialize float64")
		}
		if sample.MyBoolTrue != true {
			t.Errorf("it should initialize bool (true)")
		}
		if sample.MyBoolFalse != false {
			t.Errorf("it should initialize bool (false)")
		}
		if sample.MyString != "hello" {
			t.Errorf("it should initialize string")
		}
	})

	t.Run("no tag", func(t *testing.T) {
		if sample.StructPtrWithNoTag == nil {
			t.Errorf("it recurse into a struct with a tag")
		}
		if sample.StructPtrWithNoTag.WithDefault != "foo" {
			t.Errorf("it recurse into a struct with a tag")
		}
		if sample.StructWithNoTag.WithDefault == "foo" {
			t.Errorf("it should not recurse into a struct without a tag")
		}
	})

	t.Run("not original", func(t *testing.T) {
		if sample.StructPtrWithNoTag == nil {
			t.Errorf("it recurse into a struct with a tag")
		}
		if sample.StructPtrWithNoTag.WithDefault != "foo" {
			t.Errorf("it recurse into a struct with a tag")
		}
		if sample.StructWithNoTag.WithDefault == "foo" {
			t.Errorf("it should not recurse into a struct without a tag")
		}
	})

	t.Run("opt-out", func(t *testing.T) {
		if sample.Test1 != "Bye" {
			t.Errorf("it should be set")
		}
		if sample.Test2 != "Hello" {
			t.Errorf("it should not be set")
		}
	})

	t.Run("pointer", func(t *testing.T) {
		m := Parent{Child: &Child{Name: "Jim"}}
		_ = Set(&m)

		if m.Child.Age != 20 {
			t.Errorf("20 is expected")
		}
	})
}
