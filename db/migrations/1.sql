-- Migration
CREATE TABLE IF NOT EXISTS Migration (id INTEGER PRIMARY KEY);
INSERT INTO Migration VALUES (1);

-- App
CREATE TABLE IF NOT EXISTS App (
	id TEXT PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	description TEXT DEFAULT ''
);

-- Domain
CREATE TABLE IF NOT EXISTS Domain (
	name TEXT PRIMARY KEY NOT NULL,
	app TEXT NOT NULL,
	FOREIGN KEY (app) REFERENCES App(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- User
CREATE TABLE IF NOT EXISTS User (
	id TEXT PRIMARY KEY NOT NULL,
	name TEXT NOT NULL UNIQUE COLLATE NOCASE,
	password TEXT NOT NULL,
	email TEXT DEFAULT '' UNIQUE,
	role INTEGER NOT NULL DEFAULT 1
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_name ON User (name COLLATE NOCASE);

-- AppManager
CREATE TABLE IF NOT EXISTS AppManager (
	app TEXT NOT NULL,
	user TEXT NOT NULL,
	PRIMARY KEY (app, user),
	FOREIGN KEY (app) REFERENCES App(id) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY (user) REFERENCES User(id) ON UPDATE CASCADE ON DELETE CASCADE
);
