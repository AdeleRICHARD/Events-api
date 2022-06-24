package testdb

import "database/sql"

func TestInsert(db *sql.DB) error {
	sqlStatement := `INSERT INTO "user" (name) VALUES ('Titi')`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
