// this is the only file where we have access to database

package users

import (
	"fmt"

	"github.com/hammaad90/bookstore_users-api/datasources/mysql/users_db"
	"github.com/hammaad90/bookstore_users-api/logger"
	"github.com/hammaad90/bookstore_users-api/utils/errors"
)

const (
	errorNoRows           = ""
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?, ?,?);"
	queryGetuser          = "SELECT id, first_name, last_name, email, date_created, status, password FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name = ?, last_name = ?, email= ? where id = ?;"
	queryDeleteUser       = "DELETE FROM  users WHERE id=?;"
	queryFindUserByStatus = "SELECT, id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

// mock db
// var (
// 	userDB = make(map[int64]*User) // what does this map means?
// )

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetuser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to user by id", getErr)
		return errors.NewInternalServerError("Database error")
		// return mysql_utils.ParseError(getErr)
	}

	return nil
}

// where are using pointer to user and pointer to error
func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.NewInternalServerError("Database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to last insert id after creating a new user", err)
		return errors.NewInternalServerError("Database error")
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user statement", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find user statement", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to scan user rows into user struct", err)
			return nil, errors.NewInternalServerError("Database error")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user matching status %s found", status))
	}
	return results, nil
}
