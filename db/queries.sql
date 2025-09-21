-- name: GetLatestMigration :one
SELECT * From Migration
ORDER BY id DESC
LIMIT 1;

-- name: AddMigration :exec
INSERT INTO MIGRATION (id)
VALUES (?);

-- name: GetApps :many
SELECT * FROM App;

-- name: GetApp :one
SELECT * FROM App
WHERE id = ?;

-- name: GetAppUsers :many
SELECT
	id,
	name,
	password,
	email, 
	u.role role,
	au.role app_role
FROM User u
INNER JOIN AppUser au
	ON u.id = au.user
WHERE au.app = ?;

-- name: GetAppManagers :many
SELECT
	u.id id,
	u.name name,
	u.password password,
	u.email email,
	u.role role
FROM User u
INNER JOIN AppUser au ON
	u.id = au.user AND
	au.role = 2
WHERE au.app = ?;

-- name: CreateApp :one
INSERT INTO App (
	id,
	name,
	description
) VALUES (hex(randomblob(6)), ?, ?) RETURNING *;

-- name: GetAppDomains :many
SELECT * FROM Domain
WHERE app = ?;

-- name: CreateDomain :one
INSERT INTO Domain (
	name,
	app
) VALUES (?, ?) RETURNING *;

-- name: GetUserById :one
SELECT * FROM User
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM User
WHERE email = ?;

-- name: GetUserByUsername :one
SELECT * FROM User
WHERE name = ?;

-- name: GetUserByEmailOrUsername :one
SELECT * FROM User
WHERE
	email = @emailOrUsername OR
	name = @emailOrUsername
LIMIT 1;

-- name: CreateUser :one
INSERT INTO User (
	id,
	name,
	email,
	password
) VALUES (hex(randomblob(6)), ?, ?, ?)
RETURNING *;

-- name: GetSessions :many
SELECT * FROM Session
WHERE user = ?;

-- name: GetSessionById :one
SELECT * FROM Session
WHERE id = ?;
