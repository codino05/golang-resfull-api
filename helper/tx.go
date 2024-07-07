package helper

import "database/sql"

func CommitRollBack(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollBack := tx.Rollback()
		PanicHandler(errorRollBack)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicHandler(errorCommit)

	}
}
