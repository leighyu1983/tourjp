package services


import (
	"fmt"
    "daos"
    "entities"
	"utils"
	"strings"
	"mime/multipart"
	"github.com/satori/go.uuid"
)


func GetAllDishes() []entities.DishJson {
	queryStr := fmt.Sprintf("select * from dishes where shop_id=1");
    arr,err := daos.FetchRows(queryStr);
    if err != nil {
        panic("[services.GetAllDishes] query error:" + queryStr)
    }

	data := make([]interface{}, len(*arr))
	for i, item := range *arr {
		set := entities.DishDB{}
		util.GetObjSingle(&set, item)
		data[i] = set
	}


	r := make([]entities.DishJson, len(data))
	for i, item := range data {
		p, _ := item.(entities.DishDB)

		dishJ := entities.DishJson{}
		dishJ.Id = p.Id
		dishJ.Code = p.Code
		dishJ.CnName = p.CnName
		dishJ.JpName = p.JpName
		dishJ.CnDescription = p.CnDescription
		dishJ.JpDescription = p.JpDescription
		dishJ.JpMaterials = strings.Split(p.JpMaterials, ",")
		dishJ.CnMaterials = strings.Split(p.CnMaterials, ",")
		dishJ.ImageUrl = p.ImageUrl

		r[i] = dishJ
	}

	return r
}

func GetDishById(id string) entities.DishJson {
	r := entities.DishJson{}
	for _, dish := range GetAllDishes() {
		if(dish.Id == id) {
			return dish
		}
	}
	return r
}

func CreateDishJP(dishj *entities.DishJson) (string) {
	uid, _ := uuid.NewV4()
	id := fmt.Sprintf("%s", uid)

	dish := entities.DishDB{}
	dish.Id = fmt.Sprintf("%s", uid)
	dish.Code = dishj.Code
	dish.JpName = dishj.JpName
	dish.JpDescription = dishj.JpDescription
	dish.JpMaterials = strings.Join(dishj.JpMaterials, ",")

	insertStr := "insert into dishes(shop_id, id, code, jp_name, jp_description, jp_materials) values(?,?,?,?,?,?)" ;
	daos.InsertUpdate(insertStr, 1, id, dish.Code, dish.JpName, dish.JpDescription, dish.JpMaterials);

	return id
}


func UpdateDishCN(dish *entities.DishJson) {
	fmt.Println("-------------111-----------------")
	fmt.Println(dish)
	fmt.Println("-------------222-----------------")
	insertStr := "update dishes set cn_name=?, cn_description=?, cn_materials=? where id=? and shop_id=1";
    daos.InsertUpdate(insertStr, dish.CnName, dish.CnDescription, strings.Join(dish.CnMaterials, ","), dish.Id);
}


func UploadDishImage(file multipart.File, filename string, id string) error {
	var config entities.ConfigEntity
	var err error
	config, err = util.GetConfig();
	
	if(err != nil) {
		return err
	}

	err = util.SaveFile(file, config.ImageFolder, id + filename)
	if(err != nil) {
		return err
	}

	updateStr := "update dishes set image_url=? where shop_id=1 and id=?";
	daos.InsertUpdate(updateStr, config.ImageUrlShop + id + "_" + filename, id)

	return err
}