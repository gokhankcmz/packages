package Object

import (
	"fmt"
	"reflect"
	"strconv"
)

func GetObject(Item interface{}, kvp map[string]string) interface{} {
	v := reflect.Indirect(reflect.ValueOf(Item))
	v.Set(*route(&v, "PackageApi", kvp))
	fmt.Println(fmt.Sprintf(" v: %v", route(&v, "", kvp)))
	return v.Interface()
}
func route(field *reflect.Value, fieldName string, kvp map[string]string) *reflect.Value {
	switch field.Kind() {
	case reflect.String:
		return _string(field, fieldName, kvp)
	case reflect.Int, reflect.Int32, reflect.Int64:
		return _int(field, fieldName, kvp)
	case reflect.Bool:
		return _bool(field, fieldName, kvp)
	case reflect.Struct:
		return _struct(field, "/", fieldName, kvp)
	}
	return field
}
func _string(field *reflect.Value, fieldName string, kvp map[string]string) *reflect.Value {
	field.SetString(kvp[fieldName])
	for k, _ := range kvp {
		fmt.Println(k)
	}
	fmt.Println(fmt.Sprintf("FieldType:%v FieldName:%v Value: %v", field.Type().Name(), fieldName, kvp[fieldName]))
	return field
}

func _int(field *reflect.Value, fieldName string, kvp map[string]string) *reflect.Value {
	v, _ := strconv.ParseInt(kvp[fieldName], 10, 64)
	field.SetInt(v)
	fmt.Println(fmt.Sprintf("FieldType:%v FieldName:%v Value: %v", field.Type().Name(), fieldName, kvp[fieldName]))
	return field
}

func _bool(field *reflect.Value, fieldName string, kvp map[string]string) *reflect.Value {
	v, _ := strconv.ParseBool(kvp[fieldName])
	field.SetBool(v)
	fmt.Println(fmt.Sprintf("FieldType:%v FieldName:%v Value: %v", field.Type().Name(), fieldName, kvp[fieldName]))
	return field
}

func _struct(structV *reflect.Value, HierarchySeparator, prev string, kvp map[string]string) *reflect.Value {
	for i := 0; i < structV.NumField(); i++ {
		varName := structV.Type().Field(i).Name
		field := structV.Field(i)
		value := route(&field, prev+HierarchySeparator+varName, kvp)
		fmt.Println("struct")
		fmt.Println(fmt.Sprintf("FieldType:%v FieldName:%v Value: %v", field.Type().Name(), varName, value.Interface()))
		field.Set(*value)
	}
	return structV
}
