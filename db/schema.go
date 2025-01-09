package db

var Role = struct {
	User          int
	Manager       int
	Administrator int
}{
	User:          1,
	Manager:       2,
	Administrator: 3,
}

var AppVisibility = struct {
	Hidden  int
	Visible int
}{
	Hidden:  0,
	Visible: 1,
}

type Domain struct {
	Name string `db:"name"`
	App  string `db:"app"`
}

type App struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Visibility  int    `db:"visibility"`
	Domains     []Domain
}

type User struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Role     int    `db:"role"`
}

type AppUser struct {
	App  string `db:"app"`
	User string `db:"user"`
	Role int    `db:"role"`
}

type UserWithAppRole struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Role     int    `db:"role"`
	AppRole  int    `db:"app_role"`
}

type AppWithUsers struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Visibility  int    `db:"visibility"`
	Users       []UserWithAppRole
}

var Schema = `
CREATE TABLE IF NOT EXISTS App (
	id TEXT PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	description TEXT DEFAULT '',
	visibility INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS Domain (
	name TEXT PRIMARY KEY NOT NULL,
	app TEXT NOT NULL,
	FOREIGN KEY (app) REFERENCES App(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS User (
	id TEXT PRIMARY KEY NOT NULL,
	name TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	email TEXT DEFAULT '' UNIQUE,
	role INTEGER NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS AppUser (
	app TEXT NOT NULL,
	user TEXT NOT NULL,
	role INTEGER NOT NULL DEFAULT 1,
	PRIMARY KEY (app, user),
	FOREIGN KEY (app) REFERENCES App(id) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY (user) REFERENCES User(id) ON UPDATE CASCADE ON DELETE CASCADE
);
`
