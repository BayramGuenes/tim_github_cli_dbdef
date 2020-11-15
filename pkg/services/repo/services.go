package services

import (
	"database/sql"
	"fmt"
	"os"

	//"github.com/bayramguenes/tim_utils_numrange.git"
	//was weiss ich
	"github.com/BayramGuenes/tim_utils_numrange"
	_ "github.com/go-sql-driver/mysql"

	conf "github.com/bayramguenes/tim_github_cli_dbdef.git/internal/readConfig"

	"github.com/bayramguenes/tim_github_cli_dbdef.git/internal/data"
	"github.com/bayramguenes/tim_github_cli_dbdef.git/internal/dbinterface"
	"github.com/bayramguenes/tim_github_cli_dbdef.git/internal/readExternal"
)

// DropDatabaseTable ....
func DropDatabaseTable(iConnection, iDatabase string) {

	// create database with parametername and delete old database/tables
	dbinterface.CreateDatabaseSchemaIfNotExists(iConnection, iDatabase)

	db, err := sql.Open("mysql", iConnection)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	println("iDatabase:" + iDatabase)
	_, err = db.Exec("use " + iDatabase)
	var statement string
	if err != nil {
		fmt.Println("no database")
		statement = "create database " + iDatabase + " character set 'utf8'"
		_, err = db.Exec(statement)
		fmt.Println(statement)
	}

	statement = "drop table tim_vt_entitytype"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_entitychange"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + "Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_entitystate"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}
	statement = "drop table tim_csmsgcontent"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_entitychange__NMRANGEOFFSID"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}
	statement = "drop table tim_entitychange__NMRANGESTRTID"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	db.Close()
}

// CreateDatabaseTables 3. define new database and tables
func CreateDatabaseTables(iConnection, iDatabase string) {
	var lvTable string
	var lvFields string
	//var ltValue []string
	var lvIdxStatement data.Statement
	var ltIdxStatement []data.Statement
	// CREATE INDEX
	ltIdxStatement = make([]data.Statement, 0)

	/* ==================================================*/
	//Database
	println("define_database_table:" + iConnection)
	println("database:" + iDatabase)
	dbinterface.CreateDatabaseSchemaIfNotExists(iConnection, iDatabase)

	/* ==================================================*/
	lvTable = "tim_vt_entitytype"
	lvFields = "(" +
		"entitytype char(10) not null," +
		"text varchar(80), " +
		"primary key (entitytype))"

	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)
	dbinterface.InsertEntityTypeRow(iConnection, iDatabase, "article", "entity type article")
	/* ==================================================*/
	lvTable = "tim_entitychange"
	lvFields = "(" +
		"chgid bigint not null," +
		"refchgid bigint not null," +
		"timechange varchar(15), " +
		"primary key (chgid))"

	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)

	/* ===================================================*/
	lvTable = "tim_entitystate"
	lvFields = "(" +
		"entityid bigint, " +
		"entitytype varchar(10)," +
		"isactive tinyint," +
		"refchgid bigint," +
		"timeentitycreate varchar(15)," +
		"timeisactivestate varchar(15)," +
		"primary key (entityid,entitytype))"

	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)
	lvIdxStatement.Text = `CREATE INDEX refchgid ON tim_entitystate (refchgid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)

	_, lConfLocation := readExternal.LoadFromOSArgs()

	lvconf := conf.GetConfig(lConfLocation)
	lvsetting := readExternal.ReadSettings()
	/*println("Nr========"+lvconf.DB_adress, lvconf.DB_port,
	lvconf.DB_userpw, lvconf.DB_database)*/
	//dbadr:=os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/"
	lOnr := tim_utils_numrange.NewManager(lvsetting.DB_adress, lvsetting.DB_port,
		lvsetting.DB_username+":"+lvsetting.DB_userpw, lvconf.DB_databases.Repo)
	lOutput := lOnr.CreateNumRange("tim_entitychange", lvconf.StartOffsetTimEntityChange)
	println("lOutput:" + lOutput.Exception.ErrTxt)
	/* ==================================================*/
	lvTable = "tim_csmsgcontent"
	lvFields = "(" +
		"refchgid bigint not null," +
		"entityid  bigint not null," +
		"entitytype char(10) not null," +
		"content mediumblob, " +
		"timecreatemsg varchar(15)," +
		"primary key (refchgid))"

	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)
	lvIdxStatement.Text = `CREATE INDEX entityid ON tim_csmsgcontent (entityid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)

	/* ==================================================*/

	dbinterface.ExecuteStatement(iConnection, iDatabase, ltIdxStatement)

}
