package models

import (
	"errors"

	"github.com/atlokeshsk/event-booking/db"
	"github.com/atlokeshsk/event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users(email,password) VALUES(?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hassedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hassedPassword)
	if err != nil {
		return err
	}
	u.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ValidateCredential() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var hashedPassword string
	err := row.Scan(&u.ID, &hashedPassword)
	if err != nil {
		return errors.New(err.Error())
	}
	if !utils.CheckHashPassword(hashedPassword, u.Password) {
		return errors.New("invalid credentials")
	}
	return nil
}
