package dbinterface

import (
	"database/sql"
	"fmt"
	"os"
)

func InsertEntityTypeRow(iConnection, iDatabase string, iState, iText string) {
	db, err := sql.Open("mysql", iConnection+iDatabase)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statement := `insert into tim_vt_entitytype (entitytype,text)
                 values(?,?)
	`
	_, err = db.Exec(statement, iState, iText)
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
	db.Close()

}
