/**------------------------------------------------------------**
 * @filename query/query.go
 * @author   jinycoo - admin@jinycoo.com
 * @version  1.0.0
 * @date     2021/8/17 10:58
 * @desc     go-quake - url tag query
 **------------------------------------------------------------**/
package query

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var timeType = reflect.TypeOf(time.Time{})

var encoderType = reflect.TypeOf(new(Encoder)).Elem()

type Encoder interface {
	EncodeValues(key string, v *url.Values) error
}

func Values(v interface{}) (url.Values, error) {
	values := make(url.Values)
	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return values, nil
		}
		val = val.Elem()
	}

	if v == nil {
		return values, nil
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("query: Values() expects struct input. Got %v", val.Kind())
	}

	err := reflectValue(values, val, "")
	return values, err
}

func reflectValue(values url.Values, val reflect.Value, scope string) error {
	var embedded []reflect.Value

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous { // unexported
			continue
		}

		sv := val.Field(i)
		tag := sf.Tag.Get("url")
		if tag == "-" {
			continue
		}
		name, opts := parseTag(tag)

		if name == "" {
			if sf.Anonymous {
				v := reflect.Indirect(sv)
				if v.IsValid() && v.Kind() == reflect.Struct {
					embedded = append(embedded, v)
					continue
				}
			}

			name = sf.Name
		}

		if scope != "" {
			name = scope + "[" + name + "]"
		}

		if opts.Contains("omitempty") && isEmptyValue(sv) {
			continue
		}

		if sv.Type().Implements(encoderType) {
			if !reflect.Indirect(sv).IsValid() && sv.Type().Elem().Implements(encoderType) {
				sv = reflect.New(sv.Type().Elem())
			}

			m := sv.Interface().(Encoder)
			if err := m.EncodeValues(name, &values); err != nil {
				return err
			}
			continue
		}

		for sv.Kind() == reflect.Ptr {
			if sv.IsNil() {
				break
			}
			sv = sv.Elem()
		}

		if sv.Kind() == reflect.Slice || sv.Kind() == reflect.Array {
			var del string
			if opts.Contains("comma") {
				del = ","
			} else if opts.Contains("space") {
				del = " "
			} else if opts.Contains("semicolon") {
				del = ";"
			} else if opts.Contains("brackets") {
				name = name + "[]"
			} else {
				del = sf.Tag.Get("del")
			}

			if del != "" {
				s := new(bytes.Buffer)
				first := true
				for i := 0; i < sv.Len(); i++ {
					if first {
						first = false
					} else {
						s.WriteString(del)
					}
					s.WriteString(valueString(sv.Index(i), opts, sf))
				}
				values.Add(name, s.String())
			} else {
				for i := 0; i < sv.Len(); i++ {
					k := name
					if opts.Contains("numbered") {
						k = fmt.Sprintf("%s%d", name, i)
					}
					values.Add(k, valueString(sv.Index(i), opts, sf))
				}
			}
			continue
		}

		if sv.Type() == timeType {
			values.Add(name, valueString(sv, opts, sf))
			continue
		}

		if sv.Kind() == reflect.Struct {
			if err := reflectValue(values, sv, name); err != nil {
				return err
			}
			continue
		}

		values.Add(name, valueString(sv, opts, sf))
	}

	for _, f := range embedded {
		if err := reflectValue(values, f, scope); err != nil {
			return err
		}
	}

	return nil
}

func valueString(v reflect.Value, opts tagOptions, sf reflect.StructField) string {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}

	if v.Kind() == reflect.Bool && opts.Contains("int") {
		if v.Bool() {
			return "1"
		}
		return "0"
	}

	if v.Type() == timeType {
		t := v.Interface().(time.Time)
		if opts.Contains("unix") {
			return strconv.FormatInt(t.Unix(), 10)
		}
		if opts.Contains("unixmilli") {
			return strconv.FormatInt((t.UnixNano() / 1e6), 10)
		}
		if opts.Contains("unixnano") {
			return strconv.FormatInt(t.UnixNano(), 10)
		}
		if layout := sf.Tag.Get("layout"); layout != "" {
			return t.Format(layout)
		}
		return t.Format(time.RFC3339)
	}

	return fmt.Sprint(v.Interface())
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	type zeroable interface {
		IsZero() bool
	}

	if z, ok := v.Interface().(zeroable); ok {
		return z.IsZero()
	}

	return false
}

type tagOptions []string

func parseTag(tag string) (string, tagOptions) {
	s := strings.Split(tag, ",")
	return s[0], s[1:]
}

func (o tagOptions) Contains(option string) bool {
	for _, s := range o {
		if s == option {
			return true
		}
	}
	return false
}
