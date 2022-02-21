package KVP

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

//büyük harfle başlamayan, unexported typeler için boş döner.

func GetKVPs(Item interface{}, HierarchySeparator string, BasePrefix string, kvp map[string]string) map[string]string {
	v := reflect.ValueOf(Item)
	switch v.Kind() {
	case reflect.Ptr:
		_ptr(v, HierarchySeparator, BasePrefix, kvp)
	case reflect.Struct:
		_struct(v, HierarchySeparator, BasePrefix, kvp)
	case reflect.Map:
		_map(v, HierarchySeparator, BasePrefix, kvp)
	case reflect.Array, reflect.Slice:
		_array(v, HierarchySeparator, BasePrefix, kvp)
	case reflect.Func:
		_func(v, BasePrefix, kvp)
	default:
		if BasePrefix == "" {
			BasePrefix = "nil"
		}
		_primitive(Item, BasePrefix, kvp)
	}
	return kvp
}

func _ptr(Item reflect.Value, HierarchySeparator string, prev string, kvp map[string]string) {
	Value := reflect.Indirect(Item)
	if Value.IsValid() && Value.CanInterface() {
		GetKVPs(Value.Interface(), HierarchySeparator, prev, kvp)
	}
}
func _primitive(Item interface{}, prev string, kvp map[string]string) {
	kvp[prev] = fmt.Sprint(Item)
}

func _struct(Item reflect.Value, HierarchySeparator string, prev string, kvp map[string]string) {
	for i := 0; i < Item.NumField(); i++ {
		varName := Item.Type().Field(i).Name
		field := Item.Field(i)
		if field.CanInterface() {
			value := field.Interface()
			Key := getKey(prev, HierarchySeparator, varName)
			GetKVPs(value, HierarchySeparator, Key, kvp)
		}

	}
}

func _map(Input reflect.Value, HierarchySeparator string, prev string, kvp map[string]string) {
	for _, key := range Input.MapKeys() {
		strct := Input.MapIndex(key)
		keyValue := key.Interface()
		Key := getKey(prev, HierarchySeparator, fmt.Sprint(keyValue))
		GetKVPs(strct.Interface(), HierarchySeparator, Key, kvp)
	}
}

func _array(Input reflect.Value, HierarchySeparator string, prev string, kvp map[string]string) {
	for i := 0; i < Input.Len(); i++ {
		GetKVPs(Input.Index(i).Interface(), HierarchySeparator, getKey(prev, HierarchySeparator, fmt.Sprint(i)), kvp)
	}
}

func _func(Input reflect.Value, prev string, kvp map[string]string) {
	fullname := runtime.FuncForPC(Input.Pointer()).Name()
	pathparts := strings.Split(fullname, "/")
	kvp[prev] = fmt.Sprint(pathparts[len(pathparts)-1])
}

func getKey(prev, HierarchySeparator, varName string) string {
	if HierarchySeparator != "" && HierarchySeparator != " " {
		return prev + HierarchySeparator + varName
	}
	return varName
}
