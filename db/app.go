package db

// Populates App.Domains
func GetApp(id string) (App, error) {
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

	err := Connection().Select(&apps, "SELECT * FROM App WHERE visibility = 1")

	return apps, err
}
