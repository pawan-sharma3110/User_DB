package db

import (
	"database/sql"
	"log"
	"user/modle"
)

func Conect() *sql.DB {
	connStr := `host=localhost port=8080 user=postgres dbname=shope password=Pawan@2003 sslmode=disable`
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	println("Connected to Database successfully..")
	return db

}

func CreateShopTable(db *sql.DB) {
	query := `CREATE TABLE  IF NOT EXISTS shopkeper(
	shope_no INTEGER NOT NULL,
	name TEXT NOT NULL,
	gender TEXT NOT NULL,
	shope_type TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	
)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error in query %v", err)
	}
}

//	func InsertShopkeper(newShop modle.Shopkeeper) (saved modle.Shopkeeper) {
//		db := DbConection()
//		defer db.Close()
//		CreateUserTable(db)
//		query := (` INSERT INTO shopkeper(shope_no,name,gender,shope_type)VALUE($1,$2,$3,$4)`)
//		err := db.QueryRow(query, newShop.ShopeNo, newShop.Name, newShop.Gender, newShop.ShopeType)
//		if err != nil {
//			log.Fatalf("Error in query %v", err)
//		}
//		return newShop
//	}
func InsertShopkeeper(newShop modle.Shopkeeper) (saved modle.Shopkeeper, err error) {
	db := Conect()
	defer db.Close()
	CreateShopTable(db)
	query := `INSERT INTO shopkeper(shope_no, name, gender, shope_type) VALUES ($1, $2, $3,$4 )`
	_, err = db.Exec(query, newShop.ShopeNo, newShop.Name, newShop.Gender, newShop.ShopeType)
	if err != nil {
		log.Fatalf("Error in query: %v", err)
		return saved, err
	}
	return newShop, nil
}
