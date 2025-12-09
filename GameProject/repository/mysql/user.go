package mysql

import (
	"database/sql"
	"fmt"
	"gameproject/entity"
)

func (DB *MYSQLDB) Register(user entity.User) (entity.User, error) {
	result, error := DB.db.Exec(`insert into users (name, phone, password) value (?, ?, ?)`, user.Name, user.PhoneNumber, user.Password)
	if error != nil {
		return entity.User{}, fmt.Errorf("there is error in create user", error.Error())
	}

	id, _ := result.LastInsertId()
	user.ID = uint(id)

	return user, nil
}

func (DB *MYSQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {

	user := entity.User{}
	error := DB.ScanRows(&user, phoneNumber)

	if error != nil {
		if error == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("there is error in checking phone number is unique %w", error)
	}

	return false, nil
}

func (DB *MYSQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {

	user := entity.User{}
	error := DB.ScanRows(&user, phoneNumber)

	if error != nil {
		if error == sql.ErrNoRows {
			return entity.User{}, false, fmt.Errorf("user not found")
		}
		return entity.User{}, false, fmt.Errorf("user not found", error)

	}

	return user, true, nil

}

func (DB *MYSQLDB) ScanRows(user *entity.User, phoneNumber string) error {

	var createdAt []uint8
	row := DB.db.QueryRow(`SELECT * FROM users WHERE phone =?`, phoneNumber)
	return row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)

}
