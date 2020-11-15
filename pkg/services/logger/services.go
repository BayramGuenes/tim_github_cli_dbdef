package services

import (
	"database/sql"

	"fmt"
	"os"

	//huhu
	_ "github.com/go-sql-driver/mysql"

	"github.com/bayramguenes/tim_github_cli_dbdef.git/internal/data"
	"github.com/bayramguenes/tim_github_cli_dbdef.git/internal/dbinterface"
)

// DropDatabaseTable jjjjj
func DropDatabaseTable(iConnection, iDatabase string) {

	// create database with parametername and delete old database/tables
	dbinterface.CreateDatabaseSchemaIfNotExists(iConnection, iDatabase)

	db, err := sql.Open("mysql", iConnection)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = db.Exec("use " + iDatabase)
	var statement string
	if err != nil {
		fmt.Println("no database")
		statement = "create database " + iDatabase + " character set 'utf8'"
		_, err = db.Exec(statement)
		fmt.Println(statement)
	}
	statement = "drop table timlog"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println("Err:" + err.Error())
	}
	fmt.Println(statement)

	statement = "drop table tim_vt_apptransact"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println("Err:" + err.Error())
	}
	fmt.Println(statement)

	statement = "drop table tim_vt_svnapptransact"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println("Err:" + err.Error())
	}
	fmt.Println(statement)

	db.Close()
}

// CreateDatabaseTables 3. define new database and tables
func CreateDatabaseTables(iConnection, iDatabase string) {
	var lvTable string
	var lvFields string
	//var ltValue []string
	var lvStatement data.Statement
	var ltStatement []data.Statement

	//Database
	println("define_database_table:" + iConnection)
	dbinterface.CreateDatabaseSchemaIfNotExists(iConnection, iDatabase)

	lvTable = "timlog"
	lvFields = "(" +
		"logid bigint not null auto_increment," +
		"transactkey varchar(50), " +
		"concurrkey varchar(50) not null," +
		"apptransact varchar(50), " +
		"appclient  varchar(50), " +
		"applogging  varchar(50), " +
		"svnapptransact varchar(50), " +
		"svnapplogging varchar(50), " +
		"step varchar(50), " +
		"stepresult varchar(10), " +
		"stepdatetime varchar(20), " +
		"transactstarttime varchar(25)," +
		"transactstatus varchar(10) NOT NULL," +
		"uname varchar(50)," +
		"stepcontext varchar(1000), " +
		"primary key (logid))"
	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)

	// CREATE INDEX
	ltStatement = make([]data.Statement, 0)
	lvStatement.Text = `CREATE INDEX datetime ON timlog (transactstarttime ,stepdatetime, apptransact, svnapptransact, uname)`
	ltStatement = append(ltStatement, lvStatement)

	lvStatement.Text = `CREATE INDEX datestatus ON timlog (transactstarttime,stepdatetime,transactstatus)`
	ltStatement = append(ltStatement, lvStatement)

	lvStatement.Text = `CREATE INDEX unamedate ON timlog (uname,transactstarttime,stepdatetime)`
	ltStatement = append(ltStatement, lvStatement)

	//

	lvTable = "tim_vt_apptransact"
	lvFields = "(" +
		"name varchar(50) not null," +
		"text varchar(100), " +
		"primary key (name))"
		//
	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)
	dbinterface.InsertApptransactRow(iConnection, iDatabase, "datareceiver", "tim_ms_datareceiver")
	dbinterface.InsertApptransactRow(iConnection, iDatabase, "datarepo", "tim_ms_repo")

	//dbinterface.InsertApptransactRow(iConnection, iDatabase, "", "")
	//dbinterface.InsertApptransactRow(iConnection, iDatabase, "", "")

	lvTable = "tim_vt_svnapptransact"
	lvFields = "(" +
		"name varchar(50) not null," +
		"text varchar(100), " +
		"primary key (name))"
		//
	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)
	dbinterface.InsertSvnApptransactRow(iConnection, iDatabase, "DataPull", "Data-Pull-Action")
	dbinterface.InsertSvnApptransactRow(iConnection, iDatabase, "DataTransmitToRepo", "Transmit/Push Data To Repo")
	dbinterface.InsertSvnApptransactRow(iConnection, iDatabase, "ProvideEntityChanges", "Get EntityChanges for further processing")
	dbinterface.InsertSvnApptransactRow(iConnection, iDatabase, "SuspendDanglingImports", "Suspend dangling imports")

	dbinterface.ExecuteStatement(iConnection, iDatabase, ltStatement)

}
