package db

import (
	"database/sql"
	"log"
	"user/modle"

	_ "github.com/lib/pq"
)

func DbConection() *sql.DB {
	connStr := `host=localhost port=8080 user=postgres password=Pawan@2003 dbname=user sslmode=disable`
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

func CreateUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		gender TEXT NOT NULL,
		mobile TEXT NOT NULL,
		adult BOOL NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error during conection database %v", err)
	}

}

func InserUser(user modle.User) int64 {
	db := DbConection()
	defer db.Close()
	CreateUserTable(db)
	query := (`INSERT INTO users(first_name,last_name,gender,mobile,adult)VALUES($1,$2,$3,$4,$5)RETURNING id`)
	var id int64
	err := db.QueryRow(query, user.FirstName, user.LastName, user.Gender, user.Mobile, user.Adult).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	return id
}
func Users() (u []modle.User, err error) {
	var users []modle.User
	db := DbConection()
	defer db.Close()
	query := `SELECT id,first_name, last_name, gender, mobile, adult, created_at FROM users `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user modle.User
		err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Gender, &user.Mobile, &user.Adult, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
