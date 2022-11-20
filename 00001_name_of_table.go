package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(UpPhone, DownPhone)
}

func UpPhone(tx *sql.Tx) error {
	query := `CREATE TABLE "phone" (
    		"id" SERIAL PRIMARY KEY,
    		"user_name" VARCHAR(200),
    		"phone"	VARCHAR(20));`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func DownPhone(tx *sql.Tx) error {
	query := `DROP TABLE phone;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
