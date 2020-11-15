package dbinterface

import (
	"database/sql"
	"fmt"
	"os"
)

func InsertProducerStateRow(iConnection, iDatabase string, iState, iText string) {
	db, err := sql.Open("mysql", iConnection+iDatabase)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statement := `insert into tim_vt_stateproducer (state,text)
                 values(?,?) 
	`
	_, err = db.Exec(statement, iState, iText)
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
	db.Close()

}

func InsertPullStateRow(iConnection, iDatabase string, iState, iText string) {
	db, err := sql.Open("mysql", iConnection+iDatabase)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statement := `insert into tim_vt_pullstate (state,text)
                 values(?,?) 
	`
	_, err = db.Exec(statement, iState, iText)
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
	db.Close()

}
func InsertOrderStateRow(iConnection, iDatabase string, iState, iText string) {
	db, err := sql.Open("mysql", iConnection+iDatabase)
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}

	statement := `insert into tim_vt_orderstate (state,text)
                 values(?,?) 
	`
	_, err = db.Exec(statement, iState, iText)
	if err != nil {
		fmt.Println(err)
		//
		//os.Exit(1)
	}
	db.Close()

}
