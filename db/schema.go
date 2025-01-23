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

type Migration struct {
	Id int `db:"id"`
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
