package defaults

import (
	"errors"
	"fmt"
	"github.com/kamora/fluid"
	"math/rand/v2"
	"reflect"
	"strconv"
	"strings"
)

var (
	errInvalidType = errors.New("not a struct pointer")
)

const (
	tagName = "default"
)

// Set initializes members in a struct referenced by a pointer.
func Set(ptr interface{}) error {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return errInvalidType
	}

	v := reflect.ValueOf(ptr).Elem()
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return errInvalidType
	}

	for i := 0; i < t.NumField(); i++ {
		if defaultVal := t.Field(i).Tag.Get(tagName); defaultVal != "" || t.Field(i).Anonymous {
			if err := set(v.Field(i), defaultVal); err != nil {
				return err
			}
		}
	}

	return nil
}

func set(field reflect.Value, defaultVal string) error {
	if !field.CanSet() {
		return fmt.Errorf("failed to set field: %s", tagName)
	}

	if field.Kind() == reflect.Ptr {
		field.Set(reflect.New(field.Type().Elem()))
	}

	switch field.Kind() {
	case reflect.Bool:
		if val, err := strconv.ParseBool(defaultVal); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}

	case reflect.Int:
		if val, err := strconv.ParseInt(defaultVal, 0, strconv.IntSize); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}

	case reflect.Int8:
		if val, err := strconv.ParseInt(defaultVal, 0, 8); err == nil {
			field.Set(reflect.ValueOf(int8(val)).Convert(field.Type()))
		}

	case reflect.Int16:
		if val, err := strconv.ParseInt(defaultVal, 0, 16); err == nil {
			field.Set(reflect.ValueOf(int16(val)).Convert(field.Type()))
		}

	case reflect.Int32:
		if val, err := strconv.ParseInt(defaultVal, 0, 32); err == nil {
			field.Set(reflect.ValueOf(int32(val)).Convert(field.Type()))
		}

	case reflect.Int64:
		if val, err := strconv.ParseInt(defaultVal, 0, 64); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}

	case reflect.Uint:
		if val, err := strconv.ParseUint(defaultVal, 0, strconv.IntSize); err == nil {
			field.Set(reflect.ValueOf(uint(val)).Convert(field.Type()))
		}

	case reflect.Uint8:
		if val, err := strconv.ParseUint(defaultVal, 0, 8); err == nil {
			field.Set(reflect.ValueOf(uint8(val)).Convert(field.Type()))
		}

	case reflect.Uint16:
		if val, err := strconv.ParseUint(defaultVal, 0, 16); err == nil {
			field.Set(reflect.ValueOf(uint16(val)).Convert(field.Type()))
		}

	case reflect.Uint32:
		if val, err := strconv.ParseUint(defaultVal, 0, 32); err == nil {
			field.Set(reflect.ValueOf(uint32(val)).Convert(field.Type()))
		}

	case reflect.Uint64:
		if val, err := strconv.ParseUint(defaultVal, 0, 64); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}

	case reflect.Uintptr:
		if val, err := strconv.ParseUint(defaultVal, 0, strconv.IntSize); err == nil {
			field.Set(reflect.ValueOf(uintptr(val)).Convert(field.Type()))
		}

	case reflect.Float32:
		if val, err := strconv.ParseFloat(defaultVal, 32); err == nil {
			field.Set(reflect.ValueOf(float32(val)).Convert(field.Type()))
		}

	case reflect.Float64:
		if val, err := strconv.ParseFloat(defaultVal, 64); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}

	case reflect.String:
		field.Set(reflect.ValueOf(parse(defaultVal)).Convert(field.Type()))

	case reflect.Struct:
		return Set(field.Addr().Interface())

	case reflect.Ptr:
		return set(field.Elem(), defaultVal)

	default:
		return fmt.Errorf("unsupported type: %s", field.Kind())
	}

	return nil
}

func parse(target string) string {
	target = strings.ReplaceAll(target, "%fluid32%", fluid.Encode(uint32(rand.Int32())))
	target = strings.ReplaceAll(target, "%fluid64%", fluid.Encode(uint64(rand.Int64())))

	return target
}
