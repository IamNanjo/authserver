package db

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

func CreateApp(name string, description string) (string, error) {
	id, err := GenerateId(10)
	if err != nil {
		return id, err
	}

	tx, err := Connection().Begin()
	if err != nil {
		return id, err
	}

	tx.Exec("INSERT INTO App (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	tx.Commit()

	return id, nil
}
