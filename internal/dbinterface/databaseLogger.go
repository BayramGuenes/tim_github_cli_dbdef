package dbinterface

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InsertSvnApptransactRow(iConnection, iDatabase string, iName, iText string) {
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
func InsertApptransactRow(iConnection, iDatabase string, iName, iText string) {
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

}
*/
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
