package entities

import(
    "reflect"
    "fmt"
)

type TestObj struct {
    Id string 
    Name   string    
}

func (r *TestObj) GetTestStructural(a string) (string) {
	return "123-" + a + "-" + r.Id
}

func (o *TestObj) TryStruct(){
    name := "Name"
    value := 123

    element := reflect.ValueOf(o).Elem()
    field := element.FieldByName(name)

    if !field.IsValid() {
        fmt.Printf("No such field: %s in obj\n", name)
        return
    }

    if !field.CanSet() {
        fmt.Printf("Cannot set %s field value\n", name)
        return
    }

    fieldType := field.Type()
    val := reflect.ValueOf(value)
    if fieldType != val.Type() {
        fmt.Printf("Provided value type %s didn't match obj field type %s\n", val.Type(), fieldType)
        return
    }

    field.Set(val)
}
