package services


import (
    "daos"
	"entities"
	"fmt"
	"utils"
)


func GetAllOrders() []entities.OrderJson {
	queryStr := fmt.Sprintf(
		"SELECT * FROM orders WHERE created_on IN (SELECT max(created_on) FROM orders where shop_id=1 group by seat_no)");
    arr,err := daos.FetchRows(queryStr);
    if err != nil {
        panic("[services.GetAllOrders] query error:" + queryStr)
    }

	order := entities.OrderDB{}
    data := util.GetObjs(&order, *arr) 	
	fmt.Println(data)
	r := make([]entities.OrderJson, len(data))
	for i, item := range data {
		p, _ := item.(*entities.OrderDB)
		
		oj := entities.OrderJson{}
		oj.CreatedOn = p.CreatedOn
		oj.SeatNo = p.SeatNo
		r[i] = oj
	}

	return r
}


func GetOrderBySeatNo(seatNo string) entities.OrderConfirmed {
	
	queryStr := fmt.Sprintf("select * from orders where shop_id=1 and seat_no=?");
    arr,err := daos.FetchRows(queryStr, seatNo);
    if err != nil {
        panic("[services.GetOrderBySeatNo] query error:" + queryStr)
    }

	data := make([]interface{}, len(*arr))
	for i, item := range *arr {
		set := entities.OrderDB{}
		util.GetObjSingle(&set, item)
		data[i] = set
	}

	orderConfirm := entities.OrderConfirmed{}
	// 设置返回值的桌号
	orderConfirm.SeatNo = seatNo 
	orderConfirm.DishSets = make([]entities.OrderDishSetJson, len(data))

	//遍历当前桌订单中的项目(单品或套餐)
	for i, item := range data {
		p, _ := item.(entities.OrderDB)
		// 初始化一道菜的项目
		orderDishSet := entities.OrderDishSetJson{}
		// 设定一道菜的项目，所点的数量
		orderDishSet.Count = p.Count;
		// 获取单品的英文名
		if(p.Type == "dish") {
			dish := GetDishById(p.DishSetId)
			orderDishSet.JpName = dish.JpName
		} else {
			set := GetSetById(p.DishSetId)
			orderDishSet.JpName = set.JpName
		}
		// 将当前的菜添加到集合中
		orderConfirm.DishSets[i] = orderDishSet
	}
	return orderConfirm
}



func CreateOrder(order entities.OrderJson) {
	for _, dish := range order.Dishes {
		insertStr := "insert into orders(shop_id, seat_no, dish_set_id, type,count) values(?,?,?,?,?)" ;
		daos.InsertUpdate(insertStr, 1, order.SeatNo, dish.DishSetId, "dish", dish.Count);
	}

	for _, set := range order.Sets {
		insertStr := "insert into orders(shop_id, seat_no, dish_set_id, type,count) values(?,?,?,?,?)" ;
		daos.InsertUpdate(insertStr, 1, order.SeatNo, set.DishSetId, "set", set.Count);
	}
}

