package db

import (
	"database/sql"
	"hongbao/movies/src/user-srv/models"
	"time"
)

func SelectUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := db.Get(&user, "SELECT * FROM user WHERE `email` = ?", email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func InsertUser(userName string, password string, email string) error {

	today := time.Now().Format("2006-01-02")
	_, err := db.Exec("INSERT INTO `user`(`user_name`,`password`,`create_at`,`email`) VALUES(?,?,?,?)", userName, password, today, email)
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
	_, err := db.Exec("UPDATE `user` SET `user_name` = ? WHERE user_id = ?", userName, userId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateUserEmailProfile(email string, userId int64) error {
	_, err := db.Exec("UPDATE `user` SET `email` = ? WHERE user_id = ?", email, userId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateUserPhoneProfile(phone string, userId int64) error {
	_, err := db.Exec("UPDATE `user` SET `phone` = ? WHERE user_id = ?", phone, userId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
