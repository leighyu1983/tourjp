package services

import (
    "fmt"
    "daos"
    "entities"
    "utils"
)
 
func Case_1() {
	fmt.Println("===> case_1")
	
    arr,err := daos.FetchRows(
        "select * from companies limit 2");
    if err != nil {
        fmt.Println("query error.....")
    }

    shop := entities.Shop{}
    data := util.GetObjs(&shop, *arr)
    fmt.Println("--------data-------")
    for _, item := range data {
        p, _ := item.(*entities.Shop)
        fmt.Println(*p)
    }
    //fmt.Println(data)
}