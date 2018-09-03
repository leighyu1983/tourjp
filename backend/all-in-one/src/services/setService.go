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


func GetAllSets() []entities.SetJsonOut {
	queryStr := fmt.Sprintf("select * from sets where shop_id=1");
    arr,err := daos.FetchRows(queryStr);
    if err != nil {
        panic("[services.GetAllSets] query error:" + queryStr)
    }

	data := make([]interface{}, len(*arr))
	for i, item := range *arr {
		set := entities.SetDB{}
		util.GetObjSingle(&set, item)
		data[i] = set
	}


	r := make([]entities.SetJsonOut, len(data))
	for i, item := range data {
		p, _ := item.(entities.SetDB)

		setJ := entities.SetJsonOut{}
		setJ.Id = p.Id
		setJ.Code = p.Code
		setJ.CnName = p.CnName
		setJ.JpName = p.JpName
		setJ.CnDescription = p.CnDescription
		setJ.JpDescription = p.JpDescription
		setJ.ImageUrl = p.ImageUrl

		dishIdArray := strings.Split(p.Dishes, ",")
		setJ.Dishes = make([]entities.DishJson, len(dishIdArray))
		for j, dishId := range dishIdArray {
			setJ.Dishes[j] = GetDishById(dishId)
		}
		
		r[i] = setJ
	}

	return r
}

func GetSetById(setId string) entities.SetJsonOut {
	for _, set := range GetAllSets() {
		if(set.Id == setId) {
			return set
		}
	}
	return entities.SetJsonOut{}
}

func CreateSetJP(setj *entities.SetJsonIn) (string) {
	uid, _ := uuid.NewV4()
	setj.Id = fmt.Sprintf("%s", uid)

	set := entities.SetDB{}
	set.Id = setj.Id
	set.Code = setj.Code
	set.JpName = setj.JpName
	set.JpDescription = setj.JpDescription
	set.Dishes = strings.Join(setj.Dishes, ",")

	insertStr := "insert into sets(shop_id, id, code, jp_name, jp_description, dishes) values(?,?,?,?,?,?)" ;
	daos.InsertUpdate(insertStr, 1, setj.Id, set.Code, set.JpName, set.JpDescription, set.Dishes);

	return setj.Id
}


func UpdateSetCN(set *entities.SetJsonIn) {
	insertStr := "update sets set cn_name=?, cn_description=? where id=? and shop_id=1";
    daos.InsertUpdate(insertStr, set.CnName, set.CnDescription, set.Id);
}


func UploadSetImage(file multipart.File, filename string, id string) error {
	var config entities.ConfigEntity
	var err error
	config, err = util.GetConfig();
	
	if(err != nil) {
		return err
	}

	err = util.SaveFile(file, config.ImageFolder, id + ".jpg")
	if(err != nil) {
		return err
	}

	updateStr := "update sets set image_url=? where shop_id=1 and id=?";
	daos.InsertUpdate(updateStr, config.ImageUrlSet + id + ".jpg", id)

	return err
}