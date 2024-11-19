package main

import (
	"fmt"
	"reflect"
)

type ExpectedStruct struct {
	Name string
	Age  int
	A    int
}

type MyStruct struct {
	Name string
	Age  int
	// 假设我们期望的Age字段是int类型，但实际上是int32类型
	Age32 int32
}

func checkStructFields(inputStruct interface{}, expectedStruct interface{}) ([]string, error) {
	inputVal := reflect.ValueOf(inputStruct)
	expectedVal := reflect.ValueOf(expectedStruct)

	if inputVal.Kind() != reflect.Struct || expectedVal.Kind() != reflect.Struct {
		return nil, fmt.Errorf("both input and expected values must be structs")
	}

	var mismatchedFields []string

	for i := 0; i < expectedVal.NumField(); i++ {
		expectedField := expectedVal.Type().Field(i)
		inputField := inputVal.FieldByName(expectedField.Name)

		if !inputField.IsValid() {
			mismatchedFields = append(mismatchedFields, expectedField.Name)
			continue
		}

		if inputField.Type() != expectedField.Type {
			mismatchedFields = append(mismatchedFields, expectedField.Name)
		}
	}

	if len(mismatchedFields) > 0 {
		return mismatchedFields, fmt.Errorf("mismatched fields found: %v", mismatchedFields)
	}

	return nil, nil
}
func findEmptyFields(inputStruct interface{}) ([]string, error) {
	inputVal := reflect.ValueOf(inputStruct)

	if inputVal.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input value must be a struct")
	}

	var emptyFields []string

	for i := 0; i < inputVal.NumField(); i++ {
		field := inputVal.Field(i)
		fieldType := inputVal.Type().Field(i)

		if isEmptyValue(field) {
			emptyFields = append(emptyFields, fieldType.Name)
		}
	}

	return emptyFields, nil
}

// isEmptyValue 判断给定的反射值是否为空值。
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr, reflect.Map, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func main() {
	s := MyStruct{Name: "Alice", Age: 0}

	a, err := findEmptyFields(s)
	fmt.Println(a, err)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All fields match the expected types.")
	}
}
