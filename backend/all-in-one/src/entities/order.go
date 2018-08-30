package entities

type OrderJson struct {
    //Id string   `json:"-"  db:"id"  ui:"" `   // json:"-" 不显示
    SeatNo string   `json:"seat_no"  db:"-"  ui:"-" `
	Dishes   []OrderDishSetJson   `json:"dishes"  db:"-" ` 
    Sets     []OrderDishSetJson   `json:"sets"  db:"-" ` 
    CreatedOn string `json:"created_on"  db:"-" `
}

type OrderDishSetJson struct {
    //Id string   `json:"-"  db:"id"  ui:"" `   // json:"-" 不显示
    DishSetId string   `json:"id"  db:"-"  ui:"-" `
    Count   string   `json:"count"  db:"-"`  
    JpName   string   `json:"jp_name"  db:"-"` 
}

type OrderDB struct {
    SeatNo string   `json:"-"  db:"seat_no"  ui:"-" `
    DishSetId string   `json:"-"  db:"dish_set_id"  ui:"-" `
    Count string   `json:"-"  db:"count"  ui:"-" `
    CreatedOn string   `json:"-"  db:"created_on"  ui:"-" `
    Type string   `json:"-"  db:"type"  ui:"-" `
}


type OrderConfirmed struct {
    SeatNo string   `json:"seat_no"  db:"seat_no"  ui:"-" `
    DishSets []OrderDishSetJson   `json:"ordered"  db:"dish_set_id"  ui:"-" `
}
