package util

import(
	"fmt"
    "reflect"
)


/*
 * convert map(string, string) database data to generic object array
*/ 
func GetObjSingle(structObj interface{}, data map[string]string) {
    for dbColumnName, v := range data {
        entityFieldName := getFieldByTag(structObj, dbColumnName, v)
        setField(structObj, entityFieldName, v)
    }
}

/*
 * convert map(string, string) database data to generic object array
*/ 
func GetObjs(structObj interface{}, data []map[string]string) ([]interface{}) {
    r := make([]interface{}, len(data))

	for i, item := range data {
		for dbColumnName, v := range item {
            entityFieldName := getFieldByTag(structObj, dbColumnName, v)
			setField(structObj, entityFieldName, v)
        }
        r[i] = structObj
    }
    return r
}



/*
遍历对象的每一个字段，如果数据表列名字段（tagNameDbColumn）和当前字段的tag (db) 的值一样，返回列名对应的struct字段名
*/
func getFieldByTag(o interface{}, tagNameDbColumn string, value interface{}) (entityFieldName string) {
    for i := 0; i < reflect.TypeOf(o).Elem().NumField(); i++ {
        if reflect.TypeOf(o).Elem().Field(i).Tag.Get("db") == tagNameDbColumn {  
            return reflect.TypeOf(o).Elem().Field(i).Name
        }
    }

    return ""
}

func setField(o interface{}, name string, value interface{}) {
    field := reflect.ValueOf(o).Elem().FieldByName(name)

    if !field.IsValid() {
        //panic(fmt.Sprintf("No such field: %s in obj", name))
        return;
    }
 
    if !field.CanSet() {
        panic(fmt.Sprintf("Cannot set %s field value", name))
    }
 
    fieldType := field.Type()
    val := reflect.ValueOf(value)

    if value == "NULL" {
        return;
    }

    if fieldType != val.Type() {
		panic(fmt.Sprintf("Provided value type %s didn't match obj field type %s", val.Type(), fieldType))
    }
    
    //fmt.Println(name,"===>",value)
    field.Set(val)
}