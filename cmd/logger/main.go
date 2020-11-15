package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/bayramguenes/tim_github_cli_dbdef.git/internal/data"
	conf "github.com/bayramguenes/tim_github_cli_dbdef.git/internal/readConfig"
	"github.com/bayramguenes/tim_github_cli_dbdef.git/internal/readExternal"
	services "github.com/bayramguenes/tim_github_cli_dbdef.git/pkg/services/logger"
)

//global var
var gvConnection string
var gvDatabase string
var gvDirData string

func main() {
	lOperation, lConfLocation := readExternal.LoadFromOSArgs()
	println("main.lConfLocation:" + lConfLocation)
	gvConf := conf.GetConfig(lConfLocation)
	gvSetting := readExternal.ReadSettings() //gvSetting:=
	gvConnection = gvSetting.DB_username + ":" + gvSetting.DB_userpw + "@tcp(" + gvSetting.DB_adress + ":" + gvSetting.DB_port + ")/"
	gvDatabase = gvConf.DB_databases.Logger
	println("gvDatabase:" + gvDatabase)
	if lOperation == "create" {
		//println("create")
		services.CreateDatabaseTables(gvConnection, gvDatabase)

	} else {
		//println(" lOperation =" + lOperation)
		//return
		readingConsole(gvConf)
	}
}

// 2. user input evaluation
func readingConsole(iConf conf.ConfStruct) {
	reader := bufio.NewReader(os.Stdin)
	writeWelcome(iConf)
	println("DBHOST:" + readExternal.ReadSettings().DB_adress)
	_, lConfLocation := readExternal.LoadFromOSArgs()
	gvConf := conf.GetConfig(lConfLocation)

	println("database:", gvConf.DB_databases.Logger)

	for {
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
		var lineEnding string = data.LineEnding
		if runtime.GOOS == "windows" {
			lineEnding = data.LineEndingWindows
		}
		text = strings.Replace(text, lineEnding, "", -1)

		if text == "create" {
			services.CreateDatabaseTables(gvConnection, gvDatabase)
		} else if text == "drop" && !iConf.In_PRODUCTIONSYSTEM {
			services.DropDatabaseTable(gvConnection, gvDatabase)
		} else if text == "dropcreate" && !iConf.In_PRODUCTIONSYSTEM {
			services.DropDatabaseTable(gvConnection, gvDatabase)
			services.CreateDatabaseTables(gvConnection, gvDatabase)
		} else if text == "exit" {
			break
		} else {
			fmt.Println("keine gültige Anweisung")
		}
	}
}

func writeWelcome(iConf conf.ConfStruct) {
	fmt.Println("---------------------")
	fmt.Println("Tim Logger")
	fmt.Println("---------------------")
	fmt.Println("Eingabe:")
	fmt.Println("'create' zum Anlegen der Datenbank")
	if !iConf.In_PRODUCTIONSYSTEM {
		fmt.Println("'drop' zur Löschung der DB")
		fmt.Println("'dropcreate' zum Anlegen der Datenbank mit Löschung der DB")
	}
	fmt.Println("'exit' zum Beenden der Anwendung")
}
