package services


import (
	"fmt"
    "daos"
    "entities"
	"utils"
	"mime/multipart"
	//"github.com/satori/go.uuid"
)


func GetShop(id string) entities.Shop {
	queryStr := fmt.Sprintf("select * from shops where id='%s'", id);
    arr,err := daos.FetchRows(queryStr);
    if err != nil {
        panic("[services.GetShop] query error:" + queryStr)
    }

    shop := entities.Shop{}
    data := util.GetObjs(&shop, *arr) 	
	
	r := make([]*entities.Shop, len(data))
	for i, item := range data {
		p, _ := item.(*entities.Shop)
		r[i] = p
	}

	return *r[0]
}


func CreateShop(shop *entities.Shop) {
	//uid, _ := uuid.NewV4()
	//shop.Id = fmt.Sprintf("%s", uid)

	//insertStr := "insert into shops(id,cn_name,cn_description,jp_name,jp_description) values(?,?,?,?,?)";
	insertStr := "update shops set cn_name=?,cn_description=?,jp_name=?,jp_description =? where id=?";
    daos.InsertUpdate(insertStr, shop.CnName, shop.CnDescription,shop.JpName, shop.JpDescription, shop.Id);
}


func UploadShopImage(file multipart.File, id string) error {
	var config entities.ConfigEntity
	var err error
	config, err = util.GetConfig();
	
	if(err != nil) {
		panic(err)
	}

	err = util.SaveFile(file, config.ImageFolder, "1.jpg")
	if(err != nil) {
		panic(err)
	}

	updateStr := "update shops set image_url=? where id=?";
	fmt.Printf("[UploadShopImage] --updateStr-->'%s' --image-->'%s'", updateStr, config.ImageUrlShop + id + ".jpg")
	daos.InsertUpdate(updateStr, config.ImageUrlShop + id + ".jpg", id)

	return err
}