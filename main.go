package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	queryGetUser = "SELECT id, email FROM users WHERE id=%d;"
)

var (
	dbClient *sql.DB
)

type User struct {
	Id    int64
	Email string
}

func init() {
	var err error
	dbClient, err = sql.Open("mysql", "this is the connection string")
	if err != nil {
		panic(err)
	}
}
func main() {
	user, err := GetUser(123)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Email)
}

func GetUser(id int64) (*User, error) {
	rows, err := dbClient.Query(fmt.Sprintf(queryGetUser, id))
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("user not found")
}
