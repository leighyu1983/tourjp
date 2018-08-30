package util

import(
	"entities"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"errors"
)

func GetConfig()(entities.ConfigEntity, error) {
	data, err := ioutil.ReadFile("../resources/config.yaml")
	fmt.Println("===configuration data:\n" + string(data))
	
	config := entities.ConfigEntity{}
    yaml.Unmarshal(data, &config)
    if(config.MysqlUrl == ""){
        err = errors.New("read configuration file exception")
	}
	return config, err
}
