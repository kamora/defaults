package defaults

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

const (
	tag = "default"
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
		if defaultVal := t.Field(i).Tag.Get(tag); defaultVal != "" || t.Field(i).Anonymous {

			if !v.Field(i).CanSet() {
				return fmt.Errorf("failed to set: %s", t.Field(i).Name)
			}

			if v.Field(i).Kind() == reflect.Ptr {
				v.Field(i).Set(reflect.New(v.Field(i).Type().Elem()))
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
		return set(field.Elem(), defaultVal)
	}

	if !field.Equal(reflect.Zero(field.Type())) {
		return nil
	}

	value, err := convert(field.Type(), defaultVal)
	if err != nil {
		return err
	}

	field.Set(value)
	return nil
}

func convert(target reflect.Type, defaultVal string) (reflect.Value, error) {
	switch target.Kind() {

	case reflect.Bool:
		if val, err := strconv.ParseBool(defaultVal); err == nil {
			return reflect.ValueOf(val).Convert(target), nil
		}

	case reflect.Int:
		if val, err := strconv.ParseInt(defaultVal, 0, strconv.IntSize); err == nil {
			return reflect.ValueOf(int(val)).Convert(target), nil
		}

	case reflect.Int8:
		if val, err := strconv.ParseInt(defaultVal, 0, 8); err == nil {
			return reflect.ValueOf(int8(val)).Convert(target), nil
		}

	case reflect.Int16:
		if val, err := strconv.ParseInt(defaultVal, 0, 16); err == nil {
			return reflect.ValueOf(int16(val)).Convert(target), nil
		}

	case reflect.Int32:
		if val, err := strconv.ParseInt(defaultVal, 0, 32); err == nil {
			return reflect.ValueOf(int32(val)).Convert(target), nil
		}

	case reflect.Int64:
		if val, err := strconv.ParseInt(defaultVal, 0, 64); err == nil {
			return reflect.ValueOf(val).Convert(target), nil
		}

	case reflect.Uint:
		if val, err := strconv.ParseUint(defaultVal, 0, strconv.IntSize); err == nil {
			return reflect.ValueOf(uint(val)).Convert(target), nil
		}

	case reflect.Uint8:
		if val, err := strconv.ParseUint(defaultVal, 0, 8); err == nil {
			return reflect.ValueOf(uint8(val)).Convert(target), nil
		}

	case reflect.Uint16:
		if val, err := strconv.ParseUint(defaultVal, 0, 16); err == nil {
			return reflect.ValueOf(uint16(val)).Convert(target), nil
		}

	case reflect.Uint32:
		if val, err := strconv.ParseUint(defaultVal, 0, 32); err == nil {
			return reflect.ValueOf(uint32(val)).Convert(target), nil
		}

	case reflect.Uint64:
		if val, err := strconv.ParseUint(defaultVal, 0, 64); err == nil {
			return reflect.ValueOf(val).Convert(target), nil
		}

	case reflect.Float32:
		if val, err := strconv.ParseFloat(defaultVal, 32); err == nil {
			return reflect.ValueOf(float32(val)).Convert(target), nil
		}

	case reflect.Float64:
		if val, err := strconv.ParseFloat(defaultVal, 64); err == nil {
			return reflect.ValueOf(val).Convert(target), nil
		}

	case reflect.String:
		return reflect.ValueOf(defaultVal).Convert(target), nil

	default:
		return reflect.Zero(target), fmt.Errorf("unsupported type: %s", target.Kind().String())
	}

	return reflect.Zero(target), fmt.Errorf("invalid type: %s", target.Kind().String())
}
