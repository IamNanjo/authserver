package db

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
	Domains     []Domain
}

type User struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Role     int    `db:"role"`
}

type Role struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
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

type Session struct {
	Id        string
	User      string
	Device    string
	LastUsed  int
	ExpiresAt int
	Token     string
}

type OAuthToken struct {
	Id        int
	User      string
	Token     string
	Provider  string
	CreatedAt int
	ExpiresAt int
}
