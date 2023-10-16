package conections_pool

import (
	"database/sql"
	"fmt"
	"os"

	"server/config"

	_ "github.com/lib/pq"
)

func SetupDatabaseConnection(env *config.Env) *sql.DB {

	// Connect to the main database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env.HostDB,
		env.PortDB,
		env.User,
		env.Pass,
		env.Database)

	// Connect to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = makeMigrations(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("postgres working...")

	return db
}

func makeMigrations(db *sql.DB) error {

	// Read the SQL migration file and execute it on the database
	c, ioErr := os.ReadFile("database/migrations/sql_migrations.sql")

	if ioErr != nil {
		return ioErr
	}

	sql := string(c)
	_, err := db.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}
