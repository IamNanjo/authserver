package db

import (
	"errors"
	"github.com/IamNanjo/authserver/hash"
)

func GetUserByEmailOrUsername(emailOrUsername string) (User, error) {
	user := User{}
	err := Connection().Get(&user, `
		SELECT *
		FROM User
		WHERE email=$1
			OR name=$1
		`, emailOrUsername)

	return user, err
}

func GetAppUsers(id string) ([]UserWithAppRole, error) {
	users := []UserWithAppRole{}

	err := Connection().Select(
		&users,
		`SELECT
			id,
			name,
			password,
			email,
			u.role as role,
			au.role as app_role
		FROM User u
		INNER JOIN AppUser au
			ON u.id = au.user
		WHERE app=$1`, id,
	)

	return users, err
}

// Returns new user ID
func CreateUser(name string, email string, password string) (string, error) {
	connection := Connection()

	var remainingAttempts = 5
	var id = ""
	var err error

	for {
		if remainingAttempts == 0 {
			return id, errors.New("Could not create unique ID")
		}

		id, err = GenerateId(10)
		if err != nil {
			return id, err
		}

		err = connection.Get(nil, "SELECT 1 FROM User WHERE id = $1", id)
		if err != nil {
			break
		}

		remainingAttempts--
	}

	hashedPassword, err := hash.Hash([]byte(password), nil)
	if err != nil {
		return id, err
	}

	tx, err := connection.Begin()
	if err != nil {
		return id, err
	}

	_, err = tx.Exec("INSERT INTO User (id, name, email, password) VALUES ($1, $2, $3, $4)", id, name, email, hashedPassword)
	if err != nil {

	}

	err = tx.Commit()

	return id, err
}
