package db

func GetAppDomains(app string) ([]Domain, error) {
	domains := []Domain{}

	err := Connection().Select(&domains, "SELECT * FROM Domain WHERE app = $1", app)

	return domains, err
}

func CreateDomain(app string, name string) error {
	tx, err := Connection().Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO Domain (name, app) VALUES ($1,$2)", name, app)
	if err != nil {
		return err
	}

	return tx.Commit()
}
