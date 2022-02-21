package Object

import (
	"reflect"
	"strconv"
	"strings"
)

func GetObject(obj interface{}, HierarchySeparator, prev string, kvp map[string]string) {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		return
	}
	internal(reflect.Indirect(v), HierarchySeparator, prev, kvp)

}

func internal(value reflect.Value, HierarchySeparator, prev string, kvp map[string]string) reflect.Value {
	v := kvp[prev]
	switch value.Kind() {
	case reflect.String:
		value.SetString(v)
	case reflect.Bool:
		fv, _ := strconv.ParseBool(v)
		value.SetBool(fv)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fv, _ := strconv.ParseInt(v, 10, 64)
		value.SetInt(fv)
	case reflect.Struct:
		value.Set(_struct(value, HierarchySeparator, prev, kvp))
	case reflect.Map:
		value.Set(_map(value, HierarchySeparator, prev, kvp))
	case reflect.Array, reflect.Slice:
		value.Set(_array(value, HierarchySeparator, prev, kvp))
	}
	return value
}

func _struct(Item reflect.Value, HierarchySeparator, prev string, kvp map[string]string) reflect.Value {
	for i := 0; i < Item.NumField(); i++ {
		varName := Item.Type().Field(i).Name
		field := Item.Field(i)
		if field.CanInterface() {
			Key := getKey(prev, HierarchySeparator, varName)
			internal(field, HierarchySeparator, Key, kvp)
		}
	}
	return Item
}

func _map(Item reflect.Value, HierarchySeparator, prev string, kvp map[string]string) reflect.Value {
	mapType := Item.Type()
	elemType := mapType.Elem()
	mp := reflect.MakeMap(mapType)
	Item.Set(mp)
	for k, _ := range kvp {
		if strings.HasPrefix(k, prev) {
			key := strings.Split(strings.TrimPrefix(k, prev+HierarchySeparator), HierarchySeparator)[0]
			Key := getKey(prev, HierarchySeparator, key)
			newValue := reflect.Indirect(reflect.New(elemType))
			v := internal(newValue, HierarchySeparator, Key, kvp)
			Item.SetMapIndex(reflect.ValueOf(key), v)
		}
	}
	return Item
}

func _array(Item reflect.Value, HierarchySeparator string, prev string, kvp map[string]string) reflect.Value {
	elemType := Item.Type().Elem()
	var duplicateKeys []string
	for k, _ := range kvp {
		if strings.HasPrefix(k, prev) {
			key := strings.Split(strings.TrimPrefix(k, prev+HierarchySeparator), HierarchySeparator)[0]
			Key := getKey(prev, HierarchySeparator, key)
			if !contains(duplicateKeys, Key) {
				duplicateKeys = append(duplicateKeys, Key)
				newValue := reflect.Indirect(reflect.New(elemType))
				v := internal(newValue, HierarchySeparator, Key, kvp)
				Item.Set(reflect.Append(Item, reflect.Indirect(v)))
			}

		}
	}
	return Item
}
func getKey(prev, HierarchySeparator, varName string) string {
	if HierarchySeparator != "" && HierarchySeparator != " " {
		return prev + HierarchySeparator + varName
	}
	return varName
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
