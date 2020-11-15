package dbinterface

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/bayramguenes/tim_github_cli_dbdef.git/internal/data"
	_ "github.com/go-sql-driver/mysql"
)

// create database with parametername and delete old database/tables
func CreateDatabaseSchemaIfNotExists(iConnection string, iDatabase string) {
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
	db.Close()
}

// create single table
func CreateTable(iConnection string, iDatabase string, iTable string, iFields string) {
	db, err := sql.Open("mysql", iConnection+iDatabase)
	// fmt.Println(connection+database)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var statement string = "select 1 from " + iTable + " limit 1"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println("*NEWTAB")
		fmt.Println("Tabelle " + iTable + " existiert nicht und wird neu angelegt")
		var statement string = "create table " + iTable
		if len(iFields) > 0 {
			statement = statement + " " + iFields + " character set 'utf8'"
		}
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(statement)
	} else {
		fmt.Println("Tabelle " + iTable + " existiert bereits und wird nicht neu angelegt")
	}
	db.Close()
}

// execute generated statements from
func ExecuteStatement(iConnection, iDatabase string, iStatements []data.Statement) {
	db, err := sql.Open("mysql", iConnection+iDatabase)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < len(iStatements); i++ {
		_, err = db.Exec(iStatements[i].Text)
		if err != nil {
			fmt.Println(err)
			fmt.Println(iStatements[i].Text)
		} else {
			if len(iStatements[i].Info) > 1 {
				fmt.Println(iStatements[i].Info)
			}
		}
	}
	db.Close()
}

/*func InsertApptransactRow(iConnection, iDatabase string, iName, iText string) {
	db, err := sql.Open("mysql", iConnection+iDatabase)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statement := `insert into tim_vt_apptransact (name,text)
                 values(?,?)
	`
	_, err = db.Exec(statement, iName, iText)
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
	db.Close()

}*/

/*func InsertSvnApptransactRow(iConnection, iDatabase string, iName, iText string) {
	db, err := sql.Open("mysql", iConnection+iDatabase)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statement := `insert into tim_vt_svnapptransact (name,text)
                 values(?,?)
	`
	_, err = db.Exec(statement, iName, iText)
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
	db.Close()

}
*/
