package db

import (
	"AuthInGo/model"
	"database/sql"
	"fmt"
)

// created interface to avoid tight copling of service and repo layer
type UserRepo interface {
	Create() error
	GetUserById() error
	//GetAllUser() error -- todo
	//DeleteUserById() error -- todo
	//UpdateUserById() error -- todo
}

type UserRepoImpl struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}

func (u *UserRepoImpl) Create() error {
	fmt.Println("create in repo")

	query := "INSERT INTO user (username, email, password) VALUES (?, ?, ?)"

	result, err := u.db.Exec(query, "test1", "test1@gmail.com", "test1")
	if err != nil {
		fmt.Println("error executing insert:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("failed to create user: no rows affected")
		return fmt.Errorf("no rows inserted")
	}

	fmt.Println("user created, rows affected:", rowsAffected)
	return nil
}

func (u *UserRepoImpl) GetUserById() error {
	fmt.Println("fetching user by id")

	query := "SELECT id , username , email , password FROM user WHERE Id = ?"

	row := u.db.QueryRow(query, 1)

	user := &model.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found with given id")
			return nil
		} else {
			fmt.Println("error getting row ", err)
			return err
		}
	}

	fmt.Println("user data for id", user)

	return nil
}
