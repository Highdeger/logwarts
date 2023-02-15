package logwarts

import (
	"fmt"
	"reflect"
	"strings"
)

func DetailMap(item map[interface{}]interface{}) string {
	prefix := "  "
	r := fmt.Sprintf("\n%sdetails-map\n", prefix)
	r += stringMap(item, []string{prefix, "  "})
	return r
}

func DetailArray(item []interface{}) string {
	prefix := "  "
	r := fmt.Sprintf("\n%sdetails-array\n", prefix)
	r += stringArray(item, []string{prefix, "  "})
	return r
}

func stringMap(item map[interface{}]interface{}, prefix []string) string {
	r := ""
	i := 0
	for k, v := range item {
		key := fmt.Sprintf("'%+v'(%s)", k, reflect.TypeOf(k))

		r += strings.Join(prefix, "")
		r += stringRow(key, v, prefix, i == len(item)-1)

		i++
	}
	return r
}

func stringArray(item []interface{}, prefix []string) string {
	r := ""
	for i, v := range item {
		key := fmt.Sprintf("[%d]", i)

		r += strings.Join(prefix, "")
		r += stringRow(key, v, prefix, i == len(item)-1)
	}
	return r
}

func stringRow(key string, v interface{}, prefix []string, last bool) string {
	r := ""
	value := fmt.Sprintf("'%+v'(%s)", v, reflect.TypeOf(v))

	newPrefix := append(prefix, "│ ")
	if last {
		newPrefix = append(prefix, "  ")
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Map:
		vMap, ok := v.(map[interface{}]interface{})
		if ok {
			value = fmt.Sprintf("(%s)\n%s", reflect.Map.String(), stringMap(vMap, newPrefix))
		}
	case reflect.Slice | reflect.Array:
		vArray, ok := v.([]interface{})
		if ok {
			value = fmt.Sprintf("(%s)\n%s", reflect.Map.String(), stringArray(vArray, newPrefix))
		}
	}

	if last {
		r += fmt.Sprint("└")
	} else {
		r += fmt.Sprint("├")
	}

	r += fmt.Sprintf("%s => %s", key, value)

	if !last {
		r += fmt.Sprintln()
	}
	return r
}
