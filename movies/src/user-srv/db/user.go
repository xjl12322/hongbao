package db

import (
	"database/sql"
	"hongbao/movies/src/user-srv/models"
	"time"
)

func SelectUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err:= db.Get(user,"SELECT * FROM user WHERE `email` = ?",email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
	//row := db.QueryRow("SELECT * FROM user WHERE `email` = ?",email)
	//
	//row.Scan(user)
}

func InsertUser(userName string, password string, email string) error {

	today := time.Now().Format("2006-01-02")
	_, err := db.Exec("INSERT INTO `user`(`name`,`password`,`create_time`,`email`) VALUES(?,?,?,?)", userName, password, today, email)
	return err
}

func SelectUserByPasswordName(email string, password string) (*models.User, error) {

	user := models.User{}
	err := db.Get(&user, "SELECT * FROM user WHERE `email` = ? AND `password` = ? LIMIT 1", email, password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func UpdateUserNameProfile(userName string, userId int64) error {
	_, err := db.Exec("UPDATE `user` SET `name` = ? WHERE id = ?", userName, userId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateUserEmailProfile(email string, userId int64) error {
	_, err := db.Exec("UPDATE user SET `email` = ? WHERE id = ?", email, userId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateUserPhoneProfile(phone string, userId int64) error {
	_, err := db.Exec("UPDATE `user` SET `phone` = ? WHERE id = ?", phone, userId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
