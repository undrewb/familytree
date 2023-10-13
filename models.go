import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	ID          int
	Name        string
	Gender      string
	DateOfBirth time.Time
	DateOfDeath time.Time
	MotherID    sql.NullInt64
	FatherID    sql.NullInt64
	SpouseID    sql.NullInt64
}

type FamilyTreeDB struct {
	DB *sql.DB
}

func (db *FamilyTreeDB) CreatePersonTable() error {
	_, err := db.DB.Exec(`
		CREATE TABLE IF NOT EXISTS person (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			gender TEXT NOT NULL,
			date_of_birth TIMESTAMP NOT NULL,
			date_of_death TIMESTAMP,
			mother_id INTEGER REFERENCES person(id),
			father_id INTEGER REFERENCES person(id),
			spouse_id INTEGER REFERENCES person(id)
		)
	`)
	return err
}

// make this an initdb method
func (db *FamilyTreeDB) InitDB() error {
	db, err := sql.Open("sqlite3", "/path/to/database.db")
	if err != nil {
		// handle error
	}
	defer db.Close()
	err := db.CreatePersonTable()
	if err != nil {
		return err
	}
	return nil
}	

