package db

func GetUserByEmailOrUsername(emailOrUsername string) (User, error) {
	user := User{}
	err := Connection().Get(&user, `
		SELECT *
		FROM User
		WHERE email=$1
			OR name=$1
		`)

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
