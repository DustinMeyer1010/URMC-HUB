/* Initats the database with all tables need for the project */

CREATE TABLE IF NOT EXISTS agents (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS links (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	link TEXT NOT NULL,
	image_path TEXT NOT NULL,
	profile_id INTEGER NOT NULL,
	favorite BOOLEAN DEFAULT 0,
	FOREIGN KEY (profile_id) REFERENCES profiles(id)
);

