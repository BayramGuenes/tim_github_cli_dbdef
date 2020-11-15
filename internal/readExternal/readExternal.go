package readExternal

import (
	"os"
	"strings"
)

type SettingStruct struct {
	DB_username string
	DB_userpw   string
	DB_adress   string
	DB_port     string
}

var Setting SettingStruct

func ReadSettings() (eSetting SettingStruct) {
	eSetting.DB_username = os.Getenv("DB_USER")
	eSetting.DB_userpw = os.Getenv("DB_PASSWORD")
	eSetting.DB_adress = os.Getenv("DB_HOST")
	eSetting.DB_port = os.Getenv("DB_PORT")
	return

}

func LoadFromOSArgs() (eOperation string, eConfLocation string) {

	eOperation = ""
	eConfLocation = ""

	leng := len(os.Args)

	for i := 0; i < leng; i++ {
		//println("os.Args[i]=" + os.Args[i])
		if i > 0 {
			osparam := os.Args[i]
			splittedString := strings.Split(osparam, "=")
			var namevalues []string
			namevalues = append(namevalues, splittedString...)
			leng := len(namevalues)
			if leng > 1 {
				paramname := splittedString[0]
				paramval := splittedString[1]

				if paramname == "operation" || paramname == "Operation" || paramname == "OPERATION" {
					//paramval = "create"
					eOperation = paramval //"create"
				}
				if paramname == "confLocation" || paramname == "ConfLocation" ||
					paramname == "CONFLOCATION" {
					//paramval = "create"
					eConfLocation = paramval //"create"
				}

			}

		}
	}
	println("LoadFromOSArgs():eConfLocation=" + eConfLocation)
	return
}
