package database

import (
	"database/sql"

	"github.com/fatih/structs"
	"github.com/lib/pq"
)

const pgUser = "user"
const dbName = "db"
const sslMode = "very-full"

// Connect handles the connection to the PostgreSQL database
func Connect() (*sql.DB, error) {
	connStr := "user=" + pgUser + "dbname=" + dbName + "sslmode=" + sslMode
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetAllAppIDs queries the database to return all apps' IDs
func GetAllAppIDs(db *sql.DB) ([]int, error) {
	appids := []int{}

	rows, err := db.Query("SELECT appid FROM apps")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var appid int

		err = rows.Scan(&appid)
		if err != nil {
			return nil, err
		}

		appids = append(appids, appid)
	}
	return appids, nil
}

// InsertBatch inserts a batch of data in the given table of the given database
func InsertBatch(db *sql.DB, table string, columns []string, toInsert []interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(pq.CopyIn(table, columns...))
	if err != nil {
		return err
	}

	for _, item := range toInsert {
		_, err = stmt.Exec(structs.Values(item)...)
		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
