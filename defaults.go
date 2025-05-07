package defaults

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

const (
	tagName = "default"
)

type parser struct {
	regexp  *regexp.Regexp
	handler func(string) string
}

type configuration struct {
	Parsers map[string]parser
}

var (
	instance = configuration{
		Parsers: make(map[string]parser),
	}
)

func Configure(parsers map[string]func(string) string) error {
	reg, _ := regexp.Compile("^[A-Za-z0-9_]+$")

	for pattern, target := range parsers {
		if !reg.MatchString(pattern) {
			return fmt.Errorf("invalid pattern: %s", pattern)
		}

		instance.Parsers[pattern] = parser{
			regexp:  regexp.MustCompile(`%` + pattern + `%`),
			handler: target,
		}
	}

	return nil
}

func parse(target string) string {
	for _, pr := range instance.Parsers {
		target = pr.regexp.ReplaceAllStringFunc(target, pr.handler)
	}

	return target
}

func Set(ptr interface{}) error {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return fmt.Errorf("invalid type: %s", reflect.TypeOf(ptr).Kind().String())
	}

	v := reflect.ValueOf(ptr).Elem()
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("invalid type: %s", reflect.TypeOf(ptr).Kind().String())
	}

	for i := 0; i < t.NumField(); i++ {
		if defaultVal := t.Field(i).Tag.Get(tagName); defaultVal != "" || t.Field(i).Anonymous {

			if !v.Field(i).CanSet() {
				return fmt.Errorf("failed to set: %s", t.Field(i).Name)
			}

			if err := set(v.Field(i), parse(defaultVal)); err != nil {
				return err
			}
		}
	}

	return nil
}

func set(field reflect.Value, defaultVal string) error {
	if field.Kind() == reflect.Struct {
		return Set(field.Addr().Interface())
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

	case reflect.Ptr:
		return set(field.Elem(), defaultVal)

	default:
		return fmt.Errorf("unsupported type: %s", field.Kind())
	}

	return nil
}
