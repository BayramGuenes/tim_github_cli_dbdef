package readConfig

import (
	"encoding/json"
	"log"
	"os"
)

type DB_databases struct {
	Logger   string
	Receiver string
	Repo     string
}
type ConfStruct struct {
	DB_databases               DB_databases
	StartOffsetTimEntityChange int64
	In_PRODUCTIONSYSTEM        bool
}
type SettingStruct struct {
	DB_username string
	DB_userpw   string
	DB_adress   string
	DB_port     string
}

var Conf ConfStruct
var Setting SettingStruct

func GetConfig(iConfLocation string) ConfStruct {
	lLocation := "./config/configdborga.json"
	if len(iConfLocation) > 0 {
		lLocation = iConfLocation + "/configdborga.json"
	}
	println("ConfLocation:" + lLocation)
	file, _ := os.Open(lLocation)
	decoder := json.NewDecoder(file)
	Conf = ConfStruct{}
	err := decoder.Decode(&Conf)
	if err != nil {
		log.Fatal("JSON.Conf:" + err.Error())
	}
	//println(Conf.DB_port)
	return Conf
}
