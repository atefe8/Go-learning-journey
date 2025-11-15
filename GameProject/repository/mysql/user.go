package mysql

import (
	"database/sql"
	"fmt"
	"gameproject/entity"
)

func (DB *MYSQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {

	user := entity.User{}
	row := DB.db.QueryRow(`select * from users where phone_number= ?`, phoneNumber)
	error := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.CreatedAt)

	if error != nil {
		if error == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("there is error in checking phone number is unique %w", error)
	}

	return false, nil
}

func (DB *MYSQLDB) Register(user entity.User) (entity.User, error) {
	result, error := DB.db.Exec(`insert into users (name, phone_number) value (?, ?)`, user.Name, user.PhoneNumber)
	if error != nil {
		return entity.User{}, fmt.Errorf("there is error in create user")
	}

	id, _ := result.LastInsertId()
	user.ID = uint(id)

	return user, nil
}
