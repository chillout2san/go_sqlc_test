package gosqlctest

import (
	"context"
	"database/sql"
	"log"

	"tutorial.sqlc.dev/app/tutorial"
)

func getAuthor() tutorial.Author {
	db, err := sql.Open("mysql", "dummy")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer db.Close()

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal("error beginning transaction: ", err)
	}
	defer tx.Rollback()

	qtx := tutorial.New(db).WithTx(tx)

	author, err := qtx.GetAuthor(ctx, "dummy_id")
	if err != nil {
		log.Fatal("error getting author: ", err)
	}

	tx.Commit()
	return author
}
