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

	statement = "drop table tim_vt_stateproducer"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println("Err:" + err.Error())
	}
	fmt.Println(statement)
	statement = "drop table tim_mock_producerdata"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println("Err:" + err.Error())
	}

	statement = "drop table tim_vt_pullstate"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println("Err:" + err.Error())
	}
	fmt.Println(statement)
	statement = "drop table tim_datapullevent"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println("Err:" + err.Error())
	}
	fmt.Println(statement)

	statement = "drop table tim_vt_orderstate"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println("Err:" + err.Error())
	}
	fmt.Println(statement)
	statement = "drop table tim_imporder"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println("Err:" + err.Error())
	}

	/*statement = "drop table tim_cust_receiver"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println("Err:" + err.Error())
	}
	fmt.Println(statement)
	*/
	db.Close()
}

// 3. define new database and tables
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
	dbinterface.CreateDatabaseSchemaIfNotExists(iConnection, iDatabase)

	/* ==================================================*/
	lvTable = "tim_vt_stateproducer"
	lvFields = "(" +
		"state char(30) not null," +
		"text varchar(80), " +
		"primary key (state))"

	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)
	dbinterface.InsertProducerStateRow(iConnection, iDatabase, "pulled by Receiver", "pulled succesfully")
	dbinterface.InsertProducerStateRow(iConnection, iDatabase, "imported", "ImportProcess has been finished OK")
	dbinterface.InsertProducerStateRow(iConnection, iDatabase, "failed Import", "Processing Import has failed")

	lvTable = "tim_mock_producerdata"
	lvFields = "(" +
		"dataid int not null auto_increment," +
		"loadingfile varchar(80) not null," +
		"datapath varchar(150) not null, " +
		"numxml int, " +
		"timeproduced varchar(15)," +
		"timeprocessed varchar(15)," +
		"producerstate varchar(30)," + //pulledByReceiver, lockedForImport, startedImportProcess, finishedImportOK, finishedImportFailed
		"textprocessed varchar(1000)," +
		"primary key (dataid))"

	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)

	/* ==================================================*/
	lvTable = "tim_vt_pullstate"
	lvFields = "(" +
		"state char(10) not null," +
		"text varchar(80), " +
		"primary key (state))"
	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)
	dbinterface.InsertPullStateRow(iConnection, iDatabase, "PullOK", "pulled succesfully")
	dbinterface.InsertPullStateRow(iConnection, iDatabase, "PullFailed", "Pull failed")

	lvTable = "tim_datapullevent"
	lvFields = "(" +
		"pullid int not null auto_increment," +
		"pullstate varchar(10) not null," +
		"pullstatetxt varchar(100) not null," +
		"timedatafrom varchar(15) not null," +
		"timedatato varchar(15) not null, " +
		"primary key (pullid))"
	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)
	lvIdxStatement.Text = `CREATE INDEX timepull ON tim_datapullevent (timedatato,pullstate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)

	/* ==================================================*/
	lvTable = "tim_vt_orderstate"
	lvFields = "(" +
		"state char(30) not null," +
		"text varchar(80), " +
		"primary key (state))"
	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)

	dbinterface.InsertOrderStateRow(iConnection, iDatabase, "", "ready for processing")
	dbinterface.InsertOrderStateRow(iConnection, iDatabase, "processing", "processing")
	dbinterface.InsertOrderStateRow(iConnection, iDatabase, "imported", "processing finished succesfully")
	dbinterface.InsertOrderStateRow(iConnection, iDatabase, "failed Import", "failed processing")

	lvTable = "tim_imporder"
	lvFields = "(" +
		"orderid int not null auto_increment," +
		"loadingfile varchar(100), " +
		"datapath varchar(250), " +
		"numpanxml int, " +
		"orderstate  varchar(30), " +
		"timepulled varchar(15), " +
		"timeorderstate varchar(15), " +
		"timesendprocresult  varchar(15), " +
		"refpullid   int, " +
		"refproddataid   int, " +
		"timelastxmlproc varchar(15)," +
		"resultdetails varchar(1000)," +
		"primary key (orderid))"
	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)

	lvIdxStatement.Text = `CREATE INDEX timeprov ON tim_imporder (timepulled, orderstate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	lvIdxStatement.Text = `CREATE INDEX refprodid ON tim_imporder (refproddataid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)

	/* ================================================= */
	/*lvTable = "tim_cust_receiver"
	lvFields = "(" +
		"numconcurrentproc int not null," +
		"pullfrequenz int not null," +
		"procfrequenz int not null" +
		")"
	//println("huhu")
	dbinterface.CreateTable(iConnection, iDatabase, lvTable, lvFields)
	dbinterface.InsertNumConcProcRow(iConnection, iDatabase, 3, 30, 120)
	*/
	/* ==================================================*/
	dbinterface.ExecuteStatement(iConnection, iDatabase, ltIdxStatement)

}
