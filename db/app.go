package db

import "errors"

// Populates App.Domains
func GetAppById(id string) (App, error) {
	app := App{}

	conn := Connection()
	err := conn.Get(&app, "SELECT * FROM App WHERE id=$1 LIMIT 1", id)
	if err != nil {
		return app, err
	}
	err = conn.Select(&app.Domains, "SELECT * FROM Domain WHERE app=$1", id)
	if err != nil {
		return app, err
	}

	return app, err
}

// Does not populate App.Domains
func GetApps() ([]App, error) {
	apps := []App{}

	err := Connection().Select(&apps, "SELECT * FROM App")

	return apps, err
}

func GetAppManagers(app string) ([]User, error) {
	managers := []User{}

	err := Connection().Select(&managers,
		`SELECT
			u.id as id,
			u.name as name,
			u.password as password,
			u.email,
			u.role as role
		FROM User u
		INNER JOIN AppManager am
		ON u.id = am.user
		WHERE am.app = $1`,
		app,
	)

	return managers, err
}

// Returns the ID of the new App
func CreateApp(name string, description string) (string, error) {
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

	tx, err := connection.Begin()
	if err != nil {
		return id, err
	}

	tx.Exec("INSERT INTO App (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	tx.Commit()

	return id, nil
}
